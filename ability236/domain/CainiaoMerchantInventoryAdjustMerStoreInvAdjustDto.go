package domain


type CainiaoMerchantInventoryAdjustMerStoreInvAdjustDto struct {
    /*
        库存类型     */
    InventoryType  *int64 `json:"inventory_type,omitempty" `

    /*
        扩展属性     */
    Attribute  *string `json:"attribute,omitempty" `

    /*
        数量     */
    Quantity  *int64 `json:"quantity,omitempty" `

    /*
        外部操作唯一编码     */
    OutBizCode  *string `json:"out_biz_code,omitempty" `

    /*
        仓库编码     */
    StoreCode  *string `json:"store_code,omitempty" `

    /*
        货品id     */
    ScItemId  *int64 `json:"sc_item_id,omitempty" `

}

func (s *CainiaoMerchantInventoryAdjustMerStoreInvAdjustDto) SetInventoryType(v int64) *CainiaoMerchantInventoryAdjustMerStoreInvAdjustDto {
    s.InventoryType = &v
    return s
}
func (s *CainiaoMerchantInventoryAdjustMerStoreInvAdjustDto) SetAttribute(v string) *CainiaoMerchantInventoryAdjustMerStoreInvAdjustDto {
    s.Attribute = &v
    return s
}
func (s *CainiaoMerchantInventoryAdjustMerStoreInvAdjustDto) SetQuantity(v int64) *CainiaoMerchantInventoryAdjustMerStoreInvAdjustDto {
    s.Quantity = &v
    return s
}
func (s *CainiaoMerchantInventoryAdjustMerStoreInvAdjustDto) SetOutBizCode(v string) *CainiaoMerchantInventoryAdjustMerStoreInvAdjustDto {
    s.OutBizCode = &v
    return s
}
func (s *CainiaoMerchantInventoryAdjustMerStoreInvAdjustDto) SetStoreCode(v string) *CainiaoMerchantInventoryAdjustMerStoreInvAdjustDto {
    s.StoreCode = &v
    return s
}
func (s *CainiaoMerchantInventoryAdjustMerStoreInvAdjustDto) SetScItemId(v int64) *CainiaoMerchantInventoryAdjustMerStoreInvAdjustDto {
    s.ScItemId = &v
    return s
}
