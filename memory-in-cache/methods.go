package memoryincache

func New() map[string]int {

	return make(map[string]int)

}

func Set(mapp map[string]int, key string, value int) {
	mapp[key] = value

}

func Get(key string, mapp map[string]int) int {
	GetValue := mapp[key]

	return GetValue
}

func Delete(key string, mapp map[string]int) {
	delete(mapp, key)

}
