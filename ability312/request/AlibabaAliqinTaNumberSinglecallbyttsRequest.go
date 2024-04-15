package request

import (
        "topsdk/util"
    )

type AlibabaAliqinTaNumberSinglecallbyttsRequest struct {
    /*
        被叫号码     */
    CalledNum  *string `json:"called_num" required:"true" `
    /*
        显示号码     */
    CalledShowNum  *string `json:"called_show_num" required:"true" `
    /*
        tts文本模板code     */
    TtsCode  *string `json:"tts_code" required:"true" `
    /*
        上下文参数,tts模板含有变量的, 此参数需填写。示例:{"date":"2015年 " ,"name":"测试","extend":"回传参数"} date、name 为模板里的变量名作为key,extend为扩展信息作为回传参数的key     */
    Params  *map[string]interface{} `json:"params,omitempty" required:"false" `
}

func (s *AlibabaAliqinTaNumberSinglecallbyttsRequest) SetCalledNum(v string) *AlibabaAliqinTaNumberSinglecallbyttsRequest {
    s.CalledNum = &v
    return s
}
func (s *AlibabaAliqinTaNumberSinglecallbyttsRequest) SetCalledShowNum(v string) *AlibabaAliqinTaNumberSinglecallbyttsRequest {
    s.CalledShowNum = &v
    return s
}
func (s *AlibabaAliqinTaNumberSinglecallbyttsRequest) SetTtsCode(v string) *AlibabaAliqinTaNumberSinglecallbyttsRequest {
    s.TtsCode = &v
    return s
}
func (s *AlibabaAliqinTaNumberSinglecallbyttsRequest) SetParams(v map[string]interface{}) *AlibabaAliqinTaNumberSinglecallbyttsRequest {
    s.Params = &v
    return s
}

func (req *AlibabaAliqinTaNumberSinglecallbyttsRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CalledNum != nil) {
        paramMap["called_num"] = *req.CalledNum
    }
    if(req.CalledShowNum != nil) {
        paramMap["called_show_num"] = *req.CalledShowNum
    }
    if(req.TtsCode != nil) {
        paramMap["tts_code"] = *req.TtsCode
    }
    if(req.Params != nil) {
        paramMap["params"] = util.ConvertStruct(*req.Params)
    }
    return paramMap
}

func (req *AlibabaAliqinTaNumberSinglecallbyttsRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}