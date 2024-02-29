package main

import (
    "strings"
    "testing"
)

func TestGeneratePassword(t *testing.T) {
    // Test with lowercase letters
    *useLower = true
    password := generatePassword()
    if !strings.ContainsAny(password, lowercase) {
        t.Errorf("Expected password to contain at least one lowercase letter")
    }

    // Test with uppercase letters
    *useLower = false
    *useUpper = true
    password = generatePassword()
    if !strings.ContainsAny(password, uppercase) {
        t.Errorf("Expected password to contain at least one uppercase letter")
    }

    // Test with digits
    *useUpper = false
    *useDigits = true
    password = generatePassword()
    if !strings.ContainsAny(password, digits) {
        t.Errorf("Expected password to contain at least one digit")
    }

    // Test with special characters
    *useDigits = false
    *useSpecial = true
    password = generatePassword()
    if !strings.ContainsAny(password, special) {
        t.Errorf("Expected password to contain at least one special character")
    }

    // Test with no character set selected
    *useSpecial = false
    defer func() {
        if r := recover(); r == nil {
            t.Errorf("Expected function to panic when no character set is selected")
        }
    }()
    generatePassword()
}