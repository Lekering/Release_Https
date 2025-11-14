package https

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHandlers *HTTPHandlers
}

func NewHTTPServer(httpHandlers *HTTPHandlers) *HTTPServer {
	return &HTTPServer{
		httpHandlers: httpHandlers,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()
	router.Path("/tasks").Methods("POST").HandlerFunc(s.httpHandlers.HandlerCreateTask)
	router.Path("/tasks/{title}").Methods("GET").HandlerFunc(s.httpHandlers.HandlerGetTask)
	router.Path("/tasks").Methods("GET").Queries("completed", "false").HandlerFunc(s.httpHandlers.HandlerGetUncompletedTask)
	router.Path("/tasks").Methods("GET").HandlerFunc(s.httpHandlers.HandlerGetAllTask)
	router.Path("/tasks/{title}").Methods("PATCH").HandlerFunc(s.httpHandlers.HandlerCompleteTask)
	router.Path("/tasks/{title}").Methods("DELETE").HandlerFunc(s.httpHandlers.HandlerDeleteTask)

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}
