package setup

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// )

// func TestGetSelectedAndRejectedCandidates_WrongLevelParams(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()
// 	router.GET("/db", GetSelectedAndRejectedCandidates)
// 	request, err := http.NewRequest(http.MethodGet, "/db?level=true", nil)
// 	if err != nil {
// 		t.Fatalf("Couldn't create request: %v", err)
// 	}

// 	inspectRecorder := httptest.NewRecorder()

// 	router.ServeHTTP(inspectRecorder, request)
// 	// fmt.Println(inspectRecorder.Body)

// 	if inspectRecorder.Code == http.StatusNoContent {
// 		t.Logf("Expected to get status %d and recieved status %d\n", http.StatusNoContent, inspectRecorder.Code)
// 	} else {
// 		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusNoContent, inspectRecorder.Code)
// 	}
// }

// func TestGetSelectedAndRejectedCandidates_CorrectParamLevelOne(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()
// 	router.GET("/db", GetSelectedAndRejectedCandidates)
// 	request, err := http.NewRequest(http.MethodGet, "/db?level=1", nil)
// 	if err != nil {
// 		t.Fatalf("Couldn't create request: %v", err)
// 	}
// 	inspectRecorder := httptest.NewRecorder()

// 	router.ServeHTTP(inspectRecorder, request)
// 	//fmt.Println(inspectRecorder.Body)

// 	if inspectRecorder.Code == http.StatusOK {
// 		t.Logf("Expected to get status %d and recieved status %d\n", http.StatusOK, inspectRecorder.Code)
// 	} else {
// 		t.Fatalf("Expected to get status %d but instead go %d\n", http.StatusOK, inspectRecorder.Code)
// 	}
// }

// func TestGetSelectedAndRejectedCandidates_CorrectParamLevelTwo(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()
// 	router.GET("/db", GetSelectedAndRejectedCandidates)
// 	request, err := http.NewRequest(http.MethodGet, "/db?level=2", nil)
// 	if err != nil {
// 		t.Fatalf("Couldn't create request: %v", err)
// 	}
// 	inspectRecorder := httptest.NewRecorder()

// 	router.ServeHTTP(inspectRecorder, request)
// 	//fmt.Println(inspectRecorder.Body)

// 	if inspectRecorder.Code == http.StatusOK {
// 		t.Logf("Expected to get status %d and recieved status %d\n", http.StatusOK, inspectRecorder.Code)
// 	} else {
// 		t.Fatalf("Expected to get status %d but instead go %d\n", http.StatusOK, inspectRecorder.Code)
// 	}
// }

// func TestGetSelectedAndRejectedCandidates_IncorrectParamLevel(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()
// 	router.GET("/db", GetSelectedAndRejectedCandidates)
// 	request, err := http.NewRequest(http.MethodGet, "/db?level=4", nil)
// 	if err != nil {
// 		t.Fatalf("Couldn't create request: %v", err)
// 	}
// 	inspectRecorder := httptest.NewRecorder()

// 	router.ServeHTTP(inspectRecorder, request)
// 	//fmt.Println(inspectRecorder.Body)

// 	if inspectRecorder.Code == http.StatusBadRequest {
// 		t.Logf("Expected to get status %d and recieved status %d\n", http.StatusBadRequest, inspectRecorder.Code)
// 	} else {
// 		t.Fatalf("Expected to get status %d but instead go %d\n", http.StatusBadRequest, inspectRecorder.Code)
// 	}
// }

// func TestGetSepecificCandidateDetails_WrongParamLevel(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()
// 	router.GET("/db/:level", GetSepecificCandidateDetails)

// 	request, err := http.NewRequest(http.MethodGet, "/db/true?selected=true&count=false", nil)
// 	if err != nil {
// 		t.Fatalf("Couldn't create request: %v", err)
// 	}

// 	inspectRecorder := httptest.NewRecorder()

// 	router.ServeHTTP(inspectRecorder, request)

// 	if inspectRecorder.Code == http.StatusBadRequest {
// 		t.Logf("Expected to get status %d and recieved status %d\n", http.StatusBadRequest, inspectRecorder.Code)
// 	} else {
// 		t.Fatalf("Expected to get status %d but instead go %d\n", http.StatusBadRequest, inspectRecorder.Code)
// 	}
// }

// func TestGetSepecificCandidateDetails_LevelOneSelectTrueCountTrue(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()
// 	router.GET("/db/:level", GetSepecificCandidateDetails)

// 	request, err := http.NewRequest(http.MethodGet, "/db/1?selected=true&count=true", nil)
// 	if err != nil {
// 		t.Fatalf("Couldn't create request: %v", err)
// 	}

// 	inspectRecorder := httptest.NewRecorder()

// 	router.ServeHTTP(inspectRecorder, request)

// 	if inspectRecorder.Code == http.StatusOK {
// 		t.Logf("Expected to get status %d and recieved status %d\n", http.StatusOK, inspectRecorder.Code)
// 	} else {
// 		t.Fatalf("Expected to get status %d but instead go %d\n", http.StatusOK, inspectRecorder.Code)
// 	}
// }

