package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"strconv"
)

func main() {
	r := gin.Default()

	v1 := r.Group("api/v1") 
	{
		v1.POST("/user", PostUser)
		v1.GET("/user", GetAllUsers)
		v1.GET("/user/:id", GetUser)
		v1.PUT("/user/:id", PutUser)
		v1.DELETE("/user/:id", DeleteUser)
	}

	r.Run(":8080")
}

type User struct {
	Id uint64 `json:"id" binding:"required"`
	Name string `json:"name"`
	Gender string `json:"gender"`
}

type QueryUser struct {
	Name string `form:"name" binding:"required"`
}

// Assume it will has body: name, id, gender
// curl -X POST -d '{"name":"kai", "id": 2, "gender":"male"}' -H "Content-Type:application/json" 127.0.0.1:8080/api/v1/user
func PostUser(ctx *gin.Context) {
	var request User
	if err := ctx.ShouldBindJSON(&request); err != nil {
		fmt.Println("validate body failed: ", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Write User into DB
	
	ctx.JSON(http.StatusOK, &request)
}

// curl -X GET '127.0.0.1:8080/api/v1/user/22?name=kai'
func GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	// get user by id and name
	parsedId, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if parsedId == 0 {
		fmt.Print("id cannot be 0")
		ctx.JSON(http.StatusBadRequest, "id cannot be 0")
		return
	}
	// name, isExist := ctx.GetQuery("name")
	// if !isExist {
	// 	fmt.Println("Lack of Name Param")
	// 	ctx.JSON(http.StatusBadRequest, "Lack of Naame Param")
	// 	return
	// }
	// or
	var queries QueryUser
	if err := ctx.ShouldBindQuery(&queries); err != nil {
		fmt.Println("Lack of Name Param")
		ctx.JSON(http.StatusBadRequest, "Lack of Naame Param")	
		return
	}

	resId := parsedId
	// resName := name
	resName := queries.Name
	resGender := "male"
	res := User {
		Id: resId,
		Name: resName,
		Gender: resGender,
	} 
	ctx.JSON(http.StatusOK, res)
}

func GetAllUsers(ctx *gin.Context) {
	
}

func PutUser(ctx *gin.Context) {
	
}

func DeleteUser(ctx *gin.Context) {
	
}
