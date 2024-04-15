package domain

import (
	"github.com/xingwy/topsdk/util"
)

type AlibabaAliqinTaSmsNumQueryFcPartnerSmsDetailDto struct {
	/*
	   公共回传参数     */
	Extend *string `json:"extend,omitempty" `

	/*
	   短信接收号码     */
	RecNum *string `json:"rec_num,omitempty" `

	/*
	   错误码     */
	ResultCode *string `json:"result_code,omitempty" `

	/*
	   模板编码     */
	SmsCode *string `json:"sms_code,omitempty" `

	/*
	   短信内容     */
	SmsContent *string `json:"sms_content,omitempty" `

	/*
	   短信接收时间     */
	SmsReceiverTime *util.LocalTime `json:"sms_receiver_time,omitempty" `

	/*
	   短信发送时间     */
	SmsSendTime *util.LocalTime `json:"sms_send_time,omitempty" `

	/*
	   发送状态 1：等待回执，2：发送失败，3：发送成功     */
	SmsStatus *int64 `json:"sms_status,omitempty" `
}

func (s *AlibabaAliqinTaSmsNumQueryFcPartnerSmsDetailDto) SetExtend(v string) *AlibabaAliqinTaSmsNumQueryFcPartnerSmsDetailDto {
	s.Extend = &v
	return s
}
func (s *AlibabaAliqinTaSmsNumQueryFcPartnerSmsDetailDto) SetRecNum(v string) *AlibabaAliqinTaSmsNumQueryFcPartnerSmsDetailDto {
	s.RecNum = &v
	return s
}
func (s *AlibabaAliqinTaSmsNumQueryFcPartnerSmsDetailDto) SetResultCode(v string) *AlibabaAliqinTaSmsNumQueryFcPartnerSmsDetailDto {
	s.ResultCode = &v
	return s
}
func (s *AlibabaAliqinTaSmsNumQueryFcPartnerSmsDetailDto) SetSmsCode(v string) *AlibabaAliqinTaSmsNumQueryFcPartnerSmsDetailDto {
	s.SmsCode = &v
	return s
}
func (s *AlibabaAliqinTaSmsNumQueryFcPartnerSmsDetailDto) SetSmsContent(v string) *AlibabaAliqinTaSmsNumQueryFcPartnerSmsDetailDto {
	s.SmsContent = &v
	return s
}
func (s *AlibabaAliqinTaSmsNumQueryFcPartnerSmsDetailDto) SetSmsReceiverTime(v util.LocalTime) *AlibabaAliqinTaSmsNumQueryFcPartnerSmsDetailDto {
	s.SmsReceiverTime = &v
	return s
}
func (s *AlibabaAliqinTaSmsNumQueryFcPartnerSmsDetailDto) SetSmsSendTime(v util.LocalTime) *AlibabaAliqinTaSmsNumQueryFcPartnerSmsDetailDto {
	s.SmsSendTime = &v
	return s
}
func (s *AlibabaAliqinTaSmsNumQueryFcPartnerSmsDetailDto) SetSmsStatus(v int64) *AlibabaAliqinTaSmsNumQueryFcPartnerSmsDetailDto {
	s.SmsStatus = &v
	return s
}
