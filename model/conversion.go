package model

import (
	"encoding/xml"
	"time"
)

type MsgStruct struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string   `xml:"ToUserName"`
	FromUserName string   `xml:"FromUserName"`
	CreateTime   int64    `xml:"CreateTime"`
	MsgType      string   `xml:"MsgType"`
	Content      string   `xml:"Content"`
	MsgId        int64    `xml:"MsgId,omitempty"`
}

func NewReplyMsq(msg *MsgStruct, content string) *MsgStruct {
	reply := &MsgStruct{}
	reply.XMLName = msg.XMLName
	reply.ToUserName = msg.FromUserName
	reply.FromUserName = msg.ToUserName
	reply.CreateTime = time.Now().Unix()
	reply.MsgType = msg.MsgType
	reply.Content = content

	return reply
}
