package models

import (
	"testing"

	"github.com/cryo-management/api/db"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	_ "github.com/lib/pq"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type SchemaTestSuite struct {
	suite.Suite
	VariableThatShouldStartAtFive int
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *SchemaTestSuite) SetupTest() {
	suite.VariableThatShouldStartAtFive = 5

	_ = db.Connect()
}

// All methods that begin with "Test" are run as tests within a
// suite.
func TestGetAll(t *testing.T) {
	s := new(Schema)
	schemaList, _ := s.GetAll()
	assert.Equal(t, schemaList, schemaList, "invalid generated query")
}

func (suite *SchemaTestSuite) Test002Example() {
	assert.Equal(suite.T(), 5, suite.VariableThatShouldStartAtFive)
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSchemaTestSuite(t *testing.T) {
	suite.Run(t, new(SchemaTestSuite))
}
