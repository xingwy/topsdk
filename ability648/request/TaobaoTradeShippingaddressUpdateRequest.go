package request


type TaobaoTradeShippingaddressUpdateRequest struct {
    /*
        交易编号。     */
    Tid  *int64 `json:"tid" required:"true" `
    /*
        收货人全名。最大长度为50个字节。     */
    ReceiverName  *string `json:"receiver_name,omitempty" required:"false" `
    /*
        座机号码。最大长度为30个字节。传-1表示删除     */
    ReceiverPhone  *string `json:"receiver_phone,omitempty" required:"false" `
    /*
        移动电话。最大长度为11个字节。传-1表示删除     */
    ReceiverMobile  *string `json:"receiver_mobile,omitempty" required:"false" `
    /*
        省份。最大长度为32个字节。如：浙江     */
    ReceiverState  *string `json:"receiver_state,omitempty" required:"false" `
    /*
        城市。最大长度为32个字节。如：杭州     */
    ReceiverCity  *string `json:"receiver_city,omitempty" required:"false" `
    /*
        区/县。最大长度为32个字节。如：西湖区     */
    ReceiverDistrict  *string `json:"receiver_district,omitempty" required:"false" `
    /*
        收货地址。最大长度为228个字节。     */
    ReceiverAddress  *string `json:"receiver_address,omitempty" required:"false" `
    /*
        邮政编码。必须由6个数字组成。注：邮政编码根据地址信息自动填入，不可单独修改     */
    ReceiverZip  *string `json:"receiver_zip,omitempty" required:"false" `
    /*
        四级地址。最大长度为32个字节。如：五常街道     */
    ReceiverTown  *string `json:"receiver_town,omitempty" required:"false" `
}

func (s *TaobaoTradeShippingaddressUpdateRequest) SetTid(v int64) *TaobaoTradeShippingaddressUpdateRequest {
    s.Tid = &v
    return s
}
func (s *TaobaoTradeShippingaddressUpdateRequest) SetReceiverName(v string) *TaobaoTradeShippingaddressUpdateRequest {
    s.ReceiverName = &v
    return s
}
func (s *TaobaoTradeShippingaddressUpdateRequest) SetReceiverPhone(v string) *TaobaoTradeShippingaddressUpdateRequest {
    s.ReceiverPhone = &v
    return s
}
func (s *TaobaoTradeShippingaddressUpdateRequest) SetReceiverMobile(v string) *TaobaoTradeShippingaddressUpdateRequest {
    s.ReceiverMobile = &v
    return s
}
func (s *TaobaoTradeShippingaddressUpdateRequest) SetReceiverState(v string) *TaobaoTradeShippingaddressUpdateRequest {
    s.ReceiverState = &v
    return s
}
func (s *TaobaoTradeShippingaddressUpdateRequest) SetReceiverCity(v string) *TaobaoTradeShippingaddressUpdateRequest {
    s.ReceiverCity = &v
    return s
}
func (s *TaobaoTradeShippingaddressUpdateRequest) SetReceiverDistrict(v string) *TaobaoTradeShippingaddressUpdateRequest {
    s.ReceiverDistrict = &v
    return s
}
func (s *TaobaoTradeShippingaddressUpdateRequest) SetReceiverAddress(v string) *TaobaoTradeShippingaddressUpdateRequest {
    s.ReceiverAddress = &v
    return s
}
func (s *TaobaoTradeShippingaddressUpdateRequest) SetReceiverZip(v string) *TaobaoTradeShippingaddressUpdateRequest {
    s.ReceiverZip = &v
    return s
}
func (s *TaobaoTradeShippingaddressUpdateRequest) SetReceiverTown(v string) *TaobaoTradeShippingaddressUpdateRequest {
    s.ReceiverTown = &v
    return s
}

func (req *TaobaoTradeShippingaddressUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Tid != nil) {
        paramMap["tid"] = *req.Tid
    }
    if(req.ReceiverName != nil) {
        paramMap["receiver_name"] = *req.ReceiverName
    }
    if(req.ReceiverPhone != nil) {
        paramMap["receiver_phone"] = *req.ReceiverPhone
    }
    if(req.ReceiverMobile != nil) {
        paramMap["receiver_mobile"] = *req.ReceiverMobile
    }
    if(req.ReceiverState != nil) {
        paramMap["receiver_state"] = *req.ReceiverState
    }
    if(req.ReceiverCity != nil) {
        paramMap["receiver_city"] = *req.ReceiverCity
    }
    if(req.ReceiverDistrict != nil) {
        paramMap["receiver_district"] = *req.ReceiverDistrict
    }
    if(req.ReceiverAddress != nil) {
        paramMap["receiver_address"] = *req.ReceiverAddress
    }
    if(req.ReceiverZip != nil) {
        paramMap["receiver_zip"] = *req.ReceiverZip
    }
    if(req.ReceiverTown != nil) {
        paramMap["receiver_town"] = *req.ReceiverTown
    }
    return paramMap
}

func (req *TaobaoTradeShippingaddressUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}