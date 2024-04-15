package request


type TaobaoCrmMemberinfoUpdatePrivyRequest struct {
    /*
        用于描述会员的状态，normal表示正常，blacklist表示黑名单(不享受会员折扣).     */
    Status  *string `json:"status" required:"true" `
    /*
        会员等级，1：普通客户，2：高级会员，3：高级会员 ，4：至尊vip只有正常会员才给予升级，对于status blacklist的会员升级无效     */
    Grade  *int64 `json:"grade,omitempty" required:"false" `
    /*
        北京=1,天津=2,河北省=3,山西省=4,内蒙古自治区=5,辽宁省=6,吉林省=7,黑龙江省=8,上海=9,江苏省=10,浙江省=11,安徽省=12,福建省=13,江西省=14,山东省=15,河南省=16,湖北省=17,湖南省=18, 广东省=19,广西壮族自治区=20,海南省=21,重庆=22,四川省=23,贵州省=24,云南省=25,西藏自治区=26,陕西省=27,甘肃省=28,青海省=29,宁夏回族自治区=30,新疆维吾尔自治区=31,台湾省=32,香港特别行政区=33,澳门特别行政区=34,海外=35，约定36为清除Province设置.请注意:从2014.4.15之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线.     */
    Province  *string `json:"province,omitempty" required:"false" `
    /*
        城市.请注意:从2014.4.15之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线     */
    City  *string `json:"city,omitempty" required:"false" `
    /*
        交易笔数     */
    TradeCount  *int64 `json:"trade_count,omitempty" required:"false" `
    /*
        交易金额，单位：分     */
    TradeAmount  *int64 `json:"trade_amount,omitempty" required:"false" `
    /*
        交易关闭次数     */
    CloseTradeCount  *int64 `json:"close_trade_count,omitempty" required:"false" `
    /*
        交易关闭金额，单位：分     */
    CloseTradeAmount  *int64 `json:"close_trade_amount,omitempty" required:"false" `
    /*
        宝贝件数     */
    ItemNum  *int64 `json:"item_num,omitempty" required:"false" `
    /*
        分组的id集合字符串     */
    GroupIds  *string `json:"group_ids,omitempty" required:"false" `
    /*
        ouid     */
    Ouid  *string `json:"ouid" required:"true" `
}

func (s *TaobaoCrmMemberinfoUpdatePrivyRequest) SetStatus(v string) *TaobaoCrmMemberinfoUpdatePrivyRequest {
    s.Status = &v
    return s
}
func (s *TaobaoCrmMemberinfoUpdatePrivyRequest) SetGrade(v int64) *TaobaoCrmMemberinfoUpdatePrivyRequest {
    s.Grade = &v
    return s
}
func (s *TaobaoCrmMemberinfoUpdatePrivyRequest) SetProvince(v string) *TaobaoCrmMemberinfoUpdatePrivyRequest {
    s.Province = &v
    return s
}
func (s *TaobaoCrmMemberinfoUpdatePrivyRequest) SetCity(v string) *TaobaoCrmMemberinfoUpdatePrivyRequest {
    s.City = &v
    return s
}
func (s *TaobaoCrmMemberinfoUpdatePrivyRequest) SetTradeCount(v int64) *TaobaoCrmMemberinfoUpdatePrivyRequest {
    s.TradeCount = &v
    return s
}
func (s *TaobaoCrmMemberinfoUpdatePrivyRequest) SetTradeAmount(v int64) *TaobaoCrmMemberinfoUpdatePrivyRequest {
    s.TradeAmount = &v
    return s
}
func (s *TaobaoCrmMemberinfoUpdatePrivyRequest) SetCloseTradeCount(v int64) *TaobaoCrmMemberinfoUpdatePrivyRequest {
    s.CloseTradeCount = &v
    return s
}
func (s *TaobaoCrmMemberinfoUpdatePrivyRequest) SetCloseTradeAmount(v int64) *TaobaoCrmMemberinfoUpdatePrivyRequest {
    s.CloseTradeAmount = &v
    return s
}
func (s *TaobaoCrmMemberinfoUpdatePrivyRequest) SetItemNum(v int64) *TaobaoCrmMemberinfoUpdatePrivyRequest {
    s.ItemNum = &v
    return s
}
func (s *TaobaoCrmMemberinfoUpdatePrivyRequest) SetGroupIds(v string) *TaobaoCrmMemberinfoUpdatePrivyRequest {
    s.GroupIds = &v
    return s
}
func (s *TaobaoCrmMemberinfoUpdatePrivyRequest) SetOuid(v string) *TaobaoCrmMemberinfoUpdatePrivyRequest {
    s.Ouid = &v
    return s
}

func (req *TaobaoCrmMemberinfoUpdatePrivyRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Status != nil) {
        paramMap["status"] = *req.Status
    }
    if(req.Grade != nil) {
        paramMap["grade"] = *req.Grade
    }
    if(req.Province != nil) {
        paramMap["province"] = *req.Province
    }
    if(req.City != nil) {
        paramMap["city"] = *req.City
    }
    if(req.TradeCount != nil) {
        paramMap["trade_count"] = *req.TradeCount
    }
    if(req.TradeAmount != nil) {
        paramMap["trade_amount"] = *req.TradeAmount
    }
    if(req.CloseTradeCount != nil) {
        paramMap["close_trade_count"] = *req.CloseTradeCount
    }
    if(req.CloseTradeAmount != nil) {
        paramMap["close_trade_amount"] = *req.CloseTradeAmount
    }
    if(req.ItemNum != nil) {
        paramMap["item_num"] = *req.ItemNum
    }
    if(req.GroupIds != nil) {
        paramMap["group_ids"] = *req.GroupIds
    }
    if(req.Ouid != nil) {
        paramMap["ouid"] = *req.Ouid
    }
    return paramMap
}

func (req *TaobaoCrmMemberinfoUpdatePrivyRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}