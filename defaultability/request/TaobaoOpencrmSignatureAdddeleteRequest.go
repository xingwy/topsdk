package request


type TaobaoOpencrmSignatureAdddeleteRequest struct {
    /*
        短信签名     */
    Signature  *string `json:"signature" required:"true" `
    /*
        操作类型，1-添加，2-删除     */
    OperateType  *int64 `json:"operate_type" required:"true" `
}

func (s *TaobaoOpencrmSignatureAdddeleteRequest) SetSignature(v string) *TaobaoOpencrmSignatureAdddeleteRequest {
    s.Signature = &v
    return s
}
func (s *TaobaoOpencrmSignatureAdddeleteRequest) SetOperateType(v int64) *TaobaoOpencrmSignatureAdddeleteRequest {
    s.OperateType = &v
    return s
}

func (req *TaobaoOpencrmSignatureAdddeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Signature != nil) {
        paramMap["signature"] = *req.Signature
    }
    if(req.OperateType != nil) {
        paramMap["operate_type"] = *req.OperateType
    }
    return paramMap
}

func (req *TaobaoOpencrmSignatureAdddeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}