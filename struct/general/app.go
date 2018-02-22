package general

/*
|--------------------------------------------------------------------------
| Package Information
|--------------------------------------------------------------------------
| @author 	: Thomas
| @desc 	: base struct of App Yaml
|
*/

type App struct {
	Name        string `yaml:"name"`
	Environment string `yaml:"environment"`
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
}