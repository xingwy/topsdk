package response

import (
	"github.com/xingwy/topsdk/ability314/domain"
)

type TaobaoJushitaJdpUsersGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   总记录数
	*/
	TotalResults int64 `json:"total_results,omitempty" `
	/*
	   用户列表
	*/
	Users []domain.TaobaoJushitaJdpUsersGetJdpUser `json:"users,omitempty" `
}
