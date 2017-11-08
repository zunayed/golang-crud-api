package main

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/product/{id:[1-9]+}", a.getProduct).Methods("GET")
}
