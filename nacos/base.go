// Copyright (c) 2018-2028 Dunyu All Rights Reserved.
//
// Author      : https://www.wengold.net
// Email       : support@wengold.net
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2022/05/11   yangping     New version
// -------------------------------------------------------------------

package nacos

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/model"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/wengoldx/wing/comm"
	"github.com/wengoldx/wing/invar"
	"github.com/wengoldx/wing/logger"
	"regexp"
	"strconv"
	"strings"
)

// -------- Auto Register Define --------

// Server register informations
type ServerItem struct {
	Name     string         // Server name, same as beego app name
	Group    string         // Server group, range in [GP_BASIC, GP_IFSC, GP_DTE, GP_CWS]
	Callback ServerCallback // Server register datas changed callback
}

// Callback to listen server address and port changes
type ServerCallback func(svr string, addr string, port, httpport int)

// Register current server to nacos, you must set configs in app.conf
//	@return - *ServerStub nacos server stub instance
//
//	`NOTICE` : nacos config as follows.
//
// ----
//
//	; Nacos remote server host
//	nacossvr = "10.239.40.24"
//
//	; Server nacos group name
//	nacosgp = "group.ifsc"
//
//	[dev]
//	; Inner net wlan perfix for dev servers access
//	nacosaddr = "10.239.20."
//
//	; Inner net port for grpc access
//	nacosport = 3000
//
//	[prod]
//	; Inner net wlan perfix for prod servers access
//	nacosaddr = "10.239.40."
//
//	; Inner net port for grpc access
//	nacosport = 3000
func RegisterServer(opts ...string) *ServerStub {
	svr, group := parseOptions(opts...)

	// Local server listing ip proxy
	ipprefix := beego.AppConfig.String(configKeyAddr)
	if len(strings.Split(ipprefix, ".")) < 3 {
		panic("Invalid ip proxy prefix:" + ipprefix)
	}

	addr, err := matchProxyIP(ipprefix)
	if err != nil {
		panic("Find proxy local ip, err:" + err.Error())
	}

	// Server access port for grpc, it maybe same as httpport config
	// when the local server not support grpc but for http
	port, err := beego.AppConfig.Int(configKeyPort)
	if err != nil || port < 3000 /* remain 0 ~ 3000 */ {
		panic("Not found port number or less 3000!")
	}

	// Namespace id of local server
	ns := comm.Condition(beego.BConfig.RunMode == "prod", NS_PROD, NS_DEV).(string)

	// Generate nacos server stub and setup it
	stub := NewServerStub(ns, svr)
	if err := stub.Setup(); err != nil {
		panic(err)
	}

	// Fixed app name as nacos server name to register,
	// and pick server port from config 'nacosport' not form 'httpport' value,
	// becase it maybe support either grpc or http hanlder to accesse.
	//
	// And here not use cluster name, please keep it empty!
	app := beego.BConfig.AppName
	if err := stub.Register(app, addr, uint64(port), group); err != nil {
		panic(err)
	}

	logmsg := fmt.Sprintf("%s@%s:%v", app, addr, port)
	logger.I("Registered server on", logmsg)
	return stub
}

// Listing services address and port changes, it will call the callback
// immediately to return target service host when them allready registerd
// to service central of nacos.
//	@params servers []*ServerItem target server registry informations
func (ss *ServerStub) ListenServers(servers []*ServerItem) {
	for _, s := range servers {
		if err := ss.Subscribe(s.Name, s.OnChanged, s.Group); err != nil {
			panic("Subscribe server " + s.Name + " err:" + err.Error())
		}
	}
}

// Subscribe callback called when target service address and port changed
func (si *ServerItem) OnChanged(services []model.Instance, err error) {
	if err != nil {
		logger.E("Received server", si.Name, "change, err:", err)
		return
	}

	if len(services) > 0 {
		addr, port := services[0].Ip, services[0].Port

		// Paser httpport from metadata map if it exist
		meta, httpport := services[0].Metadata, 0
		if meta != nil {
			if hp, ok := meta[configKeyHPort]; ok {
				httpport, _ = strconv.Atoi(hp)
			}
		}

		logmsg := fmt.Sprintf("%s@%s:%v - httpport:%v", si.Name, addr, port, httpport)
		logger.I("Update server to", logmsg)
		si.Callback(si.Name, addr, int(port), httpport)
	}
}

// -------- Config Service Define --------

// Meta config informations
type MetaConfig struct {
	Group     string                        // Server group, range in [GP_BASIC, GP_IFSC, GP_DTE, GP_CWS]
	Stub      *ConfigStub                   // Nacos config client instance
	Callbacks map[string]MetaConfigCallback // Meta config changed callback maps, key is dataid
}

// Callback to listen server address and port changes
type MetaConfigCallback func(dataId, data string)

