package response

import (
)

type TaobaoTraderateImprImprwordsGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        返回类目下所有大家印象的标签
    */
    ImprWords  []string `json:"impr_words,omitempty" `
}
