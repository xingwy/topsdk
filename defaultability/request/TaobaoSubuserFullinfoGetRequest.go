package request

import (
        "topsdk/util"
    )

type TaobaoSubuserFullinfoGetRequest struct {
    /*
        子账号用户名（传参中sub_id和sub_nick至少需要其中一个，若sub_id与sub_nick同时传入并且合法，那么sub_nick优先，以sub_nick查询子账号）     */
    SubNick  *string `json:"sub_nick,omitempty" required:"false" `
    /*
        子账号ID（传参中sub_id和sub_nick至少需要其中一个，若sub_id与sub_nick同时传入并且合法，那么sub_nick优先，以sub_nick查询子账号）     */
    SubId  *int64 `json:"sub_id,omitempty" required:"false" `
    /*
        传入所需要的参数信息（若不需要获取子账号或主账号的企业邮箱地址，则无需传入该参数；若需要获取子账号或主账号的企业邮箱地址，则需要传入fields；可选参数值为subuser_email和user_email，传入其他参数值均无效；两个参数都需要则以逗号隔开传入即可，例如：subuser_email,user_email）     */
    Fields  *[]string `json:"fields,omitempty" required:"false" `
}

func (s *TaobaoSubuserFullinfoGetRequest) SetSubNick(v string) *TaobaoSubuserFullinfoGetRequest {
    s.SubNick = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetRequest) SetSubId(v int64) *TaobaoSubuserFullinfoGetRequest {
    s.SubId = &v
    return s
}
func (s *TaobaoSubuserFullinfoGetRequest) SetFields(v []string) *TaobaoSubuserFullinfoGetRequest {
    s.Fields = &v
    return s
}

func (req *TaobaoSubuserFullinfoGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.SubNick != nil) {
        paramMap["sub_nick"] = *req.SubNick
    }
    if(req.SubId != nil) {
        paramMap["sub_id"] = *req.SubId
    }
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    return paramMap
}

func (req *TaobaoSubuserFullinfoGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}