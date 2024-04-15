package request


type TaobaoLogisticsOrderCreateRequest struct {
    /*
        发货的物流公司运单号。在logis_type=OFFLINE且is_consign=true时，此项必填。     */
    MailNo  *string `json:"mail_no,omitempty" required:"false" `
    /*
        卖家旺旺号     */
    SellerWangwangId  *string `json:"seller_wangwang_id" required:"true" `
    /*
        物流发货类型：1-普通订单
不填为默认类型 1-普通订单 defalutValue��1    */
    OrderType  *int64 `json:"order_type,omitempty" required:"false" `
    /*
        发货的物流公司代码，如申通=STO，圆通=YTO。is_consign=true时，此项必填。     */
    LogisCompanyCode  *string `json:"logis_company_code,omitempty" required:"false" `
    /*
        订单的交易号码     */
    TradeId  *int64 `json:"trade_id,omitempty" required:"false" `
    /*
        创建订单同时是否进行发货，默认发货。 defalutValue��true    */
    IsConsign  *bool `json:"is_consign,omitempty" required:"false" `
    /*
        发货方式,默认为自己联系发货。可选值:ONLINE(在线下单)、OFFLINE(自己联系)。 defalutValue��OFFLINE    */
    LogisType  *string `json:"logis_type,omitempty" required:"false" `
    /*
        运费承担方式。1为买家承担运费，2为卖家承担运费，其他值为错误参数。 defalutValue��1    */
    Shipping  *int64 `json:"shipping,omitempty" required:"false" `
    /*
        运送货物的单价列表(注意：单位为分），用|号隔开     */
    ItemValues  *string `json:"item_values" required:"true" `
    /*
        运送的货物名称列表，用|号隔开     */
    GoodsNames  *string `json:"goods_names" required:"true" `
    /*
        运送货物的数量列表，用|号隔开     */
    GoodsQuantities  *string `json:"goods_quantities" required:"true" `
}

func (s *TaobaoLogisticsOrderCreateRequest) SetMailNo(v string) *TaobaoLogisticsOrderCreateRequest {
    s.MailNo = &v
    return s
}
func (s *TaobaoLogisticsOrderCreateRequest) SetSellerWangwangId(v string) *TaobaoLogisticsOrderCreateRequest {
    s.SellerWangwangId = &v
    return s
}
func (s *TaobaoLogisticsOrderCreateRequest) SetOrderType(v int64) *TaobaoLogisticsOrderCreateRequest {
    s.OrderType = &v
    return s
}
func (s *TaobaoLogisticsOrderCreateRequest) SetLogisCompanyCode(v string) *TaobaoLogisticsOrderCreateRequest {
    s.LogisCompanyCode = &v
    return s
}
func (s *TaobaoLogisticsOrderCreateRequest) SetTradeId(v int64) *TaobaoLogisticsOrderCreateRequest {
    s.TradeId = &v
    return s
}
func (s *TaobaoLogisticsOrderCreateRequest) SetIsConsign(v bool) *TaobaoLogisticsOrderCreateRequest {
    s.IsConsign = &v
    return s
}
func (s *TaobaoLogisticsOrderCreateRequest) SetLogisType(v string) *TaobaoLogisticsOrderCreateRequest {
    s.LogisType = &v
    return s
}
func (s *TaobaoLogisticsOrderCreateRequest) SetShipping(v int64) *TaobaoLogisticsOrderCreateRequest {
    s.Shipping = &v
    return s
}
func (s *TaobaoLogisticsOrderCreateRequest) SetItemValues(v string) *TaobaoLogisticsOrderCreateRequest {
    s.ItemValues = &v
    return s
}
func (s *TaobaoLogisticsOrderCreateRequest) SetGoodsNames(v string) *TaobaoLogisticsOrderCreateRequest {
    s.GoodsNames = &v
    return s
}
func (s *TaobaoLogisticsOrderCreateRequest) SetGoodsQuantities(v string) *TaobaoLogisticsOrderCreateRequest {
    s.GoodsQuantities = &v
    return s
}

func (req *TaobaoLogisticsOrderCreateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.MailNo != nil) {
        paramMap["mail_no"] = *req.MailNo
    }
    if(req.SellerWangwangId != nil) {
        paramMap["seller_wangwang_id"] = *req.SellerWangwangId
    }
    if(req.OrderType != nil) {
        paramMap["order_type"] = *req.OrderType
    }
    if(req.LogisCompanyCode != nil) {
        paramMap["logis_company_code"] = *req.LogisCompanyCode
    }
    if(req.TradeId != nil) {
        paramMap["trade_id"] = *req.TradeId
    }
    if(req.IsConsign != nil) {
        paramMap["is_consign"] = *req.IsConsign
    }
    if(req.LogisType != nil) {
        paramMap["logis_type"] = *req.LogisType
    }
    if(req.Shipping != nil) {
        paramMap["shipping"] = *req.Shipping
    }
    if(req.ItemValues != nil) {
        paramMap["item_values"] = *req.ItemValues
    }
    if(req.GoodsNames != nil) {
        paramMap["goods_names"] = *req.GoodsNames
    }
    if(req.GoodsQuantities != nil) {
        paramMap["goods_quantities"] = *req.GoodsQuantities
    }
    return paramMap
}

func (req *TaobaoLogisticsOrderCreateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}