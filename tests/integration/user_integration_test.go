package integration

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

// ============================================
// Test Suite Structure
// ============================================
type UserIntegrationSuite struct {
	suite.Suite
	ctx              context.Context
	container        testcontainers.Container
	db               *sql.DB
	connectionString string
}

// ============================================
// Setup - Jalankan SEBELUM semua test
// ============================================
func (s *UserIntegrationSuite) SetupSuite() {
	s.ctx = context.Background()

	fmt.Println("🐳 Starting PostgreSQL container...")

	// Start PostgreSQL Container
	pgContainer, err := postgres.Run(s.ctx,
		"postgres:15-alpine",
		postgres.WithDatabase("testdb"),
		postgres.WithUsername("testuser"),
		postgres.WithPassword("testpass"),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(60*time.Second)),
	)
	s.Require().NoError(err)
	s.container = pgContainer

	// Get Connection String
	connStr, err := pgContainer.ConnectionString(s.ctx, "sslmode=disable")
	s.Require().NoError(err)
	s.connectionString = connStr

	fmt.Printf("✅ PostgreSQL ready! Connection: %s\n", connStr)

	// Connect to Database
	db, err := sql.Open("postgres", connStr)
	s.Require().NoError(err)
	s.db = db

	// Verify connection
	err = db.Ping()
	s.Require().NoError(err)

	// Create table
	s.createTable()
}

// ============================================
// TearDown - Jalankan SETELAH semua test
// ============================================
func (s *UserIntegrationSuite) TearDownSuite() {
	fmt.Println("🧹 Cleaning up...")

	if s.db != nil {
		s.db.Close()
	}
	if s.container != nil {
		err := s.container.Terminate(s.ctx)
		if err != nil {
			fmt.Printf("Warning: Failed to terminate container: %v\n", err)
		}
	}
}

// ============================================
// Helper: Create Table
// ============================================
func (s *UserIntegrationSuite) createTable() {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL,
        email VARCHAR(100) UNIQUE NOT NULL,
        age INTEGER,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

	_, err := s.db.Exec(createTableSQL)
	s.Require().NoError(err)
	fmt.Println("📊 Table 'users' created")
}

// ============================================
// Helper: Cleanup Table
// ============================================
func (s *UserIntegrationSuite) cleanupTable() {
	_, err := s.db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE;")
	s.Require().NoError(err)
}

// ============================================
// TEST 1: Create User
// ============================================
func (s *UserIntegrationSuite) TestCreateUser() {
	s.cleanupTable()

	query := `INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id, name, email, age`

	var id int
	var name, email string
	var age int
	err := s.db.QueryRow(query, "John Doe", "john@example.com", 30).Scan(&id, &name, &email, &age)

	s.Require().NoError(err)
	assert.Equal(s.T(), "John Doe", name)
	assert.Equal(s.T(), "john@example.com", email)
	assert.Equal(s.T(), 30, age)
	assert.Greater(s.T(), id, 0)

	fmt.Println("✅ TestCreateUser passed")
}

// ============================================
// TEST 2: Get User by Email
// ============================================
func (s *UserIntegrationSuite) TestGetUserByEmail() {
	s.cleanupTable()

	// Insert user first
	_, err := s.db.Exec("INSERT INTO users (name, email, age) VALUES ($1, $2, $3)",
		"Jane Doe", "jane@example.com", 25)
	s.Require().NoError(err)

	// Query user
	query := `SELECT id, name, email, age FROM users WHERE email = $1`
	var id int
	var name, email string
	var age int
	err = s.db.QueryRow(query, "jane@example.com").Scan(&id, &name, &email, &age)

	s.Require().NoError(err)
	assert.Equal(s.T(), "Jane Doe", name)
	assert.Equal(s.T(), "jane@example.com", email)
	assert.Equal(s.T(), 25, age)

	fmt.Println("✅ TestGetUserByEmail passed")
}

// ============================================
// TEST 3: Update User
// ============================================
func (s *UserIntegrationSuite) TestUpdateUser() {
	s.cleanupTable()

	// Insert user
	var id int
	err := s.db.QueryRow("INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id",
		"Old Name", "old@example.com", 20).Scan(&id)
	s.Require().NoError(err)

	// Update user
	_, err = s.db.Exec("UPDATE users SET name = $1, age = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3",
		"New Name", 35, id)
	s.Require().NoError(err)

	// Verify update
	var newName string
	var newAge int
	err = s.db.QueryRow("SELECT name, age FROM users WHERE id = $1", id).Scan(&newName, &newAge)
	s.Require().NoError(err)
	assert.Equal(s.T(), "New Name", newName)
	assert.Equal(s.T(), 35, newAge)

	fmt.Println("✅ TestUpdateUser passed")
}

