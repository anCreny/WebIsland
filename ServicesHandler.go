package WebIsland

import "fmt"

type Service struct {
	serviceSample    IService
	serviceReference *IService
}

type ServicesHandler struct {
	transient map[string]IService
	scoped    map[string]*Service
	singleton map[string]*Service
}

func (this *ServicesHandler) startRequest() {
	for _, value := range this.scoped {
		var reference = value.serviceSample.New()
		value.serviceReference = &reference
		fmt.Println(value)
	}

	for _, value := range this.singleton {
		if value.serviceReference == nil {
			var reference = value.serviceSample.New()
			value.serviceReference = &reference
		}
	}
}

func NewHandler(servicesUpdateChan chan int) *ServicesHandler {
	return &ServicesHandler{
		transient: map[string]IService{},
		scoped:    map[string]*Service{},
		singleton: map[string]*Service{},
	}
}

func (this *ServicesHandler) GetTransient(name string) IService {
	if service, ok := this.transient[name]; ok {
		return service
	}

	return nil
}

func (this *ServicesHandler) AddTransient(service IService) {
	var name = service.GetName()
	if _, ok := this.transient[name]; !ok {
		this.transient[name] = service
	}
}

func (this *ServicesHandler) GetScoped(name string) *IService {
	if service, ok := this.scoped[name]; ok {
		return service.serviceReference
	}

	return nil
}

func (this *ServicesHandler) AddScoped(service IService) {
	var name = service.GetName()
	if _, ok := this.scoped[name]; !ok {
		this.scoped[name] = &Service{service, nil}
	}
}

func (this *ServicesHandler) GetSingleton(name string) *IService {
	if service, ok := this.singleton[name]; ok {
		return service.serviceReference
	}

	return nil
}

func (this *ServicesHandler) AddSingleton(service IService) {
	var name = service.GetName()
	if _, ok := this.singleton[name]; !ok {
		this.singleton[name] = &Service{service, nil}
	}
}
