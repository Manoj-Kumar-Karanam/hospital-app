package models

import (
    "gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
)

type User struct {
    gorm.Model
    Username string `gorm:"unique"`
    Password string
    Role     string // "receptionist" or "doctor"
}

func (u *User) SetPassword(pw string) error {
    hashed, err := bcrypt.GenerateFromPassword([]byte(pw), 12)
    if err != nil {
        return err
    }
    u.Password = string(hashed)
    return nil
}

func (u *User) CheckPassword(pw string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(pw))
    return err == nil
}
