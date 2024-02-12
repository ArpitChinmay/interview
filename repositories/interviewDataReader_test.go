package datareader

// import (
// 	"database/sql"
// 	"log"

// 	"github.com/ArpitChinmay/interview/models"
// 	"github.com/ArpitChinmay/interview/utils"
// 	"github.com/stretchr/testify/suite"
// )

// type dataReaderSuite struct {
// 	suite.Suite
// 	datareader      DataReader
// 	cleanupExecutor utils.TruncateTableExecutor
// }

// func (suite *dataReaderSuite) SetupSuite() {
// 	db, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/int_db_data?parseTime=true")

// 	if err != nil {
// 		log.Println("could not connect to the database...")
// 		log.Fatal(err)
// 	}
// 	dr := InitializeDataReader(db)
// 	suite.datareader = dr
// 	suite.cleanupExecutor = utils.InitTruncateTableExecutor(db)
// }

// func (suite *dataReaderSuite) TearDownTest() {
// 	defer suite.cleanupExecutor.TruncateTable([]string{"int_db_data.interview_status", "int_db_data.resume"})
// }

// func (suite *dataReaderSuite) TestReadInterviewDataForLevelOneSelecteOrRejected_Positive() {
// 	interviewData, err := suite.datareader.ReadInterviewDataForLevelOneSelecteOrRejected()
// 	suite.NoError(err, "No error when trying to retreive level one selected and rejected candidates.")
// 	suite.Equal(len(*interviewData), 0, "Length of tweets should be 0, since it is empty slice")
// 	suite.Equal(*interviewData, []models.Interview(nil), "interviewData is an empty slice")
// }

// func (suite *dataReaderSuite) TestReadInterviewDataForLevelTwoSelecteOrRejected_Positive() {
// 	interviewData, err := suite.datareader.ReadInterviewDataForLevelTwoSelecteOrRejected()
// 	suite.NoError(err, "No error when trying to retreive level one selected and rejected candidates.")
// 	suite.Equal(len(*interviewData), 0, "Length of tweets should be 0, since it is empty slice")
// 	suite.Equal(*interviewData, []models.Interview(nil), "interviewData is an empty slice")
// }

// func (suite *dataReaderSuite) TestReadInterviewDataForLevelOneSelecteCandidates_Positive() {
// 	interviewData, err := suite.datareader.ReadInterviewDataForLevelOneSelecteCandidates()
// 	suite.NoError(err, "No error when trying to retreive level one selected and rejected candidates.")
// 	suite.Equal(len(*interviewData), 0, "Length of tweets should be 0, since it is empty slice")
// 	suite.Equal(*interviewData, []models.Interview(nil), "interviewData is an empty slice")
// }

// func (suite *dataReaderSuite) TestReadInterviewDataForLevelOneRejectedCandidates_Positive() {
// 	interviewData, err := suite.datareader.ReadInterviewDataForLevelOneRejectedCandidates()
// 	suite.NoError(err, "No error when trying to retreive level one selected and rejected candidates.")
// 	suite.Equal(len(*interviewData), 0, "Length of tweets should be 0, since it is empty slice")
// 	suite.Equal(*interviewData, []models.Interview(nil), "interviewData is an empty slice")
// }

// func (suite *dataReaderSuite) TestReadInterviewDataForLevelTwoSelectedCandidates_Positive() {
// 	interviewData, err := suite.datareader.ReadInterviewDataForLevelTwoSelectedCandidates()
// 	suite.NoError(err, "No error when trying to retreive level one selected and rejected candidates.")
// 	suite.Equal(len(*interviewData), 0, "Length of tweets should be 0, since it is empty slice")
// 	suite.Equal(*interviewData, []models.Interview(nil), "interviewData is an empty slice")
// }

// func (suite *dataReaderSuite) TestReadInterviewDataForLevelTwoRejectedCandidates_Positive() {
// 	interviewData, err := suite.datareader.ReadInterviewDataForLevelTwoRejectedCandidates()
// 	suite.NoError(err, "No error when trying to retreive level one selected and rejected candidates.")
// 	suite.Equal(len(*interviewData), 0, "Length of tweets should be 0, since it is empty slice")
// 	suite.Equal(*interviewData, []models.Interview(nil), "interviewData is an empty slice")
// }

// func (suite *dataReaderSuite) TestReadInterviewDataForDMLevelSelectedCandidates_Positive() {
// 	interviewData, err := suite.datareader.ReadInterviewDataForDMLevelSelectedCandidates()
// 	suite.NoError(err, "No error when trying to retreive level one selected and rejected candidates.")
// 	suite.Equal(len(*interviewData), 0, "Length of tweets should be 0, since it is empty slice")
// 	suite.Equal(*interviewData, []models.Interview(nil), "interviewData is an empty slice")
// }

// func (suite *dataReaderSuite) TestReadInterviewDataForDMLevelRejectedCandidates_Positive() {
// 	interviewData, err := suite.datareader.ReadInterviewDataForDMLevelRejectedCandidates()
// 	suite.NoError(err, "No error when trying to retreive level one selected and rejected candidates.")
// 	suite.Equal(len(*interviewData), 0, "Length of tweets should be 0, since it is empty slice")
// 	suite.Equal(*interviewData, []models.Interview(nil), "interviewData is an empty slice")
// }
