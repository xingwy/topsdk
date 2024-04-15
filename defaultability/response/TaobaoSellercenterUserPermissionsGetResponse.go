package response

import (
	"github.com/xingwy/topsdk/defaultability/domain"
)

type TaobaoSellercenterUserPermissionsGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   权限列表
	*/
	Permissions []domain.TaobaoSellercenterUserPermissionsGetPermission `json:"permissions,omitempty" `
}
