package server

type UtilityServer interface {
	Start() error
	Stop() error
}
