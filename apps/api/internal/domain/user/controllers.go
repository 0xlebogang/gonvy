package user

type Controller interface{}

type controller struct {
	service Service
}

func NewController(s Service) Controller {
	return &controller{service: s}
}
