package goframe

// assert service satisfies Service interface
var _ Service = &service{}

// assert httpService satisfies HTTPService interface
var _ HTTPService = &httpService{}
