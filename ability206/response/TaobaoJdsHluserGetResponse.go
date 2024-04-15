package response

import (
	"github.com/xingwy/topsdk/ability206/domain"
)

type TaobaoJdsHluserGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   回流用户信息
	*/
	HlUser domain.TaobaoJdsHluserGetHlUserDO `json:"hl_user,omitempty" `
}
