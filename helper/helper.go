package helper

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/go-redis/redis"
)

var redisclient *redis.Client
var SYSID string
var databaseEV DatabaseX

// DatabaseX is a struct
type DatabaseX struct {
	Location   string // location of the database localhost, something.com, etc
	Database   string // database name
	Collection string // collection name
}

// Resultado is a struct
type Resultado struct {
	ErrorCode        string // error code
	ErrorDescription string // description
	IsSuccessful     string // Y or N
	ReturnedValue    string // Any string
}

// GetRedisPointer returns
func GetRedisPointer(bucket int) *redis.Client {

	bucket = 0

	if redisclient == nil {
		redisclient = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",     // no password set
			DB:       bucket, // use default DB
		})
	}

	return redisclient
}

// RestEnvVariables = restaurante environment variables
//
type RestEnvVariables struct {
	APIMongoDBLocation    string // location of the database localhost, something.com, etc
	APIMongoDBDatabase    string // database name
	APIAPIServerPort      string // collection name
	APIAPIServerIPAddress string // apiserver name
	WEBDebug              string // debug
	CollectionOrders      string // Collection Names
	CollectionSecurity    string // Collection Names
	CollectionDishes      string // Collection Names
	CollectionEvents      string // Collection Names
	MSAPIdishesPort       string // Microservices Port Dishes
	MSAPIordersPort       string // Microservices Port Orders
	SYSID                 string // Id of this specific microservice
}

// Readfileintostruct is
func Readfileintostruct() RestEnvVariables {
	dat, err := ioutil.ReadFile("fjapiorders.ini")
	check(err)
	fmt.Print(string(dat))

	var restenv RestEnvVariables

	json.Unmarshal(dat, &restenv)

	return restenv
}

// GetSYSID is just returning the System ID directly from file
// It is happening to enable multiple usage of Redis Keys ("SYSID" + "APIURL" for instance)
func GetSYSID() string {

	if SYSID == "" {

		dat, err := ioutil.ReadFile("fjapiorders.ini")
		check(err)
		fmt.Print(string(dat))

		var restenv RestEnvVariables

		json.Unmarshal(dat, &restenv)

		SYSID = restenv.SYSID

		return restenv.SYSID
	}

	return SYSID
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type PlayerRegistrationFile struct {
	FFA  string
	Name string
	DOB  string
}

// Getvaluefromcache returns the value of a key from cache
func Getvaluefromcache(key string) string {

	// bucket is ZERO for now
	// I am allowing it to be setup now
	rp := GetRedisPointer(0)

	sysid := GetSYSID()

	valuetoreturn, _ := rp.Get(sysid + key).Result()

	return valuetoreturn
}

// GetDBParmFromCache returns the value of a key from cache
func GetDBParmFromCache(collection string) *DatabaseX {

	database := new(DatabaseX)

	database.Collection = Getvaluefromcache(collection)
	database.Database = Getvaluefromcache("API.MongoDB.Database")
	database.Location = Getvaluefromcache("API.MongoDB.Location")

	return database
}

// Capitalfootball is
func Capitalfootball(redisclient *redis.Client) []PlayerRegistrationFile {

	file, err := os.Open("capitalfootball.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var playerlist []PlayerRegistrationFile

	scanner := bufio.NewScanner(file)

	playerlist = make([]PlayerRegistrationFile, 52)

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(scanner.Text())

		tmp := strings.Split(line, ",")

		i++
		playerlist[i] = PlayerRegistrationFile{}
		playerlist[i].FFA = strings.Trim(tmp[0], " ")
		playerlist[i].Name = strings.Trim(tmp[1], " ")
		playerlist[i].DOB = strings.Trim(tmp[2], " ")

		fmt.Println(playerlist[i].FFA)

		err = redisclient.Set(playerlist[i].FFA, playerlist[i].Name, 0).Err()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return playerlist
}
