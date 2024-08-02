package utils

import "time"

const ApiPrefixVersion string = "/api/v1"

func PrefixLogFilename() string {
	now := time.Now().Format("2006-01-02")
	return now + ".log"
}

const FolderLog string = "assets/logs/"

const DefaultLoggingFormat string = `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",` +
	`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}","form":"${form}"` +
	`"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` +
	`,"bytes_in":${bytes_in},"bytes_out":${bytes_out}}` + "\n"
