package ocfgo

import (
	"errors"
	"reflect"
)

func checkMarshalDataType(t reflect.Type) (err error) {
	if t.Kind() != reflect.Ptr {
		err = errors.New("expect Ptr type")
		return
	}
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("expect Struct type")
		return
	}
	return
}

func isSpaceOrComment(line string) bool {
	if line == "" || line[0] == '#' || line[0] == ';' {
		return true
	}
	return false
}

// 检查配置文件格式，todo
func checkFileFormat(line string) (err error) {
	return
}

func isSection(line string) bool {
	if line[0] == '[' && line[len(line)-1] == ']' {
		return true
	}
	return false
}