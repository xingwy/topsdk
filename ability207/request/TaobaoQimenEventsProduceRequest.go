package request

import (
	"github.com/xingwy/topsdk/ability207/domain"
	"github.com/xingwy/topsdk/util"
)

type TaobaoQimenEventsProduceRequest struct {
	/*
	   奇门事件列表, 最多50条     */
	Messages *[]domain.TaobaoQimenEventsProduceQimenEvent `json:"messages" required:"true" `
}

func (s *TaobaoQimenEventsProduceRequest) SetMessages(v []domain.TaobaoQimenEventsProduceQimenEvent) *TaobaoQimenEventsProduceRequest {
	s.Messages = &v
	return s
}

func (req *TaobaoQimenEventsProduceRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.Messages != nil {
		paramMap["messages"] = util.ConvertStructList(*req.Messages)
	}
	return paramMap
}

func (req *TaobaoQimenEventsProduceRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
