package configs

const (
	ModuleID          = 2
	ModuleName        = "VOAH-Official-File"
	ModuleVersion     = "0.0.1"
	ModuleDescription = "VOAH Official File"
)

type ObjectType string

const (
	SystemObject ObjectType = "system"
	RootObject   ObjectType = "root"
)

type PermissionScope string

const (
	AdminPermissionScope PermissionScope = "admin"
	EditPermissionScope  PermissionScope = "edit"
	ReadPermissionScope  PermissionScope = "read"
)

var (
	ModuleDeps             = []string{}
	ModuleObjectTypes      = []ObjectType{SystemObject, RootObject}
	ModulePermissionScopes = []PermissionScope{AdminPermissionScope, EditPermissionScope, ReadPermissionScope}
)
