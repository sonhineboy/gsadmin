package config

type Config struct {
	Db struct {
		Type         string `yaml:"type"`
		MaxIdleConns int    `yaml:"max-idle-conns"`
		MaxOpenConns int    `yaml:"max-open-conns"`
		Port         string `yaml:"port"`
		Host         string `yaml:"host"`
		TablePrefix  string `yaml:"table_prefix"`
		Database     string `yaml:"database"`
		User         string `yaml:"name"`
		PassWord     string `yaml:"password"`
	}
	MyJwt struct {
		Secret string `yaml:"secret"`
	}
	App struct {
		Host       string `yaml:"host"`
		Port       string `yaml:"port"`
		UploadFile string `yaml:"uploadFile"`
	}
}
