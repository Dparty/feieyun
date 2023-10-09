package feieyun

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type Line interface {
}

type PrinterFactory struct {
	User string
	Ukey string
	Url  string
}

func NewPrinterFactory(user, ukey, url string) PrinterFactory {
	return PrinterFactory{
		User: user,
		Ukey: ukey,
		Url:  url,
	}
}

func (p PrinterFactory) Connect(sn string) (Printer, error) {
	signer := func() (string, string) {
		itime := time.Now().Unix()
		s := fmt.Sprintf("%s%s%v", p.User, p.Ukey, itime)
		h := sha1.New()
		h.Write([]byte(s))
		bs := h.Sum(nil)
		return fmt.Sprintf("%x", bs), fmt.Sprintf("%d", itime)
	}
	var printer Printer = Printer{
		Url:    p.Url,
		Sn:     sn,
		Signer: signer,
		CommonValues: func() url.Values {
			sig, timestamp := signer()
			postValues := url.Values{}
			postValues.Add("user", p.User)
			postValues.Add("stime", timestamp)
			postValues.Add("sig", sig)
			postValues.Add("apiname", "Open_printMsg")
			postValues.Add("sn", sn)
			return postValues
		},
	}
	return printer, nil
}

type Printer struct {
	Sn           string
	Signer       func() (string, string)
	CommonValues func() url.Values
	Url          string
}

func (p Printer) Status() Status {
	client := http.Client{}
	postValues := p.CommonValues()
	res, _ := client.PostForm(p.Url, postValues)
	defer res.Body.Close()
	resBody, _ := io.ReadAll(res.Body)
	var status Status
	json.Unmarshal(resBody, &status)
	return status
}

func (p Printer) Print(content string, backurl string) {
	client := http.Client{}
	postValues := p.CommonValues()
	postValues.Add("content", content)
	postValues.Add("times", "1")
	if backurl != "" {
		postValues.Add("backurl", backurl)
	}
	res, _ := client.PostForm(p.Url, postValues)
	defer res.Body.Close()
}

type Status struct {
	Message            string  `json:"msg"`
	Ret                int     `json:"ret"`
	Data               *string `json:"data"`
	ServerExecutedTime int     `json:"serverExecutedTime"`
}
