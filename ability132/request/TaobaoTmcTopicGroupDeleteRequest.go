package request

import (
        "topsdk/util"
    )

type TaobaoTmcTopicGroupDeleteRequest struct {
    /*
        消息分组名     */
    GroupName  *string `json:"group_name" required:"true" `
    /*
        消息分组Id，一般不用填写，如果分组已经被删除，则根据问题排查工具返回的ID删除路由关系     */
    GroupId  *int64 `json:"group_id,omitempty" required:"false" `
    /*
        消息topic名称，多个以逗号(,)分割     */
    Topics  *[]string `json:"topics" required:"true" `
}

func (s *TaobaoTmcTopicGroupDeleteRequest) SetGroupName(v string) *TaobaoTmcTopicGroupDeleteRequest {
    s.GroupName = &v
    return s
}
func (s *TaobaoTmcTopicGroupDeleteRequest) SetGroupId(v int64) *TaobaoTmcTopicGroupDeleteRequest {
    s.GroupId = &v
    return s
}
func (s *TaobaoTmcTopicGroupDeleteRequest) SetTopics(v []string) *TaobaoTmcTopicGroupDeleteRequest {
    s.Topics = &v
    return s
}

func (req *TaobaoTmcTopicGroupDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.GroupName != nil) {
        paramMap["group_name"] = *req.GroupName
    }
    if(req.GroupId != nil) {
        paramMap["group_id"] = *req.GroupId
    }
    if(req.Topics != nil) {
        paramMap["topics"] = util.ConvertBasicList(*req.Topics)
    }
    return paramMap
}

func (req *TaobaoTmcTopicGroupDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}