package response

import (
	"github.com/xingwy/topsdk/ability648/domain"
)

type TaobaoTradeMemoUpdateResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   更新交易的备注信息后返回的Trade，其中可用字段为tid和modified
	*/
	Trade domain.TaobaoTradeMemoUpdateTrade `json:"trade,omitempty" `
}
