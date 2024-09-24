package configuration

type RefreshConfig struct {
	RefreshMaxRetries int `json:"refresh_max_retries"`
	RefreshSleepTime  int `json:"refresh_sleep_time"`
}
