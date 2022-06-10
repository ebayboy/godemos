package main

import (
	"encoding/gob"
	"log"
	"os"
	"time"

	lru "github.com/hashicorp/golang-lru"
	"github.com/patrickmn/go-cache"
)

var IPThreatLRUCache *lru.Cache
var IPThreatLRUCache2 *lru.Cache
var BlkDomainCache *cache.Cache

func gocacheSave() {
	gob.Register(lru.Cache{})

	BlkDomainCache = cache.New(time.Duration(3600)*time.Minute, time.Minute)

	log.Println("BlkDomainCache.ItemCount()):", BlkDomainCache.ItemCount())

	IPThreatLRUCache, err := lru.New(3)
	log.Println(IPThreatLRUCache, err)

	IPThreatLRUCache.Add("1.1.1.1", []float64{1, 507})
	IPThreatLRUCache.Add("2.2.2.2", []float64{0, 998})

	log.Println("===========")
	log.Println(IPThreatLRUCache.Keys())
	log.Println(IPThreatLRUCache.Len())
	log.Println(IPThreatLRUCache.GetOldest())

	//set cache
	BlkDomainCache.SetDefault("key1", IPThreatLRUCache)
	log.Println("BlkDomainCache.ItemCount()):", BlkDomainCache.ItemCount())

	t1, exist := BlkDomainCache.Get("key1")
	if !exist {
		log.Println("Error:", err.Error())
		return
	}
	ch := t1.(*lru.Cache)
	val, ok := ch.Get("6")
	if !ok {
		log.Panicln("error not ok!")
		return
	}
	log.Println("h.Get:", val)

	file := "./BlkDomainCache.bin"
	if err := BlkDomainCache.SaveFile(file); err != nil {
		log.Println("error SaveFile:", err.Error())
		return
	}

	if err := BlkDomainCache.LoadFile(file); err != nil {
		log.Println("LoadFil error:", err.Error())
		return
	}

	val2, exist := BlkDomainCache.Get("key1")
	if !exist {
		log.Println("not exist!!!")
		return
	}
	blkCache := val2.(lru.Cache)
	log.Println("blkCache:", blkCache)
	log.Println(blkCache.Keys())
	log.Println(blkCache.Len())
}

func saveGobFile(lruCache *lru.Cache, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	saveMap := make(map[string][]float64, lruCache.Len())
	for _, k := range lruCache.Keys() {
		v, ok := lruCache.Peek(k)
		if !ok {
			continue
		}
		saveMap[k.(string)] = v.([]float64)
	}

	enc := gob.NewEncoder(file)
	if err := enc.Encode(saveMap); err != nil {
		return err
	}

	return nil
}

func loadGobFile(lruCache *lru.Cache, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var m map[string][]float64
	dec := gob.NewDecoder(file)
	err = dec.Decode(&m)
	if err != nil {
		return err
	}

	for k, v := range m {
		lruCache.Add(k, v)
	}

	return nil
}

func main() {
	filename := "gob.bin"
	IPThreatLRUCache, err := lru.New(3)
	log.Println(IPThreatLRUCache, err)
	IPThreatLRUCache.Add("1.1.1.1", []float64{1, 507})
	IPThreatLRUCache.Add("2.2.2.2", []float64{0, 998})

	if err := saveGobFile(IPThreatLRUCache, filename); err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	lruNew, err := lru.New(3)
	if err := loadGobFile(lruNew, filename); err != nil {
		log.Println("Error: ", err.Error())
		return
	}
	l, ok := lruNew.Get("1.1.1.1")
	if !ok {
		log.Println("Error : not exist!", ok)
		return
	}
	log.Println("1.1.1.1:", l.([]float64))

	l, ok = lruNew.Get("2.2.2.2")
	if !ok {
		log.Println("Error : not exist!", ok)
		return
	}
	log.Println("2.2.2.2:", l.([]float64))

	l, ok = lruNew.Get("3.3.3.3")
	if !ok {
		log.Println("Error : not exist!", ok)
		return
	}
	log.Println("3.3.3.3:", l.([]float64))

}
