#!/bin/bash
# Setting a return value to a function


model=ExampleModel
model_l=examplemodel
model_url=example
controller=ExampleController
fields=('Title  string `gorm:"column:title"`' 'Text  string `gorm:"column:text"`' 'Author  string `gorm:"column:author"`')






touch app/Model/$model.go
touch app/Http/Controllers/$controller.go



set -eux; \
	{ \
echo 'type '$model' struct {'; \
echo ' gorm.Model'; \
echo ' ID     int    `gorm:"column:id;primary_key"`'; \
for field in "${fields[@]}"; do
echo ''$field''; \
done
echo '}'; \
} >> config/Database.go


set -eux; \
	{ \
echo 'package Model'; \
echo ''; \
echo 'import ('; \
echo '_ "fmt"'; \
echo '"gorm.io/gorm"'; \
echo ')'; \
echo ''; \
echo 'type '$model' struct {'; \
echo ' gorm.Model'; \
echo ' ID     int    `gorm:"column:id;primary_key"`'; \
for field in "${fields[@]}"; do
echo ''$field''; \
done
echo '}'; \

} > app/Model/$model.go



set -eux; \
	{ \
echo 'package Controllers'; \
echo ''; \
echo 'import ('; \
echo ' "larago/app/Model"'; \
echo ' "larago/config"'; \
echo ' "net/http"'; \
echo ''; \
echo ' "github.com/gin-contrib/sessions"'; \
echo ' "github.com/gin-gonic/gin"'; \
echo ' csrf "github.com/utrack/gin-csrf"'; \
echo ')'; \
echo ''; \
echo 'type '$model'Validation struct {'; \
echo ' Title  string `form:"title" json:"title" binding:"required"`'; \
echo ' Text   string `form:"text" json:"text"`'; \
echo ' Author string `form:"author" json:"author"`'; \
echo '}'; \
echo ''; \
echo 'func '$model'(router *gin.RouterGroup) {'; \
echo ' router.POST("/post_add", '$model'AddPost)'; \
echo ' router.POST("/list/:id/edit", Update'$model')'; \
echo ' router.GET("/list/:id/delete", Delete'$model')'; \
echo ' router.GET("/add", ViewAdd'$model')'; \
echo ' router.GET("/", View'$model'List)'; \
echo ' router.GET("/list/:id", View'$model'ListPrev)'; \
echo '}'; \
echo ''; \
echo 'func '$model'AddPost(c *gin.Context) {'; \
echo ' // Validate input'; \
echo ' var input '$model'Validation'; \
echo ''; \
echo ' if err := c.ShouldBind(&input); err != nil{'; \
echo '  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})'; \
echo '  return'; \
echo ' }'; \
echo ''; \
echo ' // Create role'; \
echo ' '$model_l' := Model.'$model'{Title: input.Title, Text: input.Text, Author: input.Author}'; \
echo ' //Gorm_SQL'; \
echo ' config.DB.Save(&'$model_l')'; \
echo ' //end_Gorm_SQL'; \
echo ''; \
echo ' c.Redirect(http.StatusFound, "/'$model_url'")'; \
echo '}'; \
echo ''; \
echo 'func Update'$model'(c *gin.Context) {'; \
echo ' // Get model if exist'; \
echo ''; \
echo ' //Gorm_SQL'; \
echo ' var model Model.'$model''; \
echo ''; \
echo ' if err := config.DB.Where("id = ?", c.Param("id")).First(&model).Error; err != nil{'; \
echo '  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})'; \
echo '  return'; \
echo ' }'; \
echo ' //end Gorm_SQL'; \
echo ''; \
echo ' // Validate input'; \
echo ' var input '$model'Validation'; \
echo ' if err := c.ShouldBind(&input); err != nil{'; \
echo '  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})'; \
echo '  return'; \
echo ' }'; \
echo ''; \
echo ' //Gorm_SQL'; \
echo ' config.DB.Model(&model).Updates(Model.'$model'{Title: input.Title, Text: input.Text, Author: input.Author})'; \
echo ' //end Gorm_SQL'; \
echo ''; \
echo ' c.Redirect(http.StatusFound, "/'$model_url'")'; \
echo ''; \
echo '}'; \
echo ''; \
echo 'func Delete'$model'(c *gin.Context) {'; \
echo ' // Get model if exist'; \
echo ''; \
echo ' //Gorm_SQL'; \
echo ' var model Model.'$model''; \
echo ''; \
echo ' if err := config.DB.Where("id = ?", c.Param("id")).First(&model).Error; err != nil{'; \
echo '  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})'; \
echo '  return'; \
echo ' }'; \
echo ''; \
echo ' config.DB.Delete(&model)'; \
echo ' //end Gorm_SQL'; \
echo ''; \
echo ' //c.JSON(http.StatusOK, gin.H{"data": true})'; \
echo ' c.Redirect(http.StatusFound, "/'$model_url'")'; \
echo ''; \
echo '}'; \
echo ''; \
echo 'func View'$model'List(c *gin.Context) {'; \
echo ''; \
echo ' //Gorm_SQL'; \
echo ' var model []Model.'$model''; \
echo ' //end Gorm_SQL'; \
echo ' session := sessions.Default(c)'; \
echo ' sessionID := session.Get("user_id")'; \
echo ' sessionName := session.Get("user_name")'; \
echo ' if sessionID == nil {'; \
echo '  //c.JSON(http.StatusForbidden, gin.H{'; \
echo '  // "message": "not authed",'; \
echo '  //})'; \
echo '  c.Redirect(http.StatusFound, "/users/login")'; \
echo '  c.Abort()'; \
echo ' }'; \
echo ' //Gorm_SQL'; \
echo ' config.DB.Find(&model)'; \
echo ' //end Gorm_SQL'; \
echo ''; \
echo ' c.HTML(http.StatusOK, "'$model_l'_list.html", gin.H{"csrf": csrf.GetToken(c), "session_id": sessionID, "session_name": sessionName, "list": model})'; \
echo '}'; \
echo ''; \
echo 'func View'$model'ListPrev(c *gin.Context) { // Get model if exist'; \
echo ''; \
echo ' var model Model.'$model''; \
echo ''; \
echo ' session := sessions.Default(c)'; \
echo ' sessionID := session.Get("user_id")'; \
echo ' sessionName := session.Get("user_name")'; \
echo ' if sessionID == nil {'; \
echo '  //c.JSON(http.StatusForbidden, gin.H{'; \
echo '  // "message": "not authed",'; \
echo '  //})'; \
echo '  c.Redirect(http.StatusFound, "/users/login")'; \
echo '  c.Abort()'; \
echo ' }'; \
echo ' //Gorm_SQL'; \
echo ' if err := config.DB.Where("id = ?", c.Param("id")).First(&model).Error; err != nil{'; \
echo '  c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})'; \
echo '  return'; \
echo ' }'; \
echo ' //end Gorm_SQL'; \
echo ''; \
echo ' //c.JSON(http.StatusOK, gin.H{"data": model })'; \
echo ' c.HTML(http.StatusOK, "'$model_l'_list_prev.html", gin.H{"csrf": csrf.GetToken(c), "session_id": sessionID, "session_name": sessionName,'; \
echo '  "id":     model.ID,'; \
echo '  "title":  model.Title,'; \
echo '  "text":   model.Text,'; \
echo '  "author": model.Author,'; \
echo ' })'; \
echo '}'; \
echo ''; \
echo 'func ViewAdd'$model'(c *gin.Context) { // Get model if exist'; \
echo ''; \
echo ' session := sessions.Default(c)'; \
echo ' sessionID := session.Get("user_id")'; \
echo ' sessionName := session.Get("user_name")'; \
echo ' if sessionID == nil {'; \
echo '  //c.JSON(http.StatusForbidden, gin.H{'; \
echo '  // "message": "not authed",'; \
echo '  //})'; \
echo '  c.Redirect(http.StatusFound, "/users/login")'; \
echo '  c.Abort()'; \
echo ' }'; \
echo ''; \
echo ' //c.JSON(http.StatusOK, gin.H{"data": model})'; \
echo ' c.HTML(http.StatusOK, "'$model_l'_add.html", gin.H{"csrf": csrf.GetToken(c), "session_id": sessionID, "session_name": sessionName})'; \
echo '}'; \
} > app/Http/Controllers/$controller.go


echo '
########################################################################################################################

Add field route in file main.go 
                                                        
'$model_l' := r.Group("/'$model_url'")          
Controllers.'$model'('$model_l'.Group("/"))     
                                                        
edit app/Http/Controllers/'$controller'.go and create views
in folder resources/views/
1. '$model_l'_list.html
2. '$model_l'_list_prev.html
3. '$model_l'_add.html
or run the command
touch resources/views/'$model_l'_add.html resources/views/'$model_l'_list_prev.html resources/views/'$model_l'_list.html
#########################################################################################################################
'
