package request


type TaobaoKfcKeywordSearchRequest struct {
    /*
        发布信息的淘宝会员名，可以不传     */
    Nick  *string `json:"nick,omitempty" required:"false" `
    /*
        应用点，分为一级应用点、二级应用点。其中一级应用点通常是指某一个系统或产品，比如淘宝的商品应用（taobao_auction）；二级应用点，是指一级应用点下的具体的分类，比如商品标题(title)、商品描述(content)。不同的二级应用可以设置不同关键词。

这里的apply参数是由一级应用点与二级应用点合起来的字符（一级应用点+"."+二级应用点），如taobao_auction.title。


通常apply参数是不需要传递的。如有特殊需求（比如特殊的过滤需求，需要自己维护一套自己词库），需传递此参数。     */
    Apply  *string `json:"apply,omitempty" required:"false" `
    /*
        需要过滤的文本信息     */
    Content  *string `json:"content" required:"true" `
}

func (s *TaobaoKfcKeywordSearchRequest) SetNick(v string) *TaobaoKfcKeywordSearchRequest {
    s.Nick = &v
    return s
}
func (s *TaobaoKfcKeywordSearchRequest) SetApply(v string) *TaobaoKfcKeywordSearchRequest {
    s.Apply = &v
    return s
}
func (s *TaobaoKfcKeywordSearchRequest) SetContent(v string) *TaobaoKfcKeywordSearchRequest {
    s.Content = &v
    return s
}

func (req *TaobaoKfcKeywordSearchRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Nick != nil) {
        paramMap["nick"] = *req.Nick
    }
    if(req.Apply != nil) {
        paramMap["apply"] = *req.Apply
    }
    if(req.Content != nil) {
        paramMap["content"] = *req.Content
    }
    return paramMap
}

func (req *TaobaoKfcKeywordSearchRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}