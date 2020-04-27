package router

import (
	"fmt"
	"net/http"
	"task-api/src/infrastructure/middleware"
	"task-api/src/interfaces"
	"task-api/src/usecase"

	"github.com/julienschmidt/httprouter"
)

func logging(h interfaces.HttpHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		fmt.Println("[ACCESS] ", r.Method, r.URL, r.Host, r.RequestURI)
		h(w, r, params)
	}
}

func Handler(sqlhandler interfaces.SQLHandler, validator usecase.Validator) *httprouter.Router {
	middlewares := middleware.New(sqlhandler)

	userController := interfaces.NewUserController(sqlhandler, validator)
	projectController := interfaces.NewProjectController(sqlhandler, validator)
	taskController := interfaces.NewTaskController(sqlhandler, validator)

	router := httprouter.New()
	/* users API */
	router.POST("/signup", logging(userController.Singup))
	router.POST("/login", logging(userController.Login))
	router.GET("/users", logging(userController.Index))
	router.GET("/users/:id", logging(userController.Show))

	/* projects API */
	router.GET("/projects", logging(middlewares.Authenticate(projectController.Index)))
	router.POST("/projects", logging(middlewares.Authenticate(projectController.Create)))
	router.DELETE("/projects/:id", logging(middlewares.Authenticate(projectController.Delete)))

	/* task API */
	router.POST("/projects/:id/tasks", logging(middlewares.Authenticate(middlewares.RequiredWriteRole(taskController.Create))))
	router.PUT("/projects/:id/tasks/:task_id", logging(middlewares.Authenticate(middlewares.RequiredWriteRole(taskController.Update))))

	return router
}
