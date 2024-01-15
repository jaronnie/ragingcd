package message

import (
	"encoding/json"
)

const (
	STDOUT = "stdout"
	STDIN  = "stdin"
	RESIZE = "resize"
	PING   = "ping"
	CLOSE  = "close"
	HOME   = "home"
)

type TerminalMessage struct {
	Operation string `json:"Op"`
	Data      string `json:"Data"`
	Rows      uint16 `json:"Rows"`
	Cols      uint16 `json:"Cols"`
}

func (m *TerminalMessage) String() string {
	result, _ := json.Marshal(m)
	return string(result)
}

func (m *TerminalMessage) Bytes() []byte {
	result, _ := json.Marshal(m)
	return result
}

func Stdout(msg string) *TerminalMessage {
	return &TerminalMessage{Operation: STDOUT, Data: msg}
}

func Stdin(frameData []byte) (*TerminalMessage, error) {
	result := &TerminalMessage{}
	if err := json.Unmarshal(frameData, result); err != nil {
		return nil, err
	}
	return result, nil
}

func Close(errMsg string) *TerminalMessage {
	return &TerminalMessage{Operation: CLOSE, Data: errMsg}
}

func Ping() *TerminalMessage {
	return &TerminalMessage{Operation: PING}
}
