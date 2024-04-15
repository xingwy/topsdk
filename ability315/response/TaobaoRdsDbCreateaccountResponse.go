package response

import (
	"github.com/xingwy/topsdk/ability315/domain"
)

type TaobaoRdsDbCreateaccountResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   result
	*/
	Result domain.TaobaoRdsDbCreateaccountResultSet `json:"result,omitempty" `
}
