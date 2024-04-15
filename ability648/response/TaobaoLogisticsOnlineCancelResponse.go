package response

import (
)

type TaobaoLogisticsOnlineCancelResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        成功与失败
    */
    IsSuccess  bool `json:"is_success,omitempty" `
    /*
        修改时间
    */
    ModifyTime  string `json:"modify_time,omitempty" `
    /*
        重新创建物流订单id
    */
    RecreatedOrderId  string `json:"recreated_order_id,omitempty" `
}
