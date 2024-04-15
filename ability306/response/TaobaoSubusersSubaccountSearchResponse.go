package response

import (
    "topsdk/ability306/domain"
)

type TaobaoSubusersSubaccountSearchResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        子账号列表
    */
    Subaccounts  []domain.TaobaoSubusersSubaccountSearchSubAccountInfo `json:"subaccounts,omitempty" `
    /*
        isv本次调用传入的页大小
    */
    PageSize  int64 `json:"page_size,omitempty" `
    /*
        本次调用删选条件下的总子账号数量，除以页大小可得出最大页数
    */
    TotalCount  int64 `json:"total_count,omitempty" `
    /*
        isv本次调用传入的页码
    */
    PageNum  int64 `json:"page_num,omitempty" `
}
