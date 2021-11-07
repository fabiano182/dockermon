package controllers

import (
	"net/http"

	"github.com/fabiano182/dockermon/docker/api"
	"github.com/gin-gonic/gin"
)

func HelloFromDockermon(c *gin.Context) {
	c.String(http.StatusOK, "Hello from Dockermon :D")
}

func ListContainers(c *gin.Context) {
	container := api.ContainerList()

	c.String(http.StatusOK, container)
}
