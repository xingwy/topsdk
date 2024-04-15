package request


type AlibabaAscpLogisticsSellerSendRequest struct {
    /*
        派送员手机号（支持座机号和400）     */
    DeliveryMobile  *string `json:"delivery_mobile" required:"true" `
    /*
        卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。如果为空，取的卖家的默认取货地址     */
    SenderId  *int64 `json:"sender_id,omitempty" required:"false" `
    /*
        feature参数格式 范例: identCode=tid1:识别码1,识别码2|tid2:识别码3;machineCode=tid3:3C机器号A,3C机器号B identCode为识别码的KEY,machineCode为3C的KEY,多个key之间用”;”分隔 “tid1:识别码1,识别码2|tid2:识别码3”为identCode对应的value。 "|"不同商品间的分隔符。 例1商品和2商品，之间就用"|"分开。 TID就是商品代表的子订单号，对应taobao.trade.fullinfo.get 接口获得的oid字段。(通过OID可以唯一定位到当前商品上) ":"TID和具体传入参数间的分隔符。冒号前表示TID,之后代表该商品的参数属性。 "," 属性间分隔符。（对应商品数量，当存在一个商品的数量超过1个时，用逗号分开）。 具体:当订单中A商品的数量为2个，其中手机串号分别为"12345","67890"。 参数格式：identCode=TIDA:12345,67890。 TIDA对应了A宝贝，冒号后用逗号分隔的"12345","67890".说明本订单A宝贝的数量为2，值分别为"12345","67890"。 当存在"|"时，就说明订单中存在多个商品，商品间用"|"分隔了开来。|"之后的内容含义同上。retailStoreId=12345，发货门店ID或仓信息。retailStoreType=STORE: 发货门店类别，STORE表示门店，WAREHOUSE表示电商仓。对于全渠道订单回传的商家，retailStoreId和retailStoreType字段为必填字段     */
    Feature  *string `json:"feature,omitempty" required:"false" `
    /*
        淘宝交易ID     */
    Tid  *string `json:"tid" required:"true" `
    /*
        需要拆单发货的子订单集合，针对的是一笔交易下有多个子订单需要分开发货的场景；1次可传人多个子订单号，子订单间用逗号隔开；为空表示不做拆单发货。     */
    SubTid  *string `json:"sub_tid,omitempty" required:"false" `
    /*
        派送员姓名     */
    DeliveryName  *string `json:"delivery_name" required:"true" `
    /*
        卖家联系人地址库ID，可以通过taobao.logistics.address.search接口查询到地址库ID。 如果为空，取的卖家的默认退货地址     */
    CancelId  *int64 `json:"cancel_id,omitempty" required:"false" `
    /*
         A（默认，核销签收模式），B（商家回传物流节点模式） defalutValue��A    */
    Mode  *string `json:"mode,omitempty" required:"false" `
}

func (s *AlibabaAscpLogisticsSellerSendRequest) SetDeliveryMobile(v string) *AlibabaAscpLogisticsSellerSendRequest {
    s.DeliveryMobile = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerSendRequest) SetSenderId(v int64) *AlibabaAscpLogisticsSellerSendRequest {
    s.SenderId = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerSendRequest) SetFeature(v string) *AlibabaAscpLogisticsSellerSendRequest {
    s.Feature = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerSendRequest) SetTid(v string) *AlibabaAscpLogisticsSellerSendRequest {
    s.Tid = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerSendRequest) SetSubTid(v string) *AlibabaAscpLogisticsSellerSendRequest {
    s.SubTid = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerSendRequest) SetDeliveryName(v string) *AlibabaAscpLogisticsSellerSendRequest {
    s.DeliveryName = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerSendRequest) SetCancelId(v int64) *AlibabaAscpLogisticsSellerSendRequest {
    s.CancelId = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerSendRequest) SetMode(v string) *AlibabaAscpLogisticsSellerSendRequest {
    s.Mode = &v
    return s
}

func (req *AlibabaAscpLogisticsSellerSendRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.DeliveryMobile != nil) {
        paramMap["delivery_mobile"] = *req.DeliveryMobile
    }
    if(req.SenderId != nil) {
        paramMap["sender_id"] = *req.SenderId
    }
    if(req.Feature != nil) {
        paramMap["feature"] = *req.Feature
    }
    if(req.Tid != nil) {
        paramMap["tid"] = *req.Tid
    }
    if(req.SubTid != nil) {
        paramMap["sub_tid"] = *req.SubTid
    }
    if(req.DeliveryName != nil) {
        paramMap["delivery_name"] = *req.DeliveryName
    }
    if(req.CancelId != nil) {
        paramMap["cancel_id"] = *req.CancelId
    }
    if(req.Mode != nil) {
        paramMap["mode"] = *req.Mode
    }
    return paramMap
}

func (req *AlibabaAscpLogisticsSellerSendRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}