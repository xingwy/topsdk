package response

import (
)

type TaobaoCrmGrouptaskCheckResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        异步任务是否完成，true表示完成
    */
    IsFinished  bool `json:"is_finished,omitempty" `
}
