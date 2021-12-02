package goframe

type Service interface {
	Name() string
}

type BackgroundService interface {
	Service
	Run() error
}

type HTTPService interface {
	Service
	Prefix() string
}

type RESTService interface {
	HTTPService
	Middleware(JSONEndpoint) JSONEndpoint
	Endpoints() map[string]map[string]JSONEndpoint
}

type JSONEndpoint func(APIContext) error
