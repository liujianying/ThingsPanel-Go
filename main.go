package main

import (
	_ "ThingsPanel-Go/initialize/log"

	_ "ThingsPanel-Go/modules/dataService"

	_ "ThingsPanel-Go/initialize/cache"
	_ "ThingsPanel-Go/initialize/psql"

	_ "ThingsPanel-Go/initialize/cron"
	_ "ThingsPanel-Go/initialize/send_message"
	_ "ThingsPanel-Go/initialize/session"
	_ "ThingsPanel-Go/initialize/validate"
	_ "ThingsPanel-Go/routers"

	services "ThingsPanel-Go/services"
	"fmt"
	"time"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"

	_ "ThingsPanel-Go/grpc/server"

	_ "ThingsPanel-Go/cron"
)

var Ticker *time.Ticker

func main() {
	// 读取服务器信息
	Ticker = time.NewTicker(time.Millisecond * 5000)
	go func() {
		for t := range Ticker.C {
			fmt.Println(t)
			var ResourcesService services.ResourcesService
			percent, _ := cpu.Percent(time.Second, false)
			cpu_str := fmt.Sprintf("%.2f", percent[0])
			memInfo, _ := mem.VirtualMemory()
			mem_str := fmt.Sprintf("%.2f", memInfo.UsedPercent)
			currentTime := fmt.Sprint(time.Now().Format("2006-01-02 15:04:05"))
			ResourcesService.Add(cpu_str, mem_str, currentTime)
		}
	}()
	// 系统初始化的images
	beego.SetStaticPath("/files/init-images", "others/init_images")
	// 静态文件路径
	beego.SetStaticPath("/files", "files")
	beego.Run()
}
