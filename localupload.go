package openload

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"uploader/filelib"
)

//LocalUpload is need to fix
func (u *User) LocalUpload() error {

	//文件路径
	filename, _ := os.Getwd()
	filename += "/xx.mp4"

	//获取文件哈希值
	sha1, err := filelib.GetSHA1(filename)
	if err != nil {
		return err
	}
	fmt.Println("sha1:", sha1)

	//打开文件句柄操作
	fh, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	defer fh.Close()

	//创建一个写缓冲区
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("file1", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}

	//iocopy
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}

	//添加表单数据
	bodyWriter.WriteField("login", u.Login)
	bodyWriter.WriteField("key", u.Key)
	bodyWriter.WriteField("folder", "2431614")
	bodyWriter.WriteField("sha1", sha1)
	bodyWriter.WriteField("httponly", "false")
	bodyWriter.Close()

	//创建请求
	url := baseURL + "/file/ul"
	req, err := http.NewRequest("POST", url, bodyBuf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	//发送请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("http error, Status:" + resp.Status)
	}
	//从响应中读取数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body))
	return nil
}
