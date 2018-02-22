package general

/*
|--------------------------------------------------------------------------
| Package Information
|--------------------------------------------------------------------------
| @author 	: Thomas
| @desc 	: base struct of Database Yaml
|
*/

type Database struct {
	Db_name 	string 	`yaml:"db_name"`
	User    	string 	`yaml:"user"`
	Password	string 	`yaml:"password"`
	Host    	string	`yaml:"host"`
	Port    	int		`yaml:"port"`
}
