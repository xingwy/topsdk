package request


type TaobaoJushitaJdpUserDeleteRequest struct {
    /*
        要删除用户的昵称     */
    Nick  *string `json:"nick,omitempty" required:"false" `
    /*
        要删除用户的openuid     */
    OpenUid  *string `json:"open_uid,omitempty" required:"false" `
}

func (s *TaobaoJushitaJdpUserDeleteRequest) SetNick(v string) *TaobaoJushitaJdpUserDeleteRequest {
    s.Nick = &v
    return s
}
func (s *TaobaoJushitaJdpUserDeleteRequest) SetOpenUid(v string) *TaobaoJushitaJdpUserDeleteRequest {
    s.OpenUid = &v
    return s
}

func (req *TaobaoJushitaJdpUserDeleteRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Nick != nil) {
        paramMap["nick"] = *req.Nick
    }
    if(req.OpenUid != nil) {
        paramMap["open_uid"] = *req.OpenUid
    }
    return paramMap
}

func (req *TaobaoJushitaJdpUserDeleteRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}