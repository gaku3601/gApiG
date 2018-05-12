package api

import (
	"fmt"

	"github.com/gaku3601/gApiG/backendapp/config"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Api struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Method     string `json:"method"`
	AccessPath string `json:"access_path"`
}

func indexApi(offset int, limit int) []*Api {
	con := fmt.Sprintf("user=%s dbname=%s host=%s port=%s sslmode=disable", "postgres", "app", config.Value.GetString("database.host"), config.Value.GetString("database.port"))
	db, _ := gorm.Open("postgres", con)
	defer db.Close()

	var list []*Api
	db.Order("id DESC").Offset(offset).Limit(limit).Find(&list)
	return list
}

func createApi(name string, path string, method string, accessPath string) (*Api, error) {
	con := fmt.Sprintf("user=%s dbname=%s host=%s port=%s sslmode=disable", "postgres", "app", config.Value.GetString("database.host"), config.Value.GetString("database.port"))
	db, _ := gorm.Open("postgres", con)
	defer db.Close()
	api := &Api{
		Name:       name,
		Path:       path,
		Method:     method,
		AccessPath: accessPath,
	}
	return api, db.Create(&api).Error
}
