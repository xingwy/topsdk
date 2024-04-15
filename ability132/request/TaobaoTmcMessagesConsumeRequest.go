package request


type TaobaoTmcMessagesConsumeRequest struct {
    /*
        用户分组名称，不传表示消费默认分组，如果应用没有设置用户分组，传入分组名称将会返回错误     */
    GroupName  *string `json:"group_name,omitempty" required:"false" `
    /*
        每次批量消费消息的条数，最小值：10；最大值：200 defalutValue��100    */
    Quantity  *int64 `json:"quantity,omitempty" required:"false" `
}

func (s *TaobaoTmcMessagesConsumeRequest) SetGroupName(v string) *TaobaoTmcMessagesConsumeRequest {
    s.GroupName = &v
    return s
}
func (s *TaobaoTmcMessagesConsumeRequest) SetQuantity(v int64) *TaobaoTmcMessagesConsumeRequest {
    s.Quantity = &v
    return s
}

func (req *TaobaoTmcMessagesConsumeRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.GroupName != nil) {
        paramMap["group_name"] = *req.GroupName
    }
    if(req.Quantity != nil) {
        paramMap["quantity"] = *req.Quantity
    }
    return paramMap
}

func (req *TaobaoTmcMessagesConsumeRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}