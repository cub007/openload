package openload

import (
	"encoding/json"
	"errors"

	"github.com/cub007/httplib"
)

//ListFolder is function used for list folders and files by folderid
func (u *User) ListFolder(folderid string) (*ListResp, error) {
	//send request
	url := baseURL + "/file/listfolder?login=" + u.Login + "&key=" + u.Key
	if folderid != "" {
		url += "&folder=" + folderid
	}
	body, err := httplib.Do("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var res ListResp
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	if res.Status != 200 {
		return nil, errors.New("error: " + res.Msg)
	}
	return &res, nil
}
