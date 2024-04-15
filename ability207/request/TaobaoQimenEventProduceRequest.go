package request


type TaobaoQimenEventProduceRequest struct {
    /*
        事件状态，如QIMEN_ERP_TRANSFER，QIMEN_ERP_CHECK     */
    Status  *string `json:"status" required:"true" `
    /*
        JSON格式扩展字段     */
    Ext  *string `json:"ext,omitempty" required:"false" `
    /*
        淘宝订单号     */
    Tid  *string `json:"tid" required:"true" `
    /*
        商家平台编码.MAIN:官方渠道,JD:京东,DD:当当,PP:拍拍,YX:易讯,EBAY:ebay,AMAZON:亚马逊,SN:苏宁,GM:国美,WPH:唯品会,JM:聚美,MGJ:蘑菇街,YT:银泰,YHD:1号店,1688:1688,POS:POS门店,OTHER:其他     */
    Platform  *string `json:"platform,omitempty" required:"false" `
    /*
        订单创建时间,数字     */
    Create  *int64 `json:"create,omitempty" required:"false" `
    /*
        外部商家名称。必须同时填写platform     */
    Nick  *string `json:"nick,omitempty" required:"false" `
    /*
        主单号对应的erp单号，转单、审单、通知配货、出库 需要填。拆单、合单场景下不用填     */
    ErpOrderId  *string `json:"erp_order_id,omitempty" required:"false" `
    /*
        淘宝子订单id（拆单、合单场景下不用填，其他场景需要回传,用英文逗号隔开）     */
    TaobaoSubOrderIds  *string `json:"taobao_sub_order_ids,omitempty" required:"false" `
    /*
        触发事件的时间     */
    EventTime  *string `json:"event_time,omitempty" required:"false" `
}

func (s *TaobaoQimenEventProduceRequest) SetStatus(v string) *TaobaoQimenEventProduceRequest {
    s.Status = &v
    return s
}
func (s *TaobaoQimenEventProduceRequest) SetExt(v string) *TaobaoQimenEventProduceRequest {
    s.Ext = &v
    return s
}
func (s *TaobaoQimenEventProduceRequest) SetTid(v string) *TaobaoQimenEventProduceRequest {
    s.Tid = &v
    return s
}
func (s *TaobaoQimenEventProduceRequest) SetPlatform(v string) *TaobaoQimenEventProduceRequest {
    s.Platform = &v
    return s
}
func (s *TaobaoQimenEventProduceRequest) SetCreate(v int64) *TaobaoQimenEventProduceRequest {
    s.Create = &v
    return s
}
func (s *TaobaoQimenEventProduceRequest) SetNick(v string) *TaobaoQimenEventProduceRequest {
    s.Nick = &v
    return s
}
func (s *TaobaoQimenEventProduceRequest) SetErpOrderId(v string) *TaobaoQimenEventProduceRequest {
    s.ErpOrderId = &v
    return s
}
func (s *TaobaoQimenEventProduceRequest) SetTaobaoSubOrderIds(v string) *TaobaoQimenEventProduceRequest {
    s.TaobaoSubOrderIds = &v
    return s
}
func (s *TaobaoQimenEventProduceRequest) SetEventTime(v string) *TaobaoQimenEventProduceRequest {
    s.EventTime = &v
    return s
}

func (req *TaobaoQimenEventProduceRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    if(req.Ext != nil) {
        paramMap["ext"] = *req.Ext
    }
    if(req.Tid != nil) {
        paramMap["tid"] = *req.Tid
    }
    if(req.Platform != nil) {
        paramMap["platform"] = *req.Platform
    }
    if(req.Create != nil) {
        paramMap["create"] = *req.Create
    }
    if(req.Nick != nil) {
        paramMap["nick"] = *req.Nick
    }
    if(req.ErpOrderId != nil) {
        paramMap["erp_order_id"] = *req.ErpOrderId
    }
    if(req.TaobaoSubOrderIds != nil) {
        paramMap["taobao_sub_order_ids"] = *req.TaobaoSubOrderIds
    }
    if(req.EventTime != nil) {
        paramMap["event_time"] = *req.EventTime
    }
    return paramMap
}

func (req *TaobaoQimenEventProduceRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}