package infrastructure

type Config struct {
	DB struct {
		DBMySQL struct {
			User   string
			Pass   string
			Host   string
			Port   string
			DBName string
		}
		DBPostgreSQL struct {
			User   string
			Pass   string
			Host   string
			Port   string
			DBName string
		}
	}
}

func NewConfig() *Config {

	c := new(Config)

	c.DB.DBMySQL.User = "user"
	c.DB.DBMySQL.Pass = "secret"
	c.DB.DBMySQL.Host = "0.0.0.0"
	c.DB.DBMySQL.Port = "3306"
	c.DB.DBMySQL.DBName = "sample"

	c.DB.DBPostgreSQL.User = "user"
	c.DB.DBPostgreSQL.Pass = "secret"
	c.DB.DBPostgreSQL.Host = "0.0.0.0"
	c.DB.DBPostgreSQL.Port = "5432"
	c.DB.DBPostgreSQL.DBName = "sample"

	return c
}
