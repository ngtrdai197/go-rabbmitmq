package user

type Business interface {
}

type business struct {
}

func NewBusiness() Business {
	return &business{}
}
