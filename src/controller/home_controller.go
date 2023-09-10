package controller

import "github.com/myyrakle/gopring/src/service"

// @Controller(/)
type HomeController struct {
	service *service.HomeService
}

// @GetMapping("/")
func (this HomeController) Index() string {
	return this.service.GetHello()
}
