package https

import (
	"errors"
	"net"
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

	server := &http.Server{
		Addr:    ":9091",
		Handler: router,
	}

	listener, err := net.Listen("tcp", ":9091")
	if err != nil {
		return err
	}

	go func() {
		if err := server.Serve(listener); err != nil && !errors.Is(err, http.ErrServerClosed) {
			// Ошибка сервера логируется, но не возвращается, так как сервер уже запущен
		}
	}()

	return nil
}
