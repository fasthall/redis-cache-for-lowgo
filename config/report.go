package config

import (
	"errors"
	"io/ioutil"
	"net"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Controller string `yaml:"controller"`
	NumDC      string `yaml:"num_dc"`
	ID         string `yaml:"id`
}

func readConfig(file string) (Config, error) {
	var config Config
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return Config{}, err
	}
	err = yaml.Unmarshal(raw, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func getHostIP() (string, error) {
	ifs, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, ifi := range ifs {
		if ifi.Name == "eth0" {
			addrs, err := ifi.Addrs()
			if err == nil {
				for _, addr := range addrs {
					if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
						if ipnet.IP.To4() != nil {
							return ipnet.IP.String(), nil
						}
					}
				}
			}
		}
	}
	return "", errors.New("IP not found")
}

func Report() (int, string, error) {
	config, err := readConfig("config/config.yaml")
	host, err := getHostIP()
	request, err := http.NewRequest("POST", "http://"+config.Controller+"/rediscache?host="+host+":6380", nil)
	if err != nil {
		return 0, "", err
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Connection", "Keep-Alive")

	var resp *http.Response
	resp, err = http.DefaultClient.Do(request)
	if err != nil {
		return 0, "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, "", err
	}

	if resp.StatusCode == http.StatusOK {
		return http.StatusOK, string(body), nil
	}
	return resp.StatusCode, string(body), errors.New("HTTP response code isn't 200")
}
