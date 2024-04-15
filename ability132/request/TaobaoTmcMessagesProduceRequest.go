package request

import (
        "topsdk/ability132/domain"
        "topsdk/util"
    )

type TaobaoTmcMessagesProduceRequest struct {
    /*
        tmc消息列表, 最多50条，元素结构与taobao.tmc.message.produce一致，用json表示的消息列表。例如：[{"content": "{\"tid\":1234554321,\"status\":\"X_LOGISTICS_PRINTED\",\"action_time\":\"2014-08-08 18:24:00\",\"seller_nick\": \"向阳aa\",\"operator\":\"小张\"}","topic": "taobao_jds_TradeTrace"},{"content": "{\"tid\":1234554321,\"status\":\"X_LOGISTICS_PRINTED\",\"action_time\":\"2014-08-08 18:24:00\",\"seller_nick\": \"向阳aa\",\"operator\":\"小张\"}","topic": "taobao_jds_TradeTrace"}]     */
    Messages  *[]domain.TaobaoTmcMessagesProduceTmcPublishMessage `json:"messages" required:"true" `
}

func (s *TaobaoTmcMessagesProduceRequest) SetMessages(v []domain.TaobaoTmcMessagesProduceTmcPublishMessage) *TaobaoTmcMessagesProduceRequest {
    s.Messages = &v
    return s
}

func (req *TaobaoTmcMessagesProduceRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Messages != nil) {
        paramMap["messages"] = util.ConvertStructList(*req.Messages)
    }
    return paramMap
}

func (req *TaobaoTmcMessagesProduceRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}