package request

type TaobaoTmcMessageProduceRequest struct {
	/*
	   消息内容的JSON表述，必须按照topic的定义来填充     */
	Content *string `json:"content" required:"true" `
	/*
	   消息的扩增属性，用json格式表示     */
	ExContent *string `json:"ex_content,omitempty" required:"false" `
	/*
	   直发消息需要传入目标appkey     */
	TargetAppkey *string `json:"target_appkey,omitempty" required:"false" `
	/*
	   目标分组，一般为default defalutValue��default    */
	TargetGroup *string `json:"target_group,omitempty" required:"false" `
	/*
	   消息类型     */
	Topic *string `json:"topic" required:"true" `
	/*
	   回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。     */
	MediaContent *[]byte `json:"media_content,omitempty" required:"false" `
	/*
	   回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。     */
	MediaContent2 *[]byte `json:"media_content2,omitempty" required:"false" `
	/*
	   回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。     */
	MediaContent3 *[]byte `json:"media_content3,omitempty" required:"false" `
	/*
	   回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。     */
	MediaContent5 *[]byte `json:"media_content5,omitempty" required:"false" `
	/*
	   回传的文件内容，目前仅支持jpg,png,bmp,gif,pdf类型的文件，文件最大1M。只有消息中有byte[]类型的数据时，才需要传此字段; 否则不需要传此字段。具体对应到沙体中的什么值，请参考消息字段说明。     */
	MediaContent4 *[]byte `json:"media_content4,omitempty" required:"false" `
	/*
	   延时参数 时间戳 预期发送时间 defalutValue��0    */
	DelayMillis *int64 `json:"delay_millis,omitempty" required:"false" `
	/*
	   提前过期 相对时间差 毫秒 defalutValue��0    */
	ExpiresMillis *int64 `json:"expires_millis,omitempty" required:"false" `
}

func (s *TaobaoTmcMessageProduceRequest) SetContent(v string) *TaobaoTmcMessageProduceRequest {
	s.Content = &v
	return s
}
func (s *TaobaoTmcMessageProduceRequest) SetExContent(v string) *TaobaoTmcMessageProduceRequest {
	s.ExContent = &v
	return s
}
func (s *TaobaoTmcMessageProduceRequest) SetTargetAppkey(v string) *TaobaoTmcMessageProduceRequest {
	s.TargetAppkey = &v
	return s
}
func (s *TaobaoTmcMessageProduceRequest) SetTargetGroup(v string) *TaobaoTmcMessageProduceRequest {
	s.TargetGroup = &v
	return s
}
func (s *TaobaoTmcMessageProduceRequest) SetTopic(v string) *TaobaoTmcMessageProduceRequest {
	s.Topic = &v
	return s
}
func (s *TaobaoTmcMessageProduceRequest) SetMediaContent(v []byte) *TaobaoTmcMessageProduceRequest {
	s.MediaContent = &v
	return s
}
func (s *TaobaoTmcMessageProduceRequest) SetMediaContent2(v []byte) *TaobaoTmcMessageProduceRequest {
	s.MediaContent2 = &v
	return s
}
func (s *TaobaoTmcMessageProduceRequest) SetMediaContent3(v []byte) *TaobaoTmcMessageProduceRequest {
	s.MediaContent3 = &v
	return s
}
func (s *TaobaoTmcMessageProduceRequest) SetMediaContent5(v []byte) *TaobaoTmcMessageProduceRequest {
	s.MediaContent5 = &v
	return s
}
func (s *TaobaoTmcMessageProduceRequest) SetMediaContent4(v []byte) *TaobaoTmcMessageProduceRequest {
	s.MediaContent4 = &v
	return s
}
func (s *TaobaoTmcMessageProduceRequest) SetDelayMillis(v int64) *TaobaoTmcMessageProduceRequest {
	s.DelayMillis = &v
	return s
}
func (s *TaobaoTmcMessageProduceRequest) SetExpiresMillis(v int64) *TaobaoTmcMessageProduceRequest {
	s.ExpiresMillis = &v
	return s
}

func (req *TaobaoTmcMessageProduceRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Content != nil {
		paramMap["content"] = *req.Content
	}
	if req.ExContent != nil {
		paramMap["ex_content"] = *req.ExContent
	}
	if req.TargetAppkey != nil {
		paramMap["target_appkey"] = *req.TargetAppkey
	}
	if req.TargetGroup != nil {
		paramMap["target_group"] = *req.TargetGroup
	}
	if req.Topic != nil {
		paramMap["topic"] = *req.Topic
	}
	if req.DelayMillis != nil {
		paramMap["delay_millis"] = *req.DelayMillis
	}
	if req.ExpiresMillis != nil {
		paramMap["expires_millis"] = *req.ExpiresMillis
	}
	return paramMap
}

func (req *TaobaoTmcMessageProduceRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	if req.MediaContent != nil {
		fileMap["media_content"] = *req.MediaContent
	}
	if req.MediaContent2 != nil {
		fileMap["media_content2"] = *req.MediaContent2
	}
	if req.MediaContent3 != nil {
		fileMap["media_content3"] = *req.MediaContent3
	}
	if req.MediaContent5 != nil {
		fileMap["media_content5"] = *req.MediaContent5
	}
	if req.MediaContent4 != nil {
		fileMap["media_content4"] = *req.MediaContent4
	}
	return fileMap
}
