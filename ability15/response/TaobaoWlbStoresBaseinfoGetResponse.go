package response

import (
    "topsdk/ability15/domain"
)

type TaobaoWlbStoresBaseinfoGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        仓库列表
    */
    StoreInfoList  []domain.TaobaoWlbStoresBaseinfoGetStoreInfo `json:"store_info_list,omitempty" `
    /*
        返回的总数
    */
    TotalCount  int64 `json:"total_count,omitempty" `
}
