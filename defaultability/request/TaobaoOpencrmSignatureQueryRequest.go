package request


type TaobaoOpencrmSignatureQueryRequest struct {
    /*
        状态，1-待提交审核，2-审核中，3-审核不通过，4-审核通过     */
    Status  *int64 `json:"status,omitempty" required:"false" `
}

func (s *TaobaoOpencrmSignatureQueryRequest) SetStatus(v int64) *TaobaoOpencrmSignatureQueryRequest {
    s.Status = &v
    return s
}

func (req *TaobaoOpencrmSignatureQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    return paramMap
}

func (req *TaobaoOpencrmSignatureQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}