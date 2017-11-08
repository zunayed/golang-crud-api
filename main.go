package main

func main() {
	a := App{}
	a.InitializeDb("postgres", "", "postgres", 5432)
	a.InitializeRouter()
	a.Run(":8080")
	defer a.DB.Close()
}
