package response

import (
)

type AlibabaGpuSchemaCatsearchResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回按类目查询spu的schema
    */
    CatSearchResult  string `json:"cat_search_result,omitempty" `
    /*
        总记录数
    */
    Total  int64 `json:"total,omitempty" `
}
