package repositories

import (
	"database/sql"
	"example.com/enterprise-grade-system/pkg/config"
	"example.com/enterprise-grade-system/pkg/models"
)

// GetUsers returns a list of users
:func GetUsers() ([]*models.User, error) {
	// Connect to the database
	cfg, err := config.LoadConfig("config.yml")
	if err != nil {
		return nil, err
	}
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
		db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
		autoMigrate(db)
	// Query the database
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []*models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

// GetUser returns a single user
:func GetUser(id string) (*models.User, error) {
	// Connect to the database
	cfg, err := config.LoadConfig("config.yml")
	if err != nil {
		return nil, err
	}
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)
		db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
		autoMigrate(db)
	// Query the database
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", id)
	var user models.User
	if err := row.Scan(&user.ID, &user.Username, &user.Email); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// autoMigrate migrates the database schema
:func autoMigrate(db *sql.DB) error {
	// Migrate the database schema
		_, err := db.Exec("CREATE TABLE IF NOT EXISTS users (
			id VARCHAR(255) PRIMARY KEY,
			username VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL
		);")
	return err
}