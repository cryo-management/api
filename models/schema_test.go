package models

import (
	"testing"

	"database/sql"
	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	// PostgresSQL library
	_ "github.com/lib/pq"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type SchemaTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
	host                          string
	port                          int
	user                          string
	password                      string
	dbname                        string
	psqlInfo                      string
	Database                      struct{}
	conn                          *sql.DB
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *SchemaTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5

	suite.host = "cryo.cdnm8viilrat.us-east-2.rds-preview.amazonaws.com"
	suite.port = 5432
	suite.user = "cryoadmin"
	suite.password = "x3FhcrWDxnxCq9p"
	suite.dbname = "cryo"

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		suite.host, suite.port, suite.user, suite.password, suite.dbname)

	suite.conn, _ = sql.Open("postgres", psqlInfo)
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *SchemaTestSuite) TestGetAll() {
	// s := new(Schema)
	// schemaList, err := s.GetAll()
	// assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}

func (suite *SchemaTestSuite) Test002Example() {
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSchemaTestSuite(t *testing.T) {
	suite.Run(t, new(SchemaTestSuite))
}
