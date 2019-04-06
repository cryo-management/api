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

//GenerateSelectQuery docs
func GenerateSelectQuery(table string, obj interface{}) string {
	t := reflect.TypeOf(obj)

	query := ""
	fields := []string{}
	translationFields := []string{}
	pivotFields := []string{}

	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("sql") != "" {
			if t.Field(i).Tag.Get("table") != "" {
				field := t.Field(i).Tag.Get("alias")
				translationFields = append(translationFields, field)
			} else {
				field := t.Field(i).Tag.Get("sql")
				tableAndField := fmt.Sprintf("%s.%s", table, field)
				fieldType := t.Field(i).Tag.Get("sqlType")
				pivotField := fmt.Sprintf("%s %s", field, fieldType)
				pivotFields = append(pivotFields, pivotField)
				fields = append(fields, tableAndField)
			}
		}
	}

	if len(translationFields) > 0 {
		translationValues := []string{}
		translationPivotFields := []string{}
		translationInFields := []string{}

		for i := 0; i < len(translationFields); i++ {
			translationField := translationFields[i]
			value := fmt.Sprintf("('%s')", translationField)
			translationInField := fmt.Sprintf("'%s'", translationField)
			translationPivotField := fmt.Sprintf("%s character varying", translationField)
			translationInFields = append(translationInFields, translationInField)
			translationValues = append(translationValues, value)
			translationPivotFields = append(translationPivotFields, translationPivotField)
			query = fmt.Sprintf("select * from crosstab($$select %s, translations.structure_field, translations.value from %s join translations on translations.structure_id = %s.id and translations.structure_field in (%s)$$, $$values %s$$) as tab (%s, %s)", strings.Join(fields, ", "), table, table, strings.Join(translationInFields, ", "), strings.Join(translationValues, ", "), strings.Join(pivotFields, ", "), strings.Join(translationPivotFields, ", "))
		}
	} else {
		query = fmt.Sprintf("select %s from %s", strings.Join(fields, ", "), table)
	}

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
