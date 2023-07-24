package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var ninjaWeapons = map[string]string{
	"ninjaStar": "Beginner Ninja Star - Damage 1",
}

// @BasePath /weapons
// Get all weapons godoc
// @Summary Get all weapons
// @Schemes
// @Description
// @Tags weapon
// @Accept json
// @Produce json
// @Success 200 {object} model.WeaponsRequest
// @Router /weapons [get]
func GetWeapons(context *gin.Context) {
	context.JSON(200, gin.H{
		"weapons": ninjaWeapons,
	})
}

// @BasePath /weapons
// Get all weapons godoc
// @Summary Get all weapons
// @Schemes
// @Description
// @Tags weapon
// @Accept json
// @Produce json
// @Success 200 {object} model.WeaponsRequest
//
//	@Param        type    query     string  false  "type"
//
// @Router /weapon [get]
func GetWeapon(context *gin.Context) {
	weaponType := context.Query("type")
	weaponName, ok := ninjaWeapons[weaponType]

	if !ok {
		context.JSON(404, gin.H{
			weaponType: "",
		})
	}
	context.JSON(200, gin.H{
		weaponType: weaponName,
	})

}

// @BasePath /weapon
// Post  weapon
// @Summary Post a weapon
// @Schemes
// @Description
// @Tags weapon
// @Accept json
// @Produce json
// @Success 200 {object} model.WeaponsRequest
// @Param        account  body      model.Weapon  true  "Add weapon"
//
//	@Param        type    query     string  false  "type"
//	@Param        name    query     string  false  "name"
//
// @Router /weapon [post]
func PostWeapon(context *gin.Context) {
	weaponType := context.Query("type")
	weaponName := context.Query("name")

	if len(weaponType) == 0 || len(weaponName) == 0 {
		context.JSON(400, gin.H{
			weaponType: weaponName,
		})
		return
	}

	if _, ok := ninjaWeapons[weaponType]; ok {
		context.JSON(409, gin.H{
			"message": "Wepaon already exist",
		})
		return
	}
	ninjaWeapons[weaponType] = weaponName

	context.JSON(201, gin.H{
		weaponType: weaponName,
	})
}

// @BasePath /weapon
// Post  weapon
// @Summary Post a weapon
// @Schemes
// @Description
// @Tags weapon
// @Accept json
// @Produce json
// @Param        type   path  string  true  "Account ID"
// @Success 200
// @Router /weapon/{type} [delete]
func DeleteWeapon(context *gin.Context) {
	weaponType := context.Param("type")
	fmt.Println(weaponType)
	if _, ok := ninjaWeapons[weaponType]; !ok {
		context.JSON(404, gin.H{
			"message": "Weapon doesn't exist",
		})
	}

	delete(ninjaWeapons, weaponType)
	context.JSON(204, gin.H{
		"message": "Weapon deleted successfully",
	})

}
