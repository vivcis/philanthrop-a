package routes

import (
	"gorm.io/gorm"
	"server/internal/adapters/api/controllers"
	"server/internal/adapters/repository"
	"server/internal/core/helpers"
	"server/internal/core/service"
)

func Injection(db *gorm.DB) {
	userRepository := repository.NewUser(db)
	userService := service.NewUserService(userRepository)

	Handler := controllers.NewHTTPHandler(userService)
	router := SetupRouter(Handler, userService)

	_ = router.Run(":" + helpers.Instance.Port)
}
