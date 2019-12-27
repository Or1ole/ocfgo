package ocfgo

import (
	"fmt"
	"reflect"
	"strings"
)

func parseMarshalData(sectionName string, t reflect.Type) {
	for i:=0; i < t.Elem().NumField(); i++ {
		field := t.Elem().Field(i)
		fmt.Println(field.Name)
	}
}

func MarshalIni(data []byte, marshalData interface{}) (err error) {
	t := reflect.TypeOf(marshalData)
	err = checkMarshalDataType(t)
	if err != nil {
		return
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if err = checkFileFormat(line); err != nil {
			return err
		}
		if isSpaceOrComment(line) {
			continue
		}
		if isSection(line) {
			sectionName := line[1 : len(line)-1]
			parseMarshalData(sectionName, t)
		}
	}
	return
}
