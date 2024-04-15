package request

import (
        "topsdk/util"
    )

type TaobaoTmcTopicGroupAddRequest struct {
    /*
        消息分组名，如果不存在，会自动创建     */
    GroupName  *string `json:"group_name" required:"true" `
    /*
        消息topic名称，多个以逗号(,)分割     */
    Topics  *[]string `json:"topics" required:"true" `
}

func (s *TaobaoTmcTopicGroupAddRequest) SetGroupName(v string) *TaobaoTmcTopicGroupAddRequest {
    s.GroupName = &v
    return s
}
func (s *TaobaoTmcTopicGroupAddRequest) SetTopics(v []string) *TaobaoTmcTopicGroupAddRequest {
    s.Topics = &v
    return s
}

func (req *TaobaoTmcTopicGroupAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.GroupName != nil) {
        paramMap["group_name"] = *req.GroupName
    }
    if(req.Topics != nil) {
        paramMap["topics"] = util.ConvertBasicList(*req.Topics)
    }
    return paramMap
}

func (req *TaobaoTmcTopicGroupAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}