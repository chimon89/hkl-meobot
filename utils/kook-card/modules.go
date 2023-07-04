package kook_card

type TextModule struct {
	Type    string          `json:"type"`
	Cols    int             `json:"cols,omitempty"`
	Content string          `json:"content,omitempty"`
	Fields  *[]FieldsModule `json:"fields,om,omitempty"`
}

type FieldsModule struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

type CardMessageModule struct {
	Type      string            `json:"type"`
	Text      *TextModule       `json:"text,omitempty"`
	Elements  *[]ElementsModule `json:"elements,omitempty"`
	Mode      string            `json:"mode,omitempty"`
	EndTime   int               `json:"endTime,omitempty"`
	Accessory *ElementsModule   `json:"accessory,omitempty"`
}

type ElementsModule struct {
	Type  string      `json:"type"`
	Theme string      `json:"theme,omitempty"`
	Value string      `json:"value"`
	Click string      `json:"click"`
	Text  *TextModule `json:"text"`
}
