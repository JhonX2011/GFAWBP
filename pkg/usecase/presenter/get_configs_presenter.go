package presenter

import mcs "github.com/JhonX2011/GFAWBP/pkg/domain/models/cross_structs"

type IGetConfigsPresenter interface {
	ResponseGetConfigs([]mcs.ConfigMember) mcs.Configurations
}
