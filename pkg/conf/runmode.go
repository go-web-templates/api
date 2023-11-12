package conf

import "os"

func envOrDefault(env string, defValue string) string {
	value := os.Getenv(env)
	if value != "" {
		return value
	}
	return defValue
}

const (
	RUN_MODE_DEVEL   = "development"
	RUN_MODE_RELEASE = "release"
)

var RunMode = envOrDefault("RUN_MODE", RUN_MODE_DEVEL)
