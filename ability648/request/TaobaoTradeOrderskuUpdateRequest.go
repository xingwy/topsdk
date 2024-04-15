package request


type TaobaoTradeOrderskuUpdateRequest struct {
    /*
        子订单编号（对于单笔订单的交易可以传交易编号）。     */
    Oid  *int64 `json:"oid" required:"true" `
    /*
        销售属性编号，可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。     */
    SkuId  *int64 `json:"sku_id,omitempty" required:"false" `
    /*
        销售属性组合串，格式：p1:v1;p2:v2，如：1627207:28329;20509:28314。可以通过taobao.item.skus.get获取订单对应的商品的所有销售属性。     */
    SkuProps  *string `json:"sku_props,omitempty" required:"false" `
}

func (s *TaobaoTradeOrderskuUpdateRequest) SetOid(v int64) *TaobaoTradeOrderskuUpdateRequest {
    s.Oid = &v
    return s
}
func (s *TaobaoTradeOrderskuUpdateRequest) SetSkuId(v int64) *TaobaoTradeOrderskuUpdateRequest {
    s.SkuId = &v
    return s
}
func (s *TaobaoTradeOrderskuUpdateRequest) SetSkuProps(v string) *TaobaoTradeOrderskuUpdateRequest {
    s.SkuProps = &v
    return s
}

func (req *TaobaoTradeOrderskuUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Oid != nil) {
        paramMap["oid"] = *req.Oid
    }
    if(req.SkuId != nil) {
        paramMap["sku_id"] = *req.SkuId
    }
    if(req.SkuProps != nil) {
        paramMap["sku_props"] = *req.SkuProps
    }
    return paramMap
}

func (req *TaobaoTradeOrderskuUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}