package vcapool

import (
	"github.com/Viva-con-Agua/vcago"
	"github.com/google/uuid"
)

func NewRole(role string, userID string) (r *vcago.Role, err error) {
	switch role {
	case "asp":
		return RoleASP(userID), err
	case "finance":
		return RoleFinance(userID), err
	case "operation":
		return RoleAction(userID), err
	case "education":
		return RoleEducation(userID), err
	case "network":
		return RoleNetwork(userID), err
	case "socialmedia":
		return RoleSocialMedia(userID), err
	case "awareness":
		return RoleAwareness(userID), err
	default:
		return nil, vcago.NewValidationError("role not supported: " + role)
	}
}

func RoleASP(userID string) *vcago.Role {
	return &vcago.Role{
		ID:     uuid.NewString(),
		Name:   "asp",
		Label:  "ASP",
		Root:   "employee;admin",
		UserID: userID,
	}
}

func RoleSupporter(userID string) *vcago.Role {
	return &vcago.Role{
		ID:     uuid.NewString(),
		Name:   "supporter",
		Label:  "Supporter",
		Root:   "system",
		UserID: userID,
	}
}

func RoleFinance(userID string) *vcago.Role {
	return &vcago.Role{
		ID:     uuid.NewString(),
		Name:   "finance",
		Label:  "Finanzen",
		Root:   "finance;employee;admin",
		UserID: userID,
	}
}
func RoleAction(userID string) *vcago.Role {
	return &vcago.Role{
		ID:     uuid.NewString(),
		Name:   "action",
		Label:  "Aktion",
		Root:   "action;employee;admin",
		UserID: userID,
	}
}
func RoleEducation(userID string) *vcago.Role {
	return &vcago.Role{
		ID:     uuid.NewString(),
		Name:   "education",
		Label:  "Bildung",
		Root:   "education;employee;admin",
		UserID: userID,
	}
}
func RoleNetwork(userID string) *vcago.Role {
	return &vcago.Role{
		ID:     uuid.NewString(),
		Name:   "network",
		Label:  "Netzwerk",
		Root:   "network;employee;admin",
		UserID: userID,
	}
}
func RoleSocialMedia(userID string) *vcago.Role {
	return &vcago.Role{
		ID:     uuid.NewString(),
		Name:   "socialmedia",
		Label:  "Social Media",
		Root:   "socialmedia;employee;admin",
		UserID: userID,
	}
}
func RoleAwareness(userID string) *vcago.Role {
	return &vcago.Role{
		ID:     uuid.NewString(),
		Name:   "awareness",
		Label:  "Awareness",
		Root:   "awareness;employee;admin",
		UserID: userID,
	}
}
