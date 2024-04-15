package request

import (
        "topsdk/util"
    )

type TaobaoTmcMessagesConfirmRequest struct {
    /*
        分组名称，不传代表默认分组     */
    GroupName  *string `json:"group_name,omitempty" required:"false" `
    /*
        处理成功的消息ID列表 最大 200个ID     */
    SMessageIds  *[]int64 `json:"s_message_ids" required:"true" `
    /*
        处理失败的消息ID列表--已废弃，无需传此字段     */
    FMessageIds  *[]int64 `json:"f_message_ids,omitempty" required:"false" `
}

func (s *TaobaoTmcMessagesConfirmRequest) SetGroupName(v string) *TaobaoTmcMessagesConfirmRequest {
    s.GroupName = &v
    return s
}
func (s *TaobaoTmcMessagesConfirmRequest) SetSMessageIds(v []int64) *TaobaoTmcMessagesConfirmRequest {
    s.SMessageIds = &v
    return s
}
func (s *TaobaoTmcMessagesConfirmRequest) SetFMessageIds(v []int64) *TaobaoTmcMessagesConfirmRequest {
    s.FMessageIds = &v
    return s
}

func (req *TaobaoTmcMessagesConfirmRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.GroupName != nil) {
        paramMap["group_name"] = *req.GroupName
    }
    if(req.SMessageIds != nil) {
        paramMap["s_message_ids"] = util.ConvertBasicList(*req.SMessageIds)
    }
    if(req.FMessageIds != nil) {
        paramMap["f_message_ids"] = util.ConvertBasicList(*req.FMessageIds)
    }
    return paramMap
}

func (req *TaobaoTmcMessagesConfirmRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}