package utils

type LibrariesRolesUser int

const (
	SUPERADMIN LibrariesRolesUser = iota
	ADMIN
	CLIENT
)

func (s LibrariesRolesUser) String() string {
	return [...]string{"SUPERADMIN", "ADMIN", "CLIENT"}[s]
}
