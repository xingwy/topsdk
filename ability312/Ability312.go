package ability312

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability312/request"
    "topsdk/ability312/response"
    "topsdk/util"
)

type Ability312 struct {
    Client *topsdk.TopClient
}

func NewAbility312(client *topsdk.TopClient) *Ability312{
    return &Ability312{client}
}

/*
    根据号码tts单呼
*/
func (ability *Ability312) AlibabaAliqinTaNumberSinglecallbyvoice(req *request.AlibabaAliqinTaNumberSinglecallbyvoiceRequest) (*response.AlibabaAliqinTaNumberSinglecallbyvoiceResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability312 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.aliqin.ta.number.singlecallbyvoice",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAliqinTaNumberSinglecallbyvoiceResponse{}
    if(err != nil){
        log.Println("alibabaAliqinTaNumberSinglecallbyvoice error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    根据号码tts单呼
*/
func (ability *Ability312) AlibabaAliqinTaNumberSinglecallbytts(req *request.AlibabaAliqinTaNumberSinglecallbyttsRequest) (*response.AlibabaAliqinTaNumberSinglecallbyttsResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability312 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.aliqin.ta.number.singlecallbytts",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAliqinTaNumberSinglecallbyttsResponse{}
    if(err != nil){
        log.Println("alibabaAliqinTaNumberSinglecallbytts error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    聚石塔语音双呼接口
*/
func (ability *Ability312) AlibabaAliqinTaVoiceNumDoublecall(req *request.AlibabaAliqinTaVoiceNumDoublecallRequest) (*response.AlibabaAliqinTaVoiceNumDoublecallResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability312 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.aliqin.ta.voice.num.doublecall",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAliqinTaVoiceNumDoublecallResponse{}
    if(err != nil){
        log.Println("alibabaAliqinTaVoiceNumDoublecall error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    短信发送
*/
func (ability *Ability312) AlibabaAliqinTaSmsNumSend(req *request.AlibabaAliqinTaSmsNumSendRequest) (*response.AlibabaAliqinTaSmsNumSendResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability312 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.aliqin.ta.sms.num.send",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAliqinTaSmsNumSendResponse{}
    if(err != nil){
        log.Println("alibabaAliqinTaSmsNumSend error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    短信查询
*/
func (ability *Ability312) AlibabaAliqinTaSmsNumQuery(req *request.AlibabaAliqinTaSmsNumQueryRequest) (*response.AlibabaAliqinTaSmsNumQueryResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability312 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("alibaba.aliqin.ta.sms.num.query",req.ToMap(),req.ToFileMap())
    var respStruct = response.AlibabaAliqinTaSmsNumQueryResponse{}
    if(err != nil){
        log.Println("alibabaAliqinTaSmsNumQuery error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    服务商存量短信模板上传
*/
func (ability *Ability312) TaobaoJstSmsTemplateReport(req *request.TaobaoJstSmsTemplateReportRequest,session string) (*response.TaobaoJstSmsTemplateReportResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability312 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.jst.sms.template.report",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoJstSmsTemplateReportResponse{}
    if(err != nil){
        log.Println("taobaoJstSmsTemplateReport error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    服务商存量短信签名上传
*/
func (ability *Ability312) TaobaoJstSmsSignnameReport(req *request.TaobaoJstSmsSignnameReportRequest,session string) (*response.TaobaoJstSmsSignnameReportResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability312 topClient is nil")
    }
    var jsonStr,err = ability.Client.ExecuteWithSession("taobao.jst.sms.signname.report",req.ToMap(),req.ToFileMap(),session)
    var respStruct = response.TaobaoJstSmsSignnameReportResponse{}
    if(err != nil){
        log.Println("taobaoJstSmsSignnameReport error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    OAID批量发送，支持明文手机号发送
*/
func (ability *Ability312) TaobaoJstSmsMessageDirectBatchsend(req *request.TaobaoJstSmsMessageDirectBatchsendRequest) (*response.TaobaoJstSmsMessageDirectBatchsendResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability312 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.jst.sms.message.direct.batchsend",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoJstSmsMessageDirectBatchsendResponse{}
    if(err != nil){
        log.Println("taobaoJstSmsMessageDirectBatchsend error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
