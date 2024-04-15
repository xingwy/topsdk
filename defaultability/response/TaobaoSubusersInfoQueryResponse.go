package response

import (
	"github.com/xingwy/topsdk/defaultability/domain"
)

type TaobaoSubusersInfoQueryResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   子账号对象，如果为空，说明没查到该子账号
	*/
	Result domain.TaobaoSubusersInfoQuerySubUserDO `json:"result,omitempty" `
}
