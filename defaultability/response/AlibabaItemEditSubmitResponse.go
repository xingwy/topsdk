package response

import (
)

type AlibabaItemEditSubmitResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        商品更新时间
    */
    UpdateTime  string `json:"update_time,omitempty" `
    /*
        商品ID
    */
    ItemId  int64 `json:"item_id,omitempty" `
    /*
        商品所属市场
    */
    Market  string `json:"market,omitempty" `
}
