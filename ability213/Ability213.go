package ability213

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability213/request"
    "topsdk/ability213/response"
    "topsdk/util"
)

type Ability213 struct {
    Client *topsdk.TopClient
}

func NewAbility213(client *topsdk.TopClient) *Ability213{
    return &Ability213{client}
}

/*
    组合商品获取接口
*/
func (ability *Ability213) TmallItemCombineGet(req *request.TmallItemCombineGetRequest,session string) (*response.TmallItemCombineGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability213 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("tmall.item.combine.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TmallItemCombineGetResponse{}
    if(err != nil){
        log.Println("tmallItemCombineGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取产品发布规则接口
*/
func (ability *Ability213) AlibabaGpuAddSchemaGet(req *request.AlibabaGpuAddSchemaGetRequest,session string) (*response.AlibabaGpuAddSchemaGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability213 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.gpu.add.schema.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaGpuAddSchemaGetResponse{}
    if(err != nil){
        log.Println("alibabaGpuAddSchemaGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    使用schema文件发布产品
*/
func (ability *Ability213) AlibabaGpuSchemaAdd(req *request.AlibabaGpuSchemaAddRequest,session string) (*response.AlibabaGpuSchemaAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability213 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.gpu.schema.add",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaGpuSchemaAddResponse{}
    if(err != nil){
        log.Println("alibabaGpuSchemaAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取产品编辑schema规则的接口
*/
func (ability *Ability213) AlibabaGpuUpdateSchemaGet(req *request.AlibabaGpuUpdateSchemaGetRequest,session string) (*response.AlibabaGpuUpdateSchemaGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability213 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.gpu.update.schema.get",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaGpuUpdateSchemaGetResponse{}
    if(err != nil){
        log.Println("alibabaGpuUpdateSchemaGet error",err)
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
func (ability *Ability213) AlibabaGpuSchemaUpdate(req *request.AlibabaGpuSchemaUpdateRequest,session string) (*response.AlibabaGpuSchemaUpdateResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability213 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.gpu.schema.update",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaGpuSchemaUpdateResponse{}
    if(err != nil){
        log.Println("alibabaGpuSchemaUpdate error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    按类目查询spu接口
*/
func (ability *Ability213) AlibabaGpuSchemaCatsearch(req *request.AlibabaGpuSchemaCatsearchRequest,session string) (*response.AlibabaGpuSchemaCatsearchResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability213 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.gpu.schema.catsearch",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaGpuSchemaCatsearchResponse{}
    if(err != nil){
        log.Println("alibabaGpuSchemaCatsearch error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
