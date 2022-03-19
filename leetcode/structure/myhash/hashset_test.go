package myhash

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyHashSet_Add(t *testing.T) {
	obj := Constructor()
	obj.Add(3)
	obj.Add(6)
	obj.Add(9)
	obj.Add(12)
	fmt.Println(obj.data)
	assert.True(t, obj.Contains(3))
	assert.True(t, obj.Contains(6))
	assert.True(t, obj.Contains(9))
	assert.True(t, obj.Contains(12))

	obj.Remove(3)
	assert.False(t, obj.Contains(3))
	assert.True(t, obj.Contains(6))
	assert.True(t, obj.Contains(9))
	assert.True(t, obj.Contains(12))

	obj.Add(3)
	obj.Remove(9)
	obj.Add(15)
	obj.Add(18)
	obj.Add(21)

	obj.Add(9)
	obj.Remove(15)
	obj.Remove(9)
	assert.False(t, obj.Contains(9))

	obj.Add(9)
}

func TestMyHashSet_Remove(t *testing.T) {
	obj := Constructor()
	obj.Add(1)
	obj.Add(2)
	obj.Add(4)
	obj.Remove(4)
	obj.Remove(5)
	obj.Remove(2)
	obj.Remove(1)
	fmt.Println(obj.data)
}

func TestMyHashSet_Contains(t *testing.T) {
	obj := Constructor()
	obj.Add(1)
	obj.Add(2)
	obj.Add(4)
	obj.Remove(1)
	con := obj.Contains(4)
	assert.True(t, con)

	obj.Add(5)
	assert.True(t, obj.Contains(5))
	obj.Remove(5)
	assert.False(t, obj.Contains(5))
	assert.False(t, obj.Contains(1))
}

/*
["MyHashSet","add","remove","add","remove","remove","add","add","add","add","remove"]
[[],           [9], [19],    [14],[19],    [9],     [0],   [3],  [4],  [0],   [9]]
*/
func TestName(t *testing.T) {
	obj := Constructor()
	obj.Add(9)
	obj.Remove(19)
	obj.Add(14)
	obj.Remove(19)
	obj.Remove(9)
	obj.Add(0)
	obj.Add(3)
	obj.Add(4)
	obj.Add(0)
	obj.Remove(9)
	fmt.Println(obj.data)
}

//["MyHashSet","add","contains","add","contains","remove","add","contains","add","add","add","add","add","add","contains","add","add","add","contains","remove","contains","contains","add","remove","add","remove","add","remove","add","contains","add","add","contains","add","add","add","add","remove","contains","add","contains","add","add","add","remove","remove","add","contains","add","add","contains","remove","add","contains","add","remove","remove","add","contains","add","contains","contains","add","add","remove","remove","add","remove","add","add","add","add","add","add","remove","remove","add","remove","add","add","add","add","contains","add","remove","remove","remove","remove","add","add","add","add","contains","add","add","add","add","add","add","add","add"]
//[[],[58],[0],[14],[58],[91],[6],[58],[66],[51],[16],[40],[52],[48],[40],[42],[85],[36],[16],[0],[43],[6],[3],[25],[99],[66],[60],[58],[97],[3],[35],[65],[40],[41],[10],[37],[65],[37],[40],[28],[60],[30],[63],[76],[90],[3],[43],[81],[61],[39],[75],[10],[55],[92],[71],[2],[20],[7],[55],[88],[39],[97],[44],[1],[51],[89],[37],[19],[3],[13],[11],[68],[18],[17],[41],[87],[48],[43],[68],[80],[35],[2],[17],[71],[90],[83],[42],[88],[16],[37],[33],[66],[59],[6],[79],[77],[14],[69],[36],[21],[40]]
func TestName2(t *testing.T) {
	comand := []string{"add", "contains", "add", "contains", "remove", "add", "contains", "add", "add", "add", "add", "add", "add", "contains", "add", "add", "add", "contains", "remove", "contains", "contains", "add", "remove", "add", "remove", "add", "remove", "add", "contains", "add", "add", "contains", "add", "add", "add", "add", "remove", "contains", "add", "contains", "add", "add", "add", "remove", "remove", "add", "contains", "add", "add", "contains", "remove", "add", "contains", "add", "remove", "remove", "add", "contains", "add", "contains", "contains", "add", "add", "remove", "remove", "add", "remove", "add", "add", "add", "add", "add", "add", "remove", "remove", "add", "remove", "add", "add", "add", "add", "contains", "add", "remove", "remove", "remove", "remove", "add", "add", "add", "add", "contains", "add", "add", "add", "add", "add", "add", "add", "add"}
	for _, cmd := range comand {
		switch cmd {

		}
	}
}
