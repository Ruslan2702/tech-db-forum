package checkers

import (
	"github.com/bozaro/tech-db-forum/tests/client"
	"github.com/go-openapi/runtime"
	http_transport "github.com/go-openapi/runtime/client"
	"log"
	"net/http"
)

type Checker struct {
	// Имя текущей проверки.
	Name string
	// Описание текущей проверки.
	Description string
	// Функция для текущей проверки.
	FnCheck func(c *client.Forum)
	// Тесты, без которых проверка не имеет смысл.
	Deps []string
}

type CheckerTransport struct {
	t runtime.ClientTransport
}

func (self *CheckerTransport) Submit(operation *runtime.ClientOperation) (interface{}, error) {
	tracker := NewValidator(operation.Context)
	operation.Client = &http.Client{Transport: tracker}
	return self.t.Submit(operation)
}

var checks []Checker

func Register(checker Checker) {
	checks = append(checks, checker)
}

func RunCheck(check Checker) string {
	cfg := client.DefaultTransportConfig().WithHost("localhost:5000").WithSchemes([]string{"http"})
	transport := http_transport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	check.FnCheck(client.New(&CheckerTransport{transport}, nil))
	return "OK"
}

func Run() {
	for _, check := range checks {
		log.Printf("=== RUN:  %s", check.Name)
		r := RunCheck(check)
		log.Printf("--- DONE: %s (%s)", check.Name, r)
	}
}
