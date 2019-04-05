package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostSchema(t *testing.T) {
	assert.HTTPSuccess(t, postSchema, "POST", "/schema", nil)

}
