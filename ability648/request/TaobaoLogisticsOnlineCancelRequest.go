package request


type TaobaoLogisticsOnlineCancelRequest struct {
    /*
        淘宝交易ID     */
    Tid  *int64 `json:"tid" required:"true" `
}

func (s *TaobaoLogisticsOnlineCancelRequest) SetTid(v int64) *TaobaoLogisticsOnlineCancelRequest {
    s.Tid = &v
    return s
}

func (req *TaobaoLogisticsOnlineCancelRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Tid != nil) {
        paramMap["tid"] = *req.Tid
    }
    return paramMap
}

func (req *TaobaoLogisticsOnlineCancelRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}