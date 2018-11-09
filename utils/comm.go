package utils

import (
	"fmt"
)

func CacuCurrentConfigFile() string {
	cfgPath := "config/config.ini"
	flag, err := CheckFileExist("config/config.local.ini")
	if err != nil {
		fmt.Println("cacuCurrentConfigFile error: ", err)
	}
	if flag {
		cfgPath = "config/config.local.ini"
	}
	return cfgPath
}
