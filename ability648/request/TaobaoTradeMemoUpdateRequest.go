package request


type TaobaoTradeMemoUpdateRequest struct {
    /*
        交易编号     */
    Tid  *int64 `json:"tid" required:"true" `
    /*
        卖家交易备注。最大长度: 1000个字节     */
    Memo  *string `json:"memo,omitempty" required:"false" `
    /*
        卖家交易备注旗帜，可选值为：0(灰色), 1(红色), 2(黄色), 3(绿色), 4(蓝色), 5(粉红色)，默认值为0     */
    Flag  *int64 `json:"flag,omitempty" required:"false" `
    /*
        是否对memo的值置空若为true，则不管传入的memo字段的值是否为空，都将会对已有的memo值清空，慎用；若用false，则会根据memo是否为空来修改memo的值：若memo为空则忽略对已有memo字段的修改，若memo非空，则使用新传入的memo覆盖已有的memo的值 defalutValue��false    */
    Reset  *bool `json:"reset,omitempty" required:"false" `
}

func (s *TaobaoTradeMemoUpdateRequest) SetTid(v int64) *TaobaoTradeMemoUpdateRequest {
    s.Tid = &v
    return s
}
func (s *TaobaoTradeMemoUpdateRequest) SetMemo(v string) *TaobaoTradeMemoUpdateRequest {
    s.Memo = &v
    return s
}
func (s *TaobaoTradeMemoUpdateRequest) SetFlag(v int64) *TaobaoTradeMemoUpdateRequest {
    s.Flag = &v
    return s
}
func (s *TaobaoTradeMemoUpdateRequest) SetReset(v bool) *TaobaoTradeMemoUpdateRequest {
    s.Reset = &v
    return s
}

func (req *TaobaoTradeMemoUpdateRequest) ToMap() map[string]interface{} {
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
    if(req.Reset != nil) {
        paramMap["reset"] = *req.Reset
    }
    return paramMap
}

func (req *TaobaoTradeMemoUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}