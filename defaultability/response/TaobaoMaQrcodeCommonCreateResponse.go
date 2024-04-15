package response

import (
	"github.com/xingwy/topsdk/defaultability/domain"
)

type TaobaoMaQrcodeCommonCreateResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   二维码对像
	*/
	Modules []domain.TaobaoMaQrcodeCommonCreateQrcodeDO `json:"modules,omitempty" `
	/*
	   执行是否成功
	*/
	Suc bool `json:"suc,omitempty" `
}
