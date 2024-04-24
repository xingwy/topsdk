package defaultability

import (
	"errors"

	"github.com/xingwy/topsdk"
	"github.com/xingwy/topsdk/defaultability/request"
	"github.com/xingwy/topsdk/defaultability/response"
	"github.com/xingwy/topsdk/util"
)

type Defaultability struct {
	Client *topsdk.TopClient
}

func NewDefaultability(client *topsdk.TopClient) *Defaultability {
	return &Defaultability{client}
}

/*
查询子订单对应的评价、追评以及语义标签
*/
func (ability *Defaultability) TmallTraderateFeedsGet(req *request.TmallTraderateFeedsGetRequest, session string) (*response.TmallTraderateFeedsGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("tmall.traderate.feeds.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TmallTraderateFeedsGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
通过商品ID获取标签列表
*/
func (ability *Defaultability) TmallTraderateItemtagsGet(req *request.TmallTraderateItemtagsGetRequest, session string) (*response.TmallTraderateItemtagsGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("tmall.traderate.itemtags.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TmallTraderateItemtagsGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询指定账户的子账号列表
*/
func (ability *Defaultability) TaobaoSellercenterSubusersGet(req *request.TaobaoSellercenterSubusersGetRequest, session string) (*response.TaobaoSellercenterSubusersGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.sellercenter.subusers.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoSellercenterSubusersGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取指定用户的权限集合
*/
func (ability *Defaultability) TaobaoSellercenterUserPermissionsGet(req *request.TaobaoSellercenterUserPermissionsGetRequest, session string) (*response.TaobaoSellercenterUserPermissionsGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.sellercenter.user.permissions.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoSellercenterUserPermissionsGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取当前会话用户出售中的商品列表
*/
func (ability *Defaultability) TaobaoItemsOnsaleGet(req *request.TaobaoItemsOnsaleGetRequest, session string) (*response.TaobaoItemsOnsaleGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.items.onsale.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoItemsOnsaleGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取商家所在分组及其已授权(广播)消息topics
*/
func (ability *Defaultability) TaobaoTmcUserGet(req *request.TaobaoTmcUserGetRequest) (*response.TaobaoTmcUserGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tmc.user.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTmcUserGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取SKU
*/
func (ability *Defaultability) TaobaoItemSkuGet(req *request.TaobaoItemSkuGetRequest, session string) (*response.TaobaoItemSkuGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.item.sku.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoItemSkuGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
根据商品ID列表获取SKU信息
*/
func (ability *Defaultability) TaobaoItemSkusGet(req *request.TaobaoItemSkusGetRequest, session string) (*response.TaobaoItemSkusGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.item.skus.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoItemSkusGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取单个商品详细信息
*/
func (ability *Defaultability) TaobaoItemSellerGet(req *request.TaobaoItemSellerGetRequest, session string) (*response.TaobaoItemSellerGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.item.seller.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoItemSellerGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
批量获取商品详细信息
*/
func (ability *Defaultability) TaobaoItemsSellerListGet(req *request.TaobaoItemsSellerListGetRequest, session string) (*response.TaobaoItemsSellerListGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.items.seller.list.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoItemsSellerListGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询买家申请的退款列表
*/
func (ability *Defaultability) TaobaoRefundsApplyGet(req *request.TaobaoRefundsApplyGetRequest, session string) (*response.TaobaoRefundsApplyGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.refunds.apply.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoRefundsApplyGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取单笔退款详情
*/
func (ability *Defaultability) TaobaoRefundGet(req *request.TaobaoRefundGetRequest, session string) (*response.TaobaoRefundGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.refund.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoRefundGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询指定的子账号的权限和角色信息
*/
func (ability *Defaultability) TaobaoSellercenterSubuserPermissionsRolesGet(req *request.TaobaoSellercenterSubuserPermissionsRolesGetRequest, session string) (*response.TaobaoSellercenterSubuserPermissionsRolesGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.sellercenter.subuser.permissions.roles.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoSellercenterSubuserPermissionsRolesGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取指定卖家的角色列表
*/
func (ability *Defaultability) TaobaoSellercenterRolesGet(req *request.TaobaoSellercenterRolesGetRequest, session string) (*response.TaobaoSellercenterRolesGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.sellercenter.roles.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoSellercenterRolesGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
子账号角色的新增（指定卖家）
*/
func (ability *Defaultability) TaobaoSellercenterRoleAdd(req *request.TaobaoSellercenterRoleAddRequest, session string) (*response.TaobaoSellercenterRoleAddResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.sellercenter.role.add", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoSellercenterRoleAddResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
卖家查询等级规则
*/
func (ability *Defaultability) TaobaoCrmGradeGet(req *request.TaobaoCrmGradeGetRequest, session string) (*response.TaobaoCrmGradeGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.crm.grade.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoCrmGradeGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
增量获取卖家会员
*/
func (ability *Defaultability) TaobaoCrmMembersIncrementGet(req *request.TaobaoCrmMembersIncrementGetRequest, session string) (*response.TaobaoCrmMembersIncrementGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.crm.members.increment.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoCrmMembersIncrementGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
卖家设置等级规则
*/
func (ability *Defaultability) TaobaoCrmGradeSet(req *request.TaobaoCrmGradeSetRequest, session string) (*response.TaobaoCrmGradeSetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.crm.grade.set", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoCrmGradeSetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询卖家的分组
*/
func (ability *Defaultability) TaobaoCrmGroupsGet(req *request.TaobaoCrmGroupsGetRequest, session string) (*response.TaobaoCrmGroupsGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.crm.groups.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoCrmGroupsGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
修改一个已经存在的分组
*/
func (ability *Defaultability) TaobaoCrmGroupUpdate(req *request.TaobaoCrmGroupUpdateRequest, session string) (*response.TaobaoCrmGroupUpdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.crm.group.update", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoCrmGroupUpdateResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
卖家创建一个分组
*/
func (ability *Defaultability) TaobaoCrmGroupAdd(req *request.TaobaoCrmGroupAddRequest, session string) (*response.TaobaoCrmGroupAddResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.crm.group.add", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoCrmGroupAddResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取卖家会员（高级查询）
*/
func (ability *Defaultability) TaobaoCrmMembersSearch(req *request.TaobaoCrmMembersSearchRequest, session string) (*response.TaobaoCrmMembersSearchResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.crm.members.search", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoCrmMembersSearchResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
编辑会员资料
*/
func (ability *Defaultability) TaobaoCrmMemberinfoUpdate(req *request.TaobaoCrmMemberinfoUpdateRequest, session string) (*response.TaobaoCrmMemberinfoUpdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.crm.memberinfo.update", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoCrmMemberinfoUpdateResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取卖家的会员（基本查询）
*/
func (ability *Defaultability) TaobaoCrmMembersGet(req *request.TaobaoCrmMembersGetRequest, session string) (*response.TaobaoCrmMembersGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.crm.members.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoCrmMembersGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
删除分组
*/
func (ability *Defaultability) TaobaoCrmGroupDelete(req *request.TaobaoCrmGroupDeleteRequest, session string) (*response.TaobaoCrmGroupDeleteResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.crm.group.delete", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoCrmGroupDeleteResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询分组任务是否完成
*/
func (ability *Defaultability) TaobaoCrmGrouptaskCheck(req *request.TaobaoCrmGrouptaskCheckRequest, session string) (*response.TaobaoCrmGrouptaskCheckResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.crm.grouptask.check", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoCrmGrouptaskCheckResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
服务平台评价查询接口
*/
func (ability *Defaultability) TaobaoFuwuScoresGet(req *request.TaobaoFuwuScoresGetRequest) (*response.TaobaoFuwuScoresGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.fuwu.scores.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoFuwuScoresGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
创建码平台常用二维码
*/
func (ability *Defaultability) TaobaoMaQrcodeCommonCreate(req *request.TaobaoMaQrcodeCommonCreateRequest, session string) (*response.TaobaoMaQrcodeCommonCreateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.ma.qrcode.common.create", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoMaQrcodeCommonCreateResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取后台供卖家发布商品的标准商品类目
*/
func (ability *Defaultability) TaobaoItemcatsGet(req *request.TaobaoItemcatsGetRequest) (*response.TaobaoItemcatsGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.itemcats.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoItemcatsGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询退款留言/凭证列表
*/
func (ability *Defaultability) TaobaoRefundMessagesGet(req *request.TaobaoRefundMessagesGetRequest, session string) (*response.TaobaoRefundMessagesGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.refund.messages.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoRefundMessagesGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
创建退款留言/凭证
*/
func (ability *Defaultability) TaobaoRefundMessageAdd(req *request.TaobaoRefundMessageAddRequest, session string) (*response.TaobaoRefundMessageAddResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.refund.message.add", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoRefundMessageAddResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
服务平台营销链接生成接口
*/
func (ability *Defaultability) TaobaoFuwuSaleLinkGen(req *request.TaobaoFuwuSaleLinkGenRequest) (*response.TaobaoFuwuSaleLinkGenResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.fuwu.sale.link.gen", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoFuwuSaleLinkGenResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
关键词过滤匹配
*/
func (ability *Defaultability) TaobaoKfcKeywordSearch(req *request.TaobaoKfcKeywordSearchRequest, session string) (*response.TaobaoKfcKeywordSearchResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.kfc.keyword.search", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoKfcKeywordSearchResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
自己联系物流发货
*/
func (ability *Defaultability) AlibabaAscpLogisticsOfflineSend(req *request.AlibabaAscpLogisticsOfflineSendRequest, session string) (*response.AlibabaAscpLogisticsOfflineSendResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("alibaba.ascp.logistics.offline.send", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.AlibabaAscpLogisticsOfflineSendResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
修改物流公司和运单号
*/
func (ability *Defaultability) AlibabaAscpLogisticsConsignResend(req *request.AlibabaAscpLogisticsConsignResendRequest, session string) (*response.AlibabaAscpLogisticsConsignResendResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("alibaba.ascp.logistics.consign.resend", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.AlibabaAscpLogisticsConsignResendResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
修改物流公司和运单号
*/
func (ability *Defaultability) AlibabaAscpLogisticsConsignModify(req *request.AlibabaAscpLogisticsConsignModifyRequest, session string) (*response.AlibabaAscpLogisticsConsignModifyResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("alibaba.ascp.logistics.consign.modify", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.AlibabaAscpLogisticsConsignModifyResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询商家被授权品牌列表和类目列表
*/
func (ability *Defaultability) TaobaoItemcatsAuthorizeGet(req *request.TaobaoItemcatsAuthorizeGetRequest, session string) (*response.TaobaoItemcatsAuthorizeGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.itemcats.authorize.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoItemcatsAuthorizeGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
得到当前会话用户库存中的商品列表
*/
func (ability *Defaultability) TaobaoItemsInventoryGet(req *request.TaobaoItemsInventoryGetRequest, session string) (*response.TaobaoItemsInventoryGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.items.inventory.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoItemsInventoryGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
根据外部ID取商品SKU
*/
func (ability *Defaultability) TaobaoSkusCustomGet(req *request.TaobaoSkusCustomGetRequest, session string) (*response.TaobaoSkusCustomGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.skus.custom.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoSkusCustomGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取指定账户的子账号简易信息列表
*/
func (ability *Defaultability) TaobaoSubusersGet(req *request.TaobaoSubusersGetRequest, session string) (*response.TaobaoSubusersGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.subusers.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoSubusersGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取指定账户子账号的详细信息
*/
func (ability *Defaultability) TaobaoSubuserFullinfoGet(req *request.TaobaoSubuserFullinfoGetRequest, session string) (*response.TaobaoSubuserFullinfoGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.subuser.fullinfo.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoSubuserFullinfoGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取指定账户的所有部门列表
*/
func (ability *Defaultability) TaobaoSubuserDepartmentsGet(req *request.TaobaoSubuserDepartmentsGetRequest, session string) (*response.TaobaoSubuserDepartmentsGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.subuser.departments.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoSubuserDepartmentsGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取指定账户的所有职务信息列表
*/
func (ability *Defaultability) TaobaoSubuserDutysGet(req *request.TaobaoSubuserDutysGetRequest, session string) (*response.TaobaoSubuserDutysGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.subuser.dutys.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoSubuserDutysGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
修改指定账户子账号的基本信息
*/
func (ability *Defaultability) TaobaoSubuserInfoUpdate(req *request.TaobaoSubuserInfoUpdateRequest, session string) (*response.TaobaoSubuserInfoUpdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.subuser.info.update", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoSubuserInfoUpdateResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
商品编辑获取schema信息
*/
func (ability *Defaultability) AlibabaItemEditSchemaGet(req *request.AlibabaItemEditSchemaGetRequest, session string) (*response.AlibabaItemEditSchemaGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("alibaba.item.edit.schema.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.AlibabaItemEditSchemaGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
商品编辑提交schema信息
*/
func (ability *Defaultability) AlibabaItemEditSubmit(req *request.AlibabaItemEditSubmitRequest, session string) (*response.AlibabaItemEditSubmitResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("alibaba.item.edit.submit", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.AlibabaItemEditSubmitResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
产品发布规则获取接口
*/
func (ability *Defaultability) TmallProductAddSchemaGet(req *request.TmallProductAddSchemaGetRequest, session string) (*response.TmallProductAddSchemaGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("tmall.product.add.schema.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TmallProductAddSchemaGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取匹配产品规则
*/
func (ability *Defaultability) TmallProductMatchSchemaGet(req *request.TmallProductMatchSchemaGetRequest, session string) (*response.TmallProductMatchSchemaGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("tmall.product.match.schema.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TmallProductMatchSchemaGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
product匹配接口
*/
func (ability *Defaultability) TmallProductSchemaMatch(req *request.TmallProductSchemaMatchRequest, session string) (*response.TmallProductSchemaMatchResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("tmall.product.schema.match", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TmallProductSchemaMatchResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
使用Schema文件发布一个产品
*/
func (ability *Defaultability) TmallProductSchemaAdd(req *request.TmallProductSchemaAddRequest, session string) (*response.TmallProductSchemaAddResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("tmall.product.schema.add", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TmallProductSchemaAddResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
根据买家nick获取买家openuid
*/
func (ability *Defaultability) TaobaoUserOpenuidGetbynick(req *request.TaobaoUserOpenuidGetbynickRequest) (*response.TaobaoUserOpenuidGetbynickResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.user.openuid.getbynick", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoUserOpenuidGetbynickResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
通过主账号登陆态分页查询子账号列表
*/
func (ability *Defaultability) TaobaoSellercenterSubusersPage(req *request.TaobaoSellercenterSubusersPageRequest, session string) (*response.TaobaoSellercenterSubusersPageResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.sellercenter.subusers.page", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoSellercenterSubusersPageResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
创建物流订单
*/
func (ability *Defaultability) TaobaoLogisticsOrderCreate(req *request.TaobaoLogisticsOrderCreateRequest, session string) (*response.TaobaoLogisticsOrderCreateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.logistics.order.create", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoLogisticsOrderCreateResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
分页获取指定账户的子账号简易信息列表（新isv建议使用taobao.sellercenter.subusers.page接口）
*/
func (ability *Defaultability) TaobaoSubusersPage(req *request.TaobaoSubusersPageRequest, session string) (*response.TaobaoSubusersPageResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.subusers.page", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoSubusersPageResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
商品编辑增量更新
*/
func (ability *Defaultability) AlibabaItemEditFastupdate(req *request.AlibabaItemEditFastupdateRequest, session string) (*response.AlibabaItemEditFastupdateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("alibaba.item.edit.fastupdate", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.AlibabaItemEditFastupdateResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
根据buyerNick获取ouid
*/
func (ability *Defaultability) TaobaoCrmHistoryOuidGet(req *request.TaobaoCrmHistoryOuidGetRequest, session string) (*response.TaobaoCrmHistoryOuidGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.crm.history.ouid.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoCrmHistoryOuidGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
根据buyerNick获取omid
*/
func (ability *Defaultability) TaobaoCrmHistoryOmidGet(req *request.TaobaoCrmHistoryOmidGetRequest, session string) (*response.TaobaoCrmHistoryOmidGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.crm.history.omid.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoCrmHistoryOmidGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询appstore应用订购关系
*/
func (ability *Defaultability) TaobaoAppstoreSubscribeGet(req *request.TaobaoAppstoreSubscribeGetRequest) (*response.TaobaoAppstoreSubscribeGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.appstore.subscribe.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoAppstoreSubscribeGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取交易订单的简易信息
*/
func (ability *Defaultability) TaobaoTradeSimpleGet(req *request.TaobaoTradeSimpleGetRequest, session string) (*response.TaobaoTradeSimpleGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.trade.simple.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTradeSimpleGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取单条订单跟踪详情
*/
func (ability *Defaultability) TaobaoJdsTradeTracesGet(req *request.TaobaoJdsTradeTracesGetRequest, session string) (*response.TaobaoJdsTradeTracesGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.jds.trade.traces.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoJdsTradeTracesGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询卖家已卖出的交易简易数据
*/
func (ability *Defaultability) TaobaoTradesSimpleSoldGet(req *request.TaobaoTradesSimpleSoldGetRequest, session string) (*response.TaobaoTradesSimpleSoldGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.trades.simple.sold.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTradesSimpleSoldGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
短信签名增加删除接口
*/
func (ability *Defaultability) TaobaoOpencrmSignatureAdddelete(req *request.TaobaoOpencrmSignatureAdddeleteRequest, session string) (*response.TaobaoOpencrmSignatureAdddeleteResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.opencrm.signature.adddelete", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoOpencrmSignatureAdddeleteResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
短信签名查询接口
*/
func (ability *Defaultability) TaobaoOpencrmSignatureQuery(req *request.TaobaoOpencrmSignatureQueryRequest, session string) (*response.TaobaoOpencrmSignatureQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.opencrm.signature.query", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoOpencrmSignatureQueryResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
物流详情查询
*/
func (ability *Defaultability) TaobaoLogisticsInstantTraceSearch(req *request.TaobaoLogisticsInstantTraceSearchRequest, session string) (*response.TaobaoLogisticsInstantTraceSearchResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.logistics.instant.trace.search", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoLogisticsInstantTraceSearchResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
节点报告获取接口
*/
func (ability *Defaultability) TaobaoOpencrmNodereportGet(req *request.TaobaoOpencrmNodereportGetRequest, session string) (*response.TaobaoOpencrmNodereportGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.opencrm.nodereport.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoOpencrmNodereportGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
查询卖家已卖出的增量交易简易数据（根据修改时间）
*/
func (ability *Defaultability) TaobaoTradesSimpleSoldIncrementGet(req *request.TaobaoTradesSimpleSoldIncrementGetRequest, session string) (*response.TaobaoTradesSimpleSoldIncrementGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.trades.simple.sold.increment.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTradesSimpleSoldIncrementGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
商家配送核销
*/
func (ability *Defaultability) AlibabaAscpLogisticsSellerWriteoff(req *request.AlibabaAscpLogisticsSellerWriteoffRequest, session string) (*response.AlibabaAscpLogisticsSellerWriteoffResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("alibaba.ascp.logistics.seller.writeoff", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.AlibabaAscpLogisticsSellerWriteoffResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
商家配送发货
*/
func (ability *Defaultability) AlibabaAscpLogisticsSellerSend(req *request.AlibabaAscpLogisticsSellerSendRequest, session string) (*response.AlibabaAscpLogisticsSellerSendResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("alibaba.ascp.logistics.seller.send", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.AlibabaAscpLogisticsSellerSendResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
商家配送核销订单列表查询
*/
func (ability *Defaultability) AlibabaAscpLogisticsSellerOrdersGet(req *request.AlibabaAscpLogisticsSellerOrdersGetRequest, session string) (*response.AlibabaAscpLogisticsSellerOrdersGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("alibaba.ascp.logistics.seller.orders.get", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.AlibabaAscpLogisticsSellerOrdersGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
数字短信模板创建
*/
func (ability *Defaultability) AlibabaAliqinFcDigitalsmsCreatetemplate(req *request.AlibabaAliqinFcDigitalsmsCreatetemplateRequest, session string) (*response.AlibabaAliqinFcDigitalsmsCreatetemplateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("alibaba.aliqin.fc.digitalsms.createtemplate", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.AlibabaAliqinFcDigitalsmsCreatetemplateResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
订购关系查询
*/
func (ability *Defaultability) TaobaoVasSubscribeGet(req *request.TaobaoVasSubscribeGetRequest) (*response.TaobaoVasSubscribeGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.vas.subscribe.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoVasSubscribeGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
订单记录导出
*/
func (ability *Defaultability) TaobaoVasOrderSearch(req *request.TaobaoVasOrderSearchRequest) (*response.TaobaoVasOrderSearchResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.vas.order.search", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoVasOrderSearchResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
订购记录导出
*/
func (ability *Defaultability) TaobaoVasSubscSearch(req *request.TaobaoVasSubscSearchRequest) (*response.TaobaoVasSubscSearchResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.vas.subsc.search", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoVasSubscSearchResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
淘宝短信签名创建接口
*/
func (ability *Defaultability) TaobaoJstSmsSignnameCreate(req *request.TaobaoJstSmsSignnameCreateRequest, session string) (*response.TaobaoJstSmsSignnameCreateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.jst.sms.signname.create", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoJstSmsSignnameCreateResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
淘宝短信签名查询
*/
func (ability *Defaultability) TaobaoJstSmsSignnameQuery(req *request.TaobaoJstSmsSignnameQueryRequest, session string) (*response.TaobaoJstSmsSignnameQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.jst.sms.signname.query", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoJstSmsSignnameQueryResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
淘宝短信签名删除
*/
func (ability *Defaultability) TaobaoJstSmsSignnameDelete(req *request.TaobaoJstSmsSignnameDeleteRequest, session string) (*response.TaobaoJstSmsSignnameDeleteResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.jst.sms.signname.delete", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoJstSmsSignnameDeleteResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
淘宝短信签名修改
*/
func (ability *Defaultability) TaobaoJstSmsSignnameModify(req *request.TaobaoJstSmsSignnameModifyRequest, session string) (*response.TaobaoJstSmsSignnameModifyResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.jst.sms.signname.modify", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoJstSmsSignnameModifyResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
淘宝短信模板创建
*/
func (ability *Defaultability) TaobaoJstSmsTemplateCreate(req *request.TaobaoJstSmsTemplateCreateRequest) (*response.TaobaoJstSmsTemplateCreateResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.jst.sms.template.create", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoJstSmsTemplateCreateResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
淘宝短信模板查询
*/
func (ability *Defaultability) TaobaoJstSmsTemplateQuery(req *request.TaobaoJstSmsTemplateQueryRequest) (*response.TaobaoJstSmsTemplateQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.jst.sms.template.query", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoJstSmsTemplateQueryResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
淘宝短信模板删除
*/
func (ability *Defaultability) TaobaoJstSmsTemplateDelete(req *request.TaobaoJstSmsTemplateDeleteRequest) (*response.TaobaoJstSmsTemplateDeleteResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.jst.sms.template.delete", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoJstSmsTemplateDeleteResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
淘宝短信模板修改
*/
func (ability *Defaultability) TaobaoJstSmsTemplateModify(req *request.TaobaoJstSmsTemplateModifyRequest) (*response.TaobaoJstSmsTemplateModifyResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.jst.sms.template.modify", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoJstSmsTemplateModifyResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
BC融合版本节点报告获取
*/
func (ability *Defaultability) TaobaoOpencrmNodereportGetbc(req *request.TaobaoOpencrmNodereportGetbcRequest, session string) (*response.TaobaoOpencrmNodereportGetbcResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.opencrm.nodereport.getbc", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoOpencrmNodereportGetbcResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
评价大家印象印象短语接口
*/
func (ability *Defaultability) TaobaoTraderateImprImprwordsGet(req *request.TaobaoTraderateImprImprwordsGetRequest) (*response.TaobaoTraderateImprImprwordsGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.traderate.impr.imprwords.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTraderateImprImprwordsGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
根据当前子账号登陆态，获取该子账号基本信息
*/
func (ability *Defaultability) TaobaoSubusersInfoQuery(req *request.TaobaoSubusersInfoQueryRequest, session string) (*response.TaobaoSubusersInfoQueryResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.subusers.info.query", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoSubusersInfoQueryResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
发布单条消息
*/
func (ability *Defaultability) TaobaoTmcMessageProduce(req *request.TaobaoTmcMessageProduceRequest) (*response.TaobaoTmcMessageProduceResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tmc.message.produce", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTmcMessageProduceResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
取消用户的消息服务
*/
func (ability *Defaultability) TaobaoTmcUserCancel(req *request.TaobaoTmcUserCancelRequest) (*response.TaobaoTmcUserCancelResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.tmc.user.cancel", req.ToMap(), req.ToFileMap())
	var respStruct = response.TaobaoTmcUserCancelResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
为已授权的用户开通消息服务
*/
func (ability *Defaultability) TaobaoTmcUserPermit(req *request.TaobaoTmcUserPermitRequest, session string) (*response.TaobaoTmcUserPermitResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.ExecuteWithSession("taobao.tmc.user.permit", req.ToMap(), req.ToFileMap(), session)
	var respStruct = response.TaobaoTmcUserPermitResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}

/*
获取淘宝系统事件
*/
func (ability *Defaultability) TimeGet(req *request.TimeGetRequest) (*response.TimeGetResponse, error) {
	if ability.Client == nil {
		return nil, errors.New("Defaultability topClient is nil")
	}
	var jsonStr, err = ability.Client.Execute("taobao.time.get", req.ToMap(), req.ToFileMap())
	var respStruct = response.TimeGetResponse{}
	if err != nil {
		return &respStruct, err
	}
	err = util.HandleJsonResponse(jsonStr, &respStruct)
	if respStruct.Body == "" || len(respStruct.Body) == 0 {
		respStruct.Body = jsonStr
	}
	return &respStruct, err
}
