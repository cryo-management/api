package models

const (
	// TableSchemas database table name
	TableSchemas string = "schemas"
	// TableTranslations database table name
	TableTranslations string = "translations"
	// TableFields database table name
	TableFields string = "fields"
	// TableUsers database table name
	TableUsers string = "users"
	// TableGroups database table name
	TableGroups string = "groups"
	// TableGroupsUsers database table name
	TableGroupsUsers string = "groups_users"
	// TableGroupsPermissions database table name
	TableGroupsPermissions string = "groups_permissions"
	// TableLookups database table name
	TableLookups string = "lookups"
	// TableLookupsOptions database table name
	TableLookupsOptions string = "lookups_options"
)

// Model interface define default methods
type Model interface {
	GetID() string
}
