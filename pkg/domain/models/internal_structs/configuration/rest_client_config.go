package configuration

import "fmt"

type RestClientConfig struct {
	BaseService
	Timeout        int    `json:"timeout"`
	MaxRetry       int    `json:"max_retry"`
	BackOffTime    int    `json:"back_off_time"`
	PathURL        string `json:"path_url"`
	TargetID       string `json:"target_id"`
	DialTimeout    int    `json:"dial_timeout"`
	EnableCache    bool   `json:"enable_cache"`
	DisableTimeout bool   `json:"disable_timeout"`
}

type HTTPRestPool struct {
	Services []RestClientConfig `json:"http_rest_pool"`
}

func (s HTTPRestPool) Get(serviceID string) (RestClientConfig, error) {
	for _, service := range s.Services {
		if service.ServiceID == serviceID {
			return service, nil
		}
	}

	return RestClientConfig{}, fmt.Errorf("service not found [%s]", serviceID)
}
