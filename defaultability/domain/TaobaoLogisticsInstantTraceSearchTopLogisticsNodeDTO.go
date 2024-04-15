package domain


type TaobaoLogisticsInstantTraceSearchTopLogisticsNodeDTO struct {
    /*
        节点描述     */
    StatusDesc  *string `json:"status_desc,omitempty" `

    /*
        当前节点发生时间     */
    StatusTime  *int64 `json:"status_time,omitempty" `

    /*
        节点枚举     */
    Action  *string `json:"action,omitempty" `

}

func (s *TaobaoLogisticsInstantTraceSearchTopLogisticsNodeDTO) SetStatusDesc(v string) *TaobaoLogisticsInstantTraceSearchTopLogisticsNodeDTO {
    s.StatusDesc = &v
    return s
}
func (s *TaobaoLogisticsInstantTraceSearchTopLogisticsNodeDTO) SetStatusTime(v int64) *TaobaoLogisticsInstantTraceSearchTopLogisticsNodeDTO {
    s.StatusTime = &v
    return s
}
func (s *TaobaoLogisticsInstantTraceSearchTopLogisticsNodeDTO) SetAction(v string) *TaobaoLogisticsInstantTraceSearchTopLogisticsNodeDTO {
    s.Action = &v
    return s
}
