package openload

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/cub007/httplib"
)

//GetFileInfo is function used for get file infomation by fileids
func (u *User) GetFileInfo(fileids []string) (*FileInfoResp, error) {

	//fileids can not be empty
	files := strings.Join(fileids, ",")
	if files == "" {
		return nil, errors.New("fileids is empty")
	}

	url := baseURL + "/file/info?file=" + files + "&login=" + u.Login + "&key=" + u.Key
	body, err := httplib.Do("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var res FileInfoResp
	if err = json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	if res.Status != 200 {
		return nil, errors.New(res.Msg)
	}
	return &res, nil
}
