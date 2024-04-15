package request


type TaobaoOpencrmNodereportGetbcRequest struct {
    /*
        节点实例ID     */
    NodeInstId  *int64 `json:"node_inst_id" required:"true" `
    /*
        1:B版   4:C版     */
    AppType  *int64 `json:"app_type,omitempty" required:"false" `
}

func (s *TaobaoOpencrmNodereportGetbcRequest) SetNodeInstId(v int64) *TaobaoOpencrmNodereportGetbcRequest {
    s.NodeInstId = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetbcRequest) SetAppType(v int64) *TaobaoOpencrmNodereportGetbcRequest {
    s.AppType = &v
    return s
}

func (req *TaobaoOpencrmNodereportGetbcRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.NodeInstId != nil) {
        paramMap["node_inst_id"] = *req.NodeInstId
    }
    if(req.AppType != nil) {
        paramMap["app_type"] = *req.AppType
    }
    return paramMap
}

func (req *TaobaoOpencrmNodereportGetbcRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}