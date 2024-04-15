package request


type TaobaoVasSubscribeGetRequest struct {
    /*
        淘宝会员名     */
    Nick  *string `json:"nick" required:"true" `
    /*
        商品编码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的商品代码     */
    ArticleCode  *string `json:"article_code" required:"true" `
}

func (s *TaobaoVasSubscribeGetRequest) SetNick(v string) *TaobaoVasSubscribeGetRequest {
    s.Nick = &v
    return s
}
func (s *TaobaoVasSubscribeGetRequest) SetArticleCode(v string) *TaobaoVasSubscribeGetRequest {
    s.ArticleCode = &v
    return s
}

func (req *TaobaoVasSubscribeGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Nick != nil) {
        paramMap["nick"] = *req.Nick
    }
    if(req.ArticleCode != nil) {
        paramMap["article_code"] = *req.ArticleCode
    }
    return paramMap
}

func (req *TaobaoVasSubscribeGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}