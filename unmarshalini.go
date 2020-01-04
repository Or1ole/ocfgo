package ocfgo

import (
	"fmt"
	"reflect"
)

func UnmarshalIni(marshalData interface{}) (data []byte, err error) {
	iniType := reflect.TypeOf(marshalData)
	iniValue := reflect.ValueOf(marshalData)
	err = checkMarshalDataType(iniType)
	if err != nil {
		return
	}
	for i := 0; i < iniType.Elem().NumField(); i++ {
		sectionTypeField := iniType.Elem().Field(i)
		sectionValueField := iniValue.Elem().Field(i)
		sectionTagName := sectionTypeField.Tag.Get("ini")
		dataSectionTagName := []byte(fmt.Sprintf("[%s]\n", sectionTagName))
		data = append(data, dataSectionTagName...)
		for j := 0; j < sectionValueField.NumField(); j++ {
			itemValueField := sectionValueField.Field(j)
			strItemTagName := sectionTypeField.Type.Field(j).Tag.Get("ini")
			strItemValue := itemValueField.Interface()
			dataItem := []byte(fmt.Sprintf("%s=%v\n", strItemTagName, strItemValue))
			data = append(data, dataItem...)
		}
	}
	return
}
