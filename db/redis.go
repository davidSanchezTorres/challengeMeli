package db

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var pool = newPool()

//CreateDNA Guarda un nuevo DNA en la base de datos
func CreateDNA(dna []string, isMutant bool) {
	client := pool.Get()
	defer client.Close()

	var dnaKey string
	for _, row := range dna {
		dnaKey += row + "+"
	}
	_, err := client.Do("SET", dnaKey, isMutant)
	if err != nil {
		panic(err)
	}
}

//GetListDNA Obtiene la lista de DNAs
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
		} else {
			notMutant++
		}
	}
	var ratio float32
	if (float32(isMutant) + float32(notMutant)) == 0 {
		return 0, 0, 0
	}
	ratio = float32(isMutant) / (float32(isMutant) + float32(notMutant))
	return isMutant, (isMutant + notMutant), ratio
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}
