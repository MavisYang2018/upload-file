package api

import (
	"net/http"
	"io"
	"encoding/json"
	"../fileutil"
)

type jsonBody struct {
	FileNames []string `json:"file_name"`
}

//
// upload file
//
// Client post example :
// {
// 	"uploadfile" : "test.txt"
// }
//
//
func Upload (w http.ResponseWriter,req *http.Request) error {
	mr, err := req.MultipartReader()
	if err != nil {
		return err
	}

	// for echo json
	jb := new(jsonBody)
	jb.FileNames = make([]string,0)

	for {
		part,err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		fn := part.FileName()
		if err := fileutil.Copy(&fn,part);err != nil {
			return err
		}

		jb.FileNames = append(jb.FileNames,fn)
		part.Close()
	}

	if err := json.NewEncoder(w).Encode(jb);err != nil {
		return err
	}

	return nil
}