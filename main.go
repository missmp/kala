package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	redislib "github.com/garyburd/redigo/redis"
	"github.com/missmp/kala/api"
	"github.com/missmp/kala/job"
	"github.com/missmp/kala/job/storage/redis"

	log "github.com/Sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
}

// The current version of kala
var Version = "0.1"

func main() {
	var db job.JobDB
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.Println("Internal Use Kala Modified Main")
	jdb := os.Getenv("JDB_ADDRESS")
	if jdb == "" {
		log.Println("Please provide JDB_ADDRESS env")
		os.Exit(1)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Println("Please provide PORT env")
		os.Exit(1)
	}

	parsedPort := fmt.Sprintf(":%s", port)
	log.Println("Going to connect db", jdb)
	db = redis.New(jdb, redislib.DialOption{}, false)
	cache := job.NewLockFreeJobCache(db)
	log.Println("Preparing cache")
	cache.Start(time.Duration(5 * time.Second))
	log.Println("Starting server on port %s", parsedPort)
	log.Fatal(api.StartServer(parsedPort, cache, db, ""))
}
