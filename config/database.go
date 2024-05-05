package config

import "database/sql" 

func DBConnection() (*sql.DB, error)