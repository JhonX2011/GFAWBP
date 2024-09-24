package configuration

type profile uint8

// -- App Profiles
const (
	AppProfileName profile = iota
	ServicesProfileName
	DatabasesProfileName
	DatabaseQueriesProfileName
)

func (s profile) String() string {
	switch s {
	case AppProfileName:
		return "app"
	case ServicesProfileName:
		return "services"
	case DatabasesProfileName:
		return "databases"
	case DatabaseQueriesProfileName:
		return "database_queries"
	default:
		return "unknown"
	}
}

// Feature Flags
const ()

// -- Services HTTP Rest Pools
const ()

// -- Database names
const ()

// -- Database queries
const ()
