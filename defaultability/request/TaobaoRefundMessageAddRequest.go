package request

type TaobaoRefundMessageAddRequest struct {
	/*
	   图片（凭证）。类型: JPG,GIF,PNG;最大为: 500K     */
	Image *[]byte `json:"image" required:"true" `
	/*
	   退款编号。     */
	RefundId *int64 `json:"refund_id" required:"true" `
	/*
	   留言内容。最大长度: 400个字节     */
	Content *string `json:"content" required:"true" `
}

func (s *TaobaoRefundMessageAddRequest) SetImage(v []byte) *TaobaoRefundMessageAddRequest {
	s.Image = &v
	return s
}
func (s *TaobaoRefundMessageAddRequest) SetRefundId(v int64) *TaobaoRefundMessageAddRequest {
	s.RefundId = &v
	return s
}
func (s *TaobaoRefundMessageAddRequest) SetContent(v string) *TaobaoRefundMessageAddRequest {
	s.Content = &v
	return s
}

func (req *TaobaoRefundMessageAddRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.RefundId != nil {
		paramMap["refund_id"] = *req.RefundId
	}
	if req.Content != nil {
		paramMap["content"] = *req.Content
	}
	return paramMap
}

func (req *TaobaoRefundMessageAddRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	if req.Image != nil {
		fileMap["image"] = *req.Image
	}
	return fileMap
}
