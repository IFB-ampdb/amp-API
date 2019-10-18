package redis

import (
	"encoding/json"
	"fmt"

	"github.com/ifbampdb/amp-core/app"

	"github.com/go-redis/redis"
)

const table = "peptides"

type peptideRepository struct {
	connection *redis.Client
}

func RedisConnection(url, password string) *redis.Client {
	fmt.Println("Connecting to Redis DB")
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: password, // no password set
		DB:       0,        // use default DB
	})
	err := client.Ping().Err()

	if err != nil {
		panic(err)
	}
	return client
}

func NewRedisPeptideRepository(connection *redis.Client) app.PeptideRepository {
	return &peptideRepository{
		connection,
	}
}

func (r *peptideRepository) Create(peptide *app.Peptide) error {
	encoded, err := json.Marshal(peptide)

	if err != nil {
		return err
	}

	r.connection.HSet(table, peptide.PdbID, encoded) //Don't expire
	return nil
}

func (r *peptideRepository) FindById(id string) (*app.Peptide, error) {
	b, err := r.connection.HGet(table, id).Bytes()

	if err != nil {
		return nil, err
	}

	t := new(app.Peptide)
	err = json.Unmarshal(b, t)

	if err != nil {
		return nil, err
	}

	return t, nil
}

func (r *peptideRepository) FindAll() (peptides []*app.Peptide, err error) {
	ts := r.connection.HGetAll(table).Val()
	for key, value := range ts {
		t := new(app.Peptide)
		err = json.Unmarshal([]byte(value), t)

		if err != nil {
			return nil, err
		}

		t.PdbID = key
		peptides = append(peptides, t)
	}
	return peptides, nil
}
