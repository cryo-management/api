package db

import (
	"fmt"
	"reflect"
	"strings"
)

//GenerateInsertQuery docs
func GenerateInsertQuery(table string, obj interface{}) (string, []interface{}) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	fields := []string{}
	args := []interface{}{}
	params := []string{}

	paramIndex := 1
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("pk") != "true" && t.Field(i).Tag.Get("external") != "true" {
			fields = append(fields, t.Field(i).Tag.Get("sql"))
			params = append(params, fmt.Sprintf("$%d", paramIndex))
			paramIndex++
			args = append(args, v.Field(i).Interface())
		}
	}

	query := fmt.Sprintf("insert into %s (%s) values (%s)", table, strings.Join(fields, ", "), strings.Join(params, ", "))

	return query, args
}

//GenerateTranslationsInsertQuery docs
func GenerateTranslationsInsertQuery(objID, langCode string, obj, trs interface{}) (string, []interface{}) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	fields := []string{}
	args := []interface{}{}
	params := []string{}

	totalFields := 0
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("external") == "true" {
			args = append(args, objID)
			args = append(args, t.Field(i).Tag.Get("type"))
			args = append(args, t.Field(i).Tag.Get("alias"))
			args = append(args, v.Field(i).Interface())
			args = append(args, langCode)
			index := 5 * totalFields
			params = append(params, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", 1+index, 2+index, 3+index, 4+index, 5+index))
			totalFields++
		}
	}

	t = reflect.TypeOf(trs)
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("pk") != "true" {
			fields = append(fields, t.Field(i).Tag.Get("sql"))
		}
	}

	query := fmt.Sprintf("insert into translations (%s) values %s", strings.Join(fields, ", "), strings.Join(params, ", "))

	return query, args
}
