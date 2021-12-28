package service

import (
	"crypto/rand"
	"errors"
	"j2ee/constant"
	"j2ee/model"
	"j2ee/serializer"
	"j2ee/util/logging"
	"time"

	"gorm.io/gorm"
)

// OrderCreateService 订单创建服务
type OrderCreateService struct {
	UserID uint `json:"user_id"`
	// 发件人与收件人
	SName    string `json:"s_name"`
	SPhone   string `json:"s_phone"`
	SAddress string `json:"s_address"`
	RName    string `json:"r_name"`
	RPhone   string `json:"r_phone"`
	RAddress string `json:"r_address"`

	// 货物信息
	Type   string `json:"type"`
	Weight int    `json:"weight"`
	Volume int    `json:"volume"`
	Value  int    `json:"value"`
	Urgent bool   `json:"urgent"`
	Note   string `json:"note"`
}

// OrderCreate 创建一条订单
func (service *OrderCreateService) OrderCreate() serializer.Response {
	order := model.Order{
		UserID: service.UserID,
		// 发件人与收件人
		SName:    service.SName,
		SPhone:   service.SPhone,
		SAddress: service.SAddress,
		RName:    service.RName,
		RPhone:   service.RPhone,
		RAddress: service.RAddress,
		// 货物信息
		Type:   service.Type,
		Weight: service.Weight,
		Volume: service.Volume,
		Value:  service.Value,
		Urgent: service.Urgent,
		Note:   service.Note,
		// 订单状态
		OrderID:  RandomID(14),
		Time:     TimeEvaluateCreated(service.SAddress, service.RAddress),
		Status:   1,
		Allocate: 0,
		Rating:   0,
	}
	code := constant.SUCCESS

	err := model.DB.Model(&model.Order{}).Create(&order).Error
	if err != nil {
		logging.Info(err)
		code = constant.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 更新用户订单数
	var userInfo model.UserInfo
	db := model.DB.Model(&model.UserInfo{}).Where("user_id = ?", service.UserID).First(&userInfo)
	err = db.Error
	if err != nil {
		logging.Info(err)
		code = constant.ERROR_UPDATING_INFO
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
			Error:  err.Error(),
		}
	}
	// 更新并保存
	userInfo.OrderTotal++
	userInfo.SendTotal++
	userInfo.SendNumber++
	err = db.Save(&userInfo).Error
	if err != nil {
		logging.Info(err)
		code = constant.ERROR_UPDATING_INFO
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
			Error:  err.Error(),
		}
	}

	// 更新收件人信息
	db = model.DB.Model(&model.UserInfo{}).Where("username = ?", service.RName).First(&userInfo)
	err = db.Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {

		} else {
			logging.Info(err)
			code = constant.ERROR_UPDATING_INFO
			return serializer.Response{
				Status: code,
				Msg:    constant.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else {
		// 更新并保存
		userInfo.MsgNumber++
		userInfo.ReceiveTotal++
		userInfo.ReceiveNumber++
		err = db.Save(&userInfo).Error
		if err != nil {
			logging.Info(err)
			code = constant.ERROR_UPDATING_INFO
			return serializer.Response{
				Status: code,
				Msg:    constant.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}

	// 管理员信息更新
	adminInfo := AdminInfoUpdateService{
		DMsgNumber:      1,
		DOrderTotal:     1,
		DUserTotal:      0,
		DTruckTotal:     0,
		DTruckAvailable: 0,
		DOrderUnhandled: 1,
	}
	err, code = adminInfo.AdminInfoUpdate()
	if err != nil {
		logging.Info(err)
		return serializer.Response{
			Status: code,
			Msg:    constant.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    constant.GetMsg(code),
		Data:   order.OrderID, // 用于返回订单详情
		// 如果Data要返回结构体，则需先定义结构体类型
	}
}

// TimeEvaluate 预估在两地之间发送，到达时的时间
func TimeEvaluate(address1, address2 string) time.Time {
	h := 0
	m := 0
	s := 10
	return time.Now().Add(time.Hour*time.Duration(h) + time.Minute*time.Duration(m) + time.Second*time.Duration(s))
}

// TimeEvaluateCreated 预估订单刚创建时，在两地之间发送，到达时的时间
func TimeEvaluateCreated(address1, address2 string) time.Time {
	h := 1
	m := 5
	return time.Now().Add(time.Hour*time.Duration(h) + time.Minute*time.Duration(m))
}

var defaultLetters = []byte("0123456789")

// RandomBytes 生成随机byte序列
func RandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		logging.Info(err)
		return nil, err
	}
	return b, nil
}

// RandomID 生成随机订单号（14位数字）
func RandomID(n int) string {
	b, _ := RandomBytes(n)
	for i := range b {
		b[i] = defaultLetters[b[i]%10]
	}
	return string(b)
}
