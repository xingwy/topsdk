package request

import (
        "topsdk/util"
    )

type TaobaoTmcUserGetRequest struct {
    /*
        需返回的字段列表，多个字段以半角逗号分隔。可选值：TmcUser结构体中的所有字段，一定要返回topic。     */
    Fields  *[]string `json:"fields" required:"true" `
    /*
        用户昵称     */
    Nick  *string `json:"nick" required:"true" `
    /*
        用户所属的平台类型，tbUIC:淘宝用户; icbu: icbu用户;ae:ae用户 defalutValue��tbUIC    */
    UserPlatform  *string `json:"user_platform,omitempty" required:"false" `
}

func (s *TaobaoTmcUserGetRequest) SetFields(v []string) *TaobaoTmcUserGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoTmcUserGetRequest) SetNick(v string) *TaobaoTmcUserGetRequest {
    s.Nick = &v
    return s
}
func (s *TaobaoTmcUserGetRequest) SetUserPlatform(v string) *TaobaoTmcUserGetRequest {
    s.UserPlatform = &v
    return s
}

func (req *TaobaoTmcUserGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    if(req.Nick != nil) {
        paramMap["nick"] = *req.Nick
    }
    if(req.UserPlatform != nil) {
        paramMap["user_platform"] = *req.UserPlatform
    }
    return paramMap
}

func (req *TaobaoTmcUserGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}