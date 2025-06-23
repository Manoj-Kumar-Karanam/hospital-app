package tests

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "fmt"
    "hospitalapp/controllers"
    "hospitalapp/config"
    "io"
    "github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {
    db := SetupTestDB()
    config.DB = db

    // Register a user first
    router := SetupRouter()
    router.POST("/register", controllers.Register)
    router.POST("/login", controllers.Login)

    regBody := `{"username": "reception1", "password": "1234", "role": "receptionist"}`
    loginBody := `{"username": "reception1", "password": "1234"}`

    // Register user
    regReq := httptest.NewRequest("POST", "/register", strings.NewReader(regBody))
    regReq.Header.Set("Content-Type", "application/json")
    regW := httptest.NewRecorder()
    router.ServeHTTP(regW, regReq)

    // Login
    loginReq := httptest.NewRequest("POST", "/login", strings.NewReader(loginBody))
    loginReq.Header.Set("Content-Type", "application/json")
    loginW := httptest.NewRecorder()
    router.ServeHTTP(loginW, loginReq)
    fmt.Println("hello")
    assert.Equal(t, http.StatusOK, loginW.Code)
    body, _ := io.ReadAll(loginW.Body)
    fmt.Println("Login Response:", string(body)) // or parse JSON to extract token
}
