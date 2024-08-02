package auth

import "database/sql"

type repository struct {
	db *sql.DB
}

func newRepository(db *sql.DB) repository {
	return repository{
		db: db,
	}
}

// Create implements RepositoryInterface
func (r repository) create(auth Auth) (err error) {
	// query database
	query := `
		INSERT INTO auth (email, password, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
	`

	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.Exec(auth.Email, auth.Password, auth.CreatedAt, auth.UpdatedAt)
	return
}

// GetByEmail implements RepositoryInterface
func (r repository) getByEmail(email string) (auth Auth, err error) {
	query := `
		SELECT 
			id, email, password
			, created_at, updated_at
		FROM auth
		WHERE email = $1
	`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return
	}

	defer stmt.Close()

	row := stmt.QueryRow(email)

	err = row.Scan(
		&auth.Id, &auth.Email, &auth.Password,
		&auth.CreatedAt, &auth.UpdatedAt,
	)

	return
}
