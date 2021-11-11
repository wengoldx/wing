// Copyright (c) 2018-2019 WING All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// -------------------------------------------------------------------

package mvc

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/wengoldx/wing/invar"
	"github.com/wengoldx/wing/logger"
	"reflect"
	"strings"
	// ----------------------------------------
	// NOTIC :
	//
	// import the follows database drivers when using WingProvider.
	//
	// _ "github.com/go-sql-driver/mysql"   // usr fot mysql
	// _ "github.com/denisenkom/go-mssqldb" // use for sql server 2017 ~ 2019
	//
	// ----------------------------------------
)

// WingProvider content provider to support database utils
type WingProvider struct {
	Conn *sql.DB
}

// ScanCallback use for scan query result from rows
type ScanCallback func(rows *sql.Rows) error

const (
	/* MySQL */
	mysqlConfigUser    = "mysql::user"     // configs key of mysql database user
	mysqlConfigPwd     = "mysql::pwd"      // configs key of mysql database password
	mysqlConfigHost    = "mysql::host"     // configs key of mysql database host and port
	mysqlConfigName    = "mysql::name"     // configs key of mysql database name
	mysqlConfigDevUser = "mysql-dev::user" // DEV : configs key of mysql database user
	mysqlConfigDevPwd  = "mysql-dev::pwd"  // DEV : configs key of mysql database password
	mysqlConfigDevHost = "mysql-dev::host" // DEV : configs key of mysql database host and port
	mysqlConfigDevName = "mysql-dev::name" // DEV : configs key of mysql database name

	/* Microsoft SQL Server */
	mssqlConfigUser    = "mssql::user"        // configs key of mssql database user
	mssqlConfigPwd     = "mssql::pwd"         // configs key of mssql database password
	mssqlConfigHost    = "mssql::host"        // configs key of mssql database server host
	mssqlConfigPort    = "mssql::port"        // configs key of mssql database port
	mssqlConfigName    = "mssql::name"        // configs key of mssql database name
	mssqlConfigTout    = "mssql::timeout"     // configs key of mssql database connect timeout
	mssqlConfigDevUser = "mssql-dev::user"    // DEV : configs key of mssql database user
	mssqlConfigDevPwd  = "mssql-dev::pwd"     // DEV : configs key of mssql database password
	mssqlConfigDevHost = "mssql-dev::host"    // DEV : configs key of mssql database server host
	mssqlConfigDevPort = "mssql-dev::port"    // DEV : configs key of mssql database port
	mssqlConfigDevName = "mssql-dev::name"    // DEV : configs key of mssql database name
	mssqlConfigDevTout = "mssql-dev::timeout" // DEV : configs key of mssql database connect timeout

	// Mysql Server database source name for local connection
	mysqldsnLocal = "%s:%s@/%s?charset=%s"

	// Mysql Server database source name for tcp connection
	mysqldsnTcp = "%s:%s@tcp(%s)/%s?charset=%s"

	// Microsoft SQL Server database source name
	mssqldsn = "server=%s;port=%d;database=%s;user id=%s;password=%s;Connection Timeout=%d;Connect Timeout=%d;"
)

var (
	// WingHelper content provider to hold database connections,
	// it will nil before mvc.OpenMySQL() called.
	WingHelper *WingProvider

	// MssqlHelper content provider to hold mssql database connections,
	// it will nil before mvc.OpenMssql() called.
	MssqlHelper *WingProvider

	// limitPageItems limit to show lits items in one page, default is 50,
	// you can use SetLimitPageItems() to change the limit value.
	limitPageItems = 50
)

// readMySQLCofnigs read mysql database params from config file,
// than verify them if empty except host.
func readMySQLCofnigs() (string, string, string, string, error) {
	user, pwd, host, name := "", "", "", ""
	if beego.BConfig.RunMode == "dev" {
		user = beego.AppConfig.String(mysqlConfigDevUser)
		pwd = beego.AppConfig.String(mysqlConfigDevPwd)
		host = beego.AppConfig.String(mysqlConfigDevHost)
		name = beego.AppConfig.String(mysqlConfigDevName)
	}

	// if curren mode is dev and not found [mysql-dev] session,
	// try to load configs from [mysql] session same as prod mode.
	invalidConfigs := (user == "" && pwd == "" && name == "")
	if invalidConfigs {
		user = beego.AppConfig.String(mysqlConfigUser)
		pwd = beego.AppConfig.String(mysqlConfigPwd)
		host = beego.AppConfig.String(mysqlConfigHost)
		name = beego.AppConfig.String(mysqlConfigName)
	}

	if user == "" || pwd == "" || name == "" {
		return "", "", "", "", invar.ErrInvalidConfigs
	}
	return user, pwd, host, name, nil
}

// readMssqlCofnigs read mssql database params from config file,
// than verify them if empty.
func readMssqlCofnigs() (string, string, string, int, string, int, error) {
	user, pwd, host, port, name, timeout := "", "", "", 0, "", 0
	if beego.BConfig.RunMode == "dev" {
		user = beego.AppConfig.String(mssqlConfigDevUser)
		pwd = beego.AppConfig.String(mssqlConfigDevPwd)
		host = beego.AppConfig.DefaultString(mssqlConfigDevHost, "127.0.0.1")
		port = beego.AppConfig.DefaultInt(mssqlConfigDevPort, 1433)
		name = beego.AppConfig.String(mssqlConfigDevName)
		timeout = beego.AppConfig.DefaultInt(mssqlConfigDevTout, 600)
	}

	// if curren mode is dev and not found [mssql-dev] session,
	// try to load configs from [mssql] session same as prod mode.
	invalidConfigs := (user == "" && pwd == "" && name == "")
	if invalidConfigs {
		user = beego.AppConfig.String(mssqlConfigUser)
		pwd = beego.AppConfig.String(mssqlConfigPwd)
		host = beego.AppConfig.DefaultString(mssqlConfigHost, "127.0.0.1")
		port = beego.AppConfig.DefaultInt(mssqlConfigPort, 1433)
		name = beego.AppConfig.String(mssqlConfigName)
		timeout = beego.AppConfig.DefaultInt(mssqlConfigTout, 600)
	}

	if user == "" || pwd == "" || name == "" {
		return "", "", "", 0, "", 0, invar.ErrInvalidConfigs
	}
	return user, pwd, host, port, name, timeout, nil
}

// OpenMySQL connect database and check ping result,
// the connections holded by mvc.WingHelper object,
// the charset maybe 'utf8' or 'utf8mb4' same as database set.
//
// NOTICE : you must config database params in /conf/app.config file as:
//	~
//	[mysql]
//	host = "127.0.0.1:3306"
//	name = "sampledb"
//	user = "root"
//	pwd  = "123456"
//	~
//
// Or, config as dev mode as:
//	~
//	[mysql-dev]
//	host = "127.0.0.1:3306"
//	name = "sampledb"
//	user = "root"
//	pwd  = "123456"
//	~
//
// Or both or them for dev and prod mode.
// It will load configs from [mssql-dev] session first, if not
// found, try agen load from [mssql] session.
func OpenMySQL(charset string) error {
	dbuser, dbpwd, dbhost, dbname, err := readMySQLCofnigs()
	if err != nil {
		return err
	}

	dsn := ""
	if len(dbhost) > 0 /* check database host whether using TCP to connect */ {
		// conntect with remote host database server
		dsn = fmt.Sprintf(mysqldsnTcp, dbuser, dbpwd, dbhost, dbname, charset)
	} else {
		// just connect local database server
		dsn = fmt.Sprintf(mysqldsnLocal, dbuser, dbpwd, dbname, charset)
	}
	logger.D("MySQL DSN:", dsn)

	// open and connect database
	con, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// check database validable
	if err = con.Ping(); err != nil {
		return err
	}

	con.SetMaxIdleConns(100)
	con.SetMaxOpenConns(100)
	WingHelper = &WingProvider{con}
	return nil
}

