package request


type TaobaoFuwuSkuGetRequest struct {
    /*
        服务code     */
    ArticleCode  *string `json:"article_code" required:"true" `
    /*
        用户的淘宝nick     */
    Nick  *string `json:"nick" required:"true" `
}

func (s *TaobaoFuwuSkuGetRequest) SetArticleCode(v string) *TaobaoFuwuSkuGetRequest {
    s.ArticleCode = &v
    return s
}
func (s *TaobaoFuwuSkuGetRequest) SetNick(v string) *TaobaoFuwuSkuGetRequest {
    s.Nick = &v
    return s
}

func (req *TaobaoFuwuSkuGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ArticleCode != nil) {
        paramMap["article_code"] = *req.ArticleCode
    }
    if(req.Nick != nil) {
        paramMap["nick"] = *req.Nick
    }
    return paramMap
}

func (req *TaobaoFuwuSkuGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}