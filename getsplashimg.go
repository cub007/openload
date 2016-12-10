package openload

import (
	"encoding/json"
	"errors"

	"github.com/cub007/httplib"
)

//GetSplashIMG is used for get splash image by fileid
func (u *User) GetSplashIMG(fileid string) (*SplashResp, error) {
	//send request
	url := baseURL + "/file/getsplash?file=" + fileid + "&login=" + u.Login + "&key=" + u.Key
	body, err := httplib.Do("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var res SplashResp
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	if res.Status != 200 {
		return nil, errors.New(res.Msg)
	}
	return &res, nil
}
