package event

type Data struct {
	Type         int    `json:"type"`
	ChannelType  string `json:"channel_type"`
	TargetID     string `json:"target_id"`
	AuthorID     string `json:"author_id"`
	Content      string `json:"content"`
	MsgID        string `json:"msg_id"`
	MsgTimestamp int    `json:"msg_timestamp"`
	Nonce        string `json:"nonce"`
	Extra        Extra  `json:"extra"`
	Challenge    string `json:"challenge"`
	VerifyToken  string `json:"verify_token"`
}

type Extra struct {
	Type         interface{} `json:"type"`
	GuildID      string      `json:"guild_id"`
	ChannelName  string      `json:"channel_name"`
	Mention      []string    `json:"mention"`
	MentionAll   bool        `json:"mention_all"`
	MentionRoles []string    `json:"mention_roles"`
	MentionHere  bool        `json:"mention_here"`
	Author       Author      `json:"author"`
	Body         ExtraBody   `json:"body"`
}

type ExtraBody struct {
	MsgID    string      `json:"msg_id"`
	UserID   string      `json:"user_id"`
	Value    string      `json:"value"`
	TargetID string      `json:"target_id"`
	UserInfo interface{} `json:"user_info"`
}

type Author struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	IdentifyNum    string `json:"identify_num"`
	Online         bool   `json:"online"`
	Avatar         string `json:"avatar"`
	VipAvatar      string `json:"vip_avatar"`
	Bot            bool   `json:"bot"`
	Status         int    `json:"status"`
	MobileVerified bool   `json:"mobile_verified"`
	Nickname       string `json:"nickname"`
	Roles          []int  `json:"roles"`
}

type Event struct {
	S int  `json:"s"`
	D Data `json:"d"`
}
