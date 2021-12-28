package service

import (
	"j2ee/constant"
	"j2ee/model"
	"j2ee/serializer"
	"j2ee/util/logging"
)

// UserRegisterService 管理用户注册服务的结构
type UserRegisterService struct {
	Username  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	Captcha   string `form:"captcha" json:"captcha" binding:"required"`
	CaptchaID string `form:"captchaID" json:"captchaID" binding:"required"`
}

// Valid 验证信息
func (service *UserRegisterService) Validate() *serializer.Response {
	var count int64
	count = 0
	code := constant.SUCCESS

	err := model.DB.Model(&model.User{}).Where("username = ?", service.Username).Count(&count).Error
	if err != nil {
		code = constant.ERROR_DATABASE
		return &serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
		}
	}
	if count > 0 {
		code = constant.ERROR_EXISTED_USER
		return &serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
		}
	}
	return nil
}

// Register 用户注册
func (service *UserRegisterService) UserRegister() serializer.Response {
	user := model.User{
		Username:  service.Username,
		Status:    model.StatusNormal,
		Authority: model.AuthorityUser,
	}
	code := constant.SUCCESS

	// 先进行验证码验证
	if code = VerifyCaptcha(service.Captcha, service.CaptchaID); code != constant.SUCCESS {
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
		}
	}

	// 信息验证
	if res := service.Validate(); res != nil {
		return *res
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		logging.Info(err)
		code = constant.ERROR_FAIL_ENCRYPTION
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
		}
	}

	// 创建用户
	if err := model.DB.Model(&model.User{}).Create(&user).Error; err != nil {
		logging.Info(err)
		code = constant.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
		}
	}

	// 创建用户其它信息
	info := UserInfoCreateService{
		UserID:   user.ID,
		Username: service.Username,
	}
	if err := info.UserInfoCreate(); err != nil {
		logging.Info(err)
		code = constant.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
		}
	}

	// 管理员信息更新
	adminInfo := AdminInfoUpdateService{
		DMsgNumber:      0,
		DOrderTotal:     0,
		DUserTotal:      1,
		DTruckTotal:     0,
		DTruckAvailable: 0,
		DOrderUnhandled: 0,
	}
	var err error
	err, code = adminInfo.AdminInfoUpdate()
	if err != nil {
		logging.Info(err)
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 创建用于展示给管理员的用户信息
	// adminInfo := AdminUserInfoCreateService{
	// 	UserID:   user.ID,
	// 	Username: service.Username,
	// }
	// if err := adminInfo.AdminUserInfoCreate(); err != nil {
	// 	logging.Info(err)
	// 	code = constant.ERROR_DATABASE
	// 	return serializer.Response{
	// 		Status: code,
	// 		Msg:    constant.GetMsg(code),
	// 	}
	// }

	return serializer.Response{
		Status: code,
		Msg:    constant.GetMsg(code),
	}
}
