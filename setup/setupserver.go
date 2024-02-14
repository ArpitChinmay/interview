package setup

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
)

type databaseconfig struct {
	Dbprovider string `yaml:"dbprovider"`
	Username   string `yaml:"username"`
	Pwd        string `yaml:"pwd"`
	Protocol   string `yaml:"protocol"`
	Address    string `yaml:"address"`
	Dbport     string `yaml:"dbport"`
	Dbname     string `yaml:"dbname"`
	ParseTime  string `yaml:"parseTime"`
}

func SetupServer() {
	log.Println("Setting up server...")
	router := gin.New()
	gin.SetMode(gin.Default().TrustedPlatform)
	db := setupDBServer()
	repos := SetupRepositories(db)
	midware := SetupMiddleware(repos)
	handler := SetupHandler(midware)

	registerRoutes(router, handler)
	router.Run(":5000")
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
	file, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatal("Error reading config file:", err)
		panic(err)
	}

	config := make(map[string]databaseconfig)
	if err := yaml.Unmarshal(file, &config); err != nil {
		log.Fatal("Error unmarshalling YAML file:", err)
	}

	connectionString := config["database"].Username + ":" + config["database"].Pwd + "@" + config["database"].Protocol + "(" + config["database"].Address + ":" + config["database"].Dbport + ")/" + config["database"].Dbname + "?" + "parseTime=" + config["database"].ParseTime
	log.Println("connetion string:", connectionString)
	return config["database"].Dbprovider, connectionString
}

func registerRoutes(router *gin.Engine, handler *HandlerGenerator) {
	router.GET("/db", handler.InterviewHandler.GetSelectedAndRejectedCandidates)
	router.GET("/db/:level/", handler.InterviewHandler.GetSepecificCandidateDetails)
	router.GET("/db/onboarded", handler.InterviewHandler.GetOnboardedCandidateDetails)
	//Kundan Kumar
	router.GET("/db/interview-db/home/offer_rolled_out_accepted", handler.InterviewHandler.GetCandidatesWithAcceptedOffers)
	router.GET("/db/interview-db/home/offer_rolled_out_awaited", handler.InterviewHandler.GetCandidatesWithAwaitedOffers)
	router.GET("/db/interview-db/home/offer_rolled_out_accepted_count", handler.InterviewHandler.GetAcceptedCandidatesCount)
	router.GET("/db/interview-db/home/offer_rolled_out_awaited_count", handler.InterviewHandler.GetAwaitedCandidatesCount)
	router.POST("/home/admin", handler.InterviewHandler.AddCandidate)
	router.PUT("/home/admin/:id", handler.InterviewHandler.UpdateCandidate)
}
