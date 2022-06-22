// Copyright (c) 2018-2028 Dunyu All Rights Reserved.
//
// Author      : https://www.wengold.net
// Email       : support@wengold.net
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// -------------------------------------------------------------------

package mvc

import (
	"encoding/json"
	"encoding/xml"
	"github.com/wengoldx/wing/logger"
	"strings"
)

// WAuthController the extend controller base on WingController to support
// auth account from http headers, the client caller must append two headers
// before post request if expect the controller method enable execute token
// authentication from header.
//
// * Authoration : It must fixed keyword as WENGOLD
//
// * Token : Authenticate JWT token responsed by login success
//
//
// `USAGE` :
//
// The validator register code of input params struct see WingController description,
// but the restful auth api of router method as follow usecase 1 and 2.
//
// ---
//
// `controller.go`
//
//	// define custom controller using header auth function
//	type AccController struct {
//		mvc.WAuthController
//	}
//
//	func init() {
//		mvc.GAuthHandlerFunc = func(token string) (string, string) {
//			// decode and verify token string, than return indecated
//			// account uuid and password.
//			return "account uuid", "account password"
//		}
//	}
//
// `USECASE 1. Auth account and Parse input params`
//
//	//	@Description Restful api bind with /login on POST method
//	//	@Param Authoration header string true "WENGOLD"
//	//	@Param Token       header string true "Authentication token"
//	//	@Param data body types.Accout true "input param description"
//	//	@Success 200 {string} "response data description"
//	//	@router /login [post]
//	func (c *AccController) AccLogin() {
//		ps := &types.Accout{}
//		c.DoAfterValidated(ps, func(uuid, pwd string) (int, interface{}) {
//			// do same business with input NO-EMPTY account uuid,
//			// directe use c and ps param in this methed.
//			// ...
//			return http.StatusOK, "Done business"
//		} , false /* not limit error message even code is 40x */)
//	}
//
// `USECASE 2. Auth account on GET http method`
//
//	//	@Description Restful api bind with /detail on GET method
//	//	@Param Authoration header string true "WENGOLD"
//	//	@Param Token       header string true "Authentication token"
//	//	@Success 200 {types.Detail} "response data description"
//	//	@router /detail [get]
//	func (c *AccController) AccDetail() {
//		if uuid, _ := c.AuthRequestHeader(); uuid != "" {
//			// use c.BindValue("fieldkey", out) parse params from url
//			c.ResponJSON(service.AccDetail(uuid))
//		}
//	}
//
// `USECASE 3. No-Auth and Use WingController`
//
//	//	@Description Restful api bind with /update on POST method
//	//	@Param data body types.UserInfo true "input param description"
//	//	@Success 200
//	//	@router /update [post]
//	func (c *AccController) AccUpdate() {
//		ps := &types.UserInfo{}
//		c.WingController.DoAfterValidated(ps, func() (int, interface{}) {
//			// directe use c and ps param in this methed.
//			// ...
//			return http.StatusOK, nil
//		} , false /* not limit error message even code is 40x */)
//	}
//
// `USECASE 4. No-Auth and Custom code`
//
//	//	@Description Restful api bind with /list on GET method
//	//	@Success 200 {object} []types.Account "response data description"
//	//	@router /list [get]
//	func (c *AccController) AccList() {
//		// do same business without auth and input params
//		c.ResponJSON(service.AccList())
//	}
type WAuthController struct {
	WingController
}

// NextFunc2 do action after input params validated, it decode token to get account uuid.
type NextFunc2 func(uuid, pwd string) (int, interface{})

// AuthFunc auth request token from http header and returen account secures.
type AuthHandlerFunc func(token string) (string, string)

// Global handler function to auth token from http header
var GAuthHandlerFunc AuthHandlerFunc

