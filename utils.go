package vonage

import "runtime"

func GetVersion() string {
	return "0.12.1"
}

func GetUserAgent() string {
	user_agent := "vonage-go/" + GetVersion() + " Go/" + runtime.Version()
	return user_agent
}
