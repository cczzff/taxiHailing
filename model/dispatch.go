package model

import (
	"sync"
)

type Dispatch struct {
	PalceChan  chan *Passenger // 乘客下单通道
	Passengers sync.Map        // 记录所有的用户
	Drivers    sync.Map        // 记录所有的司机
	ChatChan   chan ChatInfo   // 这里维护所有的聊天

}

// 聊天通道
type ChatInfo struct {
	Info     string
	Reciever interface{}
}


// 新建一个调度器
func NewDispatch() *Dispatch {
	disPatch := new(Dispatch)
	disPatch.PalceChan = make(chan *Passenger, 1024)
	disPatch.ChatChan = make(chan ChatInfo, 1024)
	// 开启消息转发器
	go disPatch.RunTransponder()
	return disPatch
}

//乘客下单
func (this *Dispatch) RegisterPlace(p *Passenger) {
	this.PalceChan <- p
	this.Passengers.Store(p, true)
}

// 司机抢单
func (this *Dispatch) RegisterRob(d *Driver) {
	this.Drivers.Store(d, true)
}


// 消息传递
func (this *Dispatch) RunTransponder() {
	for chat := range this.ChatChan {
		p, toPassenger := chat.Reciever.(*Passenger)
		if toPassenger {
			p.PChat.recieveChan <- chat.Info
		}


		d, toDriver := chat.Reciever.(*Driver)
		if toDriver {
			d.DChat.recieveChan <- chat.Info
		}

	}

}
