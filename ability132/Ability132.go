package ability132

import (
    "log"
    "errors"
    "topsdk"
    "topsdk/ability132/request"
    "topsdk/ability132/response"
    "topsdk/util"
)

type Ability132 struct {
    Client *topsdk.TopClient
}

func NewAbility132(client *topsdk.TopClient) *Ability132{
    return &Ability132{client}
}

/*
    TMC授权token
*/
func (ability *Ability132) TaobaoTmcAuthGet(req *request.TaobaoTmcAuthGetRequest) (*response.TaobaoTmcAuthGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability132 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tmc.auth.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTmcAuthGetResponse{}
    if(err != nil){
        log.Println("taobaoTmcAuthGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    删除消息topic分组路由
*/
func (ability *Ability132) TaobaoTmcTopicGroupDelete(req *request.TaobaoTmcTopicGroupDeleteRequest) (*response.TaobaoTmcTopicGroupDeleteResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability132 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tmc.topic.group.delete",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTmcTopicGroupDeleteResponse{}
    if(err != nil){
        log.Println("taobaoTmcTopicGroupDelete error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    点对点消息topic分组路由
*/
func (ability *Ability132) TaobaoTmcTopicGroupAdd(req *request.TaobaoTmcTopicGroupAddRequest) (*response.TaobaoTmcTopicGroupAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability132 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tmc.topic.group.add",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTmcTopicGroupAddResponse{}
    if(err != nil){
        log.Println("taobaoTmcTopicGroupAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    批量发送消息
*/
func (ability *Ability132) TaobaoTmcMessagesProduce(req *request.TaobaoTmcMessagesProduceRequest) (*response.TaobaoTmcMessagesProduceResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability132 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tmc.messages.produce",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTmcMessagesProduceResponse{}
    if(err != nil){
        log.Println("taobaoTmcMessagesProduce error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    获取自定义用户分组列表
*/
func (ability *Ability132) TaobaoTmcGroupsGet(req *request.TaobaoTmcGroupsGetRequest) (*response.TaobaoTmcGroupsGetResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability132 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tmc.groups.get",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTmcGroupsGetResponse{}
    if(err != nil){
        log.Println("taobaoTmcGroupsGet error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    删除指定的分组或分组下的用户
*/
func (ability *Ability132) TaobaoTmcGroupDelete(req *request.TaobaoTmcGroupDeleteRequest) (*response.TaobaoTmcGroupDeleteResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability132 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tmc.group.delete",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTmcGroupDeleteResponse{}
    if(err != nil){
        log.Println("taobaoTmcGroupDelete error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    为已开通用户添加用户分组
*/
func (ability *Ability132) TaobaoTmcGroupAdd(req *request.TaobaoTmcGroupAddRequest) (*response.TaobaoTmcGroupAddResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability132 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tmc.group.add",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTmcGroupAddResponse{}
    if(err != nil){
        log.Println("taobaoTmcGroupAdd error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    确认消费消息的状态
*/
func (ability *Ability132) TaobaoTmcMessagesConfirm(req *request.TaobaoTmcMessagesConfirmRequest) (*response.TaobaoTmcMessagesConfirmResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability132 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tmc.messages.confirm",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTmcMessagesConfirmResponse{}
    if(err != nil){
        log.Println("taobaoTmcMessagesConfirm error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
/*
    消费多条消息
*/
func (ability *Ability132) TaobaoTmcMessagesConsume(req *request.TaobaoTmcMessagesConsumeRequest) (*response.TaobaoTmcMessagesConsumeResponse,error) {
    if(ability.Client == nil) {
        return nil,errors.New("Ability132 topClient is nil")
    }
    var jsonStr,err = ability.Client.Execute("taobao.tmc.messages.consume",req.ToMap(),req.ToFileMap())
    var respStruct = response.TaobaoTmcMessagesConsumeResponse{}
    if(err != nil){
        log.Println("taobaoTmcMessagesConsume error",err)
        return &respStruct,err
    }
    err = util.HandleJsonResponse(jsonStr, &respStruct)
    if(respStruct.Body == "" || len(respStruct.Body) == 0) {
        respStruct.Body = jsonStr
    }
    return &respStruct,err
}
