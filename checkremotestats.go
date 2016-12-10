package openload

import (
	"encoding/json"
	"errors"

	"github.com/cub007/httplib"
)

//CheckRemoteUploadStatus is used for check the file status by remote upload id
func (u *User) CheckRemoteUploadStatus(remoteid string) (*RemoteStatusResp, error) {
	url := baseURL + "/remotedl/status?login=" + u.Login + "&key=" + u.Key + "&limit=100&id=" + remoteid
	body, err := httplib.Do("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var res RemoteStatusResp
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}

	if res.Status != 200 {
		return nil, errors.New(res.Msg)
	}
	return &res, nil
}
