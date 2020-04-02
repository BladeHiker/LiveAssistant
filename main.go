package main

import (
	_"LiveAssistant/backend"
	"github.com/go-qamel/qamel"
	"os"
)

func main() {
	//Create Application
	app := qamel.NewApplication(len(os.Args), os.Args)
	app.SetApplicationDisplayName("Live Assistant")

	engine := qamel.NewEngine()
	engine.Load("qrc:/res/main.qml")

	// Exec app
	app.Exec()
	//v,_:=mem.VirtualMemory()
	//fmt.Println(v)
	//
	//x, _ := net.IOCounters(true)
	//for index, a := range x {
	//	fmt.Printf("%v:%v send:%v recv:%v\n", index, a.Name, a.BytesSent, a.BytesRecv)
	//}
	//
	//parts, err := disk.Partitions(true)
	//if err != nil {
	//	fmt.Printf("get Partitions failed, err:%v\n", err)
	//	return
	//}
	//for _, part := range parts {
	//	fmt.Printf("part:%v\n", part.String())
	//	diskInfo, _ := disk.Usage(part.Mountpoint)
	//	fmt.Printf("disk info:used:%v free:%v used:%v\n", diskInfo.UsedPercent, diskInfo.Free,diskInfo.UsedPercent)
	//}
	//
	//ioStat, _ := disk.IOCounters()
	//for k, v := range ioStat {
	//	fmt.Printf("%v:%v\n", k, v)
	//}
	//
	//for {
	//	percent, _ := cpu.Percent(time.Second, false)
	//	fmt.Printf("cpu percent:%v\n", percent)
	//}
}
