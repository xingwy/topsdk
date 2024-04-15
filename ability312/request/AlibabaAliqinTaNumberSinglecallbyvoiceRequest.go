package request

import (
        "topsdk/util"
    )

type AlibabaAliqinTaNumberSinglecallbyvoiceRequest struct {
    /*
        单呼号码     */
    CalledNum  *string `json:"called_num" required:"true" `
    /*
        显示号码     */
    CalledShowNum  *string `json:"called_show_num" required:"true" `
    /*
        语音文件code     */
    VoiceCode  *string `json:"voice_code" required:"true" `
    /*
        上下文参数 示例:{"extend":"回传参数"} extend为扩展信息作为回传参数的key     */
    Params  *map[string]interface{} `json:"params,omitempty" required:"false" `
}

func (s *AlibabaAliqinTaNumberSinglecallbyvoiceRequest) SetCalledNum(v string) *AlibabaAliqinTaNumberSinglecallbyvoiceRequest {
    s.CalledNum = &v
    return s
}
func (s *AlibabaAliqinTaNumberSinglecallbyvoiceRequest) SetCalledShowNum(v string) *AlibabaAliqinTaNumberSinglecallbyvoiceRequest {
    s.CalledShowNum = &v
    return s
}
func (s *AlibabaAliqinTaNumberSinglecallbyvoiceRequest) SetVoiceCode(v string) *AlibabaAliqinTaNumberSinglecallbyvoiceRequest {
    s.VoiceCode = &v
    return s
}
func (s *AlibabaAliqinTaNumberSinglecallbyvoiceRequest) SetParams(v map[string]interface{}) *AlibabaAliqinTaNumberSinglecallbyvoiceRequest {
    s.Params = &v
    return s
}

func (req *AlibabaAliqinTaNumberSinglecallbyvoiceRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.CalledNum != nil) {
        paramMap["called_num"] = *req.CalledNum
    }
    if(req.CalledShowNum != nil) {
        paramMap["called_show_num"] = *req.CalledShowNum
    }
    if(req.VoiceCode != nil) {
        paramMap["voice_code"] = *req.VoiceCode
    }
    if(req.Params != nil) {
        paramMap["params"] = util.ConvertStruct(*req.Params)
    }
    return paramMap
}

func (req *AlibabaAliqinTaNumberSinglecallbyvoiceRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}