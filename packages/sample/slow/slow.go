package slow

import "time"

func Main(args map[string]interface{}) map[string]interface{} {
	time.Sleep(time.Minute)

	msg := make(map[string]interface{})
	msg["body"] = "Slept for 1 minute."
	return msg
}
