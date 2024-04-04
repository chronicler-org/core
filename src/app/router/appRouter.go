package appRouter

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	appMiddleware "github.com/chronicler-org/core/src/app/middleware"
	attendantRouter "github.com/chronicler-org/core/src/attendant/router"
	authMiddleware "github.com/chronicler-org/core/src/auth/middleware"
	authRouter "github.com/chronicler-org/core/src/auth/router"
	customerRouter "github.com/chronicler-org/core/src/customer/router"
	customerCareRouter "github.com/chronicler-org/core/src/customerCare/router"
	managerRouter "github.com/chronicler-org/core/src/manager/router"
	productRouter "github.com/chronicler-org/core/src/product/router"
	salesRouter "github.com/chronicler-org/core/src/sales/router"
	tagRouter "github.com/chronicler-org/core/src/tag/router"
	teamRouter "github.com/chronicler-org/core/src/team/router"
)

func InitAppRouter(app *fiber.App, db *gorm.DB, Validator *validator.Validate) {

	tagController, tagService := tagRouter.InitTagModule(db)
	productController, _ := productRouter.InitProductModule(db)
	teamController, teamService := teamRouter.InitTeamModule(db)
	managerController, managerService := managerRouter.InitManagerModule(db, teamService)
	customerController, customerService := customerRouter.InitCustomerModule(db, tagService)
	attendantController, attendantService := attendantRouter.InitAttendantModule(db, teamService)
	authRouterController, _ := authRouter.InitAuthModule(db, managerService, attendantService)

	customerCareController, customerCareService := customerCareRouter.InitCustomerCareModule(db, customerService, teamService)

	salesController, _ := salesRouter.InitSalesModule(db, customerCareService)

	validatorMiddleware := appMiddleware.Validate(Validator)
	authRouter.InitAuthRouter(app, authRouterController, validatorMiddleware)

	app.Use(authMiddleware.WithAuth(managerService, attendantService))

	tagRouter.InitTagRouter(app, tagController, validatorMiddleware)
	teamRouter.InitTeamRouter(app, teamController, validatorMiddleware)
	productRouter.InitProductRouter(app, productController, validatorMiddleware)
	managerRouter.InitManagerRouter(app, managerController, validatorMiddleware)
	customerRouter.InitCustomerRouter(app, customerController, validatorMiddleware)
	attendantRouter.InitAttendantRouter(app, attendantController, validatorMiddleware)
	customerCareRouter.InitCustomerCareRouter(app, customerCareController, validatorMiddleware)
	salesRouter.InitSalesRouter(app, salesController, validatorMiddleware)
}
