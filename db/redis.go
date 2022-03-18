package db

import (
	"fmt"
	//"log"
	"github.com/gomodule/redigo/redis"
	//"isMutant/models"
)

var pool = newPool()

func CreateDNA(dna []string, isMutant bool) {
	client := pool.Get()
	defer client.Close()

	var dnaKey string
	for _, row := range dna {
		dnaKey += row + "+"
	}
	fmt.Println(dnaKey)
	_, err := client.Do("SET", dnaKey, isMutant)
	if err != nil {
		panic(err)
	}
}

func GetListDNA() (int, int, float32) {
	client := pool.Get()
	defer client.Close()
	
	iter := 0
    keys := []string{}
    for {
        arr, err := redis.Values(client.Do("SCAN", iter, "MATCH", "*"))
        if err != nil {
            fmt.Println("Error: ", err)
        }
		iter, _ = redis.Int(arr[0], nil)
		k, _ := redis.Strings(arr[1], nil)
		keys = append(keys, k...)
        if iter == 0 {
            break
        }
    }

	var isMutant int
	var notMutant int
	for _, keyDna := range keys {
		mutant, err := client.Do("GET", keyDna)
		if err != nil {
			panic(err)
		}
		m, _ := redis.Bool(mutant, nil)
		if m {
			isMutant++
		}else{
			notMutant++
		}
	}
	fmt.Printf("%v son mutantes y %v no lo son", isMutant, notMutant)
	var ratio float32
	ratio = float32(isMutant)/(float32(isMutant)+float32(notMutant)) 
	return isMutant, (isMutant+notMutant), ratio
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle: 80,
		MaxActive: 12000,
		Dial: func () (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}
