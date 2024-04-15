package request


type TaobaoLogisticsConsignOrderCreateandsendRequest struct {
    /*
        物流公司ID     */
    CompanyId  *int64 `json:"company_id" required:"true" `
    /*
        交易流水号，淘外订单号或者商家内部交易流水号     */
    TradeId  *int64 `json:"trade_id" required:"true" `
    /*
        运单号     */
    MailNo  *string `json:"mail_no,omitempty" required:"false" `
    /*
        订单来源，值选择：30     */
    OrderSource  *int64 `json:"order_source" required:"true" `
    /*
        用户ID     */
    UserId  *int64 `json:"user_id" required:"true" `
    /*
        物流订单物流类型，值固定选择：2     */
    LogisType  *int64 `json:"logis_type" required:"true" `
    /*
        订单类型，值固定选择：30     */
    OrderType  *int64 `json:"order_type" required:"true" `
    /*
        费用承担方式 1买家承担运费 2卖家承担运费     */
    Shipping  *string `json:"shipping,omitempty" required:"false" `
    /*
        物品的json数据。     */
    ItemJsonString  *string `json:"item_json_string" required:"true" `
    /*
        发件人街道地址     */
    SAddress  *string `json:"s_address" required:"true" `
    /*
        市     */
    SCityName  *string `json:"s_city_name" required:"true" `
    /*
        发件人名称     */
    SName  *string `json:"s_name" required:"true" `
    /*
        省     */
    SProvName  *string `json:"s_prov_name" required:"true" `
    /*
        区     */
    SDistName  *string `json:"s_dist_name,omitempty" required:"false" `
    /*
        发件人区域ID     */
    SAreaId  *int64 `json:"s_area_id" required:"true" `
    /*
        电话号码     */
    STelephone  *string `json:"s_telephone,omitempty" required:"false" `
    /*
        手机号码     */
    SMobilePhone  *string `json:"s_mobile_phone,omitempty" required:"false" `
    /*
        发件人出编     */
    SZipCode  *string `json:"s_zip_code" required:"true" `
    /*
        电话号码     */
    RTelephone  *string `json:"r_telephone,omitempty" required:"false" `
    /*
        省     */
    RProvName  *string `json:"r_prov_name" required:"true" `
    /*
        收件人街道地址     */
    RAddress  *string `json:"r_address" required:"true" `
    /*
        市     */
    RCityName  *string `json:"r_city_name" required:"true" `
    /*
        手机号码     */
    RMobilePhone  *string `json:"r_mobile_phone,omitempty" required:"false" `
    /*
        区     */
    RDistName  *string `json:"r_dist_name,omitempty" required:"false" `
    /*
        收件人邮编     */
    RZipCode  *string `json:"r_zip_code" required:"true" `
    /*
        收件人名称     */
    RName  *string `json:"r_name" required:"true" `
    /*
        收件人区域ID     */
    RAreaId  *int64 `json:"r_area_id" required:"true" `
}

func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetCompanyId(v int64) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.CompanyId = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetTradeId(v int64) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.TradeId = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetMailNo(v string) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.MailNo = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetOrderSource(v int64) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.OrderSource = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetUserId(v int64) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.UserId = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetLogisType(v int64) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.LogisType = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetOrderType(v int64) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.OrderType = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetShipping(v string) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.Shipping = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetItemJsonString(v string) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.ItemJsonString = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetSAddress(v string) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.SAddress = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetSCityName(v string) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.SCityName = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetSName(v string) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.SName = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetSProvName(v string) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.SProvName = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetSDistName(v string) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.SDistName = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetSAreaId(v int64) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.SAreaId = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetSTelephone(v string) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.STelephone = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetSMobilePhone(v string) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.SMobilePhone = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetSZipCode(v string) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.SZipCode = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetRTelephone(v string) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.RTelephone = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetRProvName(v string) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.RProvName = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetRAddress(v string) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.RAddress = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetRCityName(v string) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.RCityName = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetRMobilePhone(v string) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.RMobilePhone = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetRDistName(v string) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.RDistName = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetRZipCode(v string) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.RZipCode = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetRName(v string) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.RName = &v
    return s
}
func (s *TaobaoLogisticsConsignOrderCreateandsendRequest) SetRAreaId(v int64) *TaobaoLogisticsConsignOrderCreateandsendRequest {
    s.RAreaId = &v
    return s
}

func (req *TaobaoLogisticsConsignOrderCreateandsendRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CompanyId != nil) {
        paramMap["company_id"] = *req.CompanyId
    }
    if(req.TradeId != nil) {
        paramMap["trade_id"] = *req.TradeId
    }
    if(req.MailNo != nil) {
        paramMap["mail_no"] = *req.MailNo
    }
    if(req.OrderSource != nil) {
        paramMap["order_source"] = *req.OrderSource
    }
    if(req.UserId != nil) {
        paramMap["user_id"] = *req.UserId
    }
    if(req.LogisType != nil) {
        paramMap["logis_type"] = *req.LogisType
    }
    if(req.OrderType != nil) {
        paramMap["order_type"] = *req.OrderType
    }
    if(req.Shipping != nil) {
        paramMap["shipping"] = *req.Shipping
    }
    if(req.ItemJsonString != nil) {
        paramMap["item_json_string"] = *req.ItemJsonString
    }
    if(req.SAddress != nil) {
        paramMap["s_address"] = *req.SAddress
    }
    if(req.SCityName != nil) {
        paramMap["s_city_name"] = *req.SCityName
    }
    if(req.SName != nil) {
        paramMap["s_name"] = *req.SName
    }
    if(req.SProvName != nil) {
        paramMap["s_prov_name"] = *req.SProvName
    }
    if(req.SDistName != nil) {
        paramMap["s_dist_name"] = *req.SDistName
    }
    if(req.SAreaId != nil) {
        paramMap["s_area_id"] = *req.SAreaId
    }
    if(req.STelephone != nil) {
        paramMap["s_telephone"] = *req.STelephone
    }
    if(req.SMobilePhone != nil) {
        paramMap["s_mobile_phone"] = *req.SMobilePhone
    }
    if(req.SZipCode != nil) {
        paramMap["s_zip_code"] = *req.SZipCode
    }
    if(req.RTelephone != nil) {
        paramMap["r_telephone"] = *req.RTelephone
    }
    if(req.RProvName != nil) {
        paramMap["r_prov_name"] = *req.RProvName
    }
    if(req.RAddress != nil) {
        paramMap["r_address"] = *req.RAddress
    }
    if(req.RCityName != nil) {
        paramMap["r_city_name"] = *req.RCityName
    }
    if(req.RMobilePhone != nil) {
        paramMap["r_mobile_phone"] = *req.RMobilePhone
    }
    if(req.RDistName != nil) {
        paramMap["r_dist_name"] = *req.RDistName
    }
    if(req.RZipCode != nil) {
        paramMap["r_zip_code"] = *req.RZipCode
    }
    if(req.RName != nil) {
        paramMap["r_name"] = *req.RName
    }
    if(req.RAreaId != nil) {
        paramMap["r_area_id"] = *req.RAreaId
    }
    return paramMap
}

func (req *TaobaoLogisticsConsignOrderCreateandsendRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}