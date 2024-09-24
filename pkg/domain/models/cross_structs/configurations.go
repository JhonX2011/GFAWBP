package crossstructs

type Configurations struct {
	Configs []ConfigMember `json:"configs"`
}

type ConfigMember struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
