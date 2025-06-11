package messages

import (
	"github.com/rainuxhe/webexgosdk/models"
)

type Response struct {
	Id              string              `json:"id"`
	ParentId        string              `json:"parentId"`
	RoomId          string              `json:"roomId"`
	RoomType        string              `json:"roomType"`
	ToPersonId      string              `json:"toPersonId,omitempty"`
	ToPersonEmail   string              `json:"toPersonEmail,omitempty"`
	Text            string              `json:"text"`
	Markdown        string              `json:"markdown,omitempty"`
	Html            string              `json:"html"`
	Files           []string            `json:"files,omitempty"`
	PersonId        string              `json:"personId"`
	PersonEmail     string              `json:"personEmail"`
	MentionedPeople []string            `json:"mentionedPeople,omitempty"`
	MentionedGroups []string            `json:"mentionedGroups,omitempty"`
	Attachments     []models.Attachment `json:"attachments,omitempty"`
	Created         string              `json:"created"`
	Updated         string              `json:"updated"`
	IsVoiceClip     bool                `json:"isVoiceClip"`
}

type messageReqBody struct {
	RoomId          string              `json:"roomId,omitempty"`
	ParentId        string              `json:"parentId,omitempty"`
	ToPersonId      string              `json:"toPersonId,omitempty"`
	ToPersonEmail   string              `json:"toPersonEmail,omitempty"`
	Text            string              `json:"text,omitempty"`
	Markdown        string              `json:"markdown,omitempty"`
	Files           []string            `json:"files,omitempty"`
	Attachments     []models.Attachment `json:"attachments,omitempty"`
	MentionedPeople []string            `json:"mentionedPeople,omitempty"`
	Before          string              `json:"before,omitempty"`
	BeforeMessage   string              `json:"beforeMessage,omitempty"`
	Max             int                 `json:"max,omitempty"`
}
