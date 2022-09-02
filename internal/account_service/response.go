package account_service

type Role string

const (
	TeacherRole Role = "professor"
	StudentRole Role = "aluno"
	AdminRole   Role = "admin"
)

type AuthResponse struct {
	ID   string
	Role Role
}

func (r Role) IsRoleAuthorized(rolesAuthorized []Role) bool {
	rolesAuthorized = append(rolesAuthorized, AdminRole)

	for _, role := range rolesAuthorized {
		if role == r {
			return true
		}
	}
	return false
}
