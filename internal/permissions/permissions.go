package permissions

type Permission string

const (
	ReadPermission   Permission = "read"
	WritePermission  Permission = "write"
	DeletePermission Permission = "delete"
	UpdatePermission Permission = "update"
)

type Permissions map[Permission]bool

type Role struct {
	Name        string
	Permissions Permissions
}

var (
	AdminRole = Role{
		Name: "admin",
		Permissions: Permissions{
			ReadPermission:   true,
			WritePermission:  true,
			DeletePermission: true,
			UpdatePermission: true,
		},
	}
	SocialAdminRole = Role{
		Name: "social_admin",
		Permissions: Permissions{
			ReadPermission:   true,
			WritePermission:  true,
			DeletePermission: true,
			UpdatePermission: true,
		},
	}
	MailAdminRole = Role{
		Name: "mail_admin",
		Permissions: Permissions{
			ReadPermission:   true,
			WritePermission:  true,
			DeletePermission: true,
			UpdatePermission: true,
		},
	}
	ClubAdminRole = Role{
		Name: "club_admin",
		Permissions: Permissions{
			ReadPermission:   true,
			WritePermission:  true,
			DeletePermission: true,
			UpdatePermission: true,
		},
	}
)

func (r *Role) HasPermission(p Permission) bool {
	return r.Permissions[p]
}
