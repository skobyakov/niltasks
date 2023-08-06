package controller

type Service interface {
	GetList()
}

type Controller struct {
	service Service
}

func New(s Service) *Controller {
	return &Controller{service: s}
}

func (c *Controller) GetList() {
	c.service.GetList()
}
