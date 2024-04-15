package response

import (
	"github.com/xingwy/topsdk/ability181/domain"
)

type TaobaoItemCatpropsModificationGetResponse struct {

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
	Results []domain.TaobaoItemCatpropsModificationGetPropsModificationResult `json:"results,omitempty" `
}
