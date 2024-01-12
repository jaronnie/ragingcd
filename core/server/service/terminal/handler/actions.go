package handler

import (
	"fmt"

	"github.com/jaronnie/ragingcd/core/server/service/terminal/message"
	"github.com/jaronnie/ragingcd/core/server/service/terminal/target"

	cmd "k8s.io/client-go/tools/remotecommand"
)

type MsgActionFunc func(stdin []byte, msg *message.TerminalMessage, sizeChan chan cmd.TerminalSize, target *target.Target) (int, error)

var actionFactory = map[string]MsgActionFunc{}

func init() {
	actionFactory[message.STDIN] = MsgStdin
	actionFactory[message.RESIZE] = MsgResize
	actionFactory[message.HOME] = MsgHome
	actionFactory[message.PING] = MsgEmpty
	actionFactory[message.CLOSE] = MsgEmpty
}

func MsgStdin(stdin []byte, msg *message.TerminalMessage, sizeChan chan cmd.TerminalSize, target *target.Target) (int, error) {
	count := copy(stdin, msg.Data)
	return count, nil
}

func MsgResize(stdin []byte, msg *message.TerminalMessage, sizeChan chan cmd.TerminalSize, target *target.Target) (int, error) {
	sizeChan <- cmd.TerminalSize{Width: msg.Cols, Height: msg.Rows}
	count := copy(stdin, fmt.Sprint("export TERM='xterm-256color'\n"))
	return count, nil
}

func MsgHome(stdin []byte, msg *message.TerminalMessage, sizeChan chan cmd.TerminalSize, target *target.Target) (int, error) {
	homeCMD := fmt.Sprintf("cd %s\n", target.WorkingDir)
	count := copy(stdin, homeCMD)
	return count, nil
}

func MsgEmpty(stdin []byte, msg *message.TerminalMessage, sizeChan chan cmd.TerminalSize, target *target.Target) (int, error) {
	return 0, nil
}

func MsgError(stdin []byte, msg *message.TerminalMessage, sizeChan chan cmd.TerminalSize, target *target.Target) (int, error) {
	copy(stdin, EndOfTransmission)
	return 0, fmt.Errorf("fail to do handle msg, because unknown message type '%s'", msg.Operation)
}

func MsgPing(stdin []byte, msg *message.TerminalMessage, sizeChan chan cmd.TerminalSize, target *target.Target) (int, error) {
	count := copy(stdin, NullCharacter)
	return count, nil
}

func msgAction(msgType string) MsgActionFunc {
	handler, ok := actionFactory[msgType]
	if ok {
		return handler
	}
	return MsgError
}
