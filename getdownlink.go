package openload

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/cub007/httplib"
)

//GetDownLink get download link
func (u *User) GetDownLink(fileid, ticket, captcha string) (*DownLinkResp, error) {
	//need sleep 2 Seconds
	time.Sleep(2 * time.Second)

	//send request
	url := baseURL + "/file/dl?file=" + fileid + "&ticket=" + ticket
	body, err := httplib.Do("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var res DownLinkResp
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	if res.Status != 200 {
		return nil, errors.New(res.Msg)
	}
	return &res, nil
}
