package v1alpha1

type Test struct {
	Type string `json:"type"`
}

type Case interface {
	Name() string
	Run() error
}
