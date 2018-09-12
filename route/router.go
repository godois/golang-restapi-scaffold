package route

import (
	"log"

	"github.com/ant0ine/go-json-rest/rest"
	"github.com/godois/golang-restapi-scaffold/handler"
)

// InitRouter is a router of this API
func InitRouter() *rest.Api {

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)

	router, err := rest.MakeRouter(
		rest.Get("/reminders", handler.GetAllReminders),
		rest.Post("/reminders", handler.PostReminder),
		rest.Put("/reminders", handler.PutReminder),
		rest.Get("/reminders/:id", handler.GetReminder),
		rest.Delete("/reminders/:id", handler.DeleteReminder),
	)

	if err != nil {
		log.Fatal(err)
	}

	api.SetApp(router)

	return api
}