// Generate a meta config client to get or listen configs changes
//	@return - *MetaConfig nacos config client instance
//
//	`NOTICE` : nacos config as follows.
//
// ----
//
//	; Nacos remote server host
//	nacossvr = "10.239.40.24"
//
//	; Server nacos group name
//	nacosgp = "group.ifsc"
func GenMetaConfig(opts ...string) *MetaConfig {
	svr, group := parseOptions(opts...)

	// Namespace id of meta configs
	ns := comm.Condition(beego.BConfig.RunMode == "prod", NS_PROD, NS_DEV).(string)

	// Generate nacos config stub and setup it
	stub := NewConfigStub(ns, svr)
	if err := stub.Setup(); err != nil {
		panic("Gen config stub, err:" + err.Error())
	}

	cbs := make(map[string]MetaConfigCallback)
	return &MetaConfig{
		Group: group, Stub: stub, Callbacks: cbs,
	}
}

// Get and listing the configs of indicated dataIds
func (mc *MetaConfig) ListenConfigs(dataIds []string, cb MetaConfigCallback) {
	for _, dataId := range dataIds {
		if dataId == "" || cb == nil {
			continue
		}
		mc.ListenConfig(dataId, cb)
	}
}

// Get and listing the config of indicated dataId
func (mc *MetaConfig) ListenConfig(dataId string, cb MetaConfigCallback) {
	mc.Callbacks[dataId] = cb // cache callback

	// get config first before listing
	data, err := mc.Stub.GetString(dataId, mc.Group)
	if err != nil {
		panic("Get config " + dataId + "err: " + err.Error())
	}
	cb(dataId, data)

	// listing config changes
	logger.I("Start listing config { dataId:", dataId, "group:", mc.Group, "}")
	mc.Stub.Listen(dataId, mc.Group, mc.OnChanged)
}

// Listing callback called when target configs changed
func (mc *MetaConfig) OnChanged(namespace, group, dataId, data string) {
	if (namespace != NS_DEV && namespace != NS_PROD) || group != mc.Group {
		logger.E("Invalid meta config ns:", namespace, "or group:", group)
		return
	}

	if callback, ok := mc.Callbacks[dataId]; ok {
		logger.I("Update config dataId", dataId, "to:", data)
		callback(dataId, data)
	}
}

// ---------------------------------------

// Get and check register server's nacos informations
//	@return - string nacos remote server ip
//			- string group name of local server
func NacosSvrConfigs() (string, string) {
	svr := beego.AppConfig.String(configKeySvr)
	if svr == "" {
		panic("Not found nacos server host!")
	}

	gp := beego.AppConfig.String(configKeyGroup)
	if !(gp == GP_BASIC || gp == GP_IFSC || gp == GP_DTE || gp == GP_CWS) {
		panic("Invalid register cluster group!")
	}
	return svr, gp
}

// Parse input params and check server and group names
func parseOptions(opts ...string) (string, string) {
	if cnt := len(opts); cnt >= 2 /* 0:server, 1:group */ {
		svr, gp := opts[0], opts[1]
		validgp := (gp == GP_BASIC || gp == GP_IFSC || gp == GP_DTE || gp == GP_CWS)
		if svr == "" || !validgp {
			panic("Invalid svr or group params!")
		}
		return svr, gp
	}
	return NacosSvrConfigs()
}

// Generate nacos client config, contain nacos remote server and
// current business servers configs, this client keep alive with
// 5s pingpong heartbeat and output logs on warn leven.
//
//	`NOTICE`
//
// - Remote direct nacos server need access on http://{svr}:8848/nacos
//
// - Nginx proxy vip server need access on http://{svr}:3608/nacos
func genClientParam(ns, svr string) vo.NacosClientParam {
	sc := []constant.ServerConfig{
		constant.ServerConfig{
			Scheme: "http", ContextPath: "/nacos", IpAddr: svr, Port: 3608,
		},
	}

	// logs config
	logcfg := &constant.ClientLogRollingConfig{
		MaxSize: 10, MaxBackups: 10, // max 10 files and each max 10MB
	}

	// client config
	cc := &constant.ClientConfig{
		NamespaceId:         ns,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              nacosDirLogs,
		CacheDir:            nacosDirCache,
		LogRollingConfig:    logcfg,
		LogLevel:            nacosLogLevel,
		Username:            nacosSysSecure, // secure account
		Password:            nacosSysSecure, // secure passowrd
	}

	return vo.NacosClientParam{
		ClientConfig: cc, ServerConfigs: sc,
	}
}

// Parse and return the local register IP that meets the conditions
func matchProxyIP(proxy string) (string, error) {
	segment := `.((2((5[0-5])|([0-4]\d)))|([0-1]?\d{1,2}))`
	condition := proxy + segment
	reg, err := regexp.Compile(condition)
	if err != nil {
		logger.E("Compile regular err:", err)
		return "", err
	}

	matchips, err := comm.GetLocalIPs()
	if err != nil {
		return "", err
	}

	for _, v := range matchips {
		// Find the first matched not-empty ip and return
		if ip := reg.FindString(v); ip != "" {
			return ip, nil
		}
	}

	logger.E("Compile regular err:", err)
	return "", invar.ErrNotFound
}
