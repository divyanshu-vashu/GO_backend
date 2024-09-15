package controllers

import (
	"21BRS1444_backend/database"
	"21BRS1444_backend/models"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"fmt"
	"21BRS1444_backend/config"
	
)


// ListFiles retrieves file metadata for the current user
func ListFiles(c *fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(uint)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not authenticated"})
	}

	var files []models.File
	if err := database.DB.Where("user_id = ?", userID).Find(&files).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve files"})
	}

	return c.JSON(files)
}

// ShareFile returns a public URL to access the file
func ShareFile(c *fiber.Ctx) error {
	fileID := c.Params("file_id")
	var file models.File

	if err := database.DB.First(&file, fileID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "File not found"})
	}

	return c.JSON(fiber.Map{
		"file_url": "/uploads/" + filepath.Base(file.FilePath),
	})
}


func UploadFile(c *fiber.Ctx) error {
    // Verify if the user is authenticated
	token := c.Get("Authorization")
    fmt.Println("Token received in backend:", token) // Print the token
    
    if token == "" {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User is not logged in"})
    }

    // Clean token format (remove 'Bearer ' prefix)
    if strings.HasPrefix(token, "Bearer ") {
        token = strings.TrimPrefix(token, "Bearer ")
    }

    // Parse the token
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.SecretKey), nil
	})
	

    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
    }

    claims, ok := jwtToken.Claims.(*jwt.StandardClaims)
    if !ok || !jwtToken.Valid {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
    }

    userID, err := strconv.Atoi(claims.Issuer)
    if err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token claims"})
    }

    // Check if user exists
    var user models.User
    if err := database.DB.Where("id = ?", userID).First(&user).Error; err != nil {
        return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not found"})
    }

    // Rest of the file upload logic
    file, err := c.FormFile("file")
    if err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "No file uploaded"})
    }

    if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create upload directory"})
    }

    ext := filepath.Ext(file.Filename)
    savePath := filepath.Join("uploads", file.Filename)
    if err := c.SaveFile(file, savePath); err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save file"})
    }

    fileMeta := models.File{
        UserID:     uint(userID), // Change to uint if your model uses uint for ID
        Name:       file.Filename,
        UploadDate: time.Now().Format(time.RFC3339),
        FileType:   ext,
        FilePath:   savePath,
    }

    if err := database.DB.Create(&fileMeta).Error; err != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save file metadata"})
    }

    return c.JSON(fiber.Map{
        "file_url": "/uploads/" + file.Filename,
    })
}
