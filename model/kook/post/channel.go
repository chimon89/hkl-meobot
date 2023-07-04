package post

type ChannelMessageCreate struct {
	Type         int    `json:"type,omitempty"`
	TargetID     string `json:"target_id"`
	ChatCode     string `json:"chat_code,omitempty"`
	Content      string `json:"content"`
	Quote        string `json:"quote,omitempty"`
	Nonce        string `json:"nonce,omitempty"`
	TempTargetID string `json:"temp_target_id,omitempty"`
}

type ChannelMessageUpdate struct {
	MsgID        string `json:"msg_id"`
	Content      string `json:"content"`
	Quote        string `json:"quote,omitempty"`
	TempTargetID string `json:"temp_target_id,omitempty"`
}
