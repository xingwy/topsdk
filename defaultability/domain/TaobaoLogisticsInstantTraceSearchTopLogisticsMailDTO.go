package domain


type TaobaoLogisticsInstantTraceSearchTopLogisticsMailDTO struct {
    /*
        运单号     */
    OutSid  *string `json:"out_sid,omitempty" `

    /*
        物流节点列表     */
    TraceList  *[]TaobaoLogisticsInstantTraceSearchTopLogisticsNodeDTO `json:"trace_list,omitempty" `

    /*
        物流公司     */
    CompanyName  *string `json:"company_name,omitempty" `

    /*
        交易单号     */
    Tid  *int64 `json:"tid,omitempty" `

    /*
        当前最新节点     */
    Status  *string `json:"status,omitempty" `

}

func (s *TaobaoLogisticsInstantTraceSearchTopLogisticsMailDTO) SetOutSid(v string) *TaobaoLogisticsInstantTraceSearchTopLogisticsMailDTO {
    s.OutSid = &v
    return s
}
func (s *TaobaoLogisticsInstantTraceSearchTopLogisticsMailDTO) SetTraceList(v []TaobaoLogisticsInstantTraceSearchTopLogisticsNodeDTO) *TaobaoLogisticsInstantTraceSearchTopLogisticsMailDTO {
    s.TraceList = &v
    return s
}
func (s *TaobaoLogisticsInstantTraceSearchTopLogisticsMailDTO) SetCompanyName(v string) *TaobaoLogisticsInstantTraceSearchTopLogisticsMailDTO {
    s.CompanyName = &v
    return s
}
func (s *TaobaoLogisticsInstantTraceSearchTopLogisticsMailDTO) SetTid(v int64) *TaobaoLogisticsInstantTraceSearchTopLogisticsMailDTO {
    s.Tid = &v
    return s
}
func (s *TaobaoLogisticsInstantTraceSearchTopLogisticsMailDTO) SetStatus(v string) *TaobaoLogisticsInstantTraceSearchTopLogisticsMailDTO {
    s.Status = &v
    return s
}
