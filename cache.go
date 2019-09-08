package main

type Cache interface {
	Get(key string) (val interface{}, ok bool)
	Put(key string, val interface{})
}
