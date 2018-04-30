package config

import (
	"os"
	"../cmd"
	"log"
	"encoding/json"
)

type config struct {
	// upload to file servers
	FileServerHosts []string `json:"file_server_hosts"`
	WorkDir string `json:"work_dir"`
	UploadFolder string `json:"upload_folder"`
	ListenPort string `json:"listen_port"`
	UploadMaxSize int64 `json:"upload_max_size"`
	ServerHost string `json:"server_host"`
}

var (
	Cfg *config
)

func init() {
	Cfg = new(config)
	parser()
}

//
// parser .json file
//
func parser () {
	jf,err := os.Open(cmd.P)
	if err != nil {
		log.Fatal(err)
	}
	dec := json.NewDecoder(jf)
	err = dec.Decode(Cfg)
	if err != nil {
		log.Fatal(err)
	}
}