// Get authoration and token from http header, than verify it and return account secures.
func (c *WAuthController) AuthRequestHeader() (string, string) {
	if GAuthHandlerFunc == nil {
		c.E403Denind("Controller not set global auth hander!")
		return "", ""
	}

	// check authoration secure key
	authoration := strings.ToUpper(c.Ctx.Request.Header.Get("Authoration"))
	if authoration != "WENGOLD" {
		c.E401Unauthed("Invalid header authoration: " + authoration)
		return "", ""
	}

	// get token from header and verify it
	if token := c.Ctx.Request.Header.Get("Token"); token != "" {
		if uuid, pwd := GAuthHandlerFunc(token); uuid != "" {
			logger.D("Authenticated account:", uuid)
			return uuid, pwd
		}
	}

	// token is empty or invalid, response unauthed
	c.E401Unauthed("Unauthed header token!")
	return "", ""
}

// DoAfterValidated do bussiness action after success validate the given json data.
func (c *WAuthController) DoAfterValidated(ps interface{}, nextFunc2 NextFunc2, option ...interface{}) {
	if uuid, pwd := c.AuthRequestHeader(); uuid != "" {
		isprotect := !(len(option) > 0 && !option[0].(bool))
		c.doAfterParsedOrValidated("json", ps, nextFunc2, uuid, pwd, true, isprotect)
	}
}

// DoAfterUnmarshal do bussiness action after success unmarshaled the given json data.
func (c *WAuthController) DoAfterUnmarshal(ps interface{}, nextFunc2 NextFunc2, option ...interface{}) {
	if uuid, pwd := c.AuthRequestHeader(); uuid != "" {
		isprotect := !(len(option) > 0 && !option[0].(bool))
		c.doAfterParsedOrValidated("json", ps, nextFunc2, uuid, pwd, false, isprotect)
	}
}

// DoAfterValidatedXml do bussiness action after success validate the given xml data.
func (c *WAuthController) DoAfterValidatedXml(ps interface{}, nextFunc2 NextFunc2, option ...interface{}) {
	if uuid, pwd := c.AuthRequestHeader(); uuid != "" {
		isprotect := !(len(option) > 0 && !option[0].(bool))
		c.doAfterParsedOrValidated("xml", ps, nextFunc2, uuid, pwd, true, isprotect)
	}
}

// DoAfterUnmarshalXml do bussiness action after success unmarshaled the given xml data.
func (c *WAuthController) DoAfterUnmarshalXml(ps interface{}, nextFunc2 NextFunc2, option ...interface{}) {
	if uuid, pwd := c.AuthRequestHeader(); uuid != "" {
		isprotect := !(len(option) > 0 && !option[0].(bool))
		c.doAfterParsedOrValidated("xml", ps, nextFunc2, uuid, pwd, false, isprotect)
	}
}

// doAfterValidatedInner do bussiness action after success unmarshal params or
// validate the unmarshaled json data.
func (c *WAuthController) doAfterParsedOrValidated(datatype string,
	ps interface{}, nextFunc2 NextFunc2, uuid, pwd string, isvalidate, isprotect bool) {

	// unmarshal the input params
	switch datatype {
	case "json":
		if err := json.Unmarshal(c.Ctx.Input.RequestBody, ps); err != nil {
			c.E400Unmarshal(err.Error())
			return
		}
	case "xml":
		if err := xml.Unmarshal(c.Ctx.Input.RequestBody, ps); err != nil {
			c.E400Unmarshal(err.Error())
			return
		}
	default: // current not support the jsonp and yaml parse
		c.E404Exception("Invalid data type:" + datatype)
		return
	}

	// validate input params if need
	if isvalidate {
		ensureValidatorGenerated()
		if err := Validator.Struct(ps); err != nil {
			c.E400Validate(ps, err.Error())
			return
		}
	}

	// execute business function after unmarshal and validated
	if status, resp := nextFunc2(uuid, pwd); resp != nil {
		c.responCheckState(datatype, isprotect, status, resp)
	} else {
		c.responCheckState(datatype, isprotect, status)
	}
}
