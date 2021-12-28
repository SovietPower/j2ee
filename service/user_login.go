package service

import (
	"errors"
	"j2ee/constant"
	"j2ee/model"
	"j2ee/serializer"
	"j2ee/util"
	"j2ee/util/logging"

	"gorm.io/gorm"
)

// UserLoginService 管理用户登录服务的结构
type UserLoginService struct {
	Username  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required"`
	CaptchaID string `form:"captchaID" json:"captchaID" binding:"required"`
}

// UserLogin 用户登录函数
func (service *UserLoginService) UserLogin() serializer.Response {
	var user model.User
	code := constant.SUCCESS

	// 先进行验证码验证
	if code = VerifyCaptcha(service.Captcha, service.CaptchaID); code != constant.SUCCESS {
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
		}
	}

	if err := model.DB.Where("username = ?", service.Username).First(&user).Error; err != nil {
		// 无用户记录
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logging.Info(err)
			code = constant.ERROR_NO_SUCH_USER
			return serializer.Response{
				Status: code,
				Msg:    constant.GetMsg(code),
			}
		}

		// 数据库错误
		logging.Info(err)
		code = constant.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
		}
	}

	if user.CheckPassword(service.Password) == false {
		code = constant.ERROR_WRONG_PASSWORD
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
		}
	}

	token, err := util.GenerateToken(service.Username, service.Password, user.Status, user.Authority)
	if err != nil {
		logging.Info(err)
		code = constant.ERROR_AUTH_TOKEN
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
		}
	}
	return serializer.Response{
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Status: code,
		Msg:    constant.GetMsg(code),
	}
}
