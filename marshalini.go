package ocfgo

import (
	"reflect"
	"strconv"
	"strings"
)

func parseMarshalData(line string, marshalData interface{},curSectionName *string) (err error) {
	sectionType := reflect.TypeOf(marshalData)
	sectionValue := reflect.ValueOf(marshalData)
	// 将section字段保存至curSectionName
	if isSection(line) {
		sectionConfigName := line[1 : len(line)-1] // 配置文件解析出来的sectionName
		for i := 0; i < sectionType.Elem().NumField(); i++ {
			sectionTypeField := sectionType.Elem().Field(i)    // filed.Name = MysqldSection
			sectionTagValue := sectionTypeField.Tag.Get("ini") // 传递过来结构体中的tag(mysqld|client)
			//fmt.Println(sectionTypeField.Name, sectionTypeField.Type, sectionTagValue)  // MysqldSection otest.MysqldItem mysqld
			if sectionConfigName == sectionTagValue {
				*curSectionName = sectionTypeField.Name
				return
			}
		}
	}
	eqelIndex := strings.Index(line, "=")  // 分割key=value
	itemKey := line[:eqelIndex]
	itemKey = strings.TrimSpace(itemKey)
	itemValue := line[eqelIndex+1:]
	itemValue = strings.TrimSpace(itemValue)
	sectionValueField := sectionValue.Elem().FieldByName(*curSectionName).Type()
	//fmt.Println(sectionValueField)  // MysqldItem | ClientItem

	for j:=0; j < sectionValueField.NumField(); j++ {
		itemField := sectionValueField.Field(j)
		itemTagValue := itemField.Tag.Get("ini")  // tag值
		if itemKey == itemTagValue {  // tag值与配置文件截取的key作比较
			switch itemField.Type.Kind().String() {
			case "string":
				sectionValue.Elem().FieldByName(*curSectionName).Field(j).SetString(itemValue)
				return
			case "int":
				intItemValue, toierr := strconv.ParseInt(itemValue, 10, 64)
				if err != nil {
					return toierr
				}
				sectionValue.Elem().FieldByName(*curSectionName).Field(j).SetInt(intItemValue)
				return
			}
		}
	}
	return
}

// marshalData是指针,获取结构体具体值需要调用Elem()方法
func MarshalIni(data []byte, marshalData interface{}) (err error) {
	var curSectionName string
	iniType := reflect.TypeOf(marshalData)
	err = checkMarshalDataType(iniType)
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
		err = parseMarshalData(line, marshalData, &curSectionName)
		if err != nil {
			return
		}
	}
	return
}
