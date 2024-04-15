package response

import (
    "topsdk/ability648/domain"
)

type TaobaoDeliveryTemplatesGetResponse struct {

    /*
        System request id
    */
    RequestId string `json:"request_id,omitempty" `

    /*
        System body
    */
    Body string

    /*
        运费模板列表
    */
    DeliveryTemplates  []domain.TaobaoDeliveryTemplatesGetDeliveryTemplate `json:"delivery_templates,omitempty" `
    /*
        获得到符合条件的结果总数
    */
    TotalResults  int64 `json:"total_results,omitempty" `
}
