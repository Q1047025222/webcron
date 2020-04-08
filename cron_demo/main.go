package main

import (
	"github.com/lisijie/cron"
	"log"
	"time"
)

func main() {
	log.Print("Starting...")

	c := cron.New()
	// 定时5秒，每5秒执行print5
	c.AddFunc("*/5 * * * * *", print5)
	// 定时15秒，每5秒执行print5
	c.AddFunc("*/15 * * * * *", print15)

	// 开始
	c.Start()
	defer c.Stop()

	// 定时器
	//t1 := time.NewTimer(time.Second * 10)
	//for {
	//	select {
	//	case <-t1.C:
	//		t1.Reset(time.Second * 10)
	//		print10()
	//	}
	//}

	t2 := time.NewTicker(time.Second * 10)
	defer t2.Stop()
	for {
		select {
		case <-t2.C:
			print10()
		}
	}
}

func print5() {
	log.Println("Run 5s cron")
}

func print10() {
	log.Println("Run 10s cron")
}

func print15() {
	log.Println("Run 15s cron")
}
