package ability147

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability147/request"
    "topsdk/ability147/response"
    "topsdk/util"
)

type Ability147 struct {
    Client *topsdk.TopClient
}

func NewAbility147(client *topsdk.TopClient) *Ability147{
    return &Ability147{client}
}

/*
    产品信息获取schema获取
*/
func (ability *Ability147) TmallProductSchemaGet(req *request.TmallProductSchemaGetRequest,session string) (*response.TmallProductSchemaGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability147 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.product.schema.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallProductSchemaGetResponse{}
    if(err != nil){
        log.Println("tmallProductSchemaGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商品hscode信息审核状态查询接口
*/
func (ability *Ability147) TmallItemHscodeAuditResultsQuery(req *request.TmallItemHscodeAuditResultsQueryRequest,session string) (*response.TmallItemHscodeAuditResultsQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability147 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.item.hscode.audit.results.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallItemHscodeAuditResultsQueryResponse{}
    if(err != nil){
        log.Println("tmallItemHscodeAuditResultsQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    天猫发布商品规则获取
*/
func (ability *Ability147) TmallItemAddSimpleschemaGet(req *request.TmallItemAddSimpleschemaGetRequest,session string) (*response.TmallItemAddSimpleschemaGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability147 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.item.add.simpleschema.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallItemAddSimpleschemaGetResponse{}
    if(err != nil){
        log.Println("tmallItemAddSimpleschemaGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    产品更新规则获取接口
*/
func (ability *Ability147) TmallProductUpdateSchemaGet(req *request.TmallProductUpdateSchemaGetRequest,session string) (*response.TmallProductUpdateSchemaGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability147 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.product.update.schema.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallProductUpdateSchemaGetResponse{}
    if(err != nil){
        log.Println("tmallProductUpdateSchemaGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    产品更新接口
*/
func (ability *Ability147) TmallProductSchemaUpdate(req *request.TmallProductSchemaUpdateRequest,session string) (*response.TmallProductSchemaUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability147 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.product.schema.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallProductSchemaUpdateResponse{}
    if(err != nil){
        log.Println("tmallProductSchemaUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    算法获取hscode
*/
func (ability *Ability147) TmallItemCalculateHscodeGet(req *request.TmallItemCalculateHscodeGetRequest,session string) (*response.TmallItemCalculateHscodeGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability147 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.item.calculate.hscode.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallItemCalculateHscodeGetResponse{}
    if(err != nil){
        log.Println("tmallItemCalculateHscodeGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    通过hscode获取计量单位
*/
func (ability *Ability147) TmallItemHscodeDetailGet(req *request.TmallItemHscodeDetailGetRequest,session string) (*response.TmallItemHscodeDetailGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability147 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.item.hscode.detail.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallItemHscodeDetailGetResponse{}
    if(err != nil){
        log.Println("tmallItemHscodeDetailGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
