package api

import (
	"encoding/json"
	"net/url"
)

type ApexMapRotateInfo struct {
	BattleRoyale struct {
		Current struct {
			Start             int    `json:"start"`
			End               int    `json:"end"`
			ReadableDateStart string `json:"readableDate_start"`
			ReadableDateEnd   string `json:"readableDate_end"`
			Map               string `json:"map"`
			Code              string `json:"code"`
			DurationInSecs    int    `json:"DurationInSecs"`
			DurationInMinutes int    `json:"DurationInMinutes"`
			Asset             string `json:"asset"`
			RemainingSecs     int    `json:"remainingSecs"`
			RemainingMins     int    `json:"remainingMins"`
			RemainingTimer    string `json:"remainingTimer"`
		} `json:"current"`
		Next struct {
			Start             int    `json:"start"`
			End               int    `json:"end"`
			ReadableDateStart string `json:"readableDate_start"`
			ReadableDateEnd   string `json:"readableDate_end"`
			Map               string `json:"map"`
			Code              string `json:"code"`
			DurationInSecs    int    `json:"DurationInSecs"`
			DurationInMinutes int    `json:"DurationInMinutes"`
			Asset             string `json:"asset"`
		} `json:"next"`
	} `json:"battle_royale"`
	Arenas struct {
		Current struct {
			Start             int    `json:"start"`
			End               int    `json:"end"`
			ReadableDateStart string `json:"readableDate_start"`
			ReadableDateEnd   string `json:"readableDate_end"`
			Map               string `json:"map"`
			Code              string `json:"code"`
			DurationInSecs    int    `json:"DurationInSecs"`
			DurationInMinutes int    `json:"DurationInMinutes"`
			Asset             string `json:"asset"`
			RemainingSecs     int    `json:"remainingSecs"`
			RemainingMins     int    `json:"remainingMins"`
			RemainingTimer    string `json:"remainingTimer"`
		} `json:"current"`
		Next struct {
			Start             int    `json:"start"`
			End               int    `json:"end"`
			ReadableDateStart string `json:"readableDate_start"`
			ReadableDateEnd   string `json:"readableDate_end"`
			Map               string `json:"map"`
			Code              string `json:"code"`
			DurationInSecs    int    `json:"DurationInSecs"`
			DurationInMinutes int    `json:"DurationInMinutes"`
			Asset             string `json:"asset"`
		} `json:"next"`
	} `json:"arenas"`
	Ranked struct {
		Current struct {
			Start             int    `json:"start"`
			End               int    `json:"end"`
			ReadableDateStart string `json:"readableDate_start"`
			ReadableDateEnd   string `json:"readableDate_end"`
			Map               string `json:"map"`
			Code              string `json:"code"`
			DurationInSecs    int    `json:"DurationInSecs"`
			DurationInMinutes int    `json:"DurationInMinutes"`
			Asset             string `json:"asset"`
			RemainingSecs     int    `json:"remainingSecs"`
			RemainingMins     int    `json:"remainingMins"`
			RemainingTimer    string `json:"remainingTimer"`
		} `json:"current"`
		Next struct {
			Start             int    `json:"start"`
			End               int    `json:"end"`
			ReadableDateStart string `json:"readableDate_start"`
			ReadableDateEnd   string `json:"readableDate_end"`
			Map               string `json:"map"`
			Code              string `json:"code"`
			DurationInSecs    int    `json:"DurationInSecs"`
			DurationInMinutes int    `json:"DurationInMinutes"`
			Asset             string `json:"asset"`
		} `json:"next"`
	} `json:"ranked"`
	ArenasRanked struct {
		Current struct {
			Start             int    `json:"start"`
			End               int    `json:"end"`
			ReadableDateStart string `json:"readableDate_start"`
			ReadableDateEnd   string `json:"readableDate_end"`
			Map               string `json:"map"`
			Code              string `json:"code"`
			DurationInSecs    int    `json:"DurationInSecs"`
			DurationInMinutes int    `json:"DurationInMinutes"`
			Asset             string `json:"asset"`
			RemainingSecs     int    `json:"remainingSecs"`
			RemainingMins     int    `json:"remainingMins"`
			RemainingTimer    string `json:"remainingTimer"`
		} `json:"current"`
		Next struct {
			Start             int    `json:"start"`
			End               int    `json:"end"`
			ReadableDateStart string `json:"readableDate_start"`
			ReadableDateEnd   string `json:"readableDate_end"`
			Map               string `json:"map"`
			Code              string `json:"code"`
			DurationInSecs    int    `json:"DurationInSecs"`
			DurationInMinutes int    `json:"DurationInMinutes"`
			Asset             string `json:"asset"`
		} `json:"next"`
	} `json:"arenasRanked"`
	Ltm struct {
		Current struct {
			Start             int    `json:"start"`
			End               int    `json:"end"`
			ReadableDateStart string `json:"readableDate_start"`
			ReadableDateEnd   string `json:"readableDate_end"`
			Map               string `json:"map"`
			Code              string `json:"code"`
			DurationInSecs    int    `json:"DurationInSecs"`
			DurationInMinutes int    `json:"DurationInMinutes"`
			EventName         string `json:"eventName"`
			Asset             string `json:"asset"`
			RemainingSecs     int    `json:"remainingSecs"`
			RemainingMins     int    `json:"remainingMins"`
			RemainingTimer    string `json:"remainingTimer"`
		} `json:"current"`
		Next struct {
			Start             int    `json:"start"`
			End               int    `json:"end"`
			ReadableDateStart string `json:"readableDate_start"`
			ReadableDateEnd   string `json:"readableDate_end"`
			Map               string `json:"map"`
			Code              string `json:"code"`
			DurationInSecs    int    `json:"DurationInSecs"`
			DurationInMinutes int    `json:"DurationInMinutes"`
			Asset             string `json:"asset"`
		} `json:"next"`
	} `json:"ltm"`
}

func ApexMapRotate() (*ApexMapRotateInfo, error) {
	values := make(url.Values)
	values.Add("auth", GameApiToken)
	values.Add("version", "2")
	resp, err := GameInstance().Get("/maprotation", values)
	if err != nil {
		return &ApexMapRotateInfo{}, err
	}
	var amri ApexMapRotateInfo
	json.Unmarshal(resp.Data, &amri)
	return &amri, nil
}
