package otest

import (
	"fmt"
	"github.com/Or1ole/ocfgo"
	"io/ioutil"
	"testing"
)

type MysqldItem struct {
	Basedir string `ini:"basedir"`
	Datadir string `ini:"datadir"`
	Socket string `ini:"socket"`
	Port int `ini:"port"`
}

type ClientItem struct {
	Socket string `ini:"socket"`
	Charset string `ini:"charset"`
}

type iniConfig struct {
	MysqldSection MysqldItem `ini:"mysqld"`  // 对应ini文件的section name
	ClientSection ClientItem `ini:"client"`
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
	fmt.Println(marshalData)
	fmt.Println("===========================unmarshal==========================")
	data, err := ocfgo.UnmarshalIni(&marshalData)
	err = ioutil.WriteFile("/home/sisu/xx.ini", data, 0666)
	if err != nil {
		fmt.Println("write file failed, err:", err)
		return
	}
}
