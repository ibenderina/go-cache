package golang_ninja

type Cache struct {
	m map[string]interface{}
}

func (c Cache) Set(key string, value interface{}) {
	c.m[key] = value
}

func (c Cache) Get(key string) interface{} {
	value, ok := c.m[key]
	if ok {
		return value
	}
	return nil
}

func (c Cache) Delete(key string) {
	delete(c.m, key)
}

func New() *Cache {
	return &Cache{
		m: map[string]interface{}{},
	}
}
