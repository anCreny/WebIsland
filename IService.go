package ServicesSystem

type IService interface {
	GetName() string
	New() IService
}
