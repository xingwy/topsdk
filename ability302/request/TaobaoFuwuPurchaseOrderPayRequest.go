package request


type TaobaoFuwuPurchaseOrderPayRequest struct {
    /*
        APPKEY，必填     */
    Appkey  *string `json:"appkey" required:"true" `
    /*
        设备类型，目前只支持PC，可选     */
    DeviceType  *string `json:"device_type,omitempty" required:"false" `
    /*
        订单号，与外部订单号二选一     */
    OrderId  *int64 `json:"order_id,omitempty" required:"false" `
    /*
        外部订单号，使用该参数完成查询订单等操作，与外部订单号二选一     */
    OutOrderId  *string `json:"out_order_id,omitempty" required:"false" `
}

func (s *TaobaoFuwuPurchaseOrderPayRequest) SetAppkey(v string) *TaobaoFuwuPurchaseOrderPayRequest {
    s.Appkey = &v
    return s
}
func (s *TaobaoFuwuPurchaseOrderPayRequest) SetDeviceType(v string) *TaobaoFuwuPurchaseOrderPayRequest {
    s.DeviceType = &v
    return s
}
func (s *TaobaoFuwuPurchaseOrderPayRequest) SetOrderId(v int64) *TaobaoFuwuPurchaseOrderPayRequest {
    s.OrderId = &v
    return s
}
func (s *TaobaoFuwuPurchaseOrderPayRequest) SetOutOrderId(v string) *TaobaoFuwuPurchaseOrderPayRequest {
    s.OutOrderId = &v
    return s
}

func (req *TaobaoFuwuPurchaseOrderPayRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Appkey != nil) {
        paramMap["appkey"] = *req.Appkey
    }
    if(req.DeviceType != nil) {
        paramMap["device_type"] = *req.DeviceType
    }
    if(req.OrderId != nil) {
        paramMap["order_id"] = *req.OrderId
    }
    if(req.OutOrderId != nil) {
        paramMap["out_order_id"] = *req.OutOrderId
    }
    return paramMap
}

func (req *TaobaoFuwuPurchaseOrderPayRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}