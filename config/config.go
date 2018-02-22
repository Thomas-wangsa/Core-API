package config

/*
|--------------------------------------------------------------------------
| Package Information
|--------------------------------------------------------------------------
| @author 	: Thomas
| @desc 	: all configuration of this Application
|			  import ioutil for read file
| 			  import yaml.v2 to read yaml file
|			  import route package & general struct package
|
*/

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	Route 	"BDA_API/route"
	General "BDA_API/struct/general"
)

/*
|--------------------------------------------------------------------------
| EnvBDA struct
|--------------------------------------------------------------------------
| @author 					: Thomas
| @desc						: main struct of all configuration in main package
|
*/

type EnvBDA struct {
	App General.App `yaml:"app"`
	Databases map[string]General.Database `yaml:"databases"`
	Router Route.Router
	Error error
}

var Credentials = &EnvBDA{}


/*
|--------------------------------------------------------------------------
| Function Interface Get Credentials
|--------------------------------------------------------------------------
| @author 	: Thomas
| @desc 	: Get credentials data from bda.yaml to EnvBDA struct with Credentials variable
|
*/

func (c *EnvBDA)Get_Credentials(file_name string)  {

	/*
	|--------------------------------------------------------------------------
	| Detail information
	|--------------------------------------------------------------------------
	| @author 	: Thomas
	| @desc 	: Read filename & parse that, set Error in struct if get error
	|
	*/

	credentials_file, err := ioutil.ReadFile(file_name)
		if err != nil {
			c.Error = err
		}
	err = yaml.Unmarshal(credentials_file, c)
		if err != nil {
			c.Error = err
		}

	/*
	|--------------------------------------------------------------------------
	| Detail information
	|--------------------------------------------------------------------------
	| @author 	: Thomas
	| @desc 	: get struct from route package and init the route
	|			  set the result to Credentials variable
	|
	*/

	c.Router = Route.Router{}
	c.Router.Init_Route()
	Credentials = c
}