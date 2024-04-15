package response

import (
	"github.com/xingwy/topsdk/ability680/domain"
)

type TaobaoCrmMembersGetPrivyResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   根据一定条件查询到卖家的会员
	*/
	Members []domain.TaobaoCrmMembersGetPrivyBasicMember `json:"members,omitempty" `
	/*
	   记录总数
	*/
	TotalResult int64 `json:"total_result,omitempty" `
}
