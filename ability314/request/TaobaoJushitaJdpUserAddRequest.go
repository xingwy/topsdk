package request

import (
        "topsdk/util"
    )

type TaobaoJushitaJdpUserAddRequest struct {
    /*
        已废弃，使用页面中应用的配置。用户同步的数据类型,多个用','号分割。可选值：trade,refund,item     */
    Topics  *[]string `json:"topics,omitempty" required:"false" `
    /*
        推送历史数据天数，只能为90天内，包含90天。当此参数不填时，表示以页面中应用配置的历史天数为准；如果为0表示这个用户不推送历史数据；其它表示推送的历史天数。     */
    HistoryDays  *int64 `json:"history_days,omitempty" required:"false" `
    /*
        已废弃，新版订单同步服务不要使用。同步用户数据的机器IP,必须是界面配置的IP。     */
    HostIp  *string `json:"host_ip,omitempty" required:"false" `
    /*
        已废弃，新版订单同步服务不要使用。绑定类型，用户数据同步的机器名称。     */
    HostName  *string `json:"host_name,omitempty" required:"false" `
    /*
        RDS实例名称     */
    RdsName  *string `json:"rds_name" required:"true" `
}

func (s *TaobaoJushitaJdpUserAddRequest) SetTopics(v []string) *TaobaoJushitaJdpUserAddRequest {
    s.Topics = &v
    return s
}
func (s *TaobaoJushitaJdpUserAddRequest) SetHistoryDays(v int64) *TaobaoJushitaJdpUserAddRequest {
    s.HistoryDays = &v
    return s
}
func (s *TaobaoJushitaJdpUserAddRequest) SetHostIp(v string) *TaobaoJushitaJdpUserAddRequest {
    s.HostIp = &v
    return s
}
func (s *TaobaoJushitaJdpUserAddRequest) SetHostName(v string) *TaobaoJushitaJdpUserAddRequest {
    s.HostName = &v
    return s
}
func (s *TaobaoJushitaJdpUserAddRequest) SetRdsName(v string) *TaobaoJushitaJdpUserAddRequest {
    s.RdsName = &v
    return s
}

func (req *TaobaoJushitaJdpUserAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Topics != nil) {
        paramMap["topics"] = util.ConvertBasicList(*req.Topics)
    }
    if(req.HistoryDays != nil) {
        paramMap["history_days"] = *req.HistoryDays
    }
    if(req.HostIp != nil) {
        paramMap["host_ip"] = *req.HostIp
    }
    if(req.HostName != nil) {
        paramMap["host_name"] = *req.HostName
    }
    if(req.RdsName != nil) {
        paramMap["rds_name"] = *req.RdsName
    }
    return paramMap
}

func (req *TaobaoJushitaJdpUserAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}