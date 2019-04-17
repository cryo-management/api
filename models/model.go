package models

const (
	TableSchemas           string = "schemas"
	TableTranslations      string = "translations"
	TableFields            string = "fields"
	TableUsers             string = "users"
	TableGroups            string = "groups"
	TableGroupsUsers       string = "groups_users"
	TableGroupsPermissions string = "groups_permissions"
	TableLookups           string = "lookups"
	TableLookupsOptions    string = "lookups_options"
)

type Model interface {
	GetID() string
}
