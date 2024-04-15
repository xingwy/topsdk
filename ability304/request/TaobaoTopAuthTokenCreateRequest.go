package request


type TaobaoTopAuthTokenCreateRequest struct {
    /*
        授权code，grantType==authorization_code 时需要     */
    Code  *string `json:"code" required:"true" `
    /*
        非必填，与生成code的uuid配对，使用方式参考文档     */
    Uuid  *string `json:"uuid,omitempty" required:"false" `
}

func (s *TaobaoTopAuthTokenCreateRequest) SetCode(v string) *TaobaoTopAuthTokenCreateRequest {
    s.Code = &v
    return s
}
func (s *TaobaoTopAuthTokenCreateRequest) SetUuid(v string) *TaobaoTopAuthTokenCreateRequest {
    s.Uuid = &v
    return s
}

func (req *TaobaoTopAuthTokenCreateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Code != nil) {
        paramMap["code"] = *req.Code
    }
    if(req.Uuid != nil) {
        paramMap["uuid"] = *req.Uuid
    }
    return paramMap
}

func (req *TaobaoTopAuthTokenCreateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}