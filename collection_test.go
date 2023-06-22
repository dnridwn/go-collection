package collection

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKey(t *testing.T) {
	cl := New("Test", 1, 2, 3, true)
	testCases := []struct {
		name     string
		c        Collection
		t        Entity
		expected int
	}{
		{
			name:     "Test Exists",
			c:        cl,
			t:        "Test",
			expected: 0,
		},
		{
			name:     "Test Not Exists",
			c:        cl,
			t:        99,
			expected: -1,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.c.FindKey(tt.t))
		})
	}
}

func TestFindKeys(t *testing.T) {
	cl := New("Test", 1, 2, 3, "Test", true, 0.00, -88)
	testCases := []struct {
		name     string
		c        Collection
		t        Entity
		expected []int
	}{
		{
			name:     "Test Exists",
			c:        cl,
			t:        "Test",
			expected: []int{0, 4},
		},
		{
			name:     "Test Not Exists",
			c:        cl,
			t:        -99,
			expected: []int{},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.c.FindKeys(tt.t))
		})
	}
}

func TestFind(t *testing.T) {
	cl := New(1, 2, 4, "test", true, "test 2")

	testCases := []struct {
		name     string
		c        Collection
		k        int
		expected Entity
	}{
		{
			name:     "Test exists",
			c:        cl,
			k:        3,
			expected: "test",
		},
		{
			name:     "Test not exists",
			c:        cl,
			k:        99,
			expected: nil,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.c.Find(tt.k))
		})
	}
}

func TestGet(t *testing.T) {
	testCases := []struct {
		name     string
		c        Collection
		k        int
		expected []Entity
	}{
		{
			name:     "Test has data",
			c:        New(1, 2, 4, "test", true, "test 2"),
			k:        3,
			expected: []Entity{1, 2, 4, "test", true, "test 2"},
		},
		{
			name:     "Test empty",
			c:        New(),
			k:        3,
			expected: []Entity{},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.c.Get())
		})
	}
}

func TestUnique(t *testing.T) {
	testCases := []struct {
		name     string
		c        Collection
		expected []Entity
	}{
		{
			name:     "Test has duplicates",
			c:        New("1", 2, "test", 3, 2, "test"),
			expected: []Entity{"1", 2, "test", 3},
		},
		{
			name:     "Test has not duplicates",
			c:        New("1", 2, "test", 3),
			expected: []Entity{"1", 2, "test", 3},
		},
		{

			name:     "Test empty",
			c:        New(),
			expected: []Entity{},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.c.Unique().Get())
		})
	}
}

func TestMap(t *testing.T) {
	testCases := []struct {
		name        string
		c           Collection
		mapCallback MapCallback
		expected    []Entity
	}{
		{
			name: "Test 1",
			c:    New(1, 2, 3, 4, 5),
			mapCallback: func(key int, entity Entity) Entity {
				return true
			},
			expected: []Entity{true, true, true, true, true},
		},
		{
			name: "Test 2",
			c:    New(1, 2, 3, 4, 5),
			mapCallback: func(key int, entity Entity) Entity {
				if key == 2 {
					return false
				}
				return entity
			},
			expected: []Entity{1, 2, false, 4, 5},
		},
		{
			name: "Test empty",
			c:    New(),
			mapCallback: func(key int, entity Entity) Entity {
				return entity
			},
			expected: []Entity{},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.c.Map(tt.mapCallback).Get())
		})
	}
}

func TestFilter(t *testing.T) {
	testCases := []struct {
		name           string
		c              Collection
		filterCallback FilterCallback
		expected       []Entity
	}{
		{
			name: "Test Filter Odd",
			c:    New(1, 2, 3, 4, 5, 6),
			filterCallback: func(key int, entity Entity) bool {
				return entity.(int)%2 == 1
			},
			expected: []Entity{1, 3, 5},
		},
		{
			name: "Test Filter Only String",
			c:    New(1, 2, "test", 4, 5, 6),
			filterCallback: func(key int, entity Entity) bool {
				return reflect.TypeOf(entity).Kind() == reflect.String
			},
			expected: []Entity{"test"},
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.c.Filter(tt.filterCallback).Get())
		})
	}
}
