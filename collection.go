package collection

type Entity any
type MapCallback func(key int, entity Entity) Entity
type FilterCallback func(key int, entity Entity) bool

type Collection struct {
	data []Entity
}

func (c *Collection) Find(k int) Entity {
	if k < 0 || k >= len(c.data) {
		return nil
	}
	return c.data[k]
}

func (c *Collection) Get() []Entity {
	return c.data
}

func (c *Collection) FindKey(t Entity) int {
	for k, v := range c.data {
		if v == t {
			return k
		}
	}
	return -1
}

func (c *Collection) FindKeys(t Entity) []int {
	keys := make([]int, 0)
	for k, v := range c.data {
		if v == t {
			keys = append(keys, k)
		}
	}
	return keys
}

func (c *Collection) Unique() *Collection {
	nData := make([]Entity, 0)
	seen := make(map[Entity]bool)
	for _, v := range c.data {
		if seen[v] {
			continue
		}
		seen[v] = true
		nData = append(nData, v)
	}
	c.data = nData
	return c
}

func (c *Collection) Map(callback MapCallback) *Collection {
	for k, v := range c.data {
		c.data[k] = callback(k, v)
	}
	return c
}

func (c *Collection) Filter(callback FilterCallback) *Collection {
	nData := make([]Entity, 0)
	for k, v := range c.data {
		if callback(k, v) {
			nData = append(nData, v)
		}
	}
	c.data = nData
	return c
}

func (c *Collection) First() Entity {
	if len(c.data) == 0 {
		return nil
	}
	return c.data[0]
}

func (c *Collection) Last() Entity {
	if len(c.data) == 0 {
		return nil
	}
	return c.data[len(c.data)-1]
}

func (c *Collection) Reverse() *Collection {
	nData := make([]Entity, 0)
	for i := len(c.data) - 1; i >= 0; i-- {
		nData = append(nData, c.data[i])
	}
	c.data = nData
	return c
}

func New(nData ...Entity) *Collection {
	if len(nData) == 0 {
		nData = []Entity{}
	}
	return &Collection{
		data: nData,
	}
}
