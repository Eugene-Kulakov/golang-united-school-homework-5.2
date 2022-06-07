package cache

import "time"

type StringTime struct {
	Str  string
	Time time.Time
}

type Cache struct {
	ValuesMap map[string]StringTime
}

func NewCache() Cache {
	return Cache{}
}

func (c *Cache) Get(key string) (string, bool) {
	st, found := c.ValuesMap[key]
	if !found {
		return "", false
	}
	if st.Time.After(time.Now()) {
		delete(c.ValuesMap, key)
		return "", false
	}
	return st.Str, true
}

func (c *Cache) Put(key, value string) {
	t := time.Unix(1<<63-1, 0)
	st := StringTime{value, t}
	c.ValuesMap[key] = st
}

func (c *Cache) Keys() []string {
	var ans []string
	for key, value := range c.ValuesMap {
		if value.Time.Before(time.Now()) {
			ans = append(ans, key)
		}
	}
	return ans
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	st := StringTime{value, deadline}
	c.ValuesMap[key] = st
}
