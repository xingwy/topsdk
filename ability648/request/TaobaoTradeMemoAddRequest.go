package request


type TaobaoTradeMemoAddRequest struct {
    /*
        交易编号     */
    Tid  *int64 `json:"tid" required:"true" `
    /*
        交易备注。最大长度: 1000个字节     */
    Memo  *string `json:"memo" required:"true" `
    /*
        交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0 defalutValue��0    */
    Flag  *int64 `json:"flag,omitempty" required:"false" `
    /*
        插旗备注     */
    Tag  *string `json:"tag,omitempty" required:"false" `
}

func (s *TaobaoTradeMemoAddRequest) SetTid(v int64) *TaobaoTradeMemoAddRequest {
    s.Tid = &v
    return s
}
func (s *TaobaoTradeMemoAddRequest) SetMemo(v string) *TaobaoTradeMemoAddRequest {
    s.Memo = &v
    return s
}
func (s *TaobaoTradeMemoAddRequest) SetFlag(v int64) *TaobaoTradeMemoAddRequest {
    s.Flag = &v
    return s
}
func (s *TaobaoTradeMemoAddRequest) SetTag(v string) *TaobaoTradeMemoAddRequest {
    s.Tag = &v
    return s
}

func (req *TaobaoTradeMemoAddRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Tid != nil) {
        paramMap["tid"] = *req.Tid
    }
    if(req.Memo != nil) {
        paramMap["memo"] = *req.Memo
    }
    if(req.Flag != nil) {
        paramMap["flag"] = *req.Flag
    }
    if(req.Tag != nil) {
        paramMap["tag"] = *req.Tag
    }
    return paramMap
}

func (req *TaobaoTradeMemoAddRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}