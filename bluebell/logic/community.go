package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

// GetCommunityList 查找到所有的community 并返回
func GetCommunityList() ([]*models.Community, error) {
	return mysql.GetCommunityList()
}

func GetCommunityDetail(id int64) (*models.CommunityDetail, error) {
	return mysql.GetCommunityDetailByID(id)
}
