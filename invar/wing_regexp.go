// Copyright (c) 2018-2019 WING All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// 00002       2019/06/30   zhaixing       Add function from godfs
// -------------------------------------------------------------------
package invar

import (
	"github.com/wengoldx/wing/logger"
	"regexp"
)

const (
	ExpOnlyNumbers     = "^[0-9]*$"
	ExpOnlyLowerChars  = "^[a-z]*$"
	ExpOnlyUpperChars  = "^[A-Z]*$"
	ExpOnlyCaseChars   = "^[a-zA-Z]*$"
	ExpNumOrLowerChars = "^[0-9a-z]*$"
	ExpNumOrUpperChars = "^[0-9A-Z]*$"
	ExpNumOrCaseChars  = "^[0-9a-zA-Z]*$"
)

// MatchRegexp validate the src if matched by expression
func MatchRegexp(expression, src string) bool {
	reg, err := regexp.Compile(expression)
	if err != nil {
		logger.E("Invalid regexp expression, err:", err)
		return false
	}
	return reg.MatchString(src)
}