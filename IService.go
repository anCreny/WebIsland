package WebIsland

type IService interface {
	GetName() string
	New() IService
}
