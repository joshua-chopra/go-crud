package internal

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"path/filepath"
	"runtime"
)

// Setup /*
/*
Setup func for our app. Currently
just loads the env file specified.
*/
func Setup() {
	loadEnv()
}

/* use the following to ensure that whenever setup is called,
e.g., from tests directory that we will get the package path,
go up a directory to the root of
the project and load the env.
*/
func loadEnv() {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	fmt.Println(basepath)
	//err := godotenv.Load(".env.local")
	err := godotenv.Load(filepath.Join(basepath, "../.env.local"))
	if err != nil {
		log.Fatalf("Could not load env file. Exiting program,"+
			"see error: %v\n", err)
	}
}
