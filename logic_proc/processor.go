package logic_proc

import (
	"database/sql"

	"github.com/activate_server/protocol"
)

type Processor struct {
	db *sql.DB
}

func NewProcessor(db *sql.DB) *Processor {
	return &Processor{
		db,
	}
}

func (p *Processor) Login(usr, pwd string) int {
	rows, err := p.db.Query("SELECT * FROM user WHERE user_name=? AND password=?", usr, pwd)
	if err != nil {
		return protocol.DbFailed
	}

	defer rows.Close()

	if rows.Next() {
		return protocol.OK
	} else {
		return protocol.Reject
	}
}

func (p *Processor) DeviceActivate(serialNum string) int {
	rows, err := p.db.Query("SELECT * FROM active_serial_num WHERE serial_num=?", serialNum)
	if err != nil {
		return protocol.DbFailed
	}

	defer rows.Close()

	if rows.Next() {
		return protocol.OK
	} else {
		return protocol.Reject
	}
}

func (p *Processor) DeviceHeartbeat(serialNum string) int {

	return protocol.OK
}
