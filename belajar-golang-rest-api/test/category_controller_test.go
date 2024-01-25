package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"restapi/app"
	"restapi/controller"
	"restapi/helper"
	"restapi/middleware"
	"restapi/model/domain"
	"restapi/repository"
	"restapi/services"
	"strconv"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/categories_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func truncateCategory(db *sql.DB) {
	db.Exec("truncate category")
}

const (
	mainUrl = "htpp://localhost:8000/api/categories"
)

type categoryTest struct {
	method       string
	categoryName string
	apiKey       string
}

func newCategoryRequest(url string, isBodyEmpty bool, ct *categoryTest) *http.Request {
	var request *http.Request
	if isBodyEmpty {
		request = httptest.NewRequest(ct.method, url, nil)
	} else {
		requestBody := strings.NewReader(`{"name" : "` + ct.categoryName + `"}`)
		request = httptest.NewRequest(ct.method, url, requestBody)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", ct.apiKey)

	return request
}

func readTestResponse(response io.ReadCloser) map[string]interface{} {
	var result map[string]interface{}

	body, _ := io.ReadAll(response)
	json.Unmarshal(body, &result)

	return result
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	request := newCategoryRequest(mainUrl, false, &categoryTest{
		method:       http.MethodPost,
		categoryName: "Gadget",
		apiKey:       "RAHASIA",
	})

	recoder := httptest.NewRecorder()
	router.ServeHTTP(recoder, request)
	response := recoder.Result()

	assert.Equal(t, 200, response.StatusCode)
	responseBody := readTestResponse(response.Body)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFail(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	request := newCategoryRequest(mainUrl, false, &categoryTest{
		method:       http.MethodPost,
		categoryName: "",
		apiKey:       "RAHASIA",
	})

	recoder := httptest.NewRecorder()
	router.ServeHTTP(recoder, request)
	response := recoder.Result()

	assert.Equal(t, 400, response.StatusCode)
	responseBody := readTestResponse(response.Body)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func createCategoryFromRepository(db *sql.DB, categoryDomain *domain.Category) {
	tx, _ := db.Begin()
	categoryReponsitory := repository.NewCategoryRepository()
	*categoryDomain = categoryReponsitory.Save(context.Background(), tx, *categoryDomain)
	tx.Commit()
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	categoryDomain := domain.Category{Name: "Alat kebun"}
	createCategoryFromRepository(db, &categoryDomain)

	request := newCategoryRequest(mainUrl+"/"+strconv.Itoa(categoryDomain.Id), false, &categoryTest{
		method:       http.MethodPut,
		categoryName: categoryDomain.Name,
		apiKey:       "RAHASIA",
	})

	recoder := httptest.NewRecorder()
	router.ServeHTTP(recoder, request)
	response := recoder.Result()

	assert.Equal(t, 200, response.StatusCode)
	responseBody := readTestResponse(response.Body)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, categoryDomain.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, categoryDomain.Name, responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	categoryDomain := domain.Category{Name: ""}
	createCategoryFromRepository(db, &categoryDomain)

	request := newCategoryRequest(mainUrl+"/"+strconv.Itoa(categoryDomain.Id), false, &categoryTest{
		method:       http.MethodPut,
		categoryName: categoryDomain.Name,
		apiKey:       "RAHASIA",
	})

	recoder := httptest.NewRecorder()
	router.ServeHTTP(recoder, request)
	response := recoder.Result()

	assert.Equal(t, 400, response.StatusCode)
	responseBody := readTestResponse(response.Body)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, "BAD REQUEST", responseBody["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	categoryDomain := domain.Category{Name: "Gadget"}
	createCategoryFromRepository(db, &categoryDomain)

	request := newCategoryRequest(mainUrl+"/"+strconv.Itoa(categoryDomain.Id), true, &categoryTest{
		method:       http.MethodDelete,
		categoryName: "",
		apiKey:       "RAHASIA",
	})

	recoder := httptest.NewRecorder()
	router.ServeHTTP(recoder, request)
	response := recoder.Result()

	assert.Equal(t, 200, response.StatusCode)
	responseBody := readTestResponse(response.Body)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteCategoryFail(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	categoryDomain := domain.Category{Name: "Gadget"}
	createCategoryFromRepository(db, &categoryDomain)

	request := newCategoryRequest(mainUrl+"/4000", true, &categoryTest{
		method:       http.MethodDelete,
		categoryName: "",
		apiKey:       "RAHASIA",
	})

	recoder := httptest.NewRecorder()
	router.ServeHTTP(recoder, request)
	response := recoder.Result()

	assert.Equal(t, 404, response.StatusCode)
	responseBody := readTestResponse(response.Body)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, "NOT FOUND", responseBody["status"])
}

func TestListCategoriesSuccess(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	category1 := domain.Category{
		Name: "Alat kebersihan",
	}
	createCategoryFromRepository(db, &category1)
	category2 := domain.Category{
		Name: "Alat masak",
	}
	createCategoryFromRepository(db, &category2)

	request := newCategoryRequest(mainUrl, true, &categoryTest{
		method:       http.MethodGet,
		categoryName: "",
		apiKey:       "RAHASIA",
	})

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
	responseBody := readTestResponse(response.Body)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	categories := responseBody["data"].([]interface{})

	categoryResponse1 := categories[0].(map[string]interface{})
	categoryResponse2 := categories[1].(map[string]interface{})

	assert.Equal(t, category1.Id, int(categoryResponse1["id"].(float64)))
	assert.Equal(t, category1.Name, categoryResponse1["name"].(string))

	assert.Equal(t, category2.Id, int(categoryResponse2["id"].(float64)))
	assert.Equal(t, category2.Name, categoryResponse2["name"].(string))
}

func TestGetCategoryById(t *testing.T) {
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	categoryDomain := domain.Category{
		Name: "Makanan ringan",
	}
	createCategoryFromRepository(db, &categoryDomain)

	request := newCategoryRequest(mainUrl+"/"+strconv.Itoa(categoryDomain.Id), true, &categoryTest{
		method:       http.MethodGet,
		categoryName: "",
		apiKey:       "RAHASIA",
	})

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	response := recorder.Result()

	assert.Equal(t, 200, response.StatusCode)
	responseBody := readTestResponse(response.Body)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	assert.Equal(t, categoryDomain.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, categoryDomain.Name, responseBody["data"].(map[string]interface{})["name"])
}
