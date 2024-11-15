dependencies.install:
	go get -u github.com/gin-gonic/gin
	go get -u github.com/gin-contrib/cors
	go get -u github.com/kelseyhightower/envconfig
	go get -u gorm.io/driver/postgres
	go get -u gorm.io/gorm
	go get -u github.com/docker/docker/client