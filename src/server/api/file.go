package api

import (
	"net/http"
	"../fileutil"
	"errors"
	"encoding/json"
)

var (
	IsExistError error
)

func init() {
	IsExistError = errors.New("FIle is exist")
}

//
// check file exist
//
// client post json format example :
//{
//	"fns" : ["text.txt","text1.txt"]
//}
//
// echo :
// {
// [{"Fn":"text.txt","IsExist":true},{"Fn":"text1.txt","IsExist":false}]
//
//
func FileStat (w http.ResponseWriter,req *http.Request) error {
	dec := json.NewDecoder(req.Body)
	info := new(fileutil.Info)
	if err := dec.Decode(info);err != nil {
		return err
	}
	fi := *info
	flist := fi.Fns.([]interface {})
	list := make([]string,0)
	for _,v := range flist {
		list = append(list,v.(string))
	}

	fileStats := fileutil.CheckFiles(list)

	if err := json.NewEncoder(w).Encode(fileStats);err != nil {
		return err
	}

	return nil
}
