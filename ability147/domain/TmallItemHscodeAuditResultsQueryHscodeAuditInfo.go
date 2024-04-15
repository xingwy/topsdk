package domain


type TmallItemHscodeAuditResultsQueryHscodeAuditInfo struct {
    /*
        商品ID     */
    ItemId  *int64 `json:"item_id,omitempty" `

    /*
        hscode信息当前审核状态的具体说明     */
    Reason  *string `json:"reason,omitempty" `

    /*
        商品或SKU使用的HS海关代码     */
    Hscode  *string `json:"hscode,omitempty" `

    /*
        SKU的ID     */
    SkuId  *int64 `json:"sku_id,omitempty" `

    /*
        hscode信息当前审核状态，HISTORY_ITEM：历史已上架商品，REJECT：审核未通过，AUDITING：审核中，PASS：审核通过，ERROR：获取审核状态异常     */
    Status  *string `json:"status,omitempty" `

}

func (s *TmallItemHscodeAuditResultsQueryHscodeAuditInfo) SetItemId(v int64) *TmallItemHscodeAuditResultsQueryHscodeAuditInfo {
    s.ItemId = &v
    return s
}
func (s *TmallItemHscodeAuditResultsQueryHscodeAuditInfo) SetReason(v string) *TmallItemHscodeAuditResultsQueryHscodeAuditInfo {
    s.Reason = &v
    return s
}
func (s *TmallItemHscodeAuditResultsQueryHscodeAuditInfo) SetHscode(v string) *TmallItemHscodeAuditResultsQueryHscodeAuditInfo {
    s.Hscode = &v
    return s
}
func (s *TmallItemHscodeAuditResultsQueryHscodeAuditInfo) SetSkuId(v int64) *TmallItemHscodeAuditResultsQueryHscodeAuditInfo {
    s.SkuId = &v
    return s
}
func (s *TmallItemHscodeAuditResultsQueryHscodeAuditInfo) SetStatus(v string) *TmallItemHscodeAuditResultsQueryHscodeAuditInfo {
    s.Status = &v
    return s
}
