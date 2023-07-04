package kook_card

type CardMessage struct {
	Type    string              `json:"type"`
	Theme   string              `json:"theme,omitempty"`
	Size    string              `json:"size,omitempty"`
	Modules []CardMessageModule `json:"modules"`
}
