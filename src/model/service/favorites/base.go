/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-14 15:13:27
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-29 23:58:12
 */
package favorites_service

import (
	"context"
	"fmt"
	"log"
	"nbox/src/model/dao/db"
	"time"
)

type FavoritesService interface {
	Add()
	UpdateByFID()
	DeleteByFID()
	FindListByGID()
}

type Favorites struct {
	ID        uint64
	FID       uint64
	Title     string
	Addr      string
	GroupID   uint64
	IsDel     uint8
	Version   uint64
	CreateID  uint64
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type FavoritesList []Favorites

type UpdateFavoritesParams struct {
	Title   *string
	Addr    *string
	GroupID *uint64
}

func (f *Favorites) Add(ctx context.Context, params *Favorites) (*Favorites, error) {
	favorites := &db.Favorites{
		ID:       params.ID,
		Title:    params.Title,
		Addr:     params.Addr,
		GroupID:  params.GroupID,
		IsDel:    params.IsDel,
		Version:  params.Version,
		CreateID: params.CreateID,
	}
	err := favorites.Add(ctx)
	if err != nil {
		fmt.Printf(
			"SERVICE_FAVORITES_BASE_Add_Failed.\n error:%s\n params:%+v\n", err.Error(), favorites)
		return &Favorites{}, fmt.Errorf("Favorites_FavoritesAdd_Failed")
	}
	result := &Favorites{
		ID:        favorites.ID,
		Title:     favorites.Title,
		Addr:      favorites.Title,
		GroupID:   favorites.GroupID,
		IsDel:     favorites.IsDel,
		Version:   favorites.Version,
		CreateID:  favorites.CreateID,
		CreatedAt: favorites.CreatedAt,
		UpdatedAt: favorites.UpdatedAt,
	}
	return result, nil
}

func (f *Favorites) UpdateByFID(ctx context.Context, fid uint64, params *UpdateFavoritesParams) (int64, error) {
	val := map[string]interface{}{}

	if params.Title != nil {
		val[db.FavoritesColumn.Title] = params.Title
	}
	if params.Addr != nil {
		val[db.FavoritesColumn.Addr] = params.Addr
	}

	if params.GroupID != nil {
		val[db.FavoritesColumn.GroupID] = params.GroupID
	}

	dao := db.NewFavorites()
	rowsAffected, err := dao.UpdateTodoByFID(ctx, fid, val)
	if err != nil {
		log.Printf(
			"SERVICE_FAVORITES_BASE_UpdateByFID_Failed.\n fid:%d\n params:%+v\n error:%s", fid, params, err.Error())
		return 0, fmt.Errorf("Favorites_UpdateByFID_Failed")
	}

	return rowsAffected, nil
}

func (f *Favorites) DeleteByFID(ctx context.Context, fid uint64) (int64, error) {
	dao := db.NewFavorites()
	rowsAffected, err := dao.DeleteByFID(ctx, fid)
	if err != nil {
		log.Printf(
			"SERVICE_FAVORITES_BASE_DeleteByFID_Failed.\n fid:%d\n error:%s", fid, err.Error())
		return 0, fmt.Errorf("Favorites_DeleteByFID_Failed")
	}

	return rowsAffected, nil
}

func (f *Favorites) FindListByGID(ctx context.Context, gid uint64) (*FavoritesList, error) {
	dao := db.NewFavorites()

	list, err := dao.FindListByGID(ctx, gid)
	if err != nil {
		log.Printf(
			"SERVICE_FAVORITES_BASE_FindListByGID_Failed.\n gid:%d\n error:%s", gid, err.Error())
		return &FavoritesList{}, fmt.Errorf("Favorites_FindListByGID_Failed")
	}
	var result = make(FavoritesList, 0)
	for _, item := range *list {
		if item.IsDel != 0 {
			continue
		}
		data := Favorites{
			ID:        item.ID,
			FID:       item.FID,
			Title:     item.Title,
			Addr:      item.Addr,
			GroupID:   item.GroupID,
			IsDel:     item.IsDel,
			Version:   item.Version,
			CreateID:  item.CreateID,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}

		result = append(result, data)

	}
	return &result, nil
}
