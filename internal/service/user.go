package service

import (
	"gin-gorm-web/internal/config"
	"gin-gorm-web/internal/dao"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUsers 获取所有用户
// @Summary 获取所有用户
// @Description 获取所有用户的列表
// @Tags 用户
// @Accept json
// @Produce json
// @Success 200 {array} dao.User
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var users []dao.User
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// GetUser 根据ID获取单个用户
// @Summary 根据ID获取用户
// @Description 根据用户ID获取单个用户的信息
// @Tags 用户
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} dao.User
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	id := c.Param("id")
	var user dao.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// CreateUser 创建新用户
// @Summary 创建新用户
// @Description 创建一个新的用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param user body dao.User true "用户信息"
// @Success 201 {object} dao.User
// @Failure 400 {object} map[string]string
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var user dao.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Create(&user)
	c.JSON(http.StatusCreated, user)
}

// UpdateUser 更新用户信息
// @Summary 更新用户信息
// @Description 根据用户ID更新用户的信息
// @Tags 用户
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Param user body dao.User true "用户信息"
// @Success 200 {object} dao.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user dao.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Save(&user)
	c.JSON(http.StatusOK, user)
}

// DeleteUser 删除用户
// @Summary 删除用户
// @Description 根据用户ID删除用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param id path int true "用户ID"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var user dao.User

	if err := config.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	config.DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
