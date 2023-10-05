package feieyun

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Line interface {
}

type Printer struct {
	User string
	Ukey string
	Url  string
}

func NewPrinter(user, ukey, url string) Printer {
	return Printer{
		User: user,
		Ukey: ukey,
		Url:  url,
	}
}

func (p Printer) Print(sn string, content string, backurl string) {
	sig, timestamp := p.Sig()
	client := http.Client{}
	postValues := url.Values{}
	postValues.Add("user", p.User)             //账号名
	postValues.Add("stime", timestamp)         //当前时间的秒数，请求时间
	postValues.Add("sig", sig)                 //签名
	postValues.Add("apiname", "Open_printMsg") //固定
	postValues.Add("sn", sn)                   //打印机编号
	postValues.Add("content", content)         //打印内容
	postValues.Add("times", "1")               //打印次数
	if backurl != "" {
		postValues.Add("backurl", backurl)
	}
	res, _ := client.PostForm(p.Url, postValues)
	defer res.Body.Close()
}

func (p Printer) Sig() (string, string) {
	itime := time.Now().Unix()
	s := fmt.Sprintf("%s%s%v", p.User, p.Ukey, itime)
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs), fmt.Sprintf("%d", itime)
}

func SHA1(str string) string {
	s := sha1.Sum([]byte(str))
	strsha1 := hex.EncodeToString(s[:])
	return strsha1
}
