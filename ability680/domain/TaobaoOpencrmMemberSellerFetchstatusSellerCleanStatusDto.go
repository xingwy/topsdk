package domain


type TaobaoOpencrmMemberSellerFetchstatusSellerCleanStatusDto struct {
    /*
        1-数据备份中   2-启动清理   3-清理中   4-清理完成     */
    Status  *int64 `json:"status,omitempty" `

    /*
        sellerNick     */
    SellerNick  *string `json:"seller_nick,omitempty" `

}

func (s *TaobaoOpencrmMemberSellerFetchstatusSellerCleanStatusDto) SetStatus(v int64) *TaobaoOpencrmMemberSellerFetchstatusSellerCleanStatusDto {
    s.Status = &v
    return s
}
func (s *TaobaoOpencrmMemberSellerFetchstatusSellerCleanStatusDto) SetSellerNick(v string) *TaobaoOpencrmMemberSellerFetchstatusSellerCleanStatusDto {
    s.SellerNick = &v
    return s
}