// OpenMssql connect mssql database and check ping result,
// the connections holded by mvc.MssqlHelper object,
// the charset maybe 'utf8' or 'utf8mb4' same as database set.
//
// NOTICE : you must config database params in /conf/app.config file as:
//	~
//	[mssql]
//	host    = "127.0.0.1"
//	port    = 1433
//	name    = "sampledb"
//	user    = "sa"
//	pwd     = "123456"
//	timeout = 600
//	~
//
// Or, config as dev mode as:
//	~
//	[mssql-dev]
//	host    = "127.0.0.1"
//	port    = 1433
//	name    = "sampledb"
//	user    = "sa"
//	pwd     = "123456"
//	timeout = 600
//	~
//
// Or both or them for dev and prod mode.
// It will load configs from [mssql-dev] session first, if not
// found, try agen load from [mssql] session.
func OpenMssql(charset string) error {
	user, pwd, server, port, dbn, to, err := readMssqlCofnigs()
	if err != nil {
		return err
	}

	// get connection and connect timeouts
	dts := []int{600, 600}
	if to > 0 {
		dts[0] = to
		dts[1] = to
	}

	driver := "mssql"
	dsn := fmt.Sprintf(mssqldsn, server, port, dbn, user, pwd, dts[0], dts[1])
	logger.D("SQL Server DSN:", dsn)

	// open and connect database
	con, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}

	// check database validable
	if err = con.Ping(); err != nil {
		return err
	}

	con.SetMaxIdleConns(100)
	con.SetMaxOpenConns(100)
	MssqlHelper = &WingProvider{con}
	return nil
}

// SetLimitPageItems set global setting of limit items in page,
// the input value must range in (0, 1000].
func SetLimitPageItems(limit int) {
	if limit > 0 && limit <= 1000 {
		limitPageItems = limit
	}
}

// Stub return content provider connection
func (w *WingProvider) Stub() *sql.DB {
	return w.Conn
}

// Query call sql.Query()
func (w *WingProvider) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return w.Conn.Query(query, args...)
}

// Prepare call sql.Prepare()
func (w *WingProvider) Prepare(query string) (*sql.Stmt, error) {
	return w.Conn.Prepare(query)
}

// IsEmpty call sql.Query() to check target data if exist
func (w *WingProvider) IsEmpty(query string, args ...interface{}) (bool, error) {
	rows, err := w.Conn.Query(query, args...)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	return !rows.Next(), nil
}

// QueryOne call sql.Query() to query one record
func (w *WingProvider) QueryOne(query string, cb ScanCallback, args ...interface{}) error {
	rows, err := w.Conn.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return invar.ErrNotFound
	}
	rows.Columns()
	return cb(rows)
}

// QueryArray call sql.Query() to query multi records
func (w *WingProvider) QueryArray(query string, cb ScanCallback, args ...interface{}) error {
	rows, err := w.Conn.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Columns()
		if err := cb(rows); err != nil {
			return err
		}
	}
	return nil
}

