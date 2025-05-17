package controllers

import (
	"net/http"

	"hello-go/database"
	"hello-go/models"

	"github.com/gin-gonic/gin"
)

// Portfolio Controllers
func GetAllPortfolios(c *gin.Context) {
	var portfolios []models.Portfolio
	database.DB.Find(&portfolios)
	c.JSON(http.StatusOK, gin.H{"data": portfolios})
}

func GetPortfolio(c *gin.Context) {
	var portfolio models.Portfolio
	id := c.Param("id")

	if err := database.DB.First(&portfolio, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": portfolio})
}

func CreatePortfolio(c *gin.Context) {
	var input models.Portfolio
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&input)
	c.JSON(http.StatusCreated, gin.H{"data": input})
}

func UpdatePortfolio(c *gin.Context) {
	var portfolio models.Portfolio
	id := c.Param("id")

	if err := database.DB.First(&portfolio, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}

	var input models.Portfolio
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&portfolio).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": portfolio})
}

func DeletePortfolio(c *gin.Context) {
	var portfolio models.Portfolio
	id := c.Param("id")

	if err := database.DB.First(&portfolio, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Portfolio not found"})
		return
	}

	database.DB.Delete(&portfolio)
	c.JSON(http.StatusOK, gin.H{"data": "Portfolio deleted"})
}

// BlogPost Controllers
func GetAllBlogPosts(c *gin.Context) {
	var blogPosts []models.BlogPost
	database.DB.Find(&blogPosts)
	c.JSON(http.StatusOK, gin.H{"data": blogPosts})
}

func GetBlogPost(c *gin.Context) {
	var blogPost models.BlogPost
	id := c.Param("id")

	if err := database.DB.First(&blogPost, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog post not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": blogPost})
}

func CreateBlogPost(c *gin.Context) {
	var input models.BlogPost
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&input)
	c.JSON(http.StatusCreated, gin.H{"data": input})
}

func UpdateBlogPost(c *gin.Context) {
	var blogPost models.BlogPost
	id := c.Param("id")

	if err := database.DB.First(&blogPost, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog post not found"})
		return
	}

	var input models.BlogPost
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&blogPost).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": blogPost})
}

func DeleteBlogPost(c *gin.Context) {
	var blogPost models.BlogPost
	id := c.Param("id")

	if err := database.DB.First(&blogPost, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog post not found"})
		return
	}

	database.DB.Delete(&blogPost)
	c.JSON(http.StatusOK, gin.H{"data": "Blog post deleted"})
}

// Resume Controllers
func GetAllResumes(c *gin.Context) {
	var resumes []models.Resume
	database.DB.Find(&resumes)
	c.JSON(http.StatusOK, gin.H{"data": resumes})
}

func GetResume(c *gin.Context) {
	var resume models.Resume
	id := c.Param("id")

	if err := database.DB.First(&resume, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resume entry not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": resume})
}

func CreateResume(c *gin.Context) {
	var input models.Resume
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&input)
	c.JSON(http.StatusCreated, gin.H{"data": input})
}

func UpdateResume(c *gin.Context) {
	var resume models.Resume
	id := c.Param("id")

	if err := database.DB.First(&resume, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resume entry not found"})
		return
	}

	var input models.Resume
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&resume).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": resume})
}

func DeleteResume(c *gin.Context) {
	var resume models.Resume
	id := c.Param("id")

	if err := database.DB.First(&resume, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Resume entry not found"})
		return
	}

	database.DB.Delete(&resume)
	c.JSON(http.StatusOK, gin.H{"data": "Resume entry deleted"})
}

// SocialMedia Controllers
func GetAllSocialMedia(c *gin.Context) {
	var socialMedias []models.SocialMedia
	database.DB.Find(&socialMedias)
	c.JSON(http.StatusOK, gin.H{"data": socialMedias})
}

func GetSocialMedia(c *gin.Context) {
	var socialMedia models.SocialMedia
	id := c.Param("id")

	if err := database.DB.First(&socialMedia, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Social media entry not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": socialMedia})
}

func CreateSocialMedia(c *gin.Context) {
	var input models.SocialMedia
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Create(&input)
	c.JSON(http.StatusCreated, gin.H{"data": input})
}

func UpdateSocialMedia(c *gin.Context) {
	var socialMedia models.SocialMedia
	id := c.Param("id")

	if err := database.DB.First(&socialMedia, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Social media entry not found"})
		return
	}

	var input models.SocialMedia
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&socialMedia).Updates(input)
	c.JSON(http.StatusOK, gin.H{"data": socialMedia})
}

func DeleteSocialMedia(c *gin.Context) {
	var socialMedia models.SocialMedia
	id := c.Param("id")

	if err := database.DB.First(&socialMedia, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Social media entry not found"})
		return
	}

	database.DB.Delete(&socialMedia)
	c.JSON(http.StatusOK, gin.H{"data": "Social media entry deleted"})
}
