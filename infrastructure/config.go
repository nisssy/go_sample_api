package infrastructure

type Config struct {
	DB struct {
		Production struct {
			Host     string
			UserName string
			Password string
			DbName   string
		}
		Test struct {
			Host     string
			UserName string
			Password string
			DbName   string
		}
	}
	Routing struct {
		Port string
	}
}

func NewConfig() *Config {
	c := new(Config)

	c.DB.Production.Host = "localhost"
	c.DB.Production.UserName = "username"
	c.DB.Production.Password = "password"
	c.DB.Production.DbName = "db_name"

	c.DB.Test.Host = "localhost"
	c.DB.Test.UserName = "username"
	c.DB.Test.Password = "password"
	c.DB.Test.DbName = "db_name_test"

	c.Routing.Port = ":80"

	return c
}
