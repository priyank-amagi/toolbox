package todos

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/priyank-amagi/toolbox/lib/http_helper"
)

const (
	todosRoute = "/todos"
)

//go:generate mockery --name ITodoSvc --filename=todos.go
type ITodoSvc interface {
	GetTodos() (todos []Todo, err error)
}

type TodoSvc struct {
	hostname string
	httpSvc  http_helper.IHttpSvc
}

var _ ITodoSvc = (*TodoSvc)(nil)

func NewTodoSvc(hostname string, httpSvc http_helper.IHttpSvc) *TodoSvc {
	return &TodoSvc{hostname: hostname, httpSvc: httpSvc}
}

func (tS *TodoSvc) GetTodos() ([]Todo, error) {
	res, err := http_helper.SendHTTPRequest(tS.httpSvc, "GET", tS.hostname, todosRoute, map[string]string{}, []string{}, "")
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK {
		defer res.Body.Close()
		resBody, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		var todos []Todo
		err = json.Unmarshal(resBody, &todos)
		if err != nil {
			return nil, err
		}
		return todos, nil
	}
	return nil, fmt.Errorf("failed to fetch todos from api with status_code: %d", res.StatusCode)
}
