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

type FeieyunCallback struct {
	OrderId string `form:"orderId" json:"orderId" binding:"required"`
	Status  int    `form:"status" json:"status" binding:"required"`
	Stime   int    `form:"stime" json:"stime" binding:"required"`
	Sign    string `form:"sign" json:"sign" binding:"required"`
}

type FeieyunConfig struct {
	User string
	Ukey string
	Url  string
}

type Line interface {
}

type PrinterFactory struct {
	Config FeieyunConfig
}

func NewPrinterFactory(user, ukey, url string) PrinterFactory {
	return PrinterFactory{
		Config: FeieyunConfig{
			User: user,
			Ukey: ukey,
			Url:  url,
		},
	}
}

func (p PrinterFactory) Connect(sn string) (Printer, error) {
	var printer Printer = Printer{
		Sn:     sn,
		Config: p.Config,
	}
	return printer, nil
}

type Printer struct {
	Config FeieyunConfig
	Sn     string
}

func (p Printer) CommonValues() url.Values {
	sig, itime := p.Sig()
	postValues := url.Values{}
	postValues.Add("user", p.Config.User)
	postValues.Add("stime", itime)
	postValues.Add("sig", sig)
	return postValues
}

func (p Printer) Sig() (string, string) {
	itime := time.Now().Unix()
	s := fmt.Sprintf("%s%s%v", p.Config.User, p.Config.Ukey, itime)
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs), fmt.Sprintf("%d", itime)
}

func (p Printer) Status() Status {
	client := http.Client{}
	postValues := p.CommonValues()
	postValues.Add("sn", p.Sn)
	postValues.Add("apiname", "Open_queryPrinterStatus")
	res, _ := client.PostForm(p.Config.Url, postValues)
	defer res.Body.Close()
	resBody, _ := io.ReadAll(res.Body)
	var status Status
	json.Unmarshal(resBody, &status)
	return status
}

func (p Printer) Print(content string, backurl string) {
	client := http.Client{}
	postValues := p.CommonValues()
	postValues.Add("sn", p.Sn)
	postValues.Add("apiname", "Open_printMsg")
	postValues.Add("content", content)
	postValues.Add("times", "1")
	if backurl != "" {
		postValues.Add("backurl", backurl)
	}
	res, err := client.PostForm(p.Config.Url, postValues)
	if err != nil {
		return
	}
	defer res.Body.Close()
}

type Status struct {
	Message            string  `json:"msg"`
	Ret                int     `json:"ret"`
	Data               *string `json:"data"`
	ServerExecutedTime int     `json:"serverExecutedTime"`
}
