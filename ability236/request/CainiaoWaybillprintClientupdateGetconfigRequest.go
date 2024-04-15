package request


type CainiaoWaybillprintClientupdateGetconfigRequest struct {
    /*
        服务发起机器的物理地址     */
    Mac  *string `json:"mac" required:"true" `
    /*
        服务发起机器局域网ip     */
    LanIp  *string `json:"lan_ip" required:"true" `
    /*
        当前打印客户端版本     */
    Version  *string `json:"version" required:"true" `
    /*
        当前消息版本     */
    Msgid  *int64 `json:"msgid" required:"true" `
}

func (s *CainiaoWaybillprintClientupdateGetconfigRequest) SetMac(v string) *CainiaoWaybillprintClientupdateGetconfigRequest {
    s.Mac = &v
    return s
}
func (s *CainiaoWaybillprintClientupdateGetconfigRequest) SetLanIp(v string) *CainiaoWaybillprintClientupdateGetconfigRequest {
    s.LanIp = &v
    return s
}
func (s *CainiaoWaybillprintClientupdateGetconfigRequest) SetVersion(v string) *CainiaoWaybillprintClientupdateGetconfigRequest {
    s.Version = &v
    return s
}
func (s *CainiaoWaybillprintClientupdateGetconfigRequest) SetMsgid(v int64) *CainiaoWaybillprintClientupdateGetconfigRequest {
    s.Msgid = &v
    return s
}

func (req *CainiaoWaybillprintClientupdateGetconfigRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Mac != nil) {
        paramMap["mac"] = *req.Mac
    }
    if(req.LanIp != nil) {
        paramMap["lan_ip"] = *req.LanIp
    }
    if(req.Version != nil) {
        paramMap["version"] = *req.Version
    }
    if(req.Msgid != nil) {
        paramMap["msgid"] = *req.Msgid
    }
    return paramMap
}

func (req *CainiaoWaybillprintClientupdateGetconfigRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}