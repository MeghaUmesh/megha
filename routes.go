package main

func initializeRoutes() {
	router.GET("/", index)
	router.GET("/index/registration", showregistrationPage)
	router.POST("/index/registration", register)
	router.GET("/index/upload", upload)
	router.GET("/index/login", loginpage)
	router.POST("/index/login/log", login)
	//router.GET("/index/upload",submitpic)
}
