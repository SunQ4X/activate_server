package web_admin

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/activate_server/logic_proc"
	"github.com/activate_server/protocol"
	"github.com/activate_server/webserver"
	_ "github.com/go-sql-driver/mysql"
)

type Server struct {
	*webserver.WebServer
	processor *logic_proc.Processor
}

func NewServer(addr string, proc *logic_proc.Processor) *Server {
	server := &Server{
		webserver.NewWebServer(addr),
		proc,
	}

	server.Handle("/login", http.HandlerFunc(server.Login))

	return server
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	methodUpper := strings.ToUpper(r.Method)

	switch methodUpper {
	case "GET":
		t, err := template.ParseFiles("./view/login.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t.Execute(w, nil)
	case "POST":
		var resp protocol.LoginResponse

		defer func() {
			response, _ := json.Marshal(resp)
			fmt.Fprint(w, string(response))
		}()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println("administrator login read body error:", err)
			resp.ErrorCode = protocol.ReadFailed
			return
		}

		var req protocol.LoginRequest
		err = json.Unmarshal(body, &req)
		if err != nil {
			fmt.Println("Unmarshal login request error:", err)
			resp.ErrorCode = protocol.ParseFailed
			return
		}

		resp.ErrorCode = s.processor.Login(string(req.Username), string(req.Password))
	}
}
