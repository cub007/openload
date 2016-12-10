package openload

import (
	"encoding/json"
	"errors"

	"github.com/cub007/httplib"
)

//GetAccInfo get the current account infomation
func (u *User) GetAccInfo() (*AccResp, error) {
	url := baseURL + "/account/info?login=" + u.Login + "&key=" + u.Key
	body, err := httplib.Do("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var res AccResp
	if err = json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	if res.Status != 200 {
		return nil, errors.New("user :" + u.Login + " login error")
	}
	return &res, nil

}
