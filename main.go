package main

import (
	"github.com/gen2brain/beeep"
	"github.com/spf13/cast"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

func main() {

	app := cli.NewApp()

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "time,-t",
			Value: "15",
			Usage: "time",
		},
	}

	ch := make(chan int)
	intTime := 15
	app.Action = func(c *cli.Context) error {
		strTime := c.String("-t")
		intTime = cast.ToInt(strTime)

		go func() {
			minu := 0
			for {
				time.Sleep(time.Duration(time.Minute * 5))
				minu += 1 * 5
				ch <- minu
			}
		}()

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	ticker := time.NewTicker(time.Minute * time.Duration(intTime))
	beeep.Notify("开始", cast.ToString(intTime)+"分钟", "")
	for {
		select {
		case min := <-ch:
			beeep.Notify("时间 ", cast.ToString(min)+"分钟", "")
		case <-ticker.C:
			beeep.Alert("结束 ", cast.ToString(intTime)+"分钟", "")
			return
		}
	}

}
