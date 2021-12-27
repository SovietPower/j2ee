package service

import (
	"j2ee/constant"

	"github.com/dchest/captcha"
)

// VerifyCaptcha 验证验证码输入
func VerifyCaptcha(value, captchaID string) int {
	if captcha.VerifyString(captchaID, value) {
		return constant.SUCCESS
	} else {
		return constant.ERROR_WRONG_CAPTCHA
	}
}
