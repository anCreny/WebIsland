package ServicesSystem

import "fmt"

type Service struct {
	serviceSample    IService
	serviceReference *IService
}

type Handler struct {
	transient map[string]IService
	scoped    map[string]*Service
	singleton map[string]*Service

	servicesUpdated chan int
}

func (this *Handler) StartRequest() {
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

func NewHandler(servicesUpdateChan chan int) *Handler {
	return &Handler{
		transient:       map[string]IService{},
		scoped:          map[string]*Service{},
		singleton:       map[string]*Service{},
		servicesUpdated: servicesUpdateChan,
	}
}

func (this *Handler) StartListening() {
	for {
		select {
		case <-this.servicesUpdated:
			this.StartRequest()
		}
	}
}

func (this *Handler) GetTransient(name string) IService {
	if service, ok := this.transient[name]; ok {
		return service
	}

	return nil
}

func (this *Handler) AddTransient(service IService) {
	var name = service.GetName()
	if _, ok := this.transient[name]; !ok {
		this.transient[name] = service
	}
}

func (this *Handler) GetScoped(name string) *IService {
	if service, ok := this.scoped[name]; ok {
		return service.serviceReference
	}

	return nil
}

func (this *Handler) AddScoped(service IService) {
	var name = service.GetName()
	if _, ok := this.scoped[name]; !ok {
		this.scoped[name] = &Service{service, nil}
	}
}

func (this *Handler) GetSingleton(name string) *IService {
	if service, ok := this.singleton[name]; ok {
		return service.serviceReference
	}

	return nil
}

func (this *Handler) AddSingleton(service IService) {
	var name = service.GetName()
	if _, ok := this.singleton[name]; !ok {
		this.singleton[name] = &Service{service, nil}
	}
}
