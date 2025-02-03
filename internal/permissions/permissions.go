package permissions

type Permission string

const (
	SocialMediaReadPermission   Permission = "social-media:read"
	SocialMediaWritePermission  Permission = "social-media:write"
	SocialMediaDeletePermission Permission = "social-media:detele"
	SocialMediaUpdatePermission Permission = "social-media:update"
	MailReadPermission          Permission = "mail:read"
	MailWritePermission         Permission = "mail:write"
	MailDeletePermission        Permission = "mail:delete"
	MailUpdatePermission        Permission = "mail:update"
	ClubReadPermission          Permission = "club:read"
	ClubWritePermission         Permission = "club:write"
	ClubDeletePermission        Permission = "club:delete"
	ClubUpdatePermission        Permission = "club:update"
	EventReadPermission         Permission = "event:read"
	EventWritePermission        Permission = "event:write"
	EventDeletePermission       Permission = "event:delete"
	EventUpdatePermission       Permission = "event:update"
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
			SocialMediaDeletePermission: true,
			SocialMediaReadPermission:   true,
			SocialMediaUpdatePermission: true,
			SocialMediaWritePermission:  true,
			MailDeletePermission:        true,
			MailReadPermission:          true,
			MailUpdatePermission:        true,
			MailWritePermission:         true,
			ClubDeletePermission:        true,
			ClubReadPermission:          true,
			ClubUpdatePermission:        true,
			ClubWritePermission:         true,
			EventDeletePermission:       true,
			EventReadPermission:         true,
			EventUpdatePermission:       true,
			EventWritePermission:        true,
		},
	}
	SocialAdminRole = Role{
		Name: "social_admin",
		Permissions: Permissions{
			SocialMediaDeletePermission: true,
			SocialMediaReadPermission:   true,
			SocialMediaUpdatePermission: true,
			SocialMediaWritePermission:  true,
		},
	}
	MailAdminRole = Role{
		Name: "mail_admin",
		Permissions: Permissions{
			MailDeletePermission: true,
			MailReadPermission:   true,
			MailUpdatePermission: true,
			MailWritePermission:  true,
		},
	}
	ClubAdminRole = Role{
		Name: "club_admin",
		Permissions: Permissions{
			ClubDeletePermission: true,
			ClubReadPermission:   true,
			ClubUpdatePermission: true,
			ClubWritePermission:  true,
		},
	}
)

func (r *Role) HasPermission(p Permission) bool {
	return r.Permissions[p]
}

func GetRoleWithRoleName(roleName string) *Role {
	switch roleName {
	case "admin":
		return &AdminRole
	case "social_admin":
		return &SocialAdminRole
	case "mail_admin":
		return &MailAdminRole
	case "club_admin":
		return &ClubAdminRole
	default:
		return nil
	}
}
