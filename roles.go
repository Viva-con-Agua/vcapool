package vcapool

import (
	"github.com/Viva-con-Agua/vcago"
)

func GetRole(role string) (r *vcago.Role, err error) {
	switch role {
	case "asp":
		return RoleASP(), err
	case "finance":
		return RoleFinance(), err
	case "action":
		return RoleAction(), err
	case "education":
		return RoleEducation(), err
	case "network":
		return RoleNetwork(), err
	case "socialmedia":
		return RoleSocialMedia(), err
	case "awareness":
		return RoleAwareness(), err
	default:
		return nil, vcago.NewValidationError("role not supported: " + role)
	}
}

func RoleASP() *vcago.Role {
	return &vcago.Role{
		Name:  "asp",
		Label: "ASP",
		Root:  "employee;admin",
	}
}

func RoleSupporter() *vcago.Role {
	return &vcago.Role{
		Name:  "supporter",
		Label: "Supporter",
		Root:  "system",
	}
}

func RoleFinance() *vcago.Role {
	return &vcago.Role{
		Name:  "finance",
		Label: "Finanzen",
		Root:  "finance;employee;admin",
	}
}
func RoleAction() *vcago.Role {
	return &vcago.Role{
		Name:  "action",
		Label: "Aktion",
		Root:  "action;employee;admin",
	}
}
func RoleEducation() *vcago.Role {
	return &vcago.Role{
		Name:  "education",
		Label: "Bildung",
		Root:  "education;employee;admin",
	}
}
func RoleNetwork() *vcago.Role {
	return &vcago.Role{
		Name:  "network",
		Label: "Netzwerk",
		Root:  "network;employee;admin",
	}
}
func RoleSocialMedia() *vcago.Role {
	return &vcago.Role{
		Name:  "socialmedia",
		Label: "Social Media",
		Root:  "socialmedia;employee;admin",
	}
}
func RoleAwareness() *vcago.Role {
	return &vcago.Role{
		Name:  "awareness",
		Label: "Awareness",
		Root:  "awareness;employee;admin",
	}
}
