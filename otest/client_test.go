package otest

import (
	"github.com/Or1ole/ocfgo"
	"io/ioutil"
	"testing"
)

type MysqldSection struct {
	Basedir string `ini: "basedir"`
	Datadir string `ini: "datadir"`
	Socket string `ini: "socket"`
	Port int `ini: "port"`
}

type ClientSection struct {
	Socket string `ini: "socket"`
	Charset string `ini: "charset"`
}

type iniConfig struct {
	MysqldSection `ini: "mysqld"`  // 对应ini文件的section name
	ClientSection `ini: "client"`
}

func TestReadConfig(t *testing.T)  {
	fileContent, err := ioutil.ReadFile("./my.ini")
	if err != nil {
		panic("read config file failed")
	}
	var marshalData iniConfig
	err = ocfgo.MarshalIni(fileContent, &marshalData)
	if err != nil {
		panic(err)
	}
}
