package queue

import (
	"log"
)

type Queue interface {
	Push(key interface{})
	Pop() interface{}
	Contains(key interface{}) bool
	Len() int
	Keys() []interface{}
}

type Queues struct {
	keys []interface{}
}

func New(size int) Queue {
	return nil
}

func Keys() []interface{} {
	var qu Queues
	key := qu.keys
	return key
}

func Push(key interface{}) {
	var q Queue
	var qu Queues
	if Contains(key) == true {
		log.Fatal("already queued")
	} else {
		qu.keys = append(q.Keys(), key)
	}
}

func Pop() interface{} {
	var q Queue
	x, _ := q.Keys()[len(q.Keys())-1], q.Keys()[:len(q.Keys())-1]
	return x
}

func Contains(key interface{}) bool {
	var q Queue
	if contains(q.Keys(), key) {
		return true
	} else {
		return false
	}
}

func Len() int {
	var q Queue
	length := len(q.Keys())
	return length
}

func contains(slice []interface{}, item interface{}) bool {
	set := make(map[interface{}]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
