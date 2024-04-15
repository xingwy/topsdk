package domain


type TaobaoFuwuSkuGetArticleViewResult struct {
    /*
        服务code     */
    ArticleCode  *string `json:"article_code,omitempty" `

    /*
        服务简介     */
    ArticleCommment  *string `json:"article_commment,omitempty" `

    /*
        sku详情列表     */
    ArticleItemViewUnits  *[]TaobaoFuwuSkuGetArticleItemViewUnit `json:"article_item_view_units,omitempty" `

    /*
        服务名称     */
    ArticleName  *string `json:"article_name,omitempty" `

    /*
        错误码     */
    ErrorCode  *string `json:"error_code,omitempty" `

    /*
        错误消息     */
    ErrorMsg  *string `json:"error_msg,omitempty" `

    /*
        用户淘宝nick     */
    Nick  *string `json:"nick,omitempty" `

    /*
        服务图片地址     */
    PictUrl  *string `json:"pict_url,omitempty" `

}

func (s *TaobaoFuwuSkuGetArticleViewResult) SetArticleCode(v string) *TaobaoFuwuSkuGetArticleViewResult {
    s.ArticleCode = &v
    return s
}
func (s *TaobaoFuwuSkuGetArticleViewResult) SetArticleCommment(v string) *TaobaoFuwuSkuGetArticleViewResult {
    s.ArticleCommment = &v
    return s
}
func (s *TaobaoFuwuSkuGetArticleViewResult) SetArticleItemViewUnits(v []TaobaoFuwuSkuGetArticleItemViewUnit) *TaobaoFuwuSkuGetArticleViewResult {
    s.ArticleItemViewUnits = &v
    return s
}
func (s *TaobaoFuwuSkuGetArticleViewResult) SetArticleName(v string) *TaobaoFuwuSkuGetArticleViewResult {
    s.ArticleName = &v
    return s
}
func (s *TaobaoFuwuSkuGetArticleViewResult) SetErrorCode(v string) *TaobaoFuwuSkuGetArticleViewResult {
    s.ErrorCode = &v
    return s
}
func (s *TaobaoFuwuSkuGetArticleViewResult) SetErrorMsg(v string) *TaobaoFuwuSkuGetArticleViewResult {
    s.ErrorMsg = &v
    return s
}
func (s *TaobaoFuwuSkuGetArticleViewResult) SetNick(v string) *TaobaoFuwuSkuGetArticleViewResult {
    s.Nick = &v
    return s
}
func (s *TaobaoFuwuSkuGetArticleViewResult) SetPictUrl(v string) *TaobaoFuwuSkuGetArticleViewResult {
    s.PictUrl = &v
    return s
}
