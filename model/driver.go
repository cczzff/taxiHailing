package model

import (
	"fmt"
)

type ChatModel struct {
	sendChan    chan string
	recieveChan chan string
}

type Driver struct {
	Id       string
	Name     string
	Dispatch *Dispatch
	DChat    *ChatModel
	Passenger *Passenger
}

func NewDriver(name string, dispatch *Dispatch) *Driver {
	d := &Driver{
		Dispatch: dispatch,
		Name:     name,
		DChat: &ChatModel{
			sendChan:    make(chan string, 1024),
			recieveChan: make(chan string, 1024),
		},
	}
	// 开启消息接收
	go d.Recieve()
	return d
}

// 司机抢单
func (d *Driver) Robbing() {
	d.Dispatch.RegisterRob(d)
	go func() {
		p := <-d.Dispatch.PalceChan
		d.Passenger = p
		fmt.Println(d.Name, "  接单： ", p.Name)
	}()
}


func (d *Driver) Send(info string) {
	if d.Passenger == nil {
		fmt.Println(d.Name, " 没有通话人")
		return
	}


	d.Dispatch.ChatChan <- ChatInfo{
		Info:     info,
		Reciever: d.Passenger,
	}

	fmt.Println(d.Name, " send ", info)
}

func (d *Driver) Recieve() {
	for chat := range d.DChat.recieveChan {
		fmt.Println(d.Name, " recieve ", chat)
	}

}