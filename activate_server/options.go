package activate_server

import (
	"encoding/json"
	"io/ioutil"
)

var (
	OPTIONS_CONFIG_FILE_PATH = "./options.cfg" //配置文件路径

	ADMIN_SERVE_ADDR_DEFAULT = "127.0.0.1:8088" //web_admin的服务地址
	API_SERVE_ADDR_DEFAULT   = "127.0.0.1:8089" //web_api的服务地址

	DATABASE_HOST_ADDR_DEFAULT = "127.0.0.1:3306"     //数据库地址
	DATABASE_NAME_DEFAULT      = "activate_server_db" //数据库名称
	DATABASE_USERNAME_DEFAULT  = "root"               //数据库用户名
	DATABASE_PASSWORD_DEFAULT  = "123456"             //数据库密码
)

type Options struct {
	WebAdminServeAddr string `json:"web_admin_serve_addr"`
	WebAPIServeAddr   string `json:"web_api_serve_addr"`
	DatabaseHostAddr  string `json:"database_host_addr"`
	DatabaseName      string `json:"database_name"`
	DatabaseUsername  string `json:"database_username"`
	DatabasePassword  string `json:"database_password"`
}

func NewOptions() *Options {
	return &Options{
		WebAdminServeAddr: ADMIN_SERVE_ADDR_DEFAULT,
		WebAPIServeAddr:   API_SERVE_ADDR_DEFAULT,
		DatabaseHostAddr:  DATABASE_HOST_ADDR_DEFAULT,
		DatabaseName:      DATABASE_NAME_DEFAULT,
		DatabaseUsername:  DATABASE_USERNAME_DEFAULT,
		DatabasePassword:  DATABASE_PASSWORD_DEFAULT,
	}
}

func (opts *Options) Store() {
	data, err := ioutil.ReadFile(OPTIONS_CONFIG_FILE_PATH)
	if err != nil {
		//文件不存在则创建文件并写入默认配置
		data, err := json.Marshal(opts)
		if err != nil {
			return
		}

		ioutil.WriteFile(OPTIONS_CONFIG_FILE_PATH, data, 0666)
	}

	json.Unmarshal(data, opts)
}
