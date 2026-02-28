package gaodeKit

import "github.com/richelieu042/chimera/v3/src/core/error/errorKit"

type (
	BaseResponse struct {
		// Status 值为0或1，1：成功；0：失败
		Status   string `json:"status"`
		InfoCode string `json:"infocode"`
		Info     string `json:"info"`
	}
)

func (baseResp *BaseResponse) IsSuccess() error {
	if baseResp.Status != "1" {
		return errorKit.Newf("Fail to get weather, status: %s, infocode: %s, info: %s", baseResp.Status, baseResp.InfoCode, baseResp.Info)
	}
	return nil
}
