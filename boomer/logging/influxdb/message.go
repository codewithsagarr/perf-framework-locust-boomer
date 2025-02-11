package influxdb

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type InfluxMessage struct {
	measure   string
	tagList   []string
	fieldList []string
	timestamp time.Time
}

func keyValuePair(key string, value string) string {
	return key + "=" + value
}

func escapeSpaces(s string) string {
	return strings.Join(strings.Split(s, " "), "\\ ")
}

func typedFieldValue(value string, toString bool) string {
	var typedValue string
	if toString {
		return `"` + value + `"`
	}

	if _, intErr := strconv.Atoi(value); intErr == nil {
		typedValue = value + "i"
	} else if _, floatErr := strconv.ParseFloat(value, 32); floatErr == nil {
		typedValue = value
	} else if _, boolErr := strconv.ParseBool(value); boolErr == nil {
		typedValue = value
	} else {
		typedValue = `"` + value + `"`
	}
	return typedValue
}

func (im *InfluxMessage) AddTag(tag string, value string) {
	if value != "" {
		kvp := escapeSpaces(keyValuePair(tag, value))
		im.tagList = append(im.tagList, kvp)
	}
}

func (im *InfluxMessage) _addField(field string, value string, toString bool) {
	if value != "" {
		im.fieldList = append(im.fieldList, keyValuePair(escapeSpaces(field), typedFieldValue(value, toString)))
	}
}

func (im *InfluxMessage) AddField(field string, value string) {
	im._addField(field, value, false)
}

func (im *InfluxMessage) AddStringField(field string, value string) {
	im._addField(field, value, true)
}

func (im *InfluxMessage) LineProtocol() string {
	lp := strings.Builder{}
	lp.WriteString(escapeSpaces(im.measure))

	if len(im.tagList) > 0 {
		lp.WriteString(",")
		lp.WriteString(strings.Join(im.tagList, ","))
	}

	if len(im.fieldList) == 0 {
		panic("Influxdb: line protocol requires there be at least one field")
	}

	lp.WriteString(" ")
	lp.WriteString(strings.Join(im.fieldList, ","))
	lp.WriteString(fmt.Sprintf(" %d\n", im.timestamp.Local().UnixNano()))

	return lp.String()
}
