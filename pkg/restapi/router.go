package restapi

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

type routerGroup struct {
	Path  string
	Group *gin.RouterGroup
}

func NewGroup(groupPath string) *routerGroup {
	return &routerGroup{
		groupPath,
		router.Group(groupPath),
	}
}

func (rg *routerGroup) RegisterGet(getPath string, fun gin.HandlerFunc) {
	rg.Group.GET(getPath, fun)
}

func (rg *routerGroup) RegisterPost(postPath string, fun gin.HandlerFunc) {
	rg.Group.POST(postPath, fun)
}

func Run() error {
	return router.Run(":8081")
}
