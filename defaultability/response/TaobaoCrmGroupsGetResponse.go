package response

import (
	"github.com/xingwy/topsdk/defaultability/domain"
)

type TaobaoCrmGroupsGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   查询到的当前卖家的当前页的会员
	*/
	Groups []domain.TaobaoCrmGroupsGetGroup `json:"groups,omitempty" `
	/*
	   记录总数
	*/
	TotalResult int64 `json:"total_result,omitempty" `
}
