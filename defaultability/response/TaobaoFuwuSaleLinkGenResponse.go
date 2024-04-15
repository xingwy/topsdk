package response

import (
)

type TaobaoFuwuSaleLinkGenResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        通过营销链接接口生成的营销链接短地址
    */
    Url  string `json:"url,omitempty" `
}
