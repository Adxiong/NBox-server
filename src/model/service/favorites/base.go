/*
 * @Description:
 * @version:
 * @Author: Adxiong
 * @Date: 2022-08-14 15:13:27
 * @LastEditors: Adxiong
 * @LastEditTime: 2022-08-18 23:52:56
 */
package favorites_service

import (
	"context"
	"fmt"
	"nbox/src/model/dao/db"
	"time"
)

type FavoritesService interface {
}

type Favorites struct {
	ID        uint64
	Title     string
	Addr      string
	GroupID   uint64
	IsDel     uint8
	Version   uint64
	CreateID  uint64
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

func add(ctx context.Context, params *Favorites) (*Favorites, error) {
	favoritesDao := &db.Favorites{
		ID:       params.ID,
		Title:    params.Title,
		Addr:     params.Addr,
		GroupID:  params.GroupID,
		IsDel:    params.IsDel,
		Version:  params.Version,
		CreateID: params.CreateID,
	}
	favorites, err := favoritesDao.Add(ctx)
	if err != nil {
		fmt.Println(err)
		return &Favorites{}, err
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
