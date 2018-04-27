package fileuntil

import "os"

type Path struct {
	Fn string `json:"fn"`
}

func CheckFile (p string) bool {
	_,err := os.Stat(p)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
