package activate_server

import (
	"database/sql"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/activate_server/logic_proc"
	"github.com/activate_server/web_admin"
	"github.com/activate_server/web_api"
	_ "github.com/go-sql-driver/mysql"
)

type ActivateServer struct {
	opts atomic.Value
	wrap sync.WaitGroup
}

func NewActivateServer(opts *Options) *ActivateServer {
	server := &ActivateServer{}

	server.opts.Store(opts)

	return server
}

func (s *ActivateServer) Run() {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", s.getOpts().DatabaseUsername, s.getOpts().DatabasePassword, s.getOpts().DatabaseHostAddr, s.getOpts().DatabaseName)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return
	}

	defer db.Close()

	processor := logic_proc.NewProcessor(db)

	s.Wrap(func() {
		web_admin.NewServer(s.getOpts().WebAdminServeAddr, processor).Run()
	})

	s.Wrap(func() {
		web_api.NewServer(s.getOpts().WebAPIServeAddr, processor).Run()
	})

	fmt.Println("Service started.")

	s.wrap.Wait()
}

func (s *ActivateServer) Wrap(cb func()) {
	s.wrap.Add(1)

	go func() {
		cb()
		s.wrap.Done()
	}()
}

func (s *ActivateServer) getOpts() *Options {
	return s.opts.Load().(*Options)
}
