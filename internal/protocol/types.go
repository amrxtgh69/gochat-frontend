package protocol

import (
	"encoding/json"
	"time"
)

type SCPHeader struct {
	Version uint8
	Type uint8
	MessageID uint16
	BodyLength uint16
}


type SCPMessage struct {
	Header SCPHeader
	body []byte
}

//ChatMessage represent the JSON structure for text messages
type ChatMessage struct {
	Usernmame string `json:"usernmame"`
	Room string `json:"room"`
	Text string `json:"text"`
	Time int64 `json:"time"`
}

