package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoCrmMembersIncrementGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回当前页的会员列表
    */
    Members  []domain.TaobaoCrmMembersIncrementGetBasicMember `json:"members,omitempty" `
    /*
        记录的总条数
    */
    TotalResult  int64 `json:"total_result,omitempty" `
}
