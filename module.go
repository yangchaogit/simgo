package unis

type Module interface {
	Initialize() error
	Uninitialize() error
}
