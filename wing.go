// Copyright (c) 2018-2019 WING All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// -------------------------------------------------------------------

package wing

// use for all modules pass build
import (
	_ "github.com/wengoldx/wing/apis"
	_ "github.com/wengoldx/wing/apis/chain"
	_ "github.com/wengoldx/wing/comm"
	_ "github.com/wengoldx/wing/invar"
	_ "github.com/wengoldx/wing/logger"
	_ "github.com/wengoldx/wing/mvc"
	_ "github.com/wengoldx/wing/secure"
	_ "github.com/wengoldx/wing/wechat"
)
