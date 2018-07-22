package model

import (
	"fmt"
)

type Passenger struct {
	Id          string
	Name        string
	Dispatch    *Dispatch
	PChat *ChatModel
	Driver *Driver
}

func NewPassenger(name string, dispatch *Dispatch) *Passenger {
	p := &Passenger{
		Dispatch:dispatch,
		Name:name,
		PChat:&ChatModel{
			sendChan: make(chan string, 1024),
			recieveChan: make(chan string, 1024),
		},
	}
	// 开启消息接收
	go p.Recieve()
	return p



}

func (p *Passenger) Place() {
	p.Dispatch.RegisterPlace(p)
	fmt.Println("place: ", p.Name)
}


func (p *Passenger) Recieve() {
	for chat := range p.PChat.recieveChan {
		fmt.Println(p.Name, " recieve ", chat)
	}

}

func (p *Passenger) Send(info string) {
	if p.Driver == nil {
		fmt.Println(p.Name, " 没有通话人")
		return
	}


	p.Dispatch.ChatChan <- ChatInfo{
		Info:     info,
		Reciever: p.Driver,
	}

	fmt.Println(p.Name, " send ", info)
}
