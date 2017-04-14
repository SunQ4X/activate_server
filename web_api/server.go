package web_api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/activate_server/logic_proc"
	"github.com/activate_server/protocol"
	"github.com/activate_server/webserver"
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

	server.Handle("/activate", http.HandlerFunc(server.Activate))
	server.Handle("/heartbeat", http.HandlerFunc(server.Heartbeat))

	return server
}

func (s *Server) Activate(w http.ResponseWriter, r *http.Request) {
	var resp protocol.ActiveResponse

	defer func() {
		response, err := json.Marshal(resp)
		if err != nil {
			fmt.Println("Marshal response err:", err)
		} else {
			fmt.Fprint(w, string(response))
		}
	}()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Active read body error:", err)
		resp.ErrorCode = protocol.ReadFailed
		return
	}

	var req protocol.ActiveRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println("Unmarshal active request error:", err)
		resp.ErrorCode = protocol.ParseFailed
		return
	}

	resp.ErrorCode = s.processor.DeviceActivate(string(req.SerialNum))
}

func (s *Server) Heartbeat(w http.ResponseWriter, r *http.Request) {
	var resp protocol.HeartbeatResponse

	defer func() {
		response, _ := json.Marshal(resp)
		fmt.Fprint(w, string(response))
	}()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Heartbeat read body error:", err)
		resp.ErrorCode = protocol.ReadFailed
		return
	}

	var req protocol.HeartbeatRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println("Unmarshal heartbeat request error:%v", err)
		resp.ErrorCode = protocol.ParseFailed
		return
	}

	resp.ErrorCode = s.processor.DeviceHeartbeat(string(req.SerialNum))
}
