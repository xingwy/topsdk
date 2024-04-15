package request

import (
        "topsdk/util"
    )

type TaobaoTmcGroupDeleteRequest struct {
    /*
        分组名称，分组删除后，用户的消息将会存储于默认分组中。警告：由于分组已经删除，用户之前未消费的消息将无法再获取。不能以default开头，default开头为系统默认组。     */
    GroupName  *string `json:"group_name" required:"true" `
    /*
        用户列表，不传表示删除整个分组，如果用户全部删除后，也会自动删除整个分组     */
    Nicks  *[]string `json:"nicks,omitempty" required:"false" `
    /*
        用户所属于的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户 defalutValue��tbUIC    */
    UserPlatform  *string `json:"user_platform,omitempty" required:"false" `
}

func (s *TaobaoTmcGroupDeleteRequest) SetGroupName(v string) *TaobaoTmcGroupDeleteRequest {
    s.GroupName = &v
    return s
}
func (s *TaobaoTmcGroupDeleteRequest) SetNicks(v []string) *TaobaoTmcGroupDeleteRequest {
    s.Nicks = &v
    return s
}
func (s *TaobaoTmcGroupDeleteRequest) SetUserPlatform(v string) *TaobaoTmcGroupDeleteRequest {
    s.UserPlatform = &v
    return s
}

func (req *TaobaoTmcGroupDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.GroupName != nil) {
        paramMap["group_name"] = *req.GroupName
    }
    if(req.Nicks != nil) {
        paramMap["nicks"] = util.ConvertBasicList(*req.Nicks)
    }
    if(req.UserPlatform != nil) {
        paramMap["user_platform"] = *req.UserPlatform
    }
    return paramMap
}

func (req *TaobaoTmcGroupDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}