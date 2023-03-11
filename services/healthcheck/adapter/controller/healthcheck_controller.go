package controller

import (
	"custom-boilerplate/pkg/di"
	"custom-boilerplate/pkg/restapi"
	"custom-boilerplate/services/healthcheck/domain/dto/request"
	"custom-boilerplate/services/healthcheck/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
)

func init() {
	di.RegisterBean(NewHealthCheckController)
	di.GetContainer().Invoke(
		func(hcc *HealthCheckController) {
			group := restapi.RegisterGroup("/health-check")
			group.RegisterGet("/ping", hcc.Ping)
			group.RegisterPost("/readiness", hcc.Readiness)
		})
}

type HealthCheckController struct {
	healthCheckUseCase usecase.HealthCheckUseCase
}

func NewHealthCheckController(healthCheckUseCase usecase.HealthCheckUseCase) *HealthCheckController {
	return &HealthCheckController{
		healthCheckUseCase: healthCheckUseCase,
	}
}

func (hcc *HealthCheckController) Ping(c *gin.Context) {
	message, err := hcc.healthCheckUseCase.Ping(c)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, message)
}

func (hcc *HealthCheckController) Readiness(c *gin.Context) {
	req := &request.ReadyRequest{}
	if err := c.BindJSON(req); err != nil {
		fmt.Println(err.Error())
		c.JSON(400, err)
		return
	}
	res, err := hcc.healthCheckUseCase.Readiness(c, req)
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, res)
}
