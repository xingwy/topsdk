package request

import (
        "topsdk/util"
    )

type TaobaoTmcGroupAddRequest struct {
    /*
        分组名称，同一个应用下需要保证唯一性，最长32个字符。添加分组后，消息通道会为用户的消息分配独立分组，但之前的消息还是存储于默认分组中。不能以default开头，default开头为系统默认组。     */
    GroupName  *string `json:"group_name" required:"true" `
    /*
        用户昵称列表，以半角逗号分隔，支持子账号，支持增量添加用户     */
    Nicks  *[]string `json:"nicks" required:"true" `
    /*
        用户所属于的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户 defalutValue��tbUIC    */
    UserPlatform  *string `json:"user_platform,omitempty" required:"false" `
}

func (s *TaobaoTmcGroupAddRequest) SetGroupName(v string) *TaobaoTmcGroupAddRequest {
    s.GroupName = &v
    return s
}
func (s *TaobaoTmcGroupAddRequest) SetNicks(v []string) *TaobaoTmcGroupAddRequest {
    s.Nicks = &v
    return s
}
func (s *TaobaoTmcGroupAddRequest) SetUserPlatform(v string) *TaobaoTmcGroupAddRequest {
    s.UserPlatform = &v
    return s
}

func (req *TaobaoTmcGroupAddRequest) ToMap() map[string]interface{} {
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

func (req *TaobaoTmcGroupAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}