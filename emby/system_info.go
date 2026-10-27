package emby

import (
	"fmt"
	"net/http"
)

type SystemInfo struct {
	SystemUpdateLevel                    string          `json:"SystemUpdateLevel"`
	OperatingSystemDisplayName           string          `json:"OperatingSystemDisplayName"`
	HasPendingRestart                    bool            `json:"HasPendingRestart"`
	IsShuttingDown                       bool            `json:"IsShuttingDown"`
	HasImageEnhancers                    bool            `json:"HasImageEnhancers"`
	OperatingSystem                      string          `json:"OperatingSystem"`
	SupportsLibraryMonitor               bool            `json:"SupportsLibraryMonitor"`
	SupportsLocalPortConfiguration       bool            `json:"SupportsLocalPortConfiguration"`
	SupportsWakeServer                   bool            `json:"SupportsWakeServer"`
	WebSocketPortNumber                  int             `json:"WebSocketPortNumber"`
	CanSelfRestart                       bool            `json:"CanSelfRestart"`
	CanSelfUpdate                        bool            `json:"CanSelfUpdate"`
	CanLaunchWebBrowser                  bool            `json:"CanLaunchWebBrowser"`
	HttpServerPortNumber                 int             `json:"HttpServerPortNumber"`
	SupportsHttps                        bool            `json:"SupportsHttps"`
	HttpsPortNumber                      int             `json:"HttpsPortNumber"`
	HasUpdateAvailable                   bool            `json:"HasUpdateAvailable"`
	SupportsAutoRunAtStartup             bool            `json:"SupportsAutoRunAtStartup"`
	HardwareAccelerationRequiresPremiere bool            `json:"HardwareAccelerationRequiresPremiere"`
	WakeOnLanInfo                        []WakeOnLanInfo `json:"WakeOnLanInfo"`
	IsInMaintenanceMode                  bool            `json:"IsInMaintenanceMode"`
	LocalAddresses                       []string        `json:"LocalAddresses"`
	RemoteAddresses                      []string        `json:"RemoteAddresses"`
	ServerName                           string          `json:"ServerName"`
	Version                              string          `json:"Version"`
	Id                                   string          `json:"Id"`
}

type WakeOnLanInfo struct {
	MacAddress       string `json:"MacAddress"`
	BroadcastAddress string `json:"BroadcastAddress"`
	Port             int    `json:"Port"`
}

func (c *Client) GetSystemInfo() (*SystemInfo, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/System/Info", c.URL), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(c.ctx)

	res := SystemInfo{}
	if err := c.sendRequest(req, &res); err != nil {
		return nil, err
	}

	return &res, nil
}
