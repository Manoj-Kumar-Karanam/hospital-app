package tests

import (
    "bytes"
    "encoding/json"
    "hospitalapp/config"
    "hospitalapp/controllers"
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "fmt"
   
    "github.com/stretchr/testify/assert"
)

// helper to register + login and return JWT token
func getTestToken(t *testing.T) string {
    db := SetupTestDB()
    config.DB = db

    router := SetupRouter()
    router.POST("/register", controllers.Register)
    router.POST("/login", controllers.Login)

    registerBody := `{"username": "reception1", "password": "1234", "role": "receptionist"}`
    loginBody := `{"username": "reception1", "password": "1234"}`

    // Register
    regReq := httptest.NewRequest("POST", "/register", strings.NewReader(registerBody))
    regReq.Header.Set("Content-Type", "application/json")
    regW := httptest.NewRecorder()
    router.ServeHTTP(regW, regReq)

    // Login
    loginReq := httptest.NewRequest("POST", "/login", strings.NewReader(loginBody))
    loginReq.Header.Set("Content-Type", "application/json")
    loginW := httptest.NewRecorder()
    router.ServeHTTP(loginW, loginReq)

    assert.Equal(t, http.StatusOK, loginW.Code)

    var loginResp map[string]interface{}
    json.Unmarshal(loginW.Body.Bytes(), &loginResp)
    token := loginResp["token"].(string)
    return token
}

func TestCreatePatient(t *testing.T) {
    db := SetupTestDB()
    config.DB = db

    token := getTestToken(t)

    router := SetupRouter()
    router.POST("/api/patients", controllers.CreatePatient) // add auth middleware if needed

    // create request body
    patientBody := map[string]interface{}{
        "name":        "John Doe",
        "age":         30,
        "disease":     "Flu",
        "address":     "123 Lane",
        "assigned_to": "Dr. Smith",
    }
    bodyBytes, _ := json.Marshal(patientBody)

    req := httptest.NewRequest("POST", "/api/patients", bytes.NewReader(bodyBytes))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+token)

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)
    
}

func TestGetPatients(t *testing.T) {
    db := SetupTestDB()
    config.DB = db

    token := getTestToken(t) // creates a receptionist and gets JWT

    router := SetupRouter()
    router.POST("/api/patients", controllers.CreatePatient) // patient creation
    router.GET("/api/patients", controllers.GetPatients)     // list all patients

    // First, create a patient
    patientBody := map[string]interface{}{
        "Name":        "John Doe",
        "Age":         30,
        "Disease":     "Flu",
        "Address":     "123 Lane",
        "Assigned_to": "Dr. Smith",
    }
    bodyBytes, _ := json.Marshal(patientBody)

    req := httptest.NewRequest("POST", "/api/patients", bytes.NewReader(bodyBytes))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+token)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    assert.Equal(t, http.StatusCreated, w.Code)

    // Now, get the list of patients
    getReq := httptest.NewRequest("GET", "/api/patients", nil)
    getReq.Header.Set("Authorization", "Bearer "+token)
    getW := httptest.NewRecorder()
    router.ServeHTTP(getW, getReq)
    
    assert.Equal(t, http.StatusOK, getW.Code)
    
    var response []map[string]interface{}
    err := json.Unmarshal(getW.Body.Bytes(), &response)
    assert.NoError(t, err)
    
    fmt.Println("Returned Patient:", response)
    assert.GreaterOrEqual(t, len(response), 1)
    assert.Equal(t, "John Doe", response[0]["Name"])
}

func TestUpdatePatient(t *testing.T) {
    db := SetupTestDB()
    config.DB = db
    token := getTestToken(t)
    router := SetupRouter()

    router.POST("/api/patients", controllers.CreatePatient) // patient creation
    router.PUT("/api/patients/:id", controllers.UpdatePatient)
    // Step 1: Create a patient
    patientBody := map[string]interface{}{
        "Name":        "Jane Doe",
        "Age":         28,
        "Disease":     "Cold",
        "Address":     "456 Street",
        "Assigned_to": "Dr. House",
    }
    bodyBytes, _ := json.Marshal(patientBody)

    req := httptest.NewRequest("POST", "/api/patients", bytes.NewReader(bodyBytes))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+token)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)

    var created map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &created)
    patientID := created["id"]

     fmt.Println("created :", created)
    fmt.Println(patientID)
    // Step 2: Update patient
    updatedBody := map[string]interface{}{
        "Name":        "Jane Updated",
        "Age":         28,
        "Disease":     "Recovered",
        "Address":     "456 Street",
        "Assigned_to": "Dr. House",
    }
    updateBytes, _ := json.Marshal(updatedBody)

    updateReq := httptest.NewRequest("PUT", fmt.Sprintf("/api/patients/%v", patientID), bytes.NewReader(updateBytes))
    updateReq.Header.Set("Content-Type", "application/json")
    updateReq.Header.Set("Authorization", "Bearer "+token)
    updateW := httptest.NewRecorder()
    router.ServeHTTP(updateW, updateReq)

    assert.Equal(t, http.StatusOK, updateW.Code)

    var updated map[string]interface{}
    json.Unmarshal(updateW.Body.Bytes(), &updated)
    fmt.Println("Updated patient :", updated)
    assert.Equal(t, "Jane Updated", updated["Name"])

    assert.Equal(t, "Recovered", updated["Disease"])
}

func TestDeletePatient(t *testing.T) {
    db := SetupTestDB()
    config.DB = db
    token := getTestToken(t)
    router := SetupRouter()
    router.POST("/api/patients", controllers.CreatePatient)
    router.DELETE("/api/patients/:id", controllers.DeletePatient)
    router.GET("/api/patients", controllers.DeletePatient)
    // Step 1: Create a patient to delete
    patientBody := map[string]interface{}{
        "name":        "To Be Deleted",
        "age":         50,
        "disease":     "None",
        "address":     "Unknown",
        "assigned_to": "Dr. Gone",
    }
    bodyBytes, _ := json.Marshal(patientBody)

    req := httptest.NewRequest("POST", "/api/patients", bytes.NewReader(bodyBytes))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+token)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusCreated, w.Code)

    var created map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &created)
    patientID := created["id"]

    // Step 2: Delete the patient
    deleteReq := httptest.NewRequest("DELETE", fmt.Sprintf("/api/patients/%v", patientID), nil)
    deleteReq.Header.Set("Authorization", "Bearer "+token)
    deleteW := httptest.NewRecorder()
    router.ServeHTTP(deleteW, deleteReq)

    assert.Equal(t, http.StatusOK, deleteW.Code)

    // Step 3: Confirm patient is gone
    getReq := httptest.NewRequest("GET", "/api/patients", nil)
    getReq.Header.Set("Authorization", "Bearer "+token)
    getW := httptest.NewRecorder()
    router.ServeHTTP(getW, getReq)

    var patients []map[string]interface{}
    json.Unmarshal(getW.Body.Bytes(), &patients)

    for _, p := range patients {
        if p["id"] == patientID {
            t.Fatalf("Patient %v still exists after deletion", patientID)
        }
    }
}
