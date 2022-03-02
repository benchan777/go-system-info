package main

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

// Struct to store system info
type System_Info struct {
	HostName     string
	CPUName      string
	DiskCapacity uint64
	DiskUage     uint64
	DiskFree     uint64
	RamCapacity  uint64
	RamAvailable uint64
}

func getSystemInfo(data *System_Info) {
	// Get cpu, disk, and host info
	cpuInfo, _ := cpu.Info()
	hostInfo, _ := host.Info()
	diskInfo, _ := disk.Usage("\\")
	ramInfo, _ := mem.VirtualMemory()

	// Store data into struct
	data.HostName = hostInfo.Hostname
	data.CPUName = cpuInfo[0].ModelName
	data.DiskCapacity = diskInfo.Total / 1024 / 1024
	data.DiskUage = diskInfo.Used / 1024 / 1024
	data.DiskFree = diskInfo.Free / 1024 / 1024
	data.RamCapacity = ramInfo.Total / 1024 / 1024
	data.RamAvailable = ramInfo.Available / 1024 / 1024
}
