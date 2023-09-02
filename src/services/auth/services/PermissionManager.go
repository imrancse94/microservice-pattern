package Services

import "auth/models"

func GetPermissionByUserId(userId int) models.Permission {
	return models.GetRolePageByUserId(userId)
}
