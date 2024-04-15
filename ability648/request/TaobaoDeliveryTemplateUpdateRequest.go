package request


type TaobaoDeliveryTemplateUpdateRequest struct {
    /*
        模板名称，长度不能大于50个字节     */
    Name  *string `json:"name,omitempty" required:"false" `
    /*
        需要修改的模板对应的模板ID     */
    TemplateId  *int64 `json:"template_id" required:"true" `
    /*
        可选值：0,1 <br>  说明<br>0:表示买家承担服务费;<br>1:表示卖家承担服务费     */
    Assumer  *int64 `json:"assumer,omitempty" required:"false" `
    /*
        运费方式:平邮 (post),快递公司(express),EMS (ems),货到付款(cod)结构:value1;value2;value3;value4 
如: post;express;ems;cod 
<br/><font color=red>
注意:在添加多个运费方式时,字符串中使用 ";" 分号区分。template_dests(指定地区) template_start_standards(首费标准)、template_start_fees(首费)、template_add_standards(增费标准)、template_add_fees(增费)必须与template_types的分号数量相同. 
 </font>
<br/>
<font color=blue>
普通用户：post,ems,express三种运费方式必须填写一个，不能填写cod。
货到付款用户：如果填写了cod运费方式，则post,ems,express三种运费方式也必须填写一个，如果没有填写cod则填写的运费方式中必须存在express</font>     */
    TemplateTypes  *string `json:"template_types" required:"true" `
    /*
        邮费子项涉及的地区.结构: value1;value2;value3,value4
<br>如:1,110000;1,110000;1,310000;1,320000,330000。 aredId解释(1=全国,110000=北京,310000=上海,320000=江苏,330000=浙江)
如果template_types设置为post;ems;exmpress;cod则表示为平邮(post)指定默认地区(全国)和北京地区的运费;其他的类似以分号区分一一对应
<br/>可以用taobao.areas.get接口获取.或者参考：http://www.stats.gov.cn/tjbz/xzqhdm/t20080215_402462675.htm
<br/><font color=red>每个运费方式设置的设涉及地区中必须包含全国地区（areaId=1）表示默认运费,可以只设置默认运费</font>
<br><font color=blue>注意:为多个地区指定指定不同（首费标准、首费、增费标准、增费一项不一样就算不同）的运费以逗号","区分，
template_start_standards(首费标准)、template_start_fees(首费)、
template_add_standards(增费标准)、
template_add_fees(增费)必须与template_types分号数量相同。如果为需要为多个地区指定相同运费则地区之间用“|”隔开即可。</font>     */
    TemplateDests  *string `json:"template_dests" required:"true" `
    /*
        首费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br><font color=blue>首费标准目前只能为1</br>
<br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>     */
    TemplateStartStandards  *string `json:"template_start_standards" required:"true" `
    /*
        首费：输入0.01-999.99（最多包含两位小数）
<br/><font color=blue> 首费不能为0</font><br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>     */
    TemplateStartFees  *string `json:"template_start_fees" required:"true" `
    /*
        增费标准：当valuation(记价方式)为0时输入1-9999范围内的整数<br><font color=blue>增费标准目前只能为1</font>
<br><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>     */
    TemplateAddStandards  *string `json:"template_add_standards" required:"true" `
    /*
        增费：输入0.00-999.99（最多包含两位小数）<br/><font color=blue>增费可以为0</font><br/><font color=red>输入的格式分号个数和template_types数量一致，逗号个数必须与template_dests以分号隔开之后一一对应的数量一致</font>     */
    TemplateAddFees  *string `json:"template_add_fees" required:"true" `
}

func (s *TaobaoDeliveryTemplateUpdateRequest) SetName(v string) *TaobaoDeliveryTemplateUpdateRequest {
    s.Name = &v
    return s
}
func (s *TaobaoDeliveryTemplateUpdateRequest) SetTemplateId(v int64) *TaobaoDeliveryTemplateUpdateRequest {
    s.TemplateId = &v
    return s
}
func (s *TaobaoDeliveryTemplateUpdateRequest) SetAssumer(v int64) *TaobaoDeliveryTemplateUpdateRequest {
    s.Assumer = &v
    return s
}
func (s *TaobaoDeliveryTemplateUpdateRequest) SetTemplateTypes(v string) *TaobaoDeliveryTemplateUpdateRequest {
    s.TemplateTypes = &v
    return s
}
func (s *TaobaoDeliveryTemplateUpdateRequest) SetTemplateDests(v string) *TaobaoDeliveryTemplateUpdateRequest {
    s.TemplateDests = &v
    return s
}
func (s *TaobaoDeliveryTemplateUpdateRequest) SetTemplateStartStandards(v string) *TaobaoDeliveryTemplateUpdateRequest {
    s.TemplateStartStandards = &v
    return s
}
func (s *TaobaoDeliveryTemplateUpdateRequest) SetTemplateStartFees(v string) *TaobaoDeliveryTemplateUpdateRequest {
    s.TemplateStartFees = &v
    return s
}
func (s *TaobaoDeliveryTemplateUpdateRequest) SetTemplateAddStandards(v string) *TaobaoDeliveryTemplateUpdateRequest {
    s.TemplateAddStandards = &v
    return s
}
func (s *TaobaoDeliveryTemplateUpdateRequest) SetTemplateAddFees(v string) *TaobaoDeliveryTemplateUpdateRequest {
    s.TemplateAddFees = &v
    return s
}

func (req *TaobaoDeliveryTemplateUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Name != nil) {
        paramMap["name"] = *req.Name
    }
    if(req.TemplateId != nil) {
        paramMap["template_id"] = *req.TemplateId
    }
    if(req.Assumer != nil) {
        paramMap["assumer"] = *req.Assumer
    }
    if(req.TemplateTypes != nil) {
        paramMap["template_types"] = *req.TemplateTypes
    }
    if(req.TemplateDests != nil) {
        paramMap["template_dests"] = *req.TemplateDests
    }
    if(req.TemplateStartStandards != nil) {
        paramMap["template_start_standards"] = *req.TemplateStartStandards
    }
    if(req.TemplateStartFees != nil) {
        paramMap["template_start_fees"] = *req.TemplateStartFees
    }
    if(req.TemplateAddStandards != nil) {
        paramMap["template_add_standards"] = *req.TemplateAddStandards
    }
    if(req.TemplateAddFees != nil) {
        paramMap["template_add_fees"] = *req.TemplateAddFees
    }
    return paramMap
}

func (req *TaobaoDeliveryTemplateUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}