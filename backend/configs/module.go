package configs

const (
	ModuleID          = 0
	ModuleName        = "VOAH-Template-Project"
	ModuleVersion     = "0.0.1"
	ModuleDescription = "VOAH Template Project"
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
