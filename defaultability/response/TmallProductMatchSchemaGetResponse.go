package response

import (
)

type TmallProductMatchSchemaGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回匹配product的规则文档
    */
    MatchResult  string `json:"match_result,omitempty" `
}
