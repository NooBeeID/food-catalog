package employee

import (
	"context"
	"database/sql"
)

type repository struct {
	db *sql.DB
}

// removeEmployeeById implements repositoryContract.
func (r repository) removeEmployeeById(ctx context.Context, empId int) (err error) {
	query := `
		DELETE FROM employees
		WHERE id=$1
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, empId)

	return
}

// findAllEmployees implements repositoryContract.
func (r repository) findAllEmployees(ctx context.Context) (res []Employee, err error) {
	query := `
		SELECT id, name, address, nip, created_at
		FROM employees
	`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var emp = Employee{}
		if err := rows.Scan(&emp.Id, &emp.Name, &emp.Address, &emp.NIP, &emp.CreatedAt); err != nil {
			return []Employee{}, err
		}

		res = append(res, emp)
	}

	return res, nil
}

// newEmployee implements repositoryContract.
func (r repository) newEmployee(ctx context.Context, req Employee) (err error) {
	query := `
		INSERT INTO employees (name, address, nip, created_at, updated_at)
		VALUES (
			$1, $2, $3, now(), now()
		)
	`

	stmt, err := r.db.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, req.Name, req.Address, req.NIP)
	return
}

func newRepository(db *sql.DB) repository {
	return repository{
		db: db,
	}
}
