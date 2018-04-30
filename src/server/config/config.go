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
	UploadFolder string `json:"upload_folder"`
	ListenPort string `json:"listen_port"`
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