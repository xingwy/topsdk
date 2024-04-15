package domain


type AlibabaAscpLogisticsSellerWritelogisticsnodeLocationTopDTO struct {
    /*
        省/直辖市     */
    Province  *string `json:"province,omitempty" `

    /*
        地级市     */
    City  *string `json:"city,omitempty" `

    /*
        区/县     */
    District  *string `json:"district,omitempty" `

    /*
        乡/镇/街道     */
    Town  *string `json:"town,omitempty" `

    /*
        经度，高德地图     */
    Lng  *string `json:"lng,omitempty" `

    /*
        纬度，高德地图     */
    Lat  *string `json:"lat,omitempty" `

}

func (s *AlibabaAscpLogisticsSellerWritelogisticsnodeLocationTopDTO) SetProvince(v string) *AlibabaAscpLogisticsSellerWritelogisticsnodeLocationTopDTO {
    s.Province = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerWritelogisticsnodeLocationTopDTO) SetCity(v string) *AlibabaAscpLogisticsSellerWritelogisticsnodeLocationTopDTO {
    s.City = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerWritelogisticsnodeLocationTopDTO) SetDistrict(v string) *AlibabaAscpLogisticsSellerWritelogisticsnodeLocationTopDTO {
    s.District = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerWritelogisticsnodeLocationTopDTO) SetTown(v string) *AlibabaAscpLogisticsSellerWritelogisticsnodeLocationTopDTO {
    s.Town = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerWritelogisticsnodeLocationTopDTO) SetLng(v string) *AlibabaAscpLogisticsSellerWritelogisticsnodeLocationTopDTO {
    s.Lng = &v
    return s
}
func (s *AlibabaAscpLogisticsSellerWritelogisticsnodeLocationTopDTO) SetLat(v string) *AlibabaAscpLogisticsSellerWritelogisticsnodeLocationTopDTO {
    s.Lat = &v
    return s
}
