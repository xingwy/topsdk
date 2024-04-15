package ability198

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability198/request"
    "topsdk/ability198/response"
    "topsdk/util"
)

type Ability198 struct {
    Client *topsdk.TopClient
}

func NewAbility198(client *topsdk.TopClient) *Ability198{
    return &Ability198{client}
}

/*
    SN查询接口
*/
func (ability *Ability198) AlibabaAscpLogisticsIdentcodeQuery(req *request.AlibabaAscpLogisticsIdentcodeQueryRequest,session string) (*response.AlibabaAscpLogisticsIdentcodeQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability198 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.ascp.logistics.identcode.query",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaAscpLogisticsIdentcodeQueryResponse{}
    if(err != nil){
        log.Println("alibabaAscpLogisticsIdentcodeQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    sn码上传接口
*/
func (ability *Ability198) AlibabaAscpLogisticsIdentcodeUpload(req *request.AlibabaAscpLogisticsIdentcodeUploadRequest,session string) (*response.AlibabaAscpLogisticsIdentcodeUploadResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability198 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.ascp.logistics.identcode.upload",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaAscpLogisticsIdentcodeUploadResponse{}
    if(err != nil){
        log.Println("alibabaAscpLogisticsIdentcodeUpload error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    快递公司资源列表查询接口
*/
func (ability *Ability198) AlibabaAscpLogisticsCpGet(req *request.AlibabaAscpLogisticsCpGetRequest) (*response.AlibabaAscpLogisticsCpGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability198 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.ascp.logistics.cp.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAscpLogisticsCpGetResponse{}
    if(err != nil){
        log.Println("alibabaAscpLogisticsCpGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    商家配送写入物流节点
*/
func (ability *Ability198) AlibabaAscpLogisticsSellerWritelogisticsnode(req *request.AlibabaAscpLogisticsSellerWritelogisticsnodeRequest,session string) (*response.AlibabaAscpLogisticsSellerWritelogisticsnodeResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability198 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("alibaba.ascp.logistics.seller.writelogisticsnode",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.AlibabaAscpLogisticsSellerWritelogisticsnodeResponse{}
    if(err != nil){
        log.Println("alibabaAscpLogisticsSellerWritelogisticsnode error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
