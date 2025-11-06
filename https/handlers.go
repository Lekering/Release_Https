package https

import (
	"encoding/json"
	"net/http"
	"restip/todo"
	"time"
)

type HTTPHandlers struct {
	todoList *todo.List
}

func NewHTTPHandlers(todoList *todo.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todoList,
	}
}

func (h *HTTPHandlers) HandlerCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO
	
	if err := json.NewDecoder(r.Body).Decode(*&taskDTO); err != nil{
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time: time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}
}
10:52:06

func (h *HTTPHandlers) HandlerGetTask(w http.ResponseWriter, r *http.Request) {
	
}

func (h *HTTPHandlers) HandlerGetAllTask(w http.ResponseWriter, r *http.Request) {

}

func (h *HTTPHandlers) HandlerGetUncompletedTask(w http.ResponseWriter, r *http.Request) {

}

func (h *HTTPHandlers) HandlerCompleteTask(w http.ResponseWriter, r *http.Request) {

}

func (h *HTTPHandlers) HandlerDeleteTask(w http.ResponseWriter, r *http.Request) {

}