/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-15 00:34:27
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-30 00:27:40
 */
package favorites_group_service

import (
	"context"
	"fmt"
	"log"
	"nbox/src/model/dao/db"
	"time"
)

type FavoritesGroup struct {
	ID        uint64     `json:"id"`
	GID       uint64     `json:"gid"`
	Title     string     `json:"title"`
	Creator   uint64     `json:"creator"`
	BgColor   string     `json:"bg_color"`
	IsDel     uint8      `json:"-"`
	Version   uint64     `json:"-"`
	CreatedAt *time.Time `json:"create_at"`
	UpdatedAt *time.Time `json:"update_at"`
}

type FavoritesGroupList []FavoritesGroup

type UpdateFavoritesGroupParams struct {
	Title string
}

type FavoritesGroupServices interface {
	Add()
	FindListByUID()
	UpdateByGID()
	DeleteByGID()
}

func NewFavoritesGroup() *FavoritesGroup {
	return new(FavoritesGroup)
}

func (f *FavoritesGroup) Add(ctx context.Context, params *FavoritesGroup) (*FavoritesGroup, error) {
	dbFg := db.FavoritesGroup{
		GID:     params.GID,
		Title:   params.Title,
		BgColor: params.BgColor,
		Creator: params.Creator,
		IsDel:   0,
		Version: 1,
	}

	err := dbFg.Add(ctx)
	if err != nil {
		log.Println("SERVICE_FAVORITES-GROUP_BASE_FavoritesAdd_Failed")
		log.Printf("params: %+v", dbFg)
		return &FavoritesGroup{}, fmt.Errorf("Favorites_Group_FavoritesAdd_Failed")
	}

	result := &FavoritesGroup{
		ID:        dbFg.ID,
		GID:       params.GID,
		Title:     dbFg.Title,
		BgColor:   dbFg.BgColor,
		Creator:   dbFg.Creator,
		CreatedAt: dbFg.CreatedAt,
		UpdatedAt: dbFg.UpdatedAt,
	}

	return result, nil
}

func (f *FavoritesGroup) FindListByUID(ctx context.Context, uid uint64) (*FavoritesGroupList, error) {
	dbFg := db.NewFavoritesGroup()
	list, err := dbFg.FindListByUID(ctx, uid)
	if err != nil {
		fmt.Printf(
			"SERVICE_FAVORITES-GROUP_BASE_FindListByUID_Failed. \n error:%s \n uid:%d\n",
			err.Error(), uid,
		)
		return &FavoritesGroupList{}, fmt.Errorf("Favorites_Group_FindListByUID_Failed")
	}

	if len(*list) == 0 {
		return &FavoritesGroupList{}, nil
	}
	result := make(FavoritesGroupList, 0)
	for _, item := range *list {
		if item.IsDel != 0 {
			continue
		}

		data := FavoritesGroup{
			ID:        item.ID,
			GID:       item.GID,
			Title:     item.Title,
			BgColor:   item.BgColor,
			Creator:   item.Creator,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}
		result = append(result, data)
	}
	return &result, nil
}

func (f *FavoritesGroup) UpdateByGID(ctx context.Context, params *UpdateFavoritesGroupParams, gid uint64) (int64, error) {
	var rowsAffected int64

	dbFg := db.NewFavoritesGroup()
	val := map[string]interface{}{}
	if params.Title != "" {
		val[db.FavoritesGroupColumn.Title] = params.Title
	}

	if len(val) == 0 {
		fmt.Println("SERVICE_FAVORITES-GROUP_BASE_UpdateByGID_Params_Empty")
		return rowsAffected, nil
	}
	rowsAffected, err := dbFg.UpdateByGID(ctx, gid, val)
	if err != nil {
		log.Printf("SERVICE_FAVORITES-GROUP_BASE_UpdateByGID_Failed.\nerror:%s\ngid:%d-val:%+v\n", err.Error(), gid, val)
		return 0, fmt.Errorf("Favorites_Group_UpdateByGID_Failed")
	}
	return rowsAffected, nil
}

func (f *FavoritesGroup) DeleteByGID(ctx context.Context, gid uint64) (int64, error) {
	var rowsAffected int64

	dbFg := db.NewFavoritesGroup()

	rowsAffected, err := dbFg.DeleteByGID(ctx, gid)
	if err != nil {
		log.Printf("SERVICE_FAVORITES-GROUP_BASE_DeleteByGID_Failed.\nerror:%s\ngid:%d\n", err.Error(), gid)
		return 0, fmt.Errorf("Favorites_Group_DeleteByGID_Failed")
	}
	return rowsAffected, nil
}
