package models

type Response struct {
	Id              string       `json:"id"`
	ParentId        string       `json:"parentId"`
	RoomId          string       `json:"roomId"`
	RoomType        string       `json:"roomType"`
	ToPersonId      string       `json:"toPersonId,omitempty"`
	ToPersonEmail   string       `json:"toPersonEmail,omitempty"`
	Text            string       `json:"text"`
	Markdown        string       `json:"markdown,omitempty"`
	Html            string       `json:"html"`
	Files           []string     `json:"files,omitempty"`
	PersonId        string       `json:"personId"`
	PersonEmail     string       `json:"personEmail"`
	MentionedPeople []string     `json:"mentionedPeople,omitempty"`
	MentionedGroups []string     `json:"mentionedGroups,omitempty"`
	Attachments     []Attachment `json:"attachments,omitempty"`
	Created         string       `json:"created"`
	Updated         string       `json:"updated"`
	IsVoiceClip     bool         `json:"isVoiceClip"`
}
