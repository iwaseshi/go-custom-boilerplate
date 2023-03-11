package usecase

import (
	"custom-boilerplate/services/healthcheck/domain/dto/request"
	"custom-boilerplate/services/healthcheck/domain/dto/response"

	"github.com/gin-gonic/gin"
)

type HealthCheckUseCase interface {
	Ping(c *gin.Context) (*string, error)
	Readiness(c *gin.Context, req *request.ReadyRequest) (*response.ReadyResponse, error)
}
