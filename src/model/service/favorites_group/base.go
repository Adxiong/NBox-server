/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-15 00:34:27
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-28 23:11:05
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
	Title string `json:"title"`
	isDel uint8  `json:"is_del"`
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

func (f *FavoritesGroup) Add(ctx context.Context, params FavoritesGroup) (*FavoritesGroup, error) {
	dbFg := db.FavoritesGroup{
		GID:     params.GID,
		Title:   params.Title,
		BgColor: params.BgColor,
		Creator: params.Creator,
	}

	err := dbFg.Add(ctx)
	if err != nil {
		log.Println("SERVICE_FAVORITES-GROUP_BASE_FavoritesAdd_Failed")
		log.Printf("params: %+v", dbFg)
		return &FavoritesGroup{}, fmt.Errorf("SERVICE_FAVORITES-GROUP_BASE_FavoritesAdd_Failed")
	}

	result := &FavoritesGroup{
		ID:        dbFg.ID,
		GID:       params.GID,
		Title:     dbFg.Title,
		BgColor:   dbFg.BgColor,
		IsDel:     dbFg.IsDel,
		Version:   dbFg.Version,
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
		return &FavoritesGroupList{}, fmt.Errorf("SERVICE_FAVORITES-GROUP_BASE_FindListByUID_Failed")
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
			Version:   item.Version,
			IsDel:     item.IsDel,
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

	if params.isDel == 1 {
		val[db.FavoritesGroupColumn.IsDel] = params.isDel
	}

	dbFg.UpdateByGID(ctx, gid, val)

	return rowsAffected, nil
}

func (f *FavoritesGroup) DeleteByGID() {

}
