package main

const (
	user     = "postgres"
	password = "test2test"
	dbname   = "drugcomb"
)

func main() {
	a := App{}
	a.Initialize(user, password, dbname)

	a.Run(":11204")
}
