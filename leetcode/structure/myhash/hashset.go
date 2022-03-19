package myhash

type MyHashSet struct {
	data  []int
	count int
	slot  []slot
}

type slot struct {
	indexes []int
}

var initCapacity int = 3

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	set := MyHashSet{data: make([]int, initCapacity), count: 0, slot: make([]slot, initCapacity)}
	for i := 0; i < len(set.data); i++ {
		set.data[i] = -1
		set.slot[i].indexes = []int{}
	}
	return set
}

func (this *MyHashSet) Add(key int) {
	index := getIndex(key, this)
	if this.Contains(key) {
		return
	}

	if this.data[index] == -1 {
		this.data[index] = key
		this.slot[index].indexes = append(this.slot[index].indexes, index)
		this.count++
		return
	}

	if this.count+1 > cap(this.data) {
		increase(this)
		this.Add(key)
		return
	}

	for i := index + 1; i <= cap(this.data) && i != index; i++ {
		if i == cap(this.data) {
			i = -1
			continue
		}

		if this.data[i] == -1 {
			this.data[i] = key
			this.slot[index].indexes = append(this.slot[index].indexes, i)
			this.count++
			return
		}
	}
}

func getIndex(key int, this *MyHashSet) int {
	index := key
	if index >= cap(this.data) {
		index = index % cap(this.data)
	}
	return index
}

func increase(this *MyHashSet) {
	capacity := cap(this.data) * 2
	data := make([]int, capacity)
	slots := make([]slot, capacity)
	for i := 0; i < cap(data); i++ {
		data[i] = -1
		slots[i].indexes = []int{}
	}

	for i := range this.data {
		index := this.data[i] % cap(data)
		if index == -1 {
			continue
		}
		if data[index] == -1 {
			data[index] = this.data[i]
			slots[index].indexes = append(slots[index].indexes, index)
			continue
		}

		for j := index + 1; j <= cap(data) && j != index; j++ {
			if j == cap(data) {
				j = -1
				continue
			}

			if data[j] == -1 {
				data[j] = this.data[i]
				slots[index].indexes = append(slots[index].indexes, j)
				break
			}
		}
	}

	this.data = data
	this.slot = slots

}

func (this *MyHashSet) Remove(key int) {
	index := getIndex(key, this)
	indexes := this.slot[index].indexes
	for i := range indexes {
		i2 := indexes[i]
		if i2 == -1 {
			continue
		}
		if this.data[i2] == key {
			this.data[indexes[i]] = -1
			indexes[i] = -1
		}
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	index := getIndex(key, this)
	indexes := this.slot[index].indexes
	for _, i := range indexes {
		if i == -1 {
			continue
		}
		if this.data[i] == key {
			return true
		}
	}
	return false
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