// ============================================
// TEST 4: Delete User
// ============================================
func (s *UserIntegrationSuite) TestDeleteUser() {
	s.cleanupTable()

	// Insert user
	var id int
	err := s.db.QueryRow("INSERT INTO users (name, email, age) VALUES ($1, $2, $3) RETURNING id",
		"Delete Me", "delete@example.com", 40).Scan(&id)
	s.Require().NoError(err)

	// Delete user
	_, err = s.db.Exec("DELETE FROM users WHERE id = $1", id)
	s.Require().NoError(err)

	// Verify deletion
	var count int
	err = s.db.QueryRow("SELECT COUNT(*) FROM users WHERE id = $1", id).Scan(&count)
	s.Require().NoError(err)
	assert.Equal(s.T(), 0, count)

	fmt.Println("✅ TestDeleteUser passed")
}

// ============================================
// TEST 5: Duplicate Email Constraint
// ============================================
func (s *UserIntegrationSuite) TestDuplicateEmailConstraint() {
	s.cleanupTable()

	// Insert first user
	_, err := s.db.Exec("INSERT INTO users (name, email, age) VALUES ($1, $2, $3)",
		"User 1", "duplicate@example.com", 25)
	s.Require().NoError(err)

	// Try insert duplicate email - should fail
	_, err = s.db.Exec("INSERT INTO users (name, email, age) VALUES ($1, $2, $3)",
		"User 2", "duplicate@example.com", 30)

	assert.Error(s.T(), err, "Should fail on duplicate email")
	fmt.Println("✅ TestDuplicateEmailConstraint passed")
}

// ============================================
// TEST 6: Bulk Insert Performance
// ============================================
func (s *UserIntegrationSuite) TestBulkInsertPerformance() {
	s.cleanupTable()

	startTime := time.Now()

	// Insert 100 users
	for i := 0; i < 100; i++ {
		_, err := s.db.Exec("INSERT INTO users (name, email, age) VALUES ($1, $2, $3)",
			fmt.Sprintf("User %d", i),
			fmt.Sprintf("user%d@example.com", i),
			i)
		s.Require().NoError(err)
	}

	duration := time.Since(startTime)

	// Assert: 100 inserts should complete in under 5 seconds
	assert.Less(s.T(), duration.Milliseconds(), int64(5000))

	// Verify count
	var count int
	err := s.db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	s.Require().NoError(err)
	assert.Equal(s.T(), 100, count)

	fmt.Printf("✅ TestBulkInsertPerformance passed (%dms)\n", duration.Milliseconds())
}

// ============================================
// TEST 7: Find Users by Age Range
// ============================================
func (s *UserIntegrationSuite) TestFindUsersByAgeRange() {
	s.cleanupTable()

	// Insert test data
	users := []struct {
		name  string
		email string
		age   int
	}{
		{"Alice", "alice@example.com", 20},
		{"Bob", "bob@example.com", 30},
		{"Charlie", "charlie@example.com", 40},
		{"David", "david@example.com", 50},
	}

	for _, u := range users {
		_, err := s.db.Exec("INSERT INTO users (name, email, age) VALUES ($1, $2, $3)",
			u.name, u.email, u.age)
		s.Require().NoError(err)
	}

	// Query users with age between 25 and 45
	query := `SELECT name, age FROM users WHERE age >= $1 AND age <= $2 ORDER BY age`
	rows, err := s.db.Query(query, 25, 45)
	s.Require().NoError(err)
	defer rows.Close()

	var results []struct {
		name string
		age  int
	}

	for rows.Next() {
		var name string
		var age int
		err := rows.Scan(&name, &age)
		s.Require().NoError(err)
		results = append(results, struct {
			name string
			age  int
		}{name, age})
	}

	assert.Equal(s.T(), 2, len(results))
	assert.Equal(s.T(), "Bob", results[0].name)
	assert.Equal(s.T(), 30, results[0].age)
	assert.Equal(s.T(), "Charlie", results[1].name)
	assert.Equal(s.T(), 40, results[1].age)

	fmt.Println("✅ TestFindUsersByAgeRange passed")
}

// ============================================
// Run Test Suite
// ============================================
func TestUserIntegrationSuite(t *testing.T) {
	suite.Run(t, new(UserIntegrationSuite))
}
