package ability680

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability680/request"
    "topsdk/ability680/response"
    "topsdk/util"
)

type Ability680 struct {
    Client *topsdk.TopClient
}

func NewAbility680(client *topsdk.TopClient) *Ability680{
    return &Ability680{client}
}

/*
    批量删除分组（隐私号版）
*/
func (ability *Ability680) TaobaoCrmMembersGroupsBatchdeletePrivy(req *request.TaobaoCrmMembersGroupsBatchdeletePrivyRequest,session string) (*response.TaobaoCrmMembersGroupsBatchdeletePrivyResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability680 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.crm.members.groups.batchdelete.privy",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoCrmMembersGroupsBatchdeletePrivyResponse{}
    if(err != nil){
        log.Println("taobaoCrmMembersGroupsBatchdeletePrivy error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    增量获取卖家会员
*/
func (ability *Ability680) TaobaoCrmMembersIncrementGetPrivy(req *request.TaobaoCrmMembersIncrementGetPrivyRequest,session string) (*response.TaobaoCrmMembersIncrementGetPrivyResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability680 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.crm.members.increment.get.privy",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoCrmMembersIncrementGetPrivyResponse{}
    if(err != nil){
        log.Println("taobaoCrmMembersIncrementGetPrivy error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取卖家会员(高级查询)
*/
func (ability *Ability680) TaobaoCrmMembersSearchPrivy(req *request.TaobaoCrmMembersSearchPrivyRequest,session string) (*response.TaobaoCrmMembersSearchPrivyResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability680 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.crm.members.search.privy",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoCrmMembersSearchPrivyResponse{}
    if(err != nil){
        log.Println("taobaoCrmMembersSearchPrivy error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取卖家会员(基本查询)
*/
func (ability *Ability680) TaobaoCrmMembersGetPrivy(req *request.TaobaoCrmMembersGetPrivyRequest,session string) (*response.TaobaoCrmMembersGetPrivyResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability680 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.crm.members.get.privy",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoCrmMembersGetPrivyResponse{}
    if(err != nil){
        log.Println("taobaoCrmMembersGetPrivy error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    编辑会员资料
*/
func (ability *Ability680) TaobaoCrmMemberinfoUpdatePrivy(req *request.TaobaoCrmMemberinfoUpdatePrivyRequest,session string) (*response.TaobaoCrmMemberinfoUpdatePrivyResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability680 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.crm.memberinfo.update.privy",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoCrmMemberinfoUpdatePrivyResponse{}
    if(err != nil){
        log.Println("taobaoCrmMemberinfoUpdatePrivy error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    一批会员添加分组(隐私号版）
*/
func (ability *Ability680) TaobaoCrmMembersGroupBatchaddPrivy(req *request.TaobaoCrmMembersGroupBatchaddPrivyRequest,session string) (*response.TaobaoCrmMembersGroupBatchaddPrivyResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability680 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.crm.members.group.batchadd.privy",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoCrmMembersGroupBatchaddPrivyResponse{}
    if(err != nil){
        log.Println("taobaoCrmMembersGroupBatchaddPrivy error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商家存量会员清理进度查询
*/
func (ability *Ability680) TaobaoOpencrmMemberSellerFetchstatus(req *request.TaobaoOpencrmMemberSellerFetchstatusRequest) (*response.TaobaoOpencrmMemberSellerFetchstatusResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability680 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.opencrm.member.seller.fetchstatus",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoOpencrmMemberSellerFetchstatusResponse{}
    if(err != nil){
        log.Println("taobaoOpencrmMemberSellerFetchstatus error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
