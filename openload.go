package openload

var baseURL = "https://api.openload.co/1"

//User is user struct
type User struct {
	Login string
	Key   string
}

//New create a user struct
func New(login, key string) *User {
	u := User{login, key}
	return &u
}

//AccResp is used for GetAccInfo response struct
type AccResp struct {
	Status int
	Msg    string
	Result struct {
		Extid       string
		Email       string
		SignupAt    string `json:"signup_at"`
		StorageLeft int    `json:"storage_left"`
		StorageUsed string `json:"storage_used"`
		Traffic     struct {
			Left    int
			Used24h int `json:"Used_24h"`
		}
		Balance string
	}
}

//TicketResp is used for GetDownTicket response struct
type TicketResp struct {
	Status int
	Msg    string
	Result struct {
		Ticket     string
		CaptchaURL string `json:"captcha_url"`
		CaptchaW   int    `json:"captcha_w"`
		CaptchaH   int    `json:"captcha_h"`
		WaitTime   int    `json:"wait_time"`
		WalidUntil string `json:"walid_until"`
	}
}

//DownLinkResp is used for getdownlink response struct
type DownLinkResp struct {
	Status int
	Msg    string
	Result struct {
		Name       string
		size       int
		sha1       string
		ContenType string `json:"content_type"`
		UploadAt   string `json:"upload_at"`
		URL        string
		Token      string
	}
}

//ListResp is used for ListFolder response struct
type ListResp struct {
	Status int
	Msg    string
	Result struct {
		Folders []struct {
			ID   string
			Name string
		}
		Files []struct {
			Name          string
			Sha1          string
			Folderid      string
			UploadAt      string `json:"upload_at"`
			Status        string
			Size          string
			ContentType   string `json:"content_type"`
			DownloadCount string `json:"download_count"`
			Cstatus       string
			Link          string
			Linkextid     string
		}
	}
}

//FileInfoResp is used for GetFileInfo response struct
type FileInfoResp struct {
	Status int
	Msg    string
	Result map[string]struct {
		ID          string
		Cstatus     string
		Status      int
		Name        string
		Size        string
		SHA1        string
		ContentType string `json:"Content_type"`
	}
}

//RemoteUploadResp is used for RemoteUpload response
type RemoteUploadResp struct {
	Status int
	Msg    string
	Result struct {
		ID       string
		Folderid string
	}
}

//RemoteStatusResp is used for CheckRemoteUploadStatus response
type RemoteStatusResp struct {
	Status int
	Msg    string
	Result map[int]struct {
		ID          string
		Remoteurl   string
		Status      string
		BytesLoaded string `json:"bytes_loaded"`
		BytesTotal  string `json:"bytes_total"`
		Folderid    string
		Added       string
		LastUpdate  string `json:"last_update"`
		Extid       string
		URL         string
	}
}

//SplashResp is used for GetSplashIMG response struct
type SplashResp struct {
	Status int
	Msg    string
	Result string
}
