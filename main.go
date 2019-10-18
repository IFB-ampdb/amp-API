package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/ifbampdb/amp-core/app"
	"github.com/ifbampdb/amp-core/database/psql"
	redisdb "github.com/ifbampdb/amp-core/database/redis"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {

	dbType := flag.String("database", "psql", "database type [redis, psql]")
	dbUser := flag.String("dbuser", "ampdb", "database user for authentication")
	dbPass := flag.String("dbpass", "", "database password for authentication")
	dbAddr := flag.String("dbaddr", "cockroach", "database address")
	dbPort := flag.String("dbport", "26257", "database address port")
	flag.Parse()

	var peptideRepo app.PeptideRepository

	switch *dbType {
	case "psql":
		peptideRepo = psql.NewPostgresPeptideRepository(psql.PostgresConnection(fmt.Sprintf("postgresql://%s:%s@%s:%s/ampdb", *dbUser, *dbPass, *dbAddr, *dbPort)))
	case "redis":
		peptideRepo = redisdb.NewRedisPeptideRepository(redisdb.RedisConnection(fmt.Sprintf("redis://%s:%s", *dbAddr, *dbPort), *dbPass))
	default:
		log.Fatal("Unknown database")
	}

	peptideService := app.NewPeptideService(peptideRepo)
	peptideHandler := app.NewPeptideHandler(peptideService)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/peptides", peptideHandler.Get).Methods("GET")
	router.HandleFunc("/peptides/{id}", peptideHandler.GetById).Methods("GET")
	router.HandleFunc("/peptides", peptideHandler.Create).Methods("POST")

	http.Handle("/", accessControl(router))

	errs := make(chan error, 2)
	go func() {
		log.Println("Listening on port :8080")
		errs <- http.ListenAndServe(":8080", nil)
	}()
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	log.Printf("terminated %s", <-errs)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}
