package server

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"time"
)

const QueueLen = 1000

var ServersNew = sync.Map{}

type GoState int

const (
	// Stopped server is stopped
	Stopped GoState = iota
	// Started server is started
	Started
)

type MsgAction func(msg interface{}, num int) (resp interface{}, err error)

type ScheduledTask func(num int) error

type Req struct {
	ch  chan interface{}
	msg interface{}
	err error
}

type Server struct {
	Name                      string
	Queue                     chan *Req
	State                     GoState
	ActionGoroutineNum        int
	MsgActionerExit           []chan int
	MsgActioner               MsgAction
	ScheduledTaskGoroutineNum int
	ScheduledTaskerExit       []chan int
	ScheduledTaskers          []ScheduledTask
	ScheduleTime              time.Duration
	TimedTasks                []TimedTask
}

type TimedTask struct {
	Task ScheduledTask
	Time time.Duration
}

type Option struct {
	ActionGoroutineNum int
}

func NewSvr(serverName string, action MsgAction, timedTasks []TimedTask, options *Option) (server *Server, err error) {
	_, ok := ServersNew.Load(serverName)
	if ok {
		return nil, fmt.Errorf("the same server serverName already exists:%s", serverName)
	}

	actionGoroutineNum := runtime.NumCPU() * 2
	if options != nil {
		actionGoroutineNum = options.ActionGoroutineNum
	}

	server = &Server{
		Name:                      serverName,
		ActionGoroutineNum:        actionGoroutineNum,
		Queue:                     make(chan *Req, QueueLen),
		MsgActionerExit:           []chan int{},
		MsgActioner:               action,
		ScheduledTaskGoroutineNum: len(timedTasks),
		ScheduledTaskerExit:       []chan int{},
		TimedTasks:                timedTasks,
	}
	if action == nil {
		server.ActionGoroutineNum = 0
	}
	for i := 0; i < server.ActionGoroutineNum; i++ {
		server.MsgActionerExit = append(server.MsgActionerExit, make(chan int))
	}
	for i := 0; i < server.ScheduledTaskGoroutineNum; i++ {
		server.ScheduledTaskerExit = append(server.ScheduledTaskerExit, make(chan int))
	}

	server.State = Stopped
	ServersNew.Store(serverName, server)
	return server, nil
}

func StopServer(serverName string) {
	load, ok := ServersNew.Load(serverName)
	if !ok {
		return
	}
	s := load.(*Server)
	s.stop()
	close(s.Queue)
	ServersNew.Delete(serverName)
}

func PushMsgToServer(serverName string, msg interface{}) (interface{}, error) {
	load, ok := ServersNew.Load(serverName)
	if !ok {
		return nil, fmt.Errorf("server name doesn't exist:%s", serverName)
	}
	s, ok := load.(*Server)

	if !ok {
		return nil, errors.New("empty server")
	}

	return s.receiveMsg(&Req{ch: make(chan interface{}), msg: msg, err: nil})
}

func (server *Server) receiveMsg(req *Req) (interface{}, error) {
	server.Queue <- req
	resp := <-req.ch
	return resp, req.err
}

// Go 开启服务，定时任务时间间隔：ScheduleTime
func (server *Server) Go() {

	for i := 0; i < server.ActionGoroutineNum; i++ {
		number := i
		go func() {
			for {
				select {
				case req := <-server.Queue:
					{
						if server.MsgActioner != nil {
							resp, err := server.MsgActioner(req.msg, number)
							req.err = err
							req.ch <- resp
						}
					}
				case <-server.MsgActionerExit[number]:
					return
				}
			}
		}()
	}
	for j := 0; j < server.ScheduledTaskGoroutineNum; j++ {
		number := j
		go func(number int) {
			server.TimedTasks[number].Task(number)
			for {
				select {
				case <-server.ScheduledTaskerExit[number]:
					return
				case <-time.After(server.TimedTasks[number].Time):
					server.TimedTasks[number].Task(number)
				}
			}
		}(number)
	}

	server.State = Started

	return
}

func (server *Server) stop() {
	if server.State == Started {
		for i := 0; i < server.ActionGoroutineNum; i++ {
			server.MsgActionerExit[i] <- 1
		}
		for i := 0; i < server.ScheduledTaskGoroutineNum; i++ {
			server.ScheduledTaskerExit[i] <- 1
		}
		server.State = Stopped
	}
}
