//Copyright (c) 2026 xde-dev

//Permission is hereby granted, free of charge, to any person obtaining a copy
//of this software and associated documentation files (the "Software"), to deal
//in the Software without restriction, including without limitation the rights
//to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//copies of the Software, and to permit persons to whom the Software is
//furnished to do so, subject to the following conditions:

//The above copyright notice and this permission notice shall be included in all
//copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

var cpuInfo string
var cpuCores int
var CpuLoad, TotalRam, UsingRam, RamLeft, TotalDiskSpace, DiskSpaceLeft float64
var HostName, Platform string

func GetCpuInfo() {
	i, _ := cpu.Info()
	if len(i) > 0 {
		cpuInfo = i[0].ModelName
	}
	cpuCores, _ = cpu.Counts(true)
	p, _ := cpu.Percent(time.Second, false)
	if len(p) > 0 {
		CpuLoad = p[0]
	}
}
func GetRamInfo() {
	v, _ := mem.VirtualMemory()
	TotalRam = float64(v.Total) / 1024 / 1024 / 1024
	UsingRam = float64(v.Used) / 1024 / 1024 / 1024
	RamLeft = float64(v.Available) / 1024 / 1024 / 1024
}
func GetDiskInfo() {
	DiskGeneralInfo, _ := disk.Usage("C:\\")
	TotalDiskSpace = float64(DiskGeneralInfo.Total) / 1024 / 1024 / 1024
	DiskSpaceLeft = float64(DiskGeneralInfo.Free) / 1024 / 1024 / 1024
}
func GetHostInfo() {
	GeneralHostInfo, _ := host.Info()
	HostName = GeneralHostInfo.Hostname
	Platform = GeneralHostInfo.Platform
}
func main() {
	for {
		fmt.Println("")
		fmt.Println("1. CPU info")
		fmt.Println("2. RAM info")
		fmt.Println("3. Disk info")
		fmt.Println("4. OS info")
		var choise string
		fmt.Println("Choose...")
		fmt.Scan(&choise)
		switch choise {
		case "1":
			GetCpuInfo()
			fmt.Println("your cpu info is:")
			fmt.Println(cpuInfo)
			fmt.Println("your cpu has this number of cores:")
			fmt.Println(cpuCores)
			fmt.Println("and your cpu current load is:")
			fmt.Println(CpuLoad)
			fmt.Println("_______________________________________________")

		case "2":
			GetRamInfo()
			fmt.Print("you have gbs of ram:")
			fmt.Println(TotalRam)
			fmt.Println("currently, your computer is using gbs of ram:")
			fmt.Println(UsingRam)
			fmt.Println("and your computer has gbs of ram left:")
			println(RamLeft)
			fmt.Println("_______________________________________________")
		case "3":
			GetDiskInfo()
			fmt.Println("your total disk space in gbs is:")
			fmt.Println(TotalDiskSpace)
			fmt.Println("and you have space left in gbs:")
			fmt.Println(DiskSpaceLeft)
			fmt.Println("_______________________________________________")
		case "4":
			GetHostInfo()
			fmt.Println("your computer's host name is:")
			fmt.Println(HostName)
			fmt.Println("and your system architecture is:")
			fmt.Println(Platform)
			fmt.Println("_______________________________________________")
		default:
			fmt.Println("invalid input, it may be a bug. try again")
			fmt.Println("_______________________________________________")
		}
	}
}
