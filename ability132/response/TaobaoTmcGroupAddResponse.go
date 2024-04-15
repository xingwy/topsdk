package response

import (
    "topsdk/util"
)

type TaobaoTmcGroupAddResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        创建时间
    */
    Created  util.LocalTime `json:"created,omitempty" `
    /*
        分组名称
    */
    GroupName  string `json:"group_name,omitempty" `
}
