package response

import (
)

type TmallProductSchemaMatchResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回匹配产品ID，部分类目可能返回多个产品ID，以逗号分隔。
    */
    MatchResult  string `json:"match_result,omitempty" `
}
