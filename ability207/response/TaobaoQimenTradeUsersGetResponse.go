package response

import (
	"github.com/xingwy/topsdk/ability207/domain"
)

type TaobaoQimenTradeUsersGetResponse struct {

	/*
	   System request id
	*/
	RequestId string `json:"request_id,omitempty" `

	/*
	   System body
	*/
	Body string

	/*
	   totalCount
	*/
	TotalCount int64 `json:"total_count,omitempty" `
	/*
	   modal
	*/
	Users []domain.TaobaoQimenTradeUsersGetQimenUser `json:"users,omitempty" `
}
