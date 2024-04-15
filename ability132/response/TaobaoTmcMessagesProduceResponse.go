package response

import (
    "topsdk/ability132/domain"
)

type TaobaoTmcMessagesProduceResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        是否全部成功
    */
    IsAllSuccess  bool `json:"is_all_success,omitempty" `
    /*
        发送结果，与发送时的参数顺序一致。如果is_all_success为true时，不用校验result是否成功
    */
    Results  []domain.TaobaoTmcMessagesProduceTmcProduceResult `json:"results,omitempty" `
}
