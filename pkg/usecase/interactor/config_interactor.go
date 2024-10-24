package interactor

import (
	"fmt"
	"time"

	mcs "github.com/JhonX2011/GFAWBP/pkg/domain/models/cross_structs"
	mic "github.com/JhonX2011/GFAWBP/pkg/domain/models/internal_structs/configuration"
	"github.com/JhonX2011/GFAWBP/pkg/infrastructure/configuration"
	"github.com/JhonX2011/GFAWBP/pkg/usecase/presenter"
)

const (
	defaultMaxRetries int = 5
	defaultSleepTime  int = 300
)

type configInteractor struct {
	Config             configuration.Configuration
	GetConfigPresenter presenter.IGetConfigsPresenter
}

type IConfigInteractor interface {
	Reload(int) error
	GetConfigurations() (interface{}, error)
}

func NewConfigInteractor(c configuration.Configuration, p presenter.IGetConfigsPresenter) IConfigInteractor {
	return &configInteractor{c, p}
}

func (p *configInteractor) Reload(retry int) error {
	maxRetries, sleepTime := defaultMaxRetries, defaultSleepTime
	cfg := p.Config.GetConfig()

	if refreshMaxRetries := cfg.App.Config.RefreshMaxRetries; refreshMaxRetries > 0 {
		maxRetries = refreshMaxRetries
	}

	if refreshSleepTime := cfg.App.Config.RefreshSleepTime; refreshSleepTime > 0 {
		sleepTime = refreshSleepTime
	}

	if retry > maxRetries {
		return fmt.Errorf("unable to refresh configuration after [%d] attempts", maxRetries)
	}

	err := p.Config.LoadConfig()
	if err != nil && retry <= maxRetries {
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)
		return p.Reload(retry + 1)
	}

	return err
}

// GetConfigurations TODO: improve get configs method with a json response
func (p *configInteractor) GetConfigurations() (interface{}, error) {
	var appConfig map[string]interface{}

	if _, err := p.Config.LoadJSONProfile(mic.AppProfileName.String(), &appConfig); err != nil {
		return nil, err
	}

	stringKeys := map[string]string{"APP_NAME": "GFAWBP", "ENVIRONMENT": "local", "STACK": "Go"}
	//boolKeys := map[string]bool{"MODE_DEBUG": false}

	var configs []mcs.ConfigMember
	for configName, defaultValue := range stringKeys {
		if value, ok := appConfig[configName]; ok {
			if stringValue, ok2 := value.(string); ok2 {
				defaultValue = stringValue
			}
		}
		configs = append(configs, mcs.ConfigMember{Name: configName, Value: defaultValue})
	}

	//for configName, defaultValue := range boolKeys {
	//	if value, ok := appConfig[configName]; ok {
	//		if boolValue, ok2 := value.(bool); ok2 {
	//			defaultValue = boolValue
	//		}
	//	}
	//	member := mcs.ConfigMember{Name: configName, Value: strconv.FormatBool(defaultValue)}
	//	configs = append(configs, member)
	//}

	return p.GetConfigPresenter.ResponseGetConfigs(configs), nil
}
