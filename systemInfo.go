package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

// Struct to store system info
type SystemInfo struct {
	HostName     string
	CPUName      string
	DiskCapacity uint64
	DiskUage     uint64
	DiskFree     uint64
	RamCapacity  uint64
	RamAvailable uint64
}

func getSystemInfo(data *SystemInfo) {
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

// Function to Jsonify data nd write it to a file
func saveData(data *SystemInfo) {
	jsonify, _ := json.Marshal(data)
	_ = os.WriteFile("output.json", []byte(jsonify), 0644)
}

// Homepage Route
func homePageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// fmt.Fprintf(w, "Welcome to the HomePage!")
	// fmt.Println("Endpoint Hit: homePage")
	w.Write([]byte("<h1>Hello World!</h1>"))
}

// Route
func getDataHandler(w http.ResponseWriter, r *http.Request) {
	data := SystemInfo{}
	getSystemInfo(&data)
	saveData(&data)
	json.NewEncoder(w).Encode(data)
}

// Request handler
func handleRequests(mux *http.ServeMux) {
	mux.HandleFunc("/", homePageHandler)
	mux.HandleFunc("/api", getDataHandler)
	log.Fatal(http.ListenAndServe(":3000", mux))
}

func main() {
	mux := http.NewServeMux()
	handleRequests(mux)
}
