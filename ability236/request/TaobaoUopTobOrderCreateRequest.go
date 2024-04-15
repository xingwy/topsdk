package request

import (
        "topsdk/ability236/domain"
        "topsdk/util"
    )

type TaobaoUopTobOrderCreateRequest struct {
    /*
        ERP出库对象     */
    DeliveryOrder  *domain.TaobaoUopTobOrderCreateDeliveryorder `json:"delivery_order,omitempty" required:"false" `
}

func (s *TaobaoUopTobOrderCreateRequest) SetDeliveryOrder(v domain.TaobaoUopTobOrderCreateDeliveryorder) *TaobaoUopTobOrderCreateRequest {
    s.DeliveryOrder = &v
    return s
}

func (req *TaobaoUopTobOrderCreateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.DeliveryOrder != nil) {
        paramMap["delivery_order"] = util.ConvertStruct(*req.DeliveryOrder)
    }
    return paramMap
}

func (req *TaobaoUopTobOrderCreateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}