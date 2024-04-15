package domain


type TaobaoOpencrmNodereportGetbcNodeExecuteReportDto struct {
    /*
        优惠券未使用人数     */
    CouponUnUseCount  *int64 `json:"coupon_un_use_count,omitempty" `

    /*
        优惠券已使用人数     */
    CouponUsedCount  *int64 `json:"coupon_used_count,omitempty" `

    /*
        优惠券使用中人数     */
    CouponUsingCount  *int64 `json:"coupon_using_count,omitempty" `

    /*
        失败短信计费条数     */
    FailBillCount  *int64 `json:"fail_bill_count,omitempty" `

    /*
        失败人数     */
    FailCount  *int64 `json:"fail_count,omitempty" `

    /*
        失败统计信息     */
    FailStatistics  *string `json:"fail_statistics,omitempty" `

    /*
        受限制人数     */
    LimitedCount  *int64 `json:"limited_count,omitempty" `

    /*
        节点实例ID     */
    NodeInstId  *int64 `json:"node_inst_id,omitempty" `

    /*
        sellerId     */
    SellerId  *int64 `json:"seller_id,omitempty" `

    /*
        shortLinkClickPv     */
    ShortLinkClickPv  *int64 `json:"short_link_click_pv,omitempty" `

    /*
        shortLinkClickUv     */
    ShortLinkClickUv  *int64 `json:"short_link_click_uv,omitempty" `

    /*
        成功短信计费条数     */
    SuccessBillCount  *int64 `json:"success_bill_count,omitempty" `

    /*
        成功人数     */
    SuccessCount  *int64 `json:"success_count,omitempty" `

    /*
        提交总人数     */
    TotalCount  *int64 `json:"total_count,omitempty" `

    /*
        状态未知短信计费条数     */
    UnknownStatusBillCount  *int64 `json:"unknown_status_bill_count,omitempty" `

    /*
        状态未知短信人数     */
    UnknownStatusCount  *int64 `json:"unknown_status_count,omitempty" `

}

func (s *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto) SetCouponUnUseCount(v int64) *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto {
    s.CouponUnUseCount = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto) SetCouponUsedCount(v int64) *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto {
    s.CouponUsedCount = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto) SetCouponUsingCount(v int64) *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto {
    s.CouponUsingCount = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto) SetFailBillCount(v int64) *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto {
    s.FailBillCount = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto) SetFailCount(v int64) *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto {
    s.FailCount = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto) SetFailStatistics(v string) *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto {
    s.FailStatistics = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto) SetLimitedCount(v int64) *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto {
    s.LimitedCount = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto) SetNodeInstId(v int64) *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto {
    s.NodeInstId = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto) SetSellerId(v int64) *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto {
    s.SellerId = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto) SetShortLinkClickPv(v int64) *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto {
    s.ShortLinkClickPv = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto) SetShortLinkClickUv(v int64) *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto {
    s.ShortLinkClickUv = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto) SetSuccessBillCount(v int64) *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto {
    s.SuccessBillCount = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto) SetSuccessCount(v int64) *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto {
    s.SuccessCount = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto) SetTotalCount(v int64) *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto {
    s.TotalCount = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto) SetUnknownStatusBillCount(v int64) *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto {
    s.UnknownStatusBillCount = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto) SetUnknownStatusCount(v int64) *TaobaoOpencrmNodereportGetbcNodeExecuteReportDto {
    s.UnknownStatusCount = &v
    return s
}
