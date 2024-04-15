package request


type TaobaoCrmGroupAddRequest struct {
    /*
        分组名称，每个卖家最多可以拥有100个分组     */
    GroupName  *string `json:"group_name" required:"true" `
}

func (s *TaobaoCrmGroupAddRequest) SetGroupName(v string) *TaobaoCrmGroupAddRequest {
    s.GroupName = &v
    return s
}

func (req *TaobaoCrmGroupAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.GroupName != nil) {
        paramMap["group_name"] = *req.GroupName
    }
    return paramMap
}

func (req *TaobaoCrmGroupAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}