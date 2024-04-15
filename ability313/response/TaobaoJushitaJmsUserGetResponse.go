package response

import (
	"github.com/xingwy/topsdk/ability313/domain"
)

type TaobaoJushitaJmsUserGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   同步的用户信息
	*/
	OnsUser domain.TaobaoJushitaJmsUserGetTmcUser `json:"ons_user,omitempty" `
}
