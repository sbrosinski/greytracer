package tuple

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddingTwoTuples(t *testing.T) {
	var t1 = NewTuple(3, -2, 5, 1)
	var t2 = NewTuple(-2, 3, 1, 0)
	var result = Add(t1, t2)
	assert.True(t, reflect.DeepEqual(NewTuple(1, 1, 6, 1), result))
}
