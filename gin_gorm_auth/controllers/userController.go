package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/adam-pawelek/go_exercise/tree/main/gin_and_gorm/initializers"
	"github.com/adam-pawelek/go_exercise/tree/main/gin_and_gorm/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// @BasePath /signup
// Post  weapon
// @Summary Post a weapon
// @Schemes
// @Description
// @Accept json
// @Produce json
// @Success 200
// @Param        account  body      models.Body  true  "Add weapon"
//
// @Router /signup [post]
func Signup(c *gin.Context) {
	// Get email/pass of req body

	var body struct {
		Email    string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
	}

	//Create the user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash password",
		})
		return
	}
	//Respond
	c.JSON(http.StatusOK, gin.H{
		"message": "created user",
	})
}

// @BasePath /login
// Post  weapon
// @Summary Post a weapon
// @Schemes
// @Description
// @Accept json
// @Produce json
// @Success 200
// @Param        account  body      models.Body  true  "Add weapon"
//
// @Router /login [post]
func Login(c *gin.Context) {
	//Get the email and pass of req body
	var body struct {
		Email    string
		Password string
	}
	fmt.Println("Login")
	fmt.Println(body.Email)
	fmt.Println(body.Password)

	if c.Bind(&body.Email) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	fmt.Println("CZemu?")
	// Look up requestd user
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}
	//Compare sent in pass with saved user pass hash
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	//Generate a jwt token
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECERET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}
	fmt.Println(tokenString, err)
	//send it back

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"token": tokenString,
	})

}

// @BasePath /validate
//
// @Summary Get userjkhvlkjdshakjfhaskjlfhsadkjhfdsakljhdskajfkajs
// @Schemes
// @Description
// @Accept json
// @Produce json
// @Success 200
//
// @Router /validate [get]
func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"message": user,
	})
}
