package openload

import (
	"encoding/json"
	"errors"

	"github.com/cub007/httplib"
)

//RemoteUpload is function used for remote upload file by remote url
func (u *User) RemoteUpload(remoteurl, folderid string) (*RemoteUploadResp, error) {
	//构造请求
	url := baseURL + "/remotedl/add?login=" + u.Login + "&key=" + u.Key + "&url=" + remoteurl + "&folder=" + folderid

	//发送请求
	body, err := httplib.Do("GET", url, nil)
	if err != nil {
		return nil, err
	}
	var res RemoteUploadResp
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	if res.Status != 200 {
		return nil, errors.New(res.Msg)
	}
	return &res, nil
}
