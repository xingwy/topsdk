package domain


type TaobaoLogisticsInstantTraceSearchResult struct {
    /*
        是否成功     */
    Success  *bool `json:"success,omitempty" `

    /*
        运单列表     */
    MailList  *[]TaobaoLogisticsInstantTraceSearchTopLogisticsMailDTO `json:"mail_list,omitempty" `

}

func (s *TaobaoLogisticsInstantTraceSearchResult) SetSuccess(v bool) *TaobaoLogisticsInstantTraceSearchResult {
    s.Success = &v
    return s
}
func (s *TaobaoLogisticsInstantTraceSearchResult) SetMailList(v []TaobaoLogisticsInstantTraceSearchTopLogisticsMailDTO) *TaobaoLogisticsInstantTraceSearchResult {
    s.MailList = &v
    return s
}
