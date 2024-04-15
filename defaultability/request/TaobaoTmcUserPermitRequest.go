package request

import (
        "topsdk/util"
    )

type TaobaoTmcUserPermitRequest struct {
    /*
        消息主题列表，用半角逗号分隔。当用户订阅的topic是应用订阅的子集时才需要设置，不设置表示继承应用所订阅的所有topic，一般情况建议不要设置。     */
    Topics  *[]string `json:"topics,omitempty" required:"false" `
}

func (s *TaobaoTmcUserPermitRequest) SetTopics(v []string) *TaobaoTmcUserPermitRequest {
    s.Topics = &v
    return s
}

func (req *TaobaoTmcUserPermitRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Topics != nil) {
        paramMap["topics"] = util.ConvertBasicList(*req.Topics)
    }
    return paramMap
}

func (req *TaobaoTmcUserPermitRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}