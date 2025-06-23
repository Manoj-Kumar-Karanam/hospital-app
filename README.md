
# 🏥 Hospital Management System - Backend (Go + PostgreSQL)

This project is a simple backend web application built using **Go (Golang)**, providing a hospital management system with **role-based access control**, **JWT authentication**, and **PostgreSQL** as the database.

> 🚀 Built as part of an assignment with clean structure, unit tests, and Postman documentation.

---

## 🧑‍⚕️ Roles & Access

| Role         | Permissions                    |
|--------------|--------------------------------|
| **Receptionist** | Create, Read, Update, Delete Patients |
| **Doctor**       | Read, Update Patients Only    |

---

## 🔐 Authentication

- Single `/register` and `/login` endpoint for both roles.
- Uses **JWT** for authentication.
- Authorization handled via middleware based on user roles.

---

## 📦 Tech Stack

- **Go (Golang)**
- **Gin** - HTTP web framework
- **GORM** - ORM for Go
- **PostgreSQL** - Relational database
- **JWT** - Authentication
- **dotenv** - For environment configuration
- **Postman** - For API testing
- **Unit Tests** - With `testify/assert` and separate `.env.test`

---

## 📁 Project Structure

```
hospital-app/
├── config/
│   └── db.go                        # DB initialization logic
├── controllers/
│   ├── auth_controller.go          # Register/Login logic
│   └── patient_controller.go       # CRUD for patients
├── middleware/
│   └── jwt.go           # JWT + role-based auth
├── models/
│   ├── patient.go                  # Patient model
│   └── user.go                     # User model
├── tests/
│   ├── auth_test.go                # Register/Login tests
│   ├── patient_test.go             # Patient CRUD tests
│   └── test_utils.go               # Common test setup
├── postman/
│   ├── Hospital_API_Auth_Collection.json
│   ├── Hospital_API_Doctor_Collection.json
│   └── Hospital_API_Receptionist_Collection.json
├── .env                            # App config (ignored)
├── .env.test                       # Test config (ignored)
├── go.mod / go.sum                 # Go module dependencies
└── main.go                         # Entry point

```

---

## 🔧 Getting Started

### 1. Clone the Repository
```bash
git clone https://github.com/Manoj-Kumar-Karanam/hospital-app.git
cd hospital-app
```

### 2. Create a `.env` File

```
DB_HOST=localhost
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=hospital
DB_PORT=5432
JWT_SECRET=your_jwt_secret
```

### 3. Run the App
```bash
go run main.go
```

### 4. Run Tests
```bash
go test ./tests/...
```

---

## 📬 API Documentation (Postman)

Import these into Postman:

- [Receptionist Collection](postman/Hospital_API_Receptionist_Collection.json)
- [Doctor Collection](postman/Hospital_API_Doctor_Collection.json)
- [Auth Collection](postman/Hospital_API_Auth_Collection.json)

---

## ✅ Features

- 🔐 Secure JWT-based auth
- 👨‍⚕️ Role-based access control
- 🧪 Fully tested with isolated test DB
- 🧾 Postman collection for testing
- 📂 Clean code organization

---

## 📌 Todo (Optional Enhancements)

- ✅ Swagger documentation
- 🔁 Refresh token mechanism
- 🧑‍💻 Admin role for user management
- 📄 Pagination on patient lists

---

## 👨‍💻 Author

**Manoj Kumar Karanam**  
[GitHub Profile](https://github.com/Manoj-Kumar-Karanam)

---

## 📜 License

MIT License – use freely, just don't forget to credit! 😉
