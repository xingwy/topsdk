package request


type TaobaoOpencrmNodereportGetRequest struct {
    /*
        节点实例ID     */
    NodeInstId  *int64 `json:"node_inst_id" required:"true" `
}

func (s *TaobaoOpencrmNodereportGetRequest) SetNodeInstId(v int64) *TaobaoOpencrmNodereportGetRequest {
    s.NodeInstId = &v
    return s
}

func (req *TaobaoOpencrmNodereportGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.NodeInstId != nil) {
        paramMap["node_inst_id"] = *req.NodeInstId
    }
    return paramMap
}

func (req *TaobaoOpencrmNodereportGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}