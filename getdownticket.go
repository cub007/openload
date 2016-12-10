package openload

import (
	"encoding/json"
	"errors"

	"github.com/cub007/httplib"
)

//GetDownTicket get the down ticket used for getdownlink by fileid
func (u *User) GetDownTicket(fileid string) (*TicketResp, error) {

	url := baseURL + "/file/dlticket?file=" + fileid
	body, err := httplib.Do("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var res TicketResp
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, err
	}
	if res.Status != 200 {
		return nil, errors.New(res.Msg)
	}
	return &res, nil
}
