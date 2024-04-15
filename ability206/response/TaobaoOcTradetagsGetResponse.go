package response

import (
	"github.com/xingwy/topsdk/ability206/domain"
)

type TaobaoOcTradetagsGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   返回结果
	*/
	TradeTags []domain.TaobaoOcTradetagsGetTradeTagRelationDo `json:"trade_tags,omitempty" `
}
