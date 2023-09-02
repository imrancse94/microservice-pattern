package Services

import "gateway/models"

func GetPermissionByUserId(userId int) models.Permission {
	return models.GetRolePageByUserId(userId)
}
