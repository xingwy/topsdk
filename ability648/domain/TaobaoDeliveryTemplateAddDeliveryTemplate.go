package domain

import (
        "topsdk/util"
    )

type TaobaoDeliveryTemplateAddDeliveryTemplate struct {
    /*
        模块ID     */
    TemplateId  *int64 `json:"template_id,omitempty" `

    /*
        模板创建时间     */
    Created  *util.LocalTime `json:"created,omitempty" `

}

func (s *TaobaoDeliveryTemplateAddDeliveryTemplate) SetTemplateId(v int64) *TaobaoDeliveryTemplateAddDeliveryTemplate {
    s.TemplateId = &v
    return s
}
func (s *TaobaoDeliveryTemplateAddDeliveryTemplate) SetCreated(v util.LocalTime) *TaobaoDeliveryTemplateAddDeliveryTemplate {
    s.Created = &v
    return s
}
