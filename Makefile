dependencies.install:
	go get -u github.com/gin-gonic/gin
	go get github.com/gin-contrib/cors
	go get github.com/kelseyhightower/envconfig
	go get gorm.io/driver/postgres
	go get -u gorm.io/gorm