package request


type TaobaoRdsDbGetdbRequest struct {
    /*
        账户名     */
    AccountName  *string `json:"account_name,omitempty" required:"false" `
    /*
        实例名     */
    InstanceName  *string `json:"instance_name" required:"true" `
}

func (s *TaobaoRdsDbGetdbRequest) SetAccountName(v string) *TaobaoRdsDbGetdbRequest {
    s.AccountName = &v
    return s
}
func (s *TaobaoRdsDbGetdbRequest) SetInstanceName(v string) *TaobaoRdsDbGetdbRequest {
    s.InstanceName = &v
    return s
}

func (req *TaobaoRdsDbGetdbRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.AccountName != nil) {
        paramMap["account_name"] = *req.AccountName
    }
    if(req.InstanceName != nil) {
        paramMap["instance_name"] = *req.InstanceName
    }
    return paramMap
}

func (req *TaobaoRdsDbGetdbRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}