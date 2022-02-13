package service

import (
	"log"
	"nitinaggarwal27/XM-Golang-Exercise/database"
	"nitinaggarwal27/XM-Golang-Exercise/model"

	"github.com/gin-gonic/gin"
)

//isUniqueCode : check if the company code already taken
func isUniqueCode(code string) int {
	var res model.Company
	db := database.GetDB()
	db.Where("code = ?", code).First(&res)
	return int(res.ID)
}

//isUniqueName : check if the company name already taken
func isUniqueName(name string) int {
	var res model.Company
	db := database.GetDB()
	db.Where("name = ?", name).First(&res)
	return int(res.ID)
}

//GetCompany : get a company which find by ID
func GetCompany(c *gin.Context) {
	db := database.GetDB()
	id := c.Params.ByName("id")
	var company model.Company
	db.First(&company, id)

	if company.ID != 0 {
		c.JSON(200, gin.H{
			"error":   false,
			"company": company,
		})
		return
	} else {
		c.JSON(404, gin.H{
			"error":   true,
			"message": "Company not found",
		})
	}
}

//GetCompanies : get all company data
func GetCompanies(c *gin.Context) {
	var companies []model.Company
	db := database.GetDB()
	db.Find(&companies)
	c.JSON(200, gin.H{
		"error": false,
		"list":  companies,
	})
}

//CreateCompany : create a company
func CreateCompany(c *gin.Context) {
	var company model.Company

	if err := c.BindJSON(&company); err != nil {
		log.Println("Bind JSON: ", err)
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Please enter the required fields.",
		})
		return
	}

	db := database.GetDB()
	//manually check unique constraints
	id := isUniqueName(company.Name)
	if id != 0 {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Name already taken.",
		})
		return
	}
	id = isUniqueCode(company.Code)
	if id != 0 {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Code already taken.",
		})
		return
	}

	//finally create the user
	db.Create(&company)

	if company.ID != 0 {
		c.JSON(201, gin.H{
			"error":   false,
			"message": "Company has been created successfully",
		})
	} else {
		c.JSON(500, gin.H{
			"error":   true,
			"message": "Could not create company",
		})
	}
}

//UpdateCompany : edit a company which find by ID
func UpdateCompany(c *gin.Context) {
	var company model.Company
	db := database.GetDB()
	cid := c.Params.ByName("id")
	db.First(&company, cid)

	if company.ID != 0 {
		var updatedCompany model.Company
		c.ShouldBind(&updatedCompany)

		//manually check unique constraints
		id := isUniqueName(company.Name)
		if id != 0 && id != int(company.ID) {
			c.JSON(400, gin.H{
				"error":   true,
				"message": "Name already taken.",
			})
			return
		}
		id = isUniqueCode(company.Code)
		if id != 0 && id != int(company.ID) {
			c.JSON(400, gin.H{
				"error":   true,
				"message": "Code already taken.",
			})
			return
		}

		//finally update the user
		db.Model(&company).Updates(updatedCompany)

		c.JSON(200, gin.H{
			"error":   false,
			"message": "Company has been updated successfully",
		})
	} else {
		c.JSON(404, gin.H{
			"error":   true,
			"message": "Company not found",
		})
	}
}

//DeleteCompany : delete a company which find by ID
func DeleteCompany(c *gin.Context) {
	var company model.Company
	db := database.GetDB()
	id := c.Params.ByName("id")
	db.First(&company, id)

	if company.ID != 0 {
		db.Delete(&company)
		c.JSON(200, gin.H{
			"error":   false,
			"message": "Company has been deleted successfully",
		})
	} else {
		c.JSON(404, gin.H{
			"error":   true,
			"message": "Company not found",
		})
	}
}
