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
		if t.Field(i).Tag.Get("sql") != "" && t.Field(i).Tag.Get("pk") != "true" && t.Field(i).Tag.Get("external") != "true" && t.Field(i).Tag.Get("readOnly") != "true" {
			fields = append(fields, t.Field(i).Tag.Get("sql"))
			params = append(params, fmt.Sprintf("$%d", paramIndex))
			paramIndex++
			args = append(args, v.Field(i).Interface())
		}
	}

	query := fmt.Sprintf("insert into %s (%s) values (%s)", table, strings.Join(fields, ", "), strings.Join(params, ", "))

	return query, args
}

//GenerateSelectQuery docs
func GenerateSelectQuery(table string, obj interface{}, wheres ...string) string {
	t := reflect.TypeOf(obj)

	query := ""
	where := ""
	columns := []string{}
	joins := []string{}

	if len(wheres) > 0 {
		where = "where"
	}

	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("sql") != "" && t.Field(i).Tag.Get("readOnly") != "true" {
			column := t.Field(i).Tag.Get("sql")
			if t.Field(i).Tag.Get("table") != "" {
				joinColumnAlias := t.Field(i).Tag.Get("alias")
				joinTable := t.Field(i).Tag.Get("table")
				joinOn := t.Field(i).Tag.Get("on")
				sqlJoinColumn := fmt.Sprintf("%s_%s.%s as %s", joinTable, joinColumnAlias, column, joinColumnAlias)
				sqlJoin := fmt.Sprintf("join %s %s_%s on %s", joinTable, joinTable, joinColumnAlias, joinOn)
				columns = append(columns, sqlJoinColumn)
				joins = append(joins, sqlJoin)
			} else {
				sqlColumn := fmt.Sprintf("%s.%s", table, column)
				columns = append(columns, sqlColumn)
			}
		}
	}

	query = strings.Trim(fmt.Sprintf("select %s from %s %s %s %s", strings.Join(columns, ", "), table, strings.Join(joins, " "), where, strings.Join(wheres, " ")), " ")

	return query
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

//GenerateDeleteQuery docs
func GenerateDeleteQuery(table string, wheres ...string) string {
	query := ""
	where := ""

	if len(wheres) > 0 {
		where = "where"
	}

	query = strings.Trim(fmt.Sprintf("delete from %s %s %s", table, where, strings.Join(wheres, " ")), " ")

	return query
}