// func TestGetSepecificCandidateDetails_LevelOneSelectTrueCountFalse(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()
// 	router.GET("/db/:level", GetSepecificCandidateDetails)

// 	request, err := http.NewRequest(http.MethodGet, "/db/1?selected=true&count=false", nil)
// 	if err != nil {
// 		t.Fatalf("Couldn't create request: %v", err)
// 	}

// 	inspectRecorder := httptest.NewRecorder()

// 	router.ServeHTTP(inspectRecorder, request)

// 	if inspectRecorder.Code == http.StatusOK {
// 		t.Logf("Expected to get status %d and recieved status %d\n", http.StatusOK, inspectRecorder.Code)
// 	} else {
// 		t.Fatalf("Expected to get status %d but instead go %d\n", http.StatusOK, inspectRecorder.Code)
// 	}
// }

// func TestGetSepecificCandidateDetails_LevelOneSelectFalseCountTrue(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()
// 	router.GET("/db/:level", GetSepecificCandidateDetails)

// 	request, err := http.NewRequest(http.MethodGet, "/db/1?selected=false&count=true", nil)
// 	if err != nil {
// 		t.Fatalf("Couldn't create request: %v", err)
// 	}

// 	inspectRecorder := httptest.NewRecorder()

// 	router.ServeHTTP(inspectRecorder, request)

// 	if inspectRecorder.Code == http.StatusOK {
// 		t.Logf("Expected to get status %d and recieved status %d\n", http.StatusOK, inspectRecorder.Code)
// 	} else {
// 		t.Fatalf("Expected to get status %d but instead go %d\n", http.StatusOK, inspectRecorder.Code)
// 	}
// }

// func TestGetSepecificCandidateDetails_LevelOneSelectFalseCountFalse(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()
// 	router.GET("/db/:level", GetSepecificCandidateDetails)

// 	request, err := http.NewRequest(http.MethodGet, "/db/1?selected=false&count=false", nil)
// 	if err != nil {
// 		t.Fatalf("Couldn't create request: %v", err)
// 	}

// 	inspectRecorder := httptest.NewRecorder()

// 	router.ServeHTTP(inspectRecorder, request)

// 	if inspectRecorder.Code == http.StatusOK {
// 		t.Logf("Expected to get status %d and recieved status %d\n", http.StatusOK, inspectRecorder.Code)
// 	} else {
// 		t.Fatalf("Expected to get status %d but instead go %d\n", http.StatusOK, inspectRecorder.Code)
// 	}
// }

// func TestGetSepecificCandidateDetails_LevelTwoSelectTrueCountTrue(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()
// 	router.GET("/db/:level", GetSepecificCandidateDetails)

// 	request, err := http.NewRequest(http.MethodGet, "/db/2?selected=true&count=true", nil)
// 	if err != nil {
// 		t.Fatalf("Couldn't create request: %v", err)
// 	}

// 	inspectRecorder := httptest.NewRecorder()

// 	router.ServeHTTP(inspectRecorder, request)

// 	if inspectRecorder.Code == http.StatusOK {
// 		t.Logf("Expected to get status %d and recieved status %d\n", http.StatusOK, inspectRecorder.Code)
// 	} else {
// 		t.Fatalf("Expected to get status %d but instead go %d\n", http.StatusOK, inspectRecorder.Code)
// 	}
// }

// func TestGetSepecificCandidateDetails_LevelTwoSelectTrueCountFalse(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()
// 	router.GET("/db/:level", GetSepecificCandidateDetails)

// 	request, err := http.NewRequest(http.MethodGet, "/db/2?selected=true&count=false", nil)
// 	if err != nil {
// 		t.Fatalf("Couldn't create request: %v", err)
// 	}

// 	inspectRecorder := httptest.NewRecorder()

// 	router.ServeHTTP(inspectRecorder, request)

// 	if inspectRecorder.Code == http.StatusOK {
// 		t.Logf("Expected to get status %d and recieved status %d\n", http.StatusOK, inspectRecorder.Code)
// 	} else {
// 		t.Fatalf("Expected to get status %d but instead go %d\n", http.StatusOK, inspectRecorder.Code)
// 	}
// }

// func TestGetSepecificCandidateDetails_LevelTwoSelectFalseCountTrue(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()
// 	router.GET("/db/:level", GetSepecificCandidateDetails)

// 	request, err := http.NewRequest(http.MethodGet, "/db/2?selected=false&count=true", nil)
// 	if err != nil {
// 		t.Fatalf("Couldn't create request: %v", err)
// 	}

// 	inspectRecorder := httptest.NewRecorder()

// 	router.ServeHTTP(inspectRecorder, request)

// 	if inspectRecorder.Code == http.StatusOK {
// 		t.Logf("Expected to get status %d and recieved status %d\n", http.StatusOK, inspectRecorder.Code)
// 	} else {
// 		t.Fatalf("Expected to get status %d but instead go %d\n", http.StatusOK, inspectRecorder.Code)
// 	}
// }

