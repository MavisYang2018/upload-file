package cmd

import "flag"

var (
	//config.json path
	P string
)

func init() {
	flag.StringVar(&P,"p","D:/USER/github/upload-file/src/server/config.json","config.json path")
	flag.Parse()
}
