module github.com/sonhineboy/gsadmin/service

go 1.17

require (
	github.com/dchest/captcha v1.0.0
	github.com/dgrijalva/jwt-go/v4 v4.0.0-preview1
	github.com/gin-gonic/gin v1.9.0
	github.com/go-playground/validator/v10 v10.11.2
	github.com/satori/go.uuid v1.2.0
	github.com/sonhineboy/gsadminGen v0.9.9-0.20240902074829-b56db2211f9e
	go.uber.org/zap v1.1.0
	golang.org/x/crypto v0.17.0
	golang.org/x/time v0.2.0
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
	gopkg.in/yaml.v3 v3.0.1
	gorm.io/datatypes v1.0.5
	gorm.io/driver/mysql v1.4.1
	gorm.io/gorm v1.23.10
)

require (
	github.com/bytedance/sonic v1.8.0 // indirect
	github.com/chenzhuoyu/base64x v0.0.0-20221115062448-fe3a3abad311 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.14.1 // indirect
	github.com/go-playground/universal-translator v0.18.1 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/goccy/go-json v0.10.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/cpuid/v2 v2.0.9 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-isatty v0.0.17 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ugorji/go/codec v1.2.9 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/arch v0.0.0-20210923205945-b76863e36670 // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/sys v0.15.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gorm.io/driver/postgres v0.2.4 // indirect
	gorm.io/driver/sqlite v1.4.1 // indirect
	gorm.io/driver/sqlserver v1.4.0 // indirect
)

//replace github.com/sonhineboy/gsadminGen => ../../gsadminGen
