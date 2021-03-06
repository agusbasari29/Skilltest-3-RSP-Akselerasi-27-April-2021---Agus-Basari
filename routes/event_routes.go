package routes

import (
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/database"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/handler"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/helper"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/repository"
	"github.com/agusbasari29/Skilltest-RSP-Akselerasi-2-Backend-Agus-Basari/services"
	"github.com/gin-gonic/gin"
)

type EventRoutes struct{}

func (r EventRoutes) Route() []helper.Route {
	db := database.SetupDatabaseConnection()
	eventRepo := repository.NewEventRepository(db)
	trxRepo := repository.NewTransactionRepository(db)
	userRepo := repository.NewUserRepository(db)
	eventServices := services.NewEventServices(eventRepo)
	jwtServices := services.NewJWTService()
	trxServices := services.NewTransactionServices(trxRepo)
	userServices := services.NewUserServices(userRepo)
	eventHandler := handler.NewEventHandler(eventServices, jwtServices, trxServices, userServices)

	return []helper.Route{
		{
			Path:    "/events",
			Method:  "POST",
			Handler: []gin.HandlerFunc{eventHandler.CreateEvent},
		}, {
			Path:    "/events",
			Method:  "GET",
			Handler: []gin.HandlerFunc{eventHandler.GetAllEvent},
		}, {
			Path:    "/events",
			Method:  "PUT",
			Handler: []gin.HandlerFunc{eventHandler.UpdateEvent},
		}, {
			Path:    "/events",
			Method:  "DELETE",
			Handler: []gin.HandlerFunc{eventHandler.DeleteEvent},
		}, {
			Path:    "/purchase/:id",
			Method:  "GET",
			Handler: []gin.HandlerFunc{eventHandler.MakeEventPurchase},
		}, {
			Path:    "/event_detail/:id",
			Method:  "GET",
			Handler: []gin.HandlerFunc{eventHandler.GetEventDetail},
		}, {
			Path:    "/event_release",
			Method:  "GET",
			Handler: []gin.HandlerFunc{eventHandler.GetEventByReleaseStatus},
		},
	}
}
