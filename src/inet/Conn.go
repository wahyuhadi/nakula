package inet

import (
	"encoding/json"
	"net"
	"net/http"
	"time"
)

func Check() error {
	// Set to google DNS
	hostName := "google.com"
	// Set Port
	portNum := "80"
	// Delayed time
	seconds := 5
	// Set Timeout for waiting connection
	timeOut := time.Duration(seconds) * time.Second

	_, err := net.DialTimeout("tcp", hostName+":"+portNum, timeOut)

	if err != nil {
		return err
	}
	return nil

}

var Client = &http.Client{Timeout: 10 * time.Second}

// Get Daya from server API
func GetData(url string, target interface{}) error {
	r, err := Client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
