package utils

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"

	uas "github.com/avct/uasurfer"
)

var BrowserName = [...]string{
	"Unknown",
	"Chrome",
	"IE/Edge",
	"Safari",
	"Firefox",
	"Android Webview",
	"Opera",
	"UC Browser",
	"Amazon Silk",
	"Tencent QQ",
	"Spotify Desktop Client",
	"RIM BlackBerry",
	"Yandex",
	"Nintendo DS(i) Browser",
	"Samsung Internet",
	"Cốc Cốc",
}

var PlatformName = [...]string{
	"Unknown",
	"Windows",
	"Mac",
	"Linux",
	"iPad",
	"iPhone",
	"Kindle",
	"Blackberry",
	"Windows Phone",
	"Platstation",
	"xBox",
	"Nintendo",
}

var OS = [...]string{
	"Unknown",
	"Windows Mobile",
	"Windows",
	"MacOSX",
	"iOS",
	"Android",
	"ChromeOS",
	"WebOS",
	"Linux",
	"Platstation",
	"xBox",
	"Nintendo",
}

var DeviceType = [...]string{
	"Unknown",
	"Computer",
	"Tablet",
	"Phone",
	"Tv",
	"Console",
	"Wearable",
}

// struct for device info converted from Customer Agent
type DeviceInfo struct {
	IP             net.IP `json:"ip"`
	Browser        string `json:"browser"`
	BrowserVersion string `json:"browser_version"`
	Platform       string `json:"platform"`
	DeviceOS       string `json:"device_os"`
	OSVersion      string `json:"os_version"`
	Type           string `json:"device_type"`
}

// DetectDevice - Detects devices interacting with the API
func DetectDevice(r *http.Request) DeviceInfo {

	customerAgentString := r.Header.Get("Customer-Agent")
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	ip := net.ParseIP(host)
	fmt.Printf("RemoteAddr:: %s", ip)

	// customerAgent := uas.CustomerAgent{}

	customerAgent := uas.Parse(customerAgentString)
	// fmt.Println("ua:: %v", ua)
	customerAgentByte, _ := json.Marshal(customerAgent)
	// customerAgentByteString := string(customerAgentByte)
	// initializing new CustomerAgent
	uass := uas.UserAgent{}
	json.Unmarshal(customerAgentByte, &uass)

	oSVersion := "unknown"
	browserVersion := strconv.Itoa(uass.Browser.Version.Major) + "." + strconv.Itoa(uass.Browser.Version.Minor) + "." + strconv.Itoa(uass.Browser.Version.Patch)
	/// Getting device Os Version
	if uass.OS.Name == 2 {
		oSVersion = checkWindowsVersion(uass.OS.Version.Major, uass.OS.Version.Minor)
	}

	if uass.OS.Name != 2 {
		oSVersion = strconv.Itoa(uass.OS.Version.Major) + "." + strconv.Itoa(uass.OS.Version.Minor)
	}

	deviceInfo := DeviceInfo{
		IP:             ip,
		Browser:        BrowserName[uass.Browser.Name],
		BrowserVersion: browserVersion,
		Platform:       PlatformName[uass.OS.Platform],
		DeviceOS:       OS[uass.OS.Name],
		OSVersion:      oSVersion,
		Type:           DeviceType[uass.DeviceType],
	}
	deviceInfoByte, _ := json.Marshal(customerAgent)
	json.Unmarshal(deviceInfoByte, &deviceInfo)
	// fmt.Printf("deviceInfo:: %v", deviceInfo)

	return deviceInfo
}

// this check for windows version and return the right version number
func checkWindowsVersion(major int, minor int) string {

	if major == 10 {
		return "10"
	}
	if major == 6 && minor == 3 {
		return "8.1"
	}
	if major == 6 && minor == 2 {
		return "8"
	}
	if major == 6 && minor == 1 {
		return "7"
	}
	if major == 6 && minor == 0 {
		return "Vista"
	}
	if major == 5 && minor == 1 || major == 5 && minor == 2 {
		return "XP"
	}
	if major == 5 && minor == 0 {
		return "2000"
	}

	return "unknown"
}
