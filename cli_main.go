package main

import (
	"fmt"
	"os"
	"github.com/urfave/cli"
	"encoding/json"
)

func main() {
	app := cli.NewApp()
	app.Name = "weather-cli"
	app.Usage = "天气预报小程序"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "city, c",
			Value: "上海",
			Usage: "城市中文名",
		},
		cli.StringFlag{
			Name:  "day, d",
			Value: "今天",
			Usage: "可选: 今天, 昨天, 预测",
		},
	}

	app.Action = func(c *cli.Context) error {
		city := c.String("city")
		day := c.String("day")

		var body, err = Request(apiUrl + city)
		if err != nil {
			fmt.Printf("err was %v", err)
			return nil
		}

		var r Response
		err = json.Unmarshal([]byte(body), &r)
		if err != nil {
			fmt.Printf("\nError message: %v", err)
			return nil
		}
		if r.Status != 200 {
			fmt.Printf("获取天气API出现错误, %s", r.Message)
			return nil
		}
		Print(day, r)
		return nil
	}
	app.Run(os.Args)
}
