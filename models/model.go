package models

const (
	TableSchema       string = "schemas"
	TableTranslations string = "translations"
)

type Model interface {
	GetID() string
}
