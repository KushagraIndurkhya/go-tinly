package main

func main() {

	app := new(App)
	app.Init()
	// app.Router.HandleFunc("/r/{Short}", handlers.Redirect)
	// app.Router.HandleFunc("/api/create", middleware.Auth(handlers.Create))
	// app.Router.HandleFunc("/login", handlers.HandleLogin)
	// app.Router.HandleFunc("/callback-gl", handlers.HandleCallback)
	// app.Router.HandleFunc("/api/dash", middleware.Auth(handlers.Dash))
	// app.Router.HandleFunc("/home", handlers.Index).Methods("GET")
	// app.Router.PathPrefix("/").Handler(http.FileServer(http.Dir("client/build")))
	app.Run()

}
