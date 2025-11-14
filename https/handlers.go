package https

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"restip/todo"
	"time"

	"github.com/gorilla/mux"
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

	//defer r.Body.Close()
	var taskDTO TaskDTO

	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		h.writeError(w, http.StatusBadRequest, err)
		return
	}

	if err := taskDTO.TaskDTOValidateToCreate(); err != nil {
		h.writeError(w, http.StatusBadRequest, err)
		return
	}

	task := todo.NewTask(taskDTO.Title, taskDTO.Description)
	if err := h.todoList.AddTask(task); err != nil {
		edto := NewErrorDTO(err)

		if errors.Is(err, todo.ErrTaskAlreadyExists) {
			http.Error(w, edto.ToString(), http.StatusConflict)
		} else {
			http.Error(w, edto.ToString(), http.StatusInternalServerError)
		}
		return
	}

	h.writeJSONResponse(w, http.StatusCreated, task)

}

func (h *HTTPHandlers) HandlerGetTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	task, err := h.todoList.GetTask(title)
	if err != nil {
		erdo := NewErrorDTO(err)

		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, erdo.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, erdo.ToString(), http.StatusInternalServerError)
		}
		return
	}

	h.writeJSONResponse(w, http.StatusOK, task)

}

func (h *HTTPHandlers) HandlerGetAllTask(w http.ResponseWriter, r *http.Request) {
	tasks := h.todoList.ListTasks()
	h.writeJSONResponse(w, http.StatusOK, tasks)
}

func (h *HTTPHandlers) HandlerGetUncompletedTask(w http.ResponseWriter, r *http.Request) {
	notComplitedTask := h.todoList.ListUncompletedTasks()
	h.writeJSONResponse(w, http.StatusOK, notComplitedTask)

}

func (h *HTTPHandlers) HandlerCompleteTask(w http.ResponseWriter, r *http.Request) {
	var completeDTO CompleteDTO
	if err := json.NewDecoder(r.Body).Decode(&completeDTO); err != nil {
		h.writeError(w, http.StatusBadRequest, err)
		return
	}

	title := mux.Vars(r)["title"]

	var (
		changedTask todo.Task
		err         error
	)

	if completeDTO.Complete {
		changedTask, err = h.todoList.CompleteTask(title)
	} else {
		changedTask, err = h.todoList.UncompleteTask(title)
	}
	if err != nil {
		erdo := NewErrorDTO(err)

		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, erdo.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, erdo.ToString(), http.StatusInternalServerError)
		}
		return
	}
	h.writeJSONResponse(w, http.StatusOK, changedTask)
}

func (h *HTTPHandlers) HandlerDeleteTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	if err := h.todoList.DeleteTask(title); err != nil {
		erdo := NewErrorDTO(err)

		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, erdo.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, erdo.ToString(), http.StatusInternalServerError)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *HTTPHandlers) writeError(w http.ResponseWriter, status int, err error) {
	errDTO := ErrorDTO{
		Message: err.Error(),
		Time:    time.Now(),
	}
	http.Error(w, errDTO.ToString(), status)
}

func (h *HTTPHandlers) writeJSONResponse(w http.ResponseWriter, status int, tasks any) {
	b, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(status)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response:", err)
		return
	}
}
