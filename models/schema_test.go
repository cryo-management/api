package models

import (
	"testing"

	"github.com/cryo-management/api/common"
	"github.com/cryo-management/api/db"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	_ "github.com/lib/pq"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type SchemaTestSuite struct {
	suite.Suite
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (suite *SchemaTestSuite) SetupTest() {
	_ = db.Connect()
	common.Session.User.ID = "059fa339-025c-4104-ab55-c764d3028f63"
	common.Session.User.FirstName = "Bruno"
	common.Session.User.LastName = "Piaui"
	common.Session.User.Email = "brunopiaui@gmail.com"
	common.Session.User.Language = "pt-br"
}

// All methods that begin with "Test" are run as tests within a
// suite.
func (suite *SchemaTestSuite) TestLoad() {
	s := new(Schema)
	s.Load("SC016")
	assert.Equal(suite.T(), "SC016", s.Code, "invalid generated query")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSchemaTestSuite(t *testing.T) {
	suite.Run(t, new(SchemaTestSuite))
}
