package repositories

import (
	"context"

	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/domains/messaging/models"
	"github.com/Bootcamp-Kelompok-31-BE/whatsapp-clone/shared/util"
)

type IGroupRepository interface {
	GetGroupByID(ctx context.Context, groupID uint) (*models.Group, error)
	GetGroupMemberIDs(ctx context.Context, groupID uint) ([]uint, error)
	IsUserInGroup(ctx context.Context, userID uint, groupID uint) (bool, error)
}

type GroupRepository struct {
}

func NewGroupRepository() IGroupRepository {
	return &GroupRepository{}
}

func (r *GroupRepository) GetGroupByID(ctx context.Context, groupID uint) (*models.Group, error) {
	db := util.GetDBFromContext(ctx)

	var group models.Group
	err := db.Where("id = ?", groupID).First(&group).Error
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (r *GroupRepository) GetGroupMemberIDs(ctx context.Context, groupID uint) ([]uint, error) {
	db := util.GetDBFromContext(ctx)

	var group models.Group
	err := db.
		Model(&models.Group{}).
		Preload("Users").
		Where("id = ?", groupID).
		First(&group).Error
	if err != nil {
		return nil, err
	}

	userIDs := make([]uint, len(group.Users))
	for i, user := range group.Users {
		userIDs[i] = user.ID
	}

	return userIDs, nil
}

func (r *GroupRepository) IsUserInGroup(ctx context.Context, userID uint, groupID uint) (bool, error) {
	db := util.GetDBFromContext(ctx)

	var count int64
	err := db.Model(&models.Group{}).
		Where("id = ?", groupID).
		Where("users.id = ?", userID).
		Joins("JOIN users ON groups.id = users.group_id").
		Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
