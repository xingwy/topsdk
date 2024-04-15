package domain


type TaobaoOpencrmNodereportGetNodeExecuteReportDto struct {
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

func (s *TaobaoOpencrmNodereportGetNodeExecuteReportDto) SetCouponUnUseCount(v int64) *TaobaoOpencrmNodereportGetNodeExecuteReportDto {
    s.CouponUnUseCount = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetNodeExecuteReportDto) SetCouponUsedCount(v int64) *TaobaoOpencrmNodereportGetNodeExecuteReportDto {
    s.CouponUsedCount = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetNodeExecuteReportDto) SetCouponUsingCount(v int64) *TaobaoOpencrmNodereportGetNodeExecuteReportDto {
    s.CouponUsingCount = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetNodeExecuteReportDto) SetFailBillCount(v int64) *TaobaoOpencrmNodereportGetNodeExecuteReportDto {
    s.FailBillCount = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetNodeExecuteReportDto) SetFailCount(v int64) *TaobaoOpencrmNodereportGetNodeExecuteReportDto {
    s.FailCount = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetNodeExecuteReportDto) SetFailStatistics(v string) *TaobaoOpencrmNodereportGetNodeExecuteReportDto {
    s.FailStatistics = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetNodeExecuteReportDto) SetLimitedCount(v int64) *TaobaoOpencrmNodereportGetNodeExecuteReportDto {
    s.LimitedCount = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetNodeExecuteReportDto) SetNodeInstId(v int64) *TaobaoOpencrmNodereportGetNodeExecuteReportDto {
    s.NodeInstId = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetNodeExecuteReportDto) SetSellerId(v int64) *TaobaoOpencrmNodereportGetNodeExecuteReportDto {
    s.SellerId = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetNodeExecuteReportDto) SetShortLinkClickPv(v int64) *TaobaoOpencrmNodereportGetNodeExecuteReportDto {
    s.ShortLinkClickPv = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetNodeExecuteReportDto) SetShortLinkClickUv(v int64) *TaobaoOpencrmNodereportGetNodeExecuteReportDto {
    s.ShortLinkClickUv = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetNodeExecuteReportDto) SetSuccessBillCount(v int64) *TaobaoOpencrmNodereportGetNodeExecuteReportDto {
    s.SuccessBillCount = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetNodeExecuteReportDto) SetSuccessCount(v int64) *TaobaoOpencrmNodereportGetNodeExecuteReportDto {
    s.SuccessCount = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetNodeExecuteReportDto) SetTotalCount(v int64) *TaobaoOpencrmNodereportGetNodeExecuteReportDto {
    s.TotalCount = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetNodeExecuteReportDto) SetUnknownStatusBillCount(v int64) *TaobaoOpencrmNodereportGetNodeExecuteReportDto {
    s.UnknownStatusBillCount = &v
    return s
}
func (s *TaobaoOpencrmNodereportGetNodeExecuteReportDto) SetUnknownStatusCount(v int64) *TaobaoOpencrmNodereportGetNodeExecuteReportDto {
    s.UnknownStatusCount = &v
    return s
}