// func TestGetSepecificCandidateDetails_LevelTwoSelectFalseCountFalse(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()
// 	router.GET("/db/:level", GetSepecificCandidateDetails)

// 	request, err := http.NewRequest(http.MethodGet, "/db/2?selected=false&count=False", nil)
// 	if err != nil {
// 		t.Fatalf("Couldn't create request: %v", err)
// 	}

// 	inspectRecorder := httptest.NewRecorder()

// 	router.ServeHTTP(inspectRecorder, request)

// 	if inspectRecorder.Code == http.StatusOK {
// 		t.Logf("Expected to get status %d and recieved status %d\n", http.StatusOK, inspectRecorder.Code)
// 	} else {
// 		t.Fatalf("Expected to get status %d but instead go %d\n", http.StatusOK, inspectRecorder.Code)
// 	}
// }

// func TestGetSepecificCandidateDetails_LevelThreeSelectTrueCountTrue(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()
// 	router.GET("/db/:level", GetSepecificCandidateDetails)

// 	request, err := http.NewRequest(http.MethodGet, "/db/3?selected=true&count=true", nil)
// 	if err != nil {
// 		t.Fatalf("Couldn't create request: %v", err)
// 	}

// 	inspectRecorder := httptest.NewRecorder()

// 	router.ServeHTTP(inspectRecorder, request)

// 	if inspectRecorder.Code == http.StatusOK {
// 		t.Logf("Expected to get status %d and recieved status %d\n", http.StatusOK, inspectRecorder.Code)
// 	} else {
// 		t.Fatalf("Expected to get status %d but instead go %d\n", http.StatusOK, inspectRecorder.Code)
// 	}
// }

// func TestGetSepecificCandidateDetails_LevelThreeSelectTrueCountFalse(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()
// 	router.GET("/db/:level", GetSepecificCandidateDetails)

// 	request, err := http.NewRequest(http.MethodGet, "/db/3?selected=true&count=false", nil)
// 	if err != nil {
// 		t.Fatalf("Couldn't create request: %v", err)
// 	}

// 	inspectRecorder := httptest.NewRecorder()

// 	router.ServeHTTP(inspectRecorder, request)

// 	if inspectRecorder.Code == http.StatusOK {
// 		t.Logf("Expected to get status %d and recieved status %d\n", http.StatusOK, inspectRecorder.Code)
// 	} else {
// 		t.Fatalf("Expected to get status %d but instead go %d\n", http.StatusOK, inspectRecorder.Code)
// 	}
// }

// func TestGetSepecificCandidateDetails_LevelThreeSelectFalseCountTrue(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()
// 	router.GET("/db/:level", GetSepecificCandidateDetails)

// 	request, err := http.NewRequest(http.MethodGet, "/db/3?selected=false&count=true", nil)
// 	if err != nil {
// 		t.Fatalf("Couldn't create request: %v", err)
// 	}

// 	inspectRecorder := httptest.NewRecorder()

// 	router.ServeHTTP(inspectRecorder, request)

// 	if inspectRecorder.Code == http.StatusOK {
// 		t.Logf("Expected to get status %d and recieved status %d\n", http.StatusOK, inspectRecorder.Code)
// 	} else {
// 		t.Fatalf("Expected to get status %d but instead go %d\n", http.StatusOK, inspectRecorder.Code)
// 	}
// }

// func TestGetSepecificCandidateDetails_LevelThreeSelectFalseCountfalse(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()
// 	router.GET("/db/:level", GetSepecificCandidateDetails)

// 	request, err := http.NewRequest(http.MethodGet, "/db/3?selected=false&count=false", nil)
// 	if err != nil {
// 		t.Fatalf("Couldn't create request: %v", err)
// 	}

// 	inspectRecorder := httptest.NewRecorder()

// 	router.ServeHTTP(inspectRecorder, request)

// 	if inspectRecorder.Code == http.StatusOK {
// 		t.Logf("Expected to get status %d and recieved status %d\n", http.StatusOK, inspectRecorder.Code)
// 	} else {
// 		t.Fatalf("Expected to get status %d but instead go %d\n", http.StatusOK, inspectRecorder.Code)
// 	}
// }

// func TestGetSepecificCandidateDetails_WrongRequestParams(t *testing.T) {
// 	gin.SetMode(gin.TestMode)
// 	router := gin.Default()
// 	router.GET("/db/:level", GetSepecificCandidateDetails)

// 	request, err := http.NewRequest(http.MethodGet, "/db/4?selected=false&count=false", nil)
// 	if err != nil {
// 		t.Fatalf("Couldn't create request: %v", err)
// 	}

// 	inspectRecorder := httptest.NewRecorder()

// 	router.ServeHTTP(inspectRecorder, request)

// 	if inspectRecorder.Code == http.StatusBadRequest {
// 		t.Logf("Expected to get status %d and recieved status %d\n", http.StatusBadRequest, inspectRecorder.Code)
// 	} else {
// 		t.Fatalf("Expected to get status %d but instead go %d\n", http.StatusBadRequest, inspectRecorder.Code)
// 	}
// }
