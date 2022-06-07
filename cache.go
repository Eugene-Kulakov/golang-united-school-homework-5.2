package cache

import (
	"time"
)

type StringTime struct {
	Str  string
	Time time.Time
}

type Cache struct {
	valuesMap map[string]StringTime
}

func (c *Cache) Init() {
	c.valuesMap = make(map[string]StringTime)
}

func NewCache() Cache {
	var c Cache
	c.Init()
	return c
}

func (c *Cache) Get(key string) (string, bool) {
	st, found := c.valuesMap[key]
	if !found {
		return "", false
	}
	if !st.Time.IsZero() && st.Time.After(time.Now()) {
		delete(c.valuesMap, key)
		return "", false
	}
	return st.Str, true
}

func (c *Cache) Put(key, value string) {
	var t time.Time
	st := StringTime{value, t}
	c.valuesMap[key] = st
}

func (c *Cache) Keys() []string {
	var ans []string
	for key, value := range c.valuesMap {
		if value.Time.IsZero() || value.Time.Before(time.Now()) {
			ans = append(ans, key)
		} else {
			delete(c.valuesMap, key)
		}
	}
	return ans
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	st := StringTime{value, deadline}
	c.valuesMap[key] = st
}
