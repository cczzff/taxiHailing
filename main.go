package main

import (
	"taxiHailing/model"
	"time"
)

func main() {

	dispatch := model.NewDispatch()

	p1 := model.NewPassenger("P1", dispatch)
	p2 := model.NewPassenger("P2", dispatch)
	p3 := model.NewPassenger("P3", dispatch)
	p4 := model.NewPassenger("P4", dispatch)
	p5 := model.NewPassenger("P5", dispatch)
	p6 := model.NewPassenger("P6", dispatch)
	p7 := model.NewPassenger("P7", dispatch)
	p8 := model.NewPassenger("P8", dispatch)
	p9 := model.NewPassenger("P9", dispatch)
	p10 := model.NewPassenger("P10", dispatch)

	p1.Place()
	p2.Place()
	p3.Place()
	p4.Place()
	p5.Place()
	p6.Place()
	p7.Place()
	p8.Place()
	p9.Place()
	p10.Place()


	d1 := model.NewDriver("D1", dispatch)
	d2 := model.NewDriver("D2", dispatch)
	d1.Robbing()
	d2.Robbing()
	time.Sleep(time.Second)
	d1.Send("hello ")
	d1.Send("mi fans ")

	p1.Send("hello, thank you")
	p2.Send("thank you very much")
	time.Sleep(time.Second)
}
