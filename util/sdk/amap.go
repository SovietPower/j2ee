package sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"j2ee/constant"
	"j2ee/serializer"
	"j2ee/util/logging"
	"net/http"
)

var KEY string

type District struct {
	// Level string `json:"level"`
	Name string `json:"name"`
}

type DistrictsResponse struct {
	District
	Districts []District `json:"districts"`
}

type AMapResponse struct {
	Status    string              `json:"status"`
	Districts []DistrictsResponse `json:"districts"`
}

type GetDistrictsService struct {
	Keyword string `form:"keyword"`
}

func (service GetDistrictsService) GetDistricts() serializer.Response {
	code := constant.SUCCESS

	url := "http://restapi.amap.com/v3/config/district?key=" + KEY + "&keywords=" + service.Keyword + "&subdistrict=1&extensions=base"
	println(url)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return handleError(err)
	}

	var res AMapResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		return handleError(err)
	}

	districts := make([]string, len(res.Districts[0].Districts))
	// districts[0] = "请选择"
	for i, d := range res.Districts[0].Districts {
		districts[i] = d.Name
	}

	return serializer.Response{
		Status: code,
		Msg:    constant.GetMsg(code),
		Data:   districts,
	}
}

func handleError(err error) serializer.Response {
	logging.Info(err)
	code := constant.ERROR_AMAP_QUERY
	return serializer.Response{
		Status: code,
		Msg:    constant.GetMsg(code),
		Error:  err.Error(),
	}
}
