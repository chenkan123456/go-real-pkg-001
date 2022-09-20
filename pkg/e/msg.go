package e

var MsgFlags = map[int]string{
	SUCCESS:             "成功",
	ERROR:               "错误",
	INVALID_PARAMS:      "入参数信息有误",
	ERROR_TOKEN_FAIL:    "验证失败",
	ERROR_TOKEN_TIMEOUT: "密钥过期",

	ERROR_EXIST:        "信息已存在",
	ERROR_NOTEXIST:     "信息不存在",
	ERROR_EXIST_FAIL:   "获取有误",
	ERROR_EXIST_CODE:   "编号已存在",
	ERROR_EXIST_PATH:   "地址已存在",
	ERROR_EXIST_USED:   "信息已被使用",
	ERROR_EXIST_PARENT: "该信息为父类信息",

	ERROR_EXIST_USER_FAIL: "用户验证错误，请联系管理员",

	ERROR_DELETE_PARENT_FAIL: "该信息正在被使用，无法直接被删除",
	ERROR_DELETE_FAIL:        "删除有误",
	ERROR_EDIT_FAIL:          "编辑有误",
	ERROR_ADD_FAIL:           "新增有误",

	ERROR_RESULT_DATA_FAIL:  "获取数据有误",
	ERROR_RESULT_COUNT_FAIL: "获取数量有误",

	ERROR_POINTS_FAIL: "转换坐标信息失败",

	ERROR_AD_AUTH:      "账号信息验证有误",
	ERROR_AD_AUTH_USER: "工号(邮箱前缀)不存在",
	ERROR_AD_AUTH_PWD:  "账号密码错误",

	ERROR_AD_AUTH_TEST: "测试",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
