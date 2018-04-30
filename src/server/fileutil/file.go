package fileutil

import (
	"os"
	"fmt"
	"../config"
	"io"
)

type Info struct {
	Fns interface{} `json:"fns"`
}

type FileStats struct {
	List interface{} `json:"list"`
}

type FileInfo struct {
	Fn string 
	IsExist bool 
}

func CheckFiles (ps []string) []*FileInfo {
	tmp := make([]*FileInfo,0)
	for i,_ := range ps {
		_,err := os.Stat(fmt.Sprintf("%s/%s",config.Cfg.UploadFolder,ps[i]))
		if os.IsNotExist(err) {
			tmp = append(tmp,&FileInfo{Fn:ps[i],IsExist:false})
		} else {
			tmp = append(tmp,&FileInfo{Fn:ps[i],IsExist:true})
		}
	}
	return tmp
}

func Copy (fn *string,r io.Reader) error {
	//create or overrite file in web server
	f, err := os.OpenFile(fmt.Sprintf("%s/%s",config.Cfg.UploadFolder,*fn), os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	defer f.Close()

	//write file to local
	_,err = io.Copy(f, r)
	if err != nil {
		return err
	}

	return nil
}
