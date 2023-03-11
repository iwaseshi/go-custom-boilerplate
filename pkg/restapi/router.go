package restapi

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

type routerGroup struct {
	path  string
	group *gin.RouterGroup
}

func NewGroup(groupPath string) *routerGroup {
	return &routerGroup{
		groupPath,
		router.Group(groupPath),
	}
}

func (rg *routerGroup) RegisterGet(getPath string, fun gin.HandlerFunc) {
	rg.group.GET(getPath, fun)
}

func (rg *routerGroup) RegisterPost(postPath string, fun gin.HandlerFunc) {
	rg.group.POST(postPath, fun)
}

func Run() error {
	return router.Run(":8081")
}
