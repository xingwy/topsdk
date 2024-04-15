package domain


type TaobaoWlbStoresBaseinfoGetStoreInfo struct {
    /*
        仓库真实名字     */
    RealName  *string `json:"real_name,omitempty" `

    /*
        XXX果园     */
    Name  *string `json:"name,omitempty" `

    /*
        仓库服务代码     */
    ServiceCode  *string `json:"service_code,omitempty" `

    /*
        详细地址     */
    Address  *string `json:"address,omitempty" `

}

func (s *TaobaoWlbStoresBaseinfoGetStoreInfo) SetRealName(v string) *TaobaoWlbStoresBaseinfoGetStoreInfo {
    s.RealName = &v
    return s
}
func (s *TaobaoWlbStoresBaseinfoGetStoreInfo) SetName(v string) *TaobaoWlbStoresBaseinfoGetStoreInfo {
    s.Name = &v
    return s
}
func (s *TaobaoWlbStoresBaseinfoGetStoreInfo) SetServiceCode(v string) *TaobaoWlbStoresBaseinfoGetStoreInfo {
    s.ServiceCode = &v
    return s
}
func (s *TaobaoWlbStoresBaseinfoGetStoreInfo) SetAddress(v string) *TaobaoWlbStoresBaseinfoGetStoreInfo {
    s.Address = &v
    return s
}
