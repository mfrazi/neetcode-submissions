type MyHashMap struct {
	data map[int]int
}

func Constructor() MyHashMap {
    return MyHashMap{
		data: map[int]int{},
	}
}

func (this *MyHashMap) Put(key int, value int) {
    this.data[key] = value
}

func (this *MyHashMap) Get(key int) int {
    if val, found := this.data[key]; found {
		return val
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
    delete(this.data, key)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */