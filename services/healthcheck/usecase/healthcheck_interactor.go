package usecase

import (
	"custom-boilerplate/services/healthcheck/domain/dto/request"
	"custom-boilerplate/services/healthcheck/domain/dto/response"

	"custom-boilerplate/pkg/di"

	"github.com/gin-gonic/gin"
)

func init() {
	di.RegisterBean(NewHealthCheckInteractor)
}

type HealthCheckInteractor struct {
}

func NewHealthCheckInteractor() HealthCheckUseCase {
	return &HealthCheckInteractor{}

}

func (hci *HealthCheckInteractor) Ping(c *gin.Context) (*string, error) {
	message := "pong"
	return &message, nil
}

func (hci *HealthCheckInteractor) Readiness(c *gin.Context, req *request.ReadyRequest) (*response.ReadyResponse, error) {
	var res = response.ReadyResponse{
		Message: "yeah!!",
	}
	if req.Shout != "Are you ready?" {
		res.Message = "no!"
		return &res, nil
	}
	return &res, nil
}
