package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/ArpitChinmay/interview/handlers"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

type Config struct {
	database struct {
		dbprovider string `yaml:"dbprovider"`
		username   string `yaml:"username"`
		pwd        string `yaml:"pwd"`
		protocol   string `yaml:"protocol"`
		address    string `yaml:"address"`
		dbport     string `yaml:"dbport"`
		dbname     string `yaml:"dbname"`
		parseTime  string `yaml:"parseTime"`
	}
}

type ServerObject struct {
	handler handlers.InterviewHandler
}

func SetupServer() (*ServerObject, *gin.Engine) {
	log.Println("Setting up server...")
	router := gin.New()
	gin.SetMode(gin.Default().TrustedPlatform)
	db := setupDBServer()
	registerRoutes(router)
	// create a object to use db data without exposing it to outside users.
	server := &ServerObject{handler: handlers.NewInterviewHandler(db)}
	return server, router
}

func setupDBServer() *sql.DB {
	log.Println("Setting up connection to db server...")
	dbprovider, connectionstring := getConnectionString()
	db, err := sql.Open(dbprovider, connectionstring)

	if err != nil {
		log.Println("could not connect to the database...")
		log.Fatal(err)
		return nil
	}

	log.Println("Connected to database...")
	return db
}

func getConnectionString() (string, string) {
	configFile, err := os.Open("config.yml")
	if err != nil {
		panic(err)
	}

	defer configFile.Close()

	decoder := yaml.NewDecoder(configFile)
	var databaseconfig Config
	if err := decoder.Decode(&databaseconfig); err != nil {
		panic(err)
	}

	dbprovider := databaseconfig.database.dbprovider
	connectionString := databaseconfig.database.username + ":" + databaseconfig.database.pwd + "@" + databaseconfig.database.protocol + "(" + databaseconfig.database.address + ")/" + databaseconfig.database.dbname + "?" + "parseTime=" + databaseconfig.database.parseTime

	return dbprovider, connectionString
}

func registerRoutes(router *gin.Engine) {
	router.GET("/db", GetSelectedAndRejectedCandidates)
	router.GET("/db/:level/", GetSepecificCandidateDetails)
	router.GET("/db/onboarded", GetOnboardedCandidateDetails)
	//Kundan Kumar
	router.GET("/db/interview-db/home/offer_rolled_out_accepted", GetCandidatesWithAcceptedOffers)
	router.GET("/db/interview-db/home/offer_rolled_out_awaited", GetCandidatesWithAwaitedOffers)
	router.GET("/db/interview-db/home/offer_rolled_out_accepted_count", GetAcceptedCandidatesCount)
	router.GET("/db/interview-db/home/offer_rolled_out_awaited_count", GetAwaitedCandidatesCount)
	router.POST("/home/admin", AddCandidate)
	router.PUT("/home/admin/:id", UpdateCandidate)
}
