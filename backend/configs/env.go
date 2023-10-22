package configs

var Env *MainEnv

type serverEnv struct {
	HostURL          string
	InternalHost     string
	CoreInternalHost string
	Port             int
	CSRFOrigin       string
	DataDir          string
	CoreAPIKey       string
}

type MainEnv struct {
	Server serverEnv
}
