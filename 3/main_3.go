package main

import "fmt"

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{data: make(map[string]int)}
}

func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int {
	result := make(map[string]int, len(m.data))
	for key, value := range m.data {
		result[key] = value
	}
	return result
}

func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.data[key]
	return ok
}

func (m *StringIntMap) Get(key string) (int, bool) {
	value, ok := m.data[key]
	return value, ok
}

func main() {
	m := NewStringIntMap()
	m.Add("one", 1)
	m.Add("two", 2)

	fmt.Println("one:", m.Exists("one"))
	value, ok := m.Get("two")
	fmt.Println("two:", value, ok)
	fmt.Println("copy:", m.Copy())

	m.Remove("one")
	fmt.Println("after remove:", m.Copy())
}
