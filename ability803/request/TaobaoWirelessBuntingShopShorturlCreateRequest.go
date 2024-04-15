package request


type TaobaoWirelessBuntingShopShorturlCreateRequest struct {
    /*
        店铺id     */
    ShopId  *string `json:"shop_id,omitempty" required:"false" `
}

func (s *TaobaoWirelessBuntingShopShorturlCreateRequest) SetShopId(v string) *TaobaoWirelessBuntingShopShorturlCreateRequest {
    s.ShopId = &v
    return s
}

func (req *TaobaoWirelessBuntingShopShorturlCreateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ShopId != nil) {
        paramMap["shop_id"] = *req.ShopId
    }
    return paramMap
}

func (req *TaobaoWirelessBuntingShopShorturlCreateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}