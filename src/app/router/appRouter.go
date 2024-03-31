package appRouter

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	attendantRouter "github.com/chronicler-org/core/src/attendant/router"
	authMiddleware "github.com/chronicler-org/core/src/auth/middleware"
	authRouter "github.com/chronicler-org/core/src/auth/router"
	customerRouter "github.com/chronicler-org/core/src/customer/router"
	customerCareRouter "github.com/chronicler-org/core/src/customerCare/router"
	managerRouter "github.com/chronicler-org/core/src/manager/router"
	productRouter "github.com/chronicler-org/core/src/product/router"
	tagRouter "github.com/chronicler-org/core/src/tag/router"
	teamRouter "github.com/chronicler-org/core/src/team/router"
)

func InitAppRouter(app *fiber.App, db *gorm.DB) {

	tagController, tagService := tagRouter.InitTagModule(db)
	productController, _ := productRouter.InitProductModule(db)
	teamController, teamService := teamRouter.InitTeamModule(db)
	managerController, managerService := managerRouter.InitManagerModule(db, teamService)
	customerController, customerService := customerRouter.InitCustomerModule(db, tagService)
	attendantController, attendantService := attendantRouter.InitAttendantModule(db, teamService)
	authRouterController, _ := authRouter.InitAuthModule(db, managerService, attendantService)

	customerCareController, _ := customerCareRouter.InitCustomerCareModule(db, customerService, teamService)

	authRouter.InitAuthRouter(app, authRouterController)

	app.Use(authMiddleware.WithAuth(managerService, attendantService))
	tagRouter.InitTagRouter(app, tagController)
	teamRouter.InitTeamRouter(app, teamController)
	productRouter.InitProductRouter(app, productController)
	managerRouter.InitManagerRouter(app, managerController)
	customerRouter.InitCustomerRouter(app, customerController)
	attendantRouter.InitAttendantRouter(app, attendantController)
	customerCareRouter.InitCustomerCareRouter(app, customerCareController)
}
