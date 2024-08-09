package utils

import "time"

const ApiPrefixVersion string = "/api/v1"

func PrefixLogFilename() string {
	now := time.Now().Format("2006-01-02")
	return now + ".log"
}

const FolderLog string = "assets/logs/"
