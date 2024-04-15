package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoCrmMembersSearchResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        根据一定条件查询的卖家会员
    */
    Members  []domain.TaobaoCrmMembersSearchCrmMember `json:"members,omitempty" `
    /*
        记录的总条数
    */
    TotalResult  int64 `json:"total_result,omitempty" `
}
