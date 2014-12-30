package cache

import (
	"fmt"
	"testing"
)

type testElement struct {
	k, v string
}

func (el *testElement) Key() string {
	return el.k
}

func (el *testElement) String() string {
	return el.v
}

func elval(el CacheElement) string {
	if el != nil {
		if xstringer, ok := el.(fmt.Stringer); ok {
			return xstringer.String()
		}
	}

	return ""
}

func TestLRUCacheInsertion(t *testing.T) {
	cache := NewLRUCache(1)
	cache.See(&testElement{"foo", "bar"})
	x := cache.Get("foo")

	if xval := elval(x); xval != "bar" {
		t.Fatal("expected \"bar\", got", xval)
	}
}

func TestLRUCacheOverwrite(t *testing.T) {
	cache := NewLRUCache(1)

	cache.See(&testElement{"foo", "bar"})
	x := cache.Get("foo")
	if xval := elval(x); xval != "bar" {
		t.Fatal("expected \"bar\", got", xval)
	}

	cache.See(&testElement{"foo", "notbar"})
	x = cache.Get("foo")
	if xval := elval(x); xval != "notbar" {
		t.Fatal("expected \"notbar\", got", xval)
	}
}

func TestLRUCacheExpiry(t *testing.T) {
	cache := NewLRUCache(1)

	cache.See(&testElement{"foo", "bar"})
	x := cache.Get("foo")
	if xval := elval(x); xval != "bar" {
		t.Fatal("expected \"bar\", got", xval)
	}

	cache.See(&testElement{"bar", "baz"})
	x = cache.Get("bar")
	if xval := elval(x); xval != "baz" {
		t.Fatal("expected \"baz\", got", xval)
	}

	x = cache.Get("foo")
	if x != nil {
		t.Fatal("expected nil, got", x)
	}
}

func TestLRUCacheInvalidCapacity(t *testing.T) {
	c0 := NewLRUCache(-1)
	if c0 != nil {
		t.Fatal("non-nil cache on invalid cache size")
	}

	c1 := NewLRUCache(0)
	if c1 != nil {
		t.Fatal("non-nil cache on invalid cache size")
	}
}

func TestLRUCacheSeeNil(t *testing.T) {
	cache := NewLRUCache(1)
	cache.See(nil)
}
