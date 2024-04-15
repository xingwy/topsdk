package request


type TaobaoLogisticsInstantTraceSearchRequest struct {
    /*
        运单号     */
    MailNo  *string `json:"mail_no,omitempty" required:"false" `
    /*
        是否拆单     */
    IsSplit  *int64 `json:"is_split,omitempty" required:"false" `
    /*
        子订单列表     */
    SubTid  *string `json:"sub_tid,omitempty" required:"false" `
    /*
        交易单号     */
    Tid  *int64 `json:"tid" required:"true" `
}

func (s *TaobaoLogisticsInstantTraceSearchRequest) SetMailNo(v string) *TaobaoLogisticsInstantTraceSearchRequest {
    s.MailNo = &v
    return s
}
func (s *TaobaoLogisticsInstantTraceSearchRequest) SetIsSplit(v int64) *TaobaoLogisticsInstantTraceSearchRequest {
    s.IsSplit = &v
    return s
}
func (s *TaobaoLogisticsInstantTraceSearchRequest) SetSubTid(v string) *TaobaoLogisticsInstantTraceSearchRequest {
    s.SubTid = &v
    return s
}
func (s *TaobaoLogisticsInstantTraceSearchRequest) SetTid(v int64) *TaobaoLogisticsInstantTraceSearchRequest {
    s.Tid = &v
    return s
}

func (req *TaobaoLogisticsInstantTraceSearchRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.MailNo != nil) {
        paramMap["mail_no"] = *req.MailNo
    }
    if(req.IsSplit != nil) {
        paramMap["is_split"] = *req.IsSplit
    }
    if(req.SubTid != nil) {
        paramMap["sub_tid"] = *req.SubTid
    }
    if(req.Tid != nil) {
        paramMap["tid"] = *req.Tid
    }
    return paramMap
}

func (req *TaobaoLogisticsInstantTraceSearchRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}