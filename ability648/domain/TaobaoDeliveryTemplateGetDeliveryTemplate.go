package domain

import (
	"github.com/xingwy/topsdk/util"
)

type TaobaoDeliveryTemplateGetDeliveryTemplate struct {
	/*
	   模块ID     */
	TemplateId *int64 `json:"template_id,omitempty" `

	/*
	   模板名称，长度不能超过25个字节     */
	Name *string `json:"name,omitempty" `

	/*
		        可选值：0,1,2,3<br>,说明如下<br>
		1)买家承担运费的模版<br>
		0：买家承担服务费<br>
		1: 卖家承担服务费<br>
		2)卖家承担运费的模版<br>
		2:卖家承担运费的模版（集市），卖家承担服务费<br>
		3:卖家承担运费的模版（天猫），卖家承担服务费<br>     */
	Assumer *int64 `json:"assumer,omitempty" `

	/*
		        可选值：0
		说明：
		0:表示按宝贝件数计算运费
		<br/><br/>
		1:表示按宝贝重量计算运费
		<br/><br/>
		3:表示按宝贝体积计算运费     */
	Valuation *int64 `json:"valuation,omitempty" `

	/*
	   运费模板中运费详细信息对象，包含默认运费和指定地区运费     */
	FeeList *[]TaobaoDeliveryTemplateGetTopFee `json:"fee_list,omitempty" `

	/*
	   物流服务模板支持增值服务列表，多个用分号隔开。cod:货到付款 mops：刷卡付款     */
	Supports *string `json:"supports,omitempty" `

	/*
	   模板创建时间     */
	Created *util.LocalTime `json:"created,omitempty" `

	/*
	   模板修改时间     */
	Modified *util.LocalTime `json:"modified,omitempty" `

	/*
	   运费模板上设置的发货地址     */
	Address *string `json:"address,omitempty" `

	/*
	   该模板上设置的卖家发货地址区域ID，如：address为浙江省杭州市西湖去文三路XX号那么这个consign_area_id的值就是西湖区的ID     */
	ConsignAreaId *int64 `json:"consign_area_id,omitempty" `
}

func (s *TaobaoDeliveryTemplateGetDeliveryTemplate) SetTemplateId(v int64) *TaobaoDeliveryTemplateGetDeliveryTemplate {
	s.TemplateId = &v
	return s
}
func (s *TaobaoDeliveryTemplateGetDeliveryTemplate) SetName(v string) *TaobaoDeliveryTemplateGetDeliveryTemplate {
	s.Name = &v
	return s
}
func (s *TaobaoDeliveryTemplateGetDeliveryTemplate) SetAssumer(v int64) *TaobaoDeliveryTemplateGetDeliveryTemplate {
	s.Assumer = &v
	return s
}
func (s *TaobaoDeliveryTemplateGetDeliveryTemplate) SetValuation(v int64) *TaobaoDeliveryTemplateGetDeliveryTemplate {
	s.Valuation = &v
	return s
}
func (s *TaobaoDeliveryTemplateGetDeliveryTemplate) SetFeeList(v []TaobaoDeliveryTemplateGetTopFee) *TaobaoDeliveryTemplateGetDeliveryTemplate {
	s.FeeList = &v
	return s
}
func (s *TaobaoDeliveryTemplateGetDeliveryTemplate) SetSupports(v string) *TaobaoDeliveryTemplateGetDeliveryTemplate {
	s.Supports = &v
	return s
}
func (s *TaobaoDeliveryTemplateGetDeliveryTemplate) SetCreated(v util.LocalTime) *TaobaoDeliveryTemplateGetDeliveryTemplate {
	s.Created = &v
	return s
}
func (s *TaobaoDeliveryTemplateGetDeliveryTemplate) SetModified(v util.LocalTime) *TaobaoDeliveryTemplateGetDeliveryTemplate {
	s.Modified = &v
	return s
}
func (s *TaobaoDeliveryTemplateGetDeliveryTemplate) SetAddress(v string) *TaobaoDeliveryTemplateGetDeliveryTemplate {
	s.Address = &v
	return s
}
func (s *TaobaoDeliveryTemplateGetDeliveryTemplate) SetConsignAreaId(v int64) *TaobaoDeliveryTemplateGetDeliveryTemplate {
	s.ConsignAreaId = &v
	return s
}
