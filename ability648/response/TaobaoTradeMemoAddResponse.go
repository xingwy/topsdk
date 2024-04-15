package response

import (
	"github.com/xingwy/topsdk/ability648/domain"
)

type TaobaoTradeMemoAddResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   对一笔交易添加备注后返回其对应的Trade，Trade中可用的返回字段有tid和created
	*/
	Trade domain.TaobaoTradeMemoAddTrade `json:"trade,omitempty" `
}
