package main

/*
|--------------------------------------------------------------------------
| Package Information
|--------------------------------------------------------------------------
| @author 	: Thomas, thomas.wangsa@gmail.com
| @desc 	: Main Function of Go
|			  import standard library & config from package config
|
*/

import (
	"log"
	"os"
	"runtime/debug"
	Config "BDA_API/config"
)

/*
|--------------------------------------------------------------------------
| Variable Constant Information
|--------------------------------------------------------------------------
| @author Thomas
| @defaultGCPercent				: set defaultGCPercent to 40,ref =https://golang.org/pkg/runtime/debug/#SetGCPercent
| @name_of_credentials_file		: name of credentials file
|
*/

const defaultGCPercent 			= 40
const name_of_credentials_file	= "env.yaml"

/*
|--------------------------------------------------------------------------
| Function main()
|--------------------------------------------------------------------------
| @author 	: Thomas
| @return 	: void
| @desc 	: Execute os.Exit if return > 0 from Server()
|
*/

func main() {
	go os.Exit(server());
}

/*
|--------------------------------------------------------------------------
| Function server()
|--------------------------------------------------------------------------
| @author 	: Thomas
| @return 	: int (0 or 1)
| @desc 	: run all config and depedencies in golang,
| 			  return 0 is everything ok, return 1 and log if there is an error load data
|
*/

func server() int {

	/*
	|--------------------------------------------------------------------------
	| Detail information
	|--------------------------------------------------------------------------
	| @author 				: Thomas
	| @os.Getenv("GOPATH")	: check the existing of GOPATH environment
	| @os.Getenv("GOGC")	: defaultGCPercent is the value used to to call SetGCPercent if the GOGC
	|						  environment variable is not set or empty. The value here is intended to hit
	|						  the sweet spot between memory utilization and GC effort. It is lower than the
	| 						  usual default of 100 as a lot of the heap in OMS is used to cache
	|						  memory chunks, which have a lifetime of hours if not days or weeks.
	| credentials()			: get all credentials data
	|
	*/

	if os.Getenv("GOPATH") == "" {
		log.Println("ERROR : Environment of GOPATH is not set")
		return 1
	}

	if os.Getenv("GOGC") == "" {
		debug.SetGCPercent(defaultGCPercent)
	}

	err := credentials()
		if(err != nil) {
			log.Println("ERROR : " + err.Error())
			return 1
		}
	return 0
}

/*
|--------------------------------------------------------------------------
| Function credentials()
|--------------------------------------------------------------------------
| @author 					: Thomas
| @return 					: error or nil
| @desc						: get all credentials data
|
*/

func credentials() error {

	/*
	|--------------------------------------------------------------------------
	| Function credentials()
	|--------------------------------------------------------------------------
	| @author 					: Thomas
	| @var credentials 			: init from pointer of envBDA struct in config package
	| @Get_Credentials()		: interface method to check yaml file and parse the data
	| @credentials.Router.Run	: Run the API
	|
	*/

	var credentials  = &Config.EnvBDA{}
	credentials.Get_Credentials(name_of_credentials_file)

		if(credentials.Error != nil) {
			return credentials.Error
		}

	credentials.Router.Run(credentials.App)
	return nil
}