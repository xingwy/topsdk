package response

import (
    "topsdk/ability680/domain"
)

type TaobaoCrmMembersSearchPrivyResponse struct {

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
    Members  []domain.TaobaoCrmMembersSearchPrivyCrmMember `json:"members,omitempty" `
    /*
        记录的总条数
    */
    TotalResult  int64 `json:"total_result,omitempty" `
}
