# cache
--
    import "cache"


## Usage

#### type Cache

```go
type Cache interface {
	// Observe an element for caching
	See(el CacheElement)
	// returns an element from the cache based on it's key, or nil
	// if element does not exist in cache
	Get(key string) CacheElement
}
```


#### type CacheElement

```go
type CacheElement interface {
	Key() string
}
```


#### type LRUCache

```go
type LRUCache struct {
}
```

Cache which uses a Least Recently Used (LRU) eviction policy

#### func  NewLRUCache

```go
func NewLRUCache(capacity int) *LRUCache
```
Returns a new LRU Cache

the cache is safe for concurrent use

#### func (*LRUCache) Get

```go
func (c *LRUCache) Get(key string) (el CacheElement)
```

#### func (*LRUCache) See

```go
func (c *LRUCache) See(el CacheElement)
```
