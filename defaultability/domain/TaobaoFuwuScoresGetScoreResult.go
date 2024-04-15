package domain

import (
        "topsdk/util"
    )

type TaobaoFuwuScoresGetScoreResult struct {
    /*
        是否实际付费 1-实际付费 2-实际未付费     */
    IsPay  *int64 `json:"is_pay,omitempty" `

    /*
        稳定性评分     */
    StabilityScore  *string `json:"stability_score,omitempty" `

    /*
        交片速度     */
    RapidScore  *string `json:"rapid_score,omitempty" `

    /*
        平均分     */
    AvgScore  *string `json:"avg_score,omitempty" `

    /*
        是否为有效评分 1-有效评分 2-无效评分     */
    IsValid  *int64 `json:"is_valid,omitempty" `

    /*
        评价id     */
    Id  *int64 `json:"id,omitempty" `

    /*
        专业性评分     */
    ProfScore  *string `json:"prof_score,omitempty" `

    /*
        评价时间     */
    GmtCreate  *util.LocalTime `json:"gmt_create,omitempty" `

    /*
        服务规格code     */
    ItemCode  *string `json:"item_code,omitempty" `

    /*
        服务规格名称     */
    ItemName  *string `json:"item_name,omitempty" `

    /*
        服务态度评分     */
    AttitudeScore  *string `json:"attitude_score,omitempty" `

    /*
        描述相符     */
    MatchedScore  *string `json:"matched_score,omitempty" `

    /*
        评论内容     */
    Suggestion  *string `json:"suggestion,omitempty" `

    /*
        服务code     */
    ServiceCode  *string `json:"service_code,omitempty" `

    /*
        评价人用户昵称     */
    UserNick  *string `json:"user_nick,omitempty" `

    /*
        易用性评分     */
    EasyuseScore  *string `json:"easyuse_score,omitempty" `

}

func (s *TaobaoFuwuScoresGetScoreResult) SetIsPay(v int64) *TaobaoFuwuScoresGetScoreResult {
    s.IsPay = &v
    return s
}
func (s *TaobaoFuwuScoresGetScoreResult) SetStabilityScore(v string) *TaobaoFuwuScoresGetScoreResult {
    s.StabilityScore = &v
    return s
}
func (s *TaobaoFuwuScoresGetScoreResult) SetRapidScore(v string) *TaobaoFuwuScoresGetScoreResult {
    s.RapidScore = &v
    return s
}
func (s *TaobaoFuwuScoresGetScoreResult) SetAvgScore(v string) *TaobaoFuwuScoresGetScoreResult {
    s.AvgScore = &v
    return s
}
func (s *TaobaoFuwuScoresGetScoreResult) SetIsValid(v int64) *TaobaoFuwuScoresGetScoreResult {
    s.IsValid = &v
    return s
}
func (s *TaobaoFuwuScoresGetScoreResult) SetId(v int64) *TaobaoFuwuScoresGetScoreResult {
    s.Id = &v
    return s
}
func (s *TaobaoFuwuScoresGetScoreResult) SetProfScore(v string) *TaobaoFuwuScoresGetScoreResult {
    s.ProfScore = &v
    return s
}
func (s *TaobaoFuwuScoresGetScoreResult) SetGmtCreate(v util.LocalTime) *TaobaoFuwuScoresGetScoreResult {
    s.GmtCreate = &v
    return s
}
func (s *TaobaoFuwuScoresGetScoreResult) SetItemCode(v string) *TaobaoFuwuScoresGetScoreResult {
    s.ItemCode = &v
    return s
}
func (s *TaobaoFuwuScoresGetScoreResult) SetItemName(v string) *TaobaoFuwuScoresGetScoreResult {
    s.ItemName = &v
    return s
}
func (s *TaobaoFuwuScoresGetScoreResult) SetAttitudeScore(v string) *TaobaoFuwuScoresGetScoreResult {
    s.AttitudeScore = &v
    return s
}
func (s *TaobaoFuwuScoresGetScoreResult) SetMatchedScore(v string) *TaobaoFuwuScoresGetScoreResult {
    s.MatchedScore = &v
    return s
}
func (s *TaobaoFuwuScoresGetScoreResult) SetSuggestion(v string) *TaobaoFuwuScoresGetScoreResult {
    s.Suggestion = &v
    return s
}
func (s *TaobaoFuwuScoresGetScoreResult) SetServiceCode(v string) *TaobaoFuwuScoresGetScoreResult {
    s.ServiceCode = &v
    return s
}
func (s *TaobaoFuwuScoresGetScoreResult) SetUserNick(v string) *TaobaoFuwuScoresGetScoreResult {
    s.UserNick = &v
    return s
}
func (s *TaobaoFuwuScoresGetScoreResult) SetEasyuseScore(v string) *TaobaoFuwuScoresGetScoreResult {
    s.EasyuseScore = &v
    return s
}
