package presenter

import mcs "github.com/JhonX2011/GFAWBP/pkg/domain/models/cross_structs"

type getConfigsPresenter struct{}

type IGetConfigsPresenter interface {
	ResponseGetConfigs([]mcs.ConfigMember) mcs.Configurations
}

func NewGetConfigsPresenter() IGetConfigsPresenter {
	return &getConfigsPresenter{}
}

func (g getConfigsPresenter) ResponseGetConfigs(configs []mcs.ConfigMember) mcs.Configurations {
	return mcs.Configurations{
		Configs: configs,
	}
}
