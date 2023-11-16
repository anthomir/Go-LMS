package routes

import (
	"github.com/anthomir/GoProject/handlers/chapterHandler"
	"github.com/anthomir/GoProject/handlers/courseHandler"
	"github.com/anthomir/GoProject/handlers/groupHandler"
	"github.com/anthomir/GoProject/handlers/userHandler"
	"github.com/anthomir/GoProject/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.Static("/public", "./public")

	// Define routes for user-related operations
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/signup", userHandler.UserSignUp) // Anyone
		userRoutes.POST("/login", userHandler.UserLogin) // Anyone
		userRoutes.GET("/profile", middleware.RequireUserAuth, userHandler.GetUserProfile)// User
		userRoutes.GET("/teachers", middleware.RequireUserAuth, userHandler.FindAllTeachers)// User
	}

	courseRoutes := router.Group("/course")
	{
		courseRoutes.POST("/", middleware.RequireTeacherAuth, courseHandler.Create) // Teacher
		// courseRoutes.PUT("/", middleware.RequireTeacherAuth, courseHandler.Create) // Teacher

		// ************************ File Upload ************************ //
		courseRoutes.POST("/private-upload", courseHandler.PrivateVideoUploadHandler) // Teacher
		courseRoutes.POST("/private-stream", courseHandler.PrivateVideoStreamHandler) // Users (That Subscribed to this Course)
		courseRoutes.POST("/public-upload", courseHandler.UploadHandler) // Teacher
		// ************************ *********** ************************ //


		courseRoutes.POST("/review/:id", middleware.RequireUserAuth, courseHandler.ReviewById) // User
		courseRoutes.POST("/subscribe/:id", middleware.RequireUserAuth, courseHandler.SubscribeToCourse) // User
		courseRoutes.GET("/", middleware.RequireUserAuth, courseHandler.SearchCoursesHandler) // User
		courseRoutes.GET("/:id", middleware.RequireUserAuth, courseHandler.FindById) // User
		courseRoutes.GET("/user/:id", middleware.RequireUserAuth, courseHandler.FindByUserId) // User

		courseRoutes.DELETE("/", middleware.RequireTeacherAuth, courseHandler.DeleteAll) // Teacher 
	}

	chapterRoutes := router.Group("/chapter")
	{
		chapterRoutes.POST("/", middleware.RequireTeacherAuth, chapterHandler.Create) // Teacher

		chapterRoutes.GET("/course/:id", middleware.RequireUserAuth, chapterHandler.FindAllByCourseId) // User
		chapterRoutes.GET("/:id", middleware.RequireUserAuth, chapterHandler.FindById) // User

		chapterRoutes.DELETE("/", middleware.RequireTeacherAuth, chapterHandler.DeleteAll) // Teacher 
	}

	groupRoutes := router.Group("/group")
	{
		groupRoutes.POST("/", middleware.RequireUserAuth, groupHandler.Create ) // User
		groupRoutes.PUT("/add/:id", middleware.RequireUserAuth, groupHandler.AddUsersToGroup ) // GroupAdmin
		groupRoutes.PUT("/remove/:id", middleware.RequireUserAuth, groupHandler.AddUsersToGroup ) // GroupAdmin

		groupRoutes.GET("/details/:id", middleware.RequireUserAuth, groupHandler.GetGroupDetailsHandler) // User
		groupRoutes.GET("/", middleware.RequireUserAuth, groupHandler.FindByUser ) // User

		groupRoutes.DELETE("/", middleware.RequireUserAuth, groupHandler.DeleteById ) // User
		groupRoutes.DELETE("/all", middleware.RequireAdminAuth, groupHandler.DeleteAll ) // Admin
	}
}
