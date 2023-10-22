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

type databaseEnv struct {
	Host     string // Database host
	Port     int    // Database port
	User     string // Database user
	Password string // Database password
	DBName   string // Database name
}

type fileStoreEnv struct {
	USE_S3          bool
	S3              s3Env
	FileSizeLimitMB int
	TempDataDir     string
}

type s3Env struct {
	AcessKeyID      string
	SecretAccessKey string
	Bucket          string
	Region          string
}

type MainEnv struct {
	Server   serverEnv
	File     fileStoreEnv
	Database databaseEnv
}
