package response

import (
    "topsdk/defaultability/domain"
)

type TaobaoFuwuScoresGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        评价流水记录
    */
    ScoreResult  []domain.TaobaoFuwuScoresGetScoreResult `json:"score_result,omitempty" `
}
