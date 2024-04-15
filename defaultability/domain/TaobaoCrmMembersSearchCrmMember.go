package domain

import (
        "topsdk/util"
    )

type TaobaoCrmMembersSearchCrmMember struct {
    /*
        购买的宝贝件数     */
    ItemNum  *int64 `json:"item_num,omitempty" `

    /*
        交易关闭的宝贝件数     */
    ItemCloseCount  *int64 `json:"item_close_count,omitempty" `

    /*
        交易关闭的金额     */
    CloseTradeAmount  *string `json:"close_trade_amount,omitempty" `

    /*
        显示会员的状态，normal正常，blacklist黑名单     */
    Status  *string `json:"status,omitempty" `

    /*
        会员拥有的所有分组     */
    GroupIds  *string `json:"group_ids,omitempty" `

    /*
        关系来源，1交易成功，2未成交，3卖家主动吸纳     */
    RelationSource  *int64 `json:"relation_source,omitempty" `

    /*
        交易关闭的的笔数     */
    CloseTradeCount  *int64 `json:"close_trade_count,omitempty" `

    /*
        最后交易时间     */
    LastTradeTime  *util.LocalTime `json:"last_trade_time,omitempty" `

    /*
        最后一次交易的订单号.注:该字段从2014.4.23之后不再返回.     */
    BizOrderId  *int64 `json:"biz_order_id,omitempty" `

    /*
        城市.请注意:从2014.4.23之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&amp;path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线.     */
    City  *string `json:"city,omitempty" `

    /*
        买家昵称     */
    BuyerNick  *string `json:"buyer_nick,omitempty" `

    /*
        openuid     */
    Ouid  *string `json:"ouid,omitempty" `

    /*
        北京=1,天津=2,河北省=3,山西省=4,内蒙古自治区=5,辽宁省=6,吉林省=7,黑龙江省=8,上海=9,江苏省=10,浙江省=11,安徽省=12,福建省=13,江西省=14,山东省=15,河南省=16,湖北省=17,湖南省=18, 广东省=19,广西壮族自治区=20,海南省=21,重庆=22,四川省=23,贵州省=24,云南省=25,西藏自治区26,陕西省=27,甘肃省=28,青海省=29,宁夏回族自治区=30,新疆维吾尔自治区=31,台湾省=32,香港特别行政区=33,澳门特别行政区=34,海外=35.注:请注意:从2014.4.23之后,省市将采用地区标准码,请通过物流API taobao.areas.get接口获取,参考:http://api.taobao.com/apidoc/api.htm?spm=0.0.0.0.nOOF9g&amp;path=cid:7-apiId:59.API对于老的省市代码兼容会逐步下线.     */
    Province  *int64 `json:"province,omitempty" `

    /*
        会员等级，0未非会员，其余对应等级名称取grade_name     */
    Grade  *int64 `json:"grade,omitempty" `

    /*
        交易成功的金额     */
    TradeAmount  *string `json:"trade_amount,omitempty" `

    /*
        交易成功笔数     */
    TradeCount  *int64 `json:"trade_count,omitempty" `

    /*
        平均客单价.     */
    AvgPrice  *string `json:"avg_price,omitempty" `

    /*
        等级名称     */
    GradeName  *string `json:"grade_name,omitempty" `

}

func (s *TaobaoCrmMembersSearchCrmMember) SetItemNum(v int64) *TaobaoCrmMembersSearchCrmMember {
    s.ItemNum = &v
    return s
}
func (s *TaobaoCrmMembersSearchCrmMember) SetItemCloseCount(v int64) *TaobaoCrmMembersSearchCrmMember {
    s.ItemCloseCount = &v
    return s
}
func (s *TaobaoCrmMembersSearchCrmMember) SetCloseTradeAmount(v string) *TaobaoCrmMembersSearchCrmMember {
    s.CloseTradeAmount = &v
    return s
}
func (s *TaobaoCrmMembersSearchCrmMember) SetStatus(v string) *TaobaoCrmMembersSearchCrmMember {
    s.Status = &v
    return s
}
func (s *TaobaoCrmMembersSearchCrmMember) SetGroupIds(v string) *TaobaoCrmMembersSearchCrmMember {
    s.GroupIds = &v
    return s
}
func (s *TaobaoCrmMembersSearchCrmMember) SetRelationSource(v int64) *TaobaoCrmMembersSearchCrmMember {
    s.RelationSource = &v
    return s
}
func (s *TaobaoCrmMembersSearchCrmMember) SetCloseTradeCount(v int64) *TaobaoCrmMembersSearchCrmMember {
    s.CloseTradeCount = &v
    return s
}
func (s *TaobaoCrmMembersSearchCrmMember) SetLastTradeTime(v util.LocalTime) *TaobaoCrmMembersSearchCrmMember {
    s.LastTradeTime = &v
    return s
}
func (s *TaobaoCrmMembersSearchCrmMember) SetBizOrderId(v int64) *TaobaoCrmMembersSearchCrmMember {
    s.BizOrderId = &v
    return s
}
func (s *TaobaoCrmMembersSearchCrmMember) SetCity(v string) *TaobaoCrmMembersSearchCrmMember {
    s.City = &v
    return s
}
func (s *TaobaoCrmMembersSearchCrmMember) SetBuyerNick(v string) *TaobaoCrmMembersSearchCrmMember {
    s.BuyerNick = &v
    return s
}
func (s *TaobaoCrmMembersSearchCrmMember) SetOuid(v string) *TaobaoCrmMembersSearchCrmMember {
    s.Ouid = &v
    return s
}
func (s *TaobaoCrmMembersSearchCrmMember) SetProvince(v int64) *TaobaoCrmMembersSearchCrmMember {
    s.Province = &v
    return s
}
func (s *TaobaoCrmMembersSearchCrmMember) SetGrade(v int64) *TaobaoCrmMembersSearchCrmMember {
    s.Grade = &v
    return s
}
func (s *TaobaoCrmMembersSearchCrmMember) SetTradeAmount(v string) *TaobaoCrmMembersSearchCrmMember {
    s.TradeAmount = &v
    return s
}
func (s *TaobaoCrmMembersSearchCrmMember) SetTradeCount(v int64) *TaobaoCrmMembersSearchCrmMember {
    s.TradeCount = &v
    return s
}
func (s *TaobaoCrmMembersSearchCrmMember) SetAvgPrice(v string) *TaobaoCrmMembersSearchCrmMember {
    s.AvgPrice = &v
    return s
}
func (s *TaobaoCrmMembersSearchCrmMember) SetGradeName(v string) *TaobaoCrmMembersSearchCrmMember {
    s.GradeName = &v
    return s
}
