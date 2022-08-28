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
