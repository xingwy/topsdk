package domain


type TaobaoJdsTradesStatisticsGetTradeStat struct {
    /*
        数量     */
    Count  *int64 `json:"count,omitempty" `

    /*
        状态名称     */
    Status  *string `json:"status,omitempty" `

}

func (s *TaobaoJdsTradesStatisticsGetTradeStat) SetCount(v int64) *TaobaoJdsTradesStatisticsGetTradeStat {
    s.Count = &v
    return s
}
func (s *TaobaoJdsTradesStatisticsGetTradeStat) SetStatus(v string) *TaobaoJdsTradesStatisticsGetTradeStat {
    s.Status = &v
    return s
}
