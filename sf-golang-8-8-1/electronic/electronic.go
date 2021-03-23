package electronic

const (
	phoneTypeStation    = "station"
	phoneTypeSmartphone = "smartphone"
)

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type StationPhone interface {
	ButtonsCount() int
}

type Smartphone interface {
	OS() string
}

type phone struct {
	brand     string
	model     string
	phoneType string
}

func (p phone) Brand() string {
	return p.brand
}

func (p phone) Model() string {
	return p.model
}

func (p phone) Type() string {
	return p.phoneType
}
