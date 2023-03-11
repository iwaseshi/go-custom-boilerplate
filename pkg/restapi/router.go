package restapi

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

type groups struct {
	Path  string
	Group *gin.RouterGroup
}

func RegisterGroup(groupPath string) *groups {
	return &groups{
		groupPath,
		router.Group(groupPath),
	}
}

func (g *groups) RegisterGet(getPath string, fun gin.HandlerFunc) {
	g.Group.GET(getPath, fun)
}

func (g *groups) RegisterPost(postPath string, fun gin.HandlerFunc) {
	g.Group.POST(postPath, fun)
}

func Run() error {
	return router.Run(":8081")
}
