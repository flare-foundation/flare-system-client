package utils

import (
	"fmt"
	"testing"
)

type cacheTestKey struct {
	s string
	i int
}

func TestCache(t *testing.T) {
	cache := NewCache[cacheTestKey, int]()

	// insert 100 items
	for i := 0; i < 100; i++ {
		// random string
		key := cacheTestKey{fmt.Sprintf("string-%d", i), i}
		value := -i
		cache.Add(key, value)
	}

	// check 10 items
	for i := 0; i < 100; i += 10 {
		value, ok := cache.Get(cacheTestKey{fmt.Sprintf("string-%d", i), i})
		if !ok {
			t.Fatalf("Expected key %d to exist", i)
		}
		if value != -i {
			t.Fatalf("Expected value %d, got %d", -i, value)
		}
	}

	_, ok := cache.Get(cacheTestKey{"string-100", 100})
	if ok {
		t.Fatalf("Expected key %d to not exist", 100)
	}

	cache.RemoveAccessed()

	for i := 0; i < 100; i++ {
		value, ok := cache.Get(cacheTestKey{fmt.Sprintf("string-%d", i), i})
		if i%10 == 0 {
			if ok {
				t.Fatalf("Expected key %d to not exist", i)
			}
		} else {
			if !ok {
				t.Fatalf("Expected key %d to exist", i)
			}
			if value != -i {
				t.Fatalf("Expected value %d, got %d", -i, value)
			}
		}
	}

}
