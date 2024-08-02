package employee

import (
	"context"
	"log"
	"time"
)

type repositoryContract interface {
	newEmployee(ctx context.Context, req Employee) (err error)
	findAllEmployees(ctx context.Context) (res []Employee, err error)
	removeEmployeeById(ctx context.Context, empId int) (err error)
}

type service struct {
	repo repositoryContract
}

func newService(repo repositoryContract) service {
	return service{
		repo: repo,
	}
}

func (s service) createNewEmployee(ctx context.Context, req createNewEmployeeRequest) (err error) {
	var emp = Employee{
		Name:    req.Name,
		Address: req.Address,
		NIP:     req.NIP,
	}

	err = s.repo.newEmployee(ctx, emp)
	if err != nil {
		log.Println("[createNewEmployee, newEmployee] error :", err)
		return err
	}
	return
}

func (s service) listEmployees(ctx context.Context) (employees []listEmployeeResponse, err error) {
	resp, err := s.repo.findAllEmployees(ctx)
	if err != nil {
		log.Println("[listEmployees, findAllEmployees] error :", err)
		return []listEmployeeResponse{}, err
	}

	for _, res := range resp {
		var createdAt = res.CreatedAt.Format(time.DateOnly)
		emp := listEmployeeResponse{
			Id:        res.Id,
			Name:      res.Name,
			NIP:       res.NIP,
			Address:   res.Address,
			CreatedAt: createdAt,
		}
		employees = append(employees, emp)
	}

	return employees, nil
}

func (s service) removeEmployeeById(ctx context.Context, id int) (err error) {
	if err = s.repo.removeEmployeeById(ctx, id); err != nil {
		log.Println("[removeEmployeeById, removeEmployeeById] error :", err)
		return
	}
	return
}
