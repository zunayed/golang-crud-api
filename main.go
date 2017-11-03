package main

func main() {
	a := App{}

	// TODO - load from conf or env variable
	a.InitializeDb("postgres", "", "postgres")
	a.InitializeRouter()

	a.Run(":8080")
}
