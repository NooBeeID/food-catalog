package menu

import (
	"context"
	"log"
	"time"
)

type repositoryContract interface {
	insertMenu(ctx context.Context, model Menu) (err error)
	findAll(ctx context.Context) (model []Menu, err error)
	findById(ctx context.Context, id int) (model Menu, err error)
}

type service struct {
	repo repositoryContract
}

func newService(repo repositoryContract) service {
	return service{
		repo: repo,
	}
}

func (s service) createMenu(ctx context.Context, req createMenuRequest) (err error) {
	var model = Menu{
		Name:      req.Name,
		Category:  req.Category,
		Desc:      req.Desc,
		Price:     req.Price,
		CreatedAt: time.Now(), // kita bikin default nya adalah waktu saat ini
		UpdatedAt: time.Now(),
	}

	if err = s.repo.insertMenu(ctx, model); err != nil {
		log.Println("[createMenu, insertMenu] error :", err)
		return
	}
	return
}

func (s service) getListMenus(ctx context.Context) (list []listMenuResponse, err error) {
	menus, err := s.repo.findAll(ctx)
	if err != nil {
		log.Println("[getListMenus, findAll] error :", err)
		return
	}

	if len(menus) == 0 {
		return list, nil
	}

	for _, menu := range menus {
		resp := listMenuResponse{
			Id:       menu.Id,
			Price:    menu.Price,
			Desc:     menu.Desc,
			Name:     menu.Name,
			Category: menu.Category,
		}

		list = append(list, resp)
	}

	return list, nil
}
func (s service) getMenuById(ctx context.Context, id int) (resp singleMenuResponse, err error) {
	menu, err := s.repo.findById(ctx, id)
	if err != nil {
		log.Println("[getListMenus, findAll] error :", err)
		return
	}

	resp = singleMenuResponse{
		Id:        menu.Id,
		Price:     menu.Price,
		Desc:      menu.Desc,
		Name:      menu.Name,
		Category:  menu.Category,
		CreatedAt: menu.CreatedAt,
		UpdatedAt: menu.UpdatedAt,
	}

	return resp, nil
}
