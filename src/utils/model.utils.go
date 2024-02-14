package utils

type ResponseStandard struct {
	CorrelationId string        `json:"correlationId"`
	Code          string        `json:"resultCode"`
	Success       string        `json:"success,omitempty"`
	Message       string        `json:"message"`
	Result        interface{}   `json:"data,omitempty"`
	ErrorMessage  string        `json:"error,omitempty"`
	Error         *[]ErrorParam `json:"errorParam,omitempty"`
	Errors        *[]Errors     `json:"errors,omitempty"`
}

type ErrorParam struct {
	Param   string `json:"param"`
	Message string `json:"message"`
}

type Errors struct {
	ErrorCode string `json:"errorCode"`
	ErrorType string `json:"errorType"`
	MessageTh string `json:"messageTh,omitempty"`
	MessageEn string `json:"messageEn,omitempty"`
}

// Environment environment
type Environment struct {
	Build    *Build            `mapstructure:"build"`
	Service  *Service          `mapstructure:"service"`
	Log      *Log              `mapstructure:"log"`
	Endpoint *[]Endpoint       `mapstructure:"endpoint"`
	Database *[]Database       `mapstructure:"db"`
	Secret   map[string]string `mapstructure:"secret,omitempty"`
	Notify   map[string]string `mapstructure:"notify,omitempty"`
}

type Build struct {
	Date   string `mapstructure:"date"`
	Number string `mapstructure:"number"`
}

type Service struct {
	Name            string `mapstructure:"name"`
	Port            string `mapstructure:"port"`
	Domain          string `mapstructure:"domain"`
	Endpoint        string `mapstructure:"endpoint"`
	InfoEndpoint    string `mapstructure:"infoEndpoint"`
	ProfileEndpoint string `mapstructure:"profileEndpoint"`
}

type Endpoint struct {
	Name               string   `mapstructure:"name"`
	Url                string   `mapstructure:"url"`
	Timeout            int      `mapstructure:"timeout"`
	InsecureSkipVerify bool     `mapstructure:"insecureSkipVerify"`
	Username           string   `mapstructure:"username"`
	Password           string   `mapstructure:"password"`
	Params             *[]Param `mapstructure:"params"`
}

type Log struct {
	Kibana *Kibana `mapstructure:"kibana"`
	Format string  `mapstructure:"format"`
}

type Kibana struct {
	Suffix string `mapstructure:"suffix"`
}

type Param struct {
	Name  string `mapstructure:"name"`
	Value string `mapstructure:"value"`
}

type Database struct {
	Type        string `mapstructure:"type"`
	Uri         string `mapstructure:"uri"`
	DbName      string `mapstructure:"dbName"`
	Username    string `mapstructure:"username"`
	Password    string `mapstructure:"password"`
	MaxOpenConn int    `mapstructure:"dbmaxopenconn"`
	MaxIdleConn int    `mapstructure:"dbmaxidleconn"`
}

type SystemConfig struct {
	Build          Build             `mapstructure:"build"`
	Service        SystemService     `mapstructure:"systemService"`
	Log            SystemLog         `mapstructure:"log"`
	Database       SystemDatabase    `mapstructure:"database"`
	Secret         map[string]string `mapstructure:"secret,omitempty"`
	Notify         map[string]string `mapstructure:"notify,omitempty"`
	ApiConfig      EtcdApiConfig
	Authentication JWTConfig `mapstructure:"jwt,omitempty"`
}

type JWTConfig struct {
	SecretKey string `mapstructure:"secret"`
}

type SystemService struct {
	Name            string `mapstructure:"name"`
	Endpoint        string `mapstructure:"endpoint"`
	Port            string `mapstructure:"port"`
	InfoEndpoint    string `mapstructure:"infoEndpoint"`
	ProfileEndpoint string `mapstructure:"profileEndpoint"`
}

type SystemLog struct {
	Kibana EtcdKibana `mapstructure:"kibana"`
	Format string     `mapstructure:"format"`
}
type EtcdKibana struct {
	Suffix string `mapstructure:"suffix"`
}

type SystemDatabase struct {
	MongoDB   EtcdDatabase   `mapstructure:"mongodb"`
	Redis     EtcdDatabase   `mapstructure:"redis"`
	OracleDB  EtcdDatabase   `mapstructure:"oracledb"`
	Sqlserver EtcdDatabase   `mapstructure:"sqlserver"`
	AppRedis  []EtcdDatabase `mapstructure:"appRedis"`
	Mysql     EtcdDatabase   `mapstructure:"mysql"`
}

type EtcdDatabase struct {
	Uri       string `mapstructure:"uri"`
	DbName    string `mapstructure:"dbName"`
	Encrypt   string `mapstructure:"encrypt"`
	UserName  string `mapstructure:"username"`
	Password  string `mapstructure:"password"`
	TnsString string `mapstructure:"tnsString"`
}

type EtcdApiConfig struct {
	Authentication     Authentication
	Url                string
	Timeout            int
	Endpoints          map[string]string
	Params             map[string]string
	InsecureSkipVerify bool
}

type Authentication struct {
	ApiKey string
	Bearer string
	Token  string
}

type SystemConfigDatabaseList struct {
	Build         Build                     `mapstructure:"build"`
	Service       SystemService             `mapstructure:"systemService"`
	Log           SystemLog                 `mapstructure:"log"`
	Database      map[string]EtcdDatabase   `mapstructure:"database"`
	DatabaseArray map[string][]EtcdDatabase `mapstructure:"databaseArray"`
	Secret        map[string]string         `mapstructure:"secret,omitempty"`
	Notify        map[string]string         `mapstructure:"notify,omitempty"`
	ApiConfig     EtcdApiConfig
}
