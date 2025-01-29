package config

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	// Verifica a conexão com o banco de dados
	connStr := LoadConfig()
	db := Init(connStr)
	defer db.Close()
	
	require.NoError(t, db.Ping(), "Error connecting to the database")
}
