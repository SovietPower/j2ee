package constant

// MsgFlags 状态码map
var MsgFlags = map[int]string{
	SUCCESS:                    "ok",
	UPDATE_PASSWORD_SUCCESS:    "修改密码成功",
	ERROR:                      "fail",
	INVALID_PARAMS:             "请求参数错误",
	ERROR_EXISTED_USER:         "已存在该用户",
	ERROR_NO_SUCH_USER:         "该用户不存在",
	ERROR_WRONG_PASSWORD:       "帐号或密码错误",
	ERROR_WRONG_CAPTCHA:        "验证码错误",
	ERROR_FAIL_GETTING_CAPTCHA: "获取验证码失败，请重试",
	ERROR_FAIL_ENCRYPTION:      "密码加密失败，请重试",
	ERROR_NO_SUCH_ADDRESS:      "该收货地址不存在",
	ERROR_NO_SUCH_ORDER:        "该订单不存在",

	ERROR_AUTH_CHECK_TOKEN_FAIL:       "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:    "Token已超时",
	ERROR_AUTH_TOKEN:                  "Token生成失败",
	ERROR_AUTH:                        "Token错误",
	ERROR_AUTH_INSUFFICIENT_AUTHORITY: "权限不足",
	ERROR_READ_FILE:                   "读文件失败",
	ERROR_SEND_EMAIL:                  "发送邮件失败",
	ERROR_CALL_API:                    "调用接口失败",
	ERROR_UNMARSHAL_JSON:              "解码JSON失败",

	ERROR_DATABASE:      "数据库操作出错，请重试",
	ERROR_NO_PERMISSION: "无访问权限",
	ERROR_UPDATING_INFO: "更新用户信息失败",

	ERROR_OSS: "OSS配置错误",
}

// GetMsg 获取状态码对应信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