// Insert call sql.Prepare() and stmt.Exec() to insert a new record
func (w *WingProvider) Insert(query string, args ...interface{}) (int64, error) {
	stmt, err := w.Conn.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(args...)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// Execute call sql.Prepare() and stmt.Exec() to update or delete records
func (w *WingProvider) Execute(query string, args ...interface{}) error {
	_, err := w.ExecuteWithResult(query, args...)
	return err
}

// Execute call sql.Prepare() and stmt.Exec() to update or delete records
func (w *WingProvider) ExecuteWithResult(query string, args ...interface{}) (int64, error) {
	stmt, err := w.Conn.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(args...)
	if err != nil {
		return 0, err
	}
	return w.Affected(result)
}

// AppendLike append like keyword end of sql string,
// DON'T call it after AppendLimit()
func (w *WingProvider) AppendLike(query, filed, keyword string, and ...bool) string {
	if len(and) > 0 && and[0] {
		return query + " AND " + filed + " LIKE '%%" + keyword + "%%'"
	}
	return query + " WHERE " + filed + " LIKE '%%" + keyword + "%%'"
}

// AppendLimit append page limitation end of sql string,
// DON'T call it before AppendLick()
func (w *WingProvider) AppendLimit(query string, page int) string {
	offset, items := page*limitPageItems, limitPageItems
	return query + " LIMIT " + fmt.Sprintf("%d, %d", offset, items)
}

// AppendLikeLimit append like keyword and limit end of sql string
func (w *WingProvider) AppendLikeLimit(query, filed, keyword string, page int, and ...bool) string {
	return w.AppendLimit(w.AppendLike(query, filed, keyword, and...), page)
}

// CheckAffected append page limitation end of sql string
func (w *WingProvider) Affected(result sql.Result) (int64, error) {
	row, err := result.RowsAffected()
	if err != nil || row == 0 {
		return 0, invar.ErrNotChanged
	}
	return row, nil
}

// FormatSets format update sets for sql update
//
// [CODE:]
//	sets := w.FormatSets(struct {
//		StringFiled string
//		EmptyString string
//		BlankString string
//		TrimString  string
//		IntFiled    int
//		I32Filed    int32
//		I64Filed    int64
//		F32Filed    float32
//		F64Filed    float64
//		BoolFiled   bool
//	}{"string", "", " ", " trim ", 123, 32, 64, 32.123, 64.123, true})
//	// sets: stringfiled='string', trimstring='trim', intfiled=123, i32filed=32, i64filed=64, f32filed=32.123, f64filed=64.123, boolfiled=true
//	logger.I("sets:", sets)
// [CODE]
func (w *WingProvider) FormatSets(updates interface{}) string {
	sets := []string{}
	keys, values := reflect.TypeOf(updates), reflect.ValueOf(updates)
	for i := 0; i < keys.NumField(); i++ {
		name := strings.ToLower(keys.Field(i).Name)
		if name == "" {
			continue
		}

		value := values.Field(i).Interface()
		switch value.(type) {
		case bool:
			sets = append(sets, fmt.Sprintf(name+"=%v", value))
		case invar.Bool:
			boolvalue := value.(invar.Bool)
			if boolvalue != invar.BNone {
				truevalue := (boolvalue == invar.BTrue)
				sets = append(sets, fmt.Sprintf(name+"=%v", truevalue))
			}
		case string:
			trimvalue := strings.Trim(value.(string), " ")
			if trimvalue != "" { // filter empty string fields
				sets = append(sets, fmt.Sprintf(name+"='%s'", trimvalue))
			}
		case int, int8, int16, int32, int64, float32, float64,
			invar.Status, invar.Box, invar.Role, invar.Limit, invar.Lang, invar.Kind:
			if fmt.Sprintf("%v", value) != "0" { // filter 0 fields
				sets = append(sets, fmt.Sprintf(name+"=%v", value))
			}
		}
	}
	return strings.Join(sets, ", ")
}

// Atomicity call sql.Begin() , tx.Rollback() and tx.Commit() to start one transaction
// All operations in a transaction are either completed or not completed. They will not end in an intermediate link.
// If an error occurs during the execution of the transaction, it will be rolled back to the state before the transaction starts
//
// [CODE:]
//		args := make(map[string][]interface{})
//		args[query] = []interface{}{...arg}
// [CODE]
func (w *WingProvider) Atomicity(args map[string][]interface{}) error {
	atomic, err := w.Conn.Begin()
	if err != nil {
		return err
	}
	defer atomic.Rollback()
	for query, arg := range args {
		result, err := atomic.Exec(query, arg...)
		if err != nil {
			return err
		}
		if _, err = w.Affected(result); err != nil {
			return err
		}
	}
	return atomic.Commit()
}
