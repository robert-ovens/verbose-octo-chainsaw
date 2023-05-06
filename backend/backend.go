package backend

type Backend interface {
	GetInstances() ([]Instance, error)
}

type Instance struct {
	Image    string
	SwapSize string
	Type     string
	Label    string
	Id       string
}

type BackendImpl struct {
	Instances []Instance
}

func NewBackendImpl() BackendImpl {
	return BackendImpl{}
}

func (b BackendImpl) GetInstances() ([]Instance, error) {

	b.Instances = make([]Instance, 10)
	b.Instances[0] = Instance{
		Image:    "sdfs",
		SwapSize: "ghf",
		Type:     "hgf",
		Label:    "gf",
		Id:       "Sd",
	}
	return b.Instances, nil
}
