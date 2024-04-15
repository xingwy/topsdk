package response

import (
    "topsdk/ability312/domain"
)

type AlibabaAliqinTaSmsNumQueryResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        当前页码
    */
    CurrentPage  int64 `json:"current_page,omitempty" `
    /*
        每页数量
    */
    PageSize  int64 `json:"page_size,omitempty" `
    /*
        总量
    */
    TotalCount  int64 `json:"total_count,omitempty" `
    /*
        总页数
    */
    TotalPage  int64 `json:"total_page,omitempty" `
    /*
        1
    */
    Values  []domain.AlibabaAliqinTaSmsNumQueryFcPartnerSmsDetailDto `json:"values,omitempty" `
}
