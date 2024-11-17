package route

import (
	"kompre/api/handler"
	"kompre/api/repository"
	"kompre/api/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Register(db *gorm.DB, e *echo.Echo) {
	repo := repository.NewCrudRepository(db)
	service := service.NewCrudService(repo)
	h := handler.NewCrudHandler(service)

	user := e.Group("/crud")
	user.POST("/create", h.Create)
	user.GET("/get", h.Get)
	user.PUT("/update/:id", h.Update)
	user.DELETE("/delete/:id", h.Delete)
}
