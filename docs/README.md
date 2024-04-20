使用了swagger插件生成接口文档,github链接:https://github.com/swaggo/gin-swagger

# gin-swagger

gin middleware to automatically generate RESTful API documentation with Swagger 2.0.

## Usage

### Start using it

1.Add comments to your API source code, See Declarative Comments Format. 

2.Download Swag for Go by using:
```go get -u github.com/swaggo/swag/cmd/swag```
Starting in Go 1.17, installing executables with go get is deprecated. go install may be used instead:
```go install github.com/swaggo/swag/cmd/swag@latest```

3.Run the Swag at your Go project root path(for instance ~/root/go-project-name), Swag will parse comments and generate required files(docs folder and docs/doc.go) at ~/root/go-project-name/docs.
```swag init```

4.Download gin-swagger by using:
```
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
```

Import following in your code:
```
import "github.com/swaggo/gin-swagger" // gin-swagger middleware
import "github.com/swaggo/files" // swagger embed files
```

### Canonical example:

Now assume you have implemented a simple api as following:
```
// A get function which returns a hello world string by json
func Helloworld(g *gin.Context)  {
g.JSON(http.StatusOK,"helloworld")
}
```

So how to use gin-swagger on api above? Just follow the following guide.

1.Add Comments for apis and main function with gin-swagger rules like following:
```
// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context)  {
g.JSON(http.StatusOK,"helloworld")
}
```
2.Use swag init command to generate a docs, docs generated will be stored at docs/.
3.import the docs like this: I assume your project named github.com/go-project-name/docs.
```
import (
docs "github.com/go-project-name/docs"
)
```
4.build your application and after that, go to http://localhost:8080/swagger/index.html ,you to see your Swagger UI.
5.The full code and folder relatives here:
```
package main

import (
"github.com/gin-gonic/gin"
docs "github.com/go-project-name/docs"
swaggerfiles "github.com/swaggo/files"
ginSwagger "github.com/swaggo/gin-swagger"
"net/http"
)
// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context)  {
g.JSON(http.StatusOK,"helloworld")
}

func main()  {
r := gin.Default()
docs.SwaggerInfo.BasePath = "/api/v1"
v1 := r.Group("/api/v1")
{
eg := v1.Group("/example")
{
eg.GET("/helloworld",Helloworld)
}
}
r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
r.Run(":8080")

}
```

Demo project tree, swag init is run at relative .
```
.
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
└── main.go
```

## Project with Nested Directory

```
.
├── cmd
│   └── ginsimple
│       └── main.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── go.mod
├── go.sum
└── internal
├── handlers
│   ├── helloWorld.go
│   └── userHandler.go
└── models
├── profile.go
└── user.go
```

Inorder generate swagger docs for projects with nested directories run the following command
```
swag init -g ./cmd/ginsimple/main.go -o cmd/docs
```
```-o``` will set the auto generated file to the specified path

## Multiple APIs

This feature was introduced in swag v1.7.9

## Configuration

You can configure Swagger using different configuration options
```
func main() {
r := gin.New()

	ginSwagger.WrapHandler(swaggerfiles.Handler,
		ginSwagger.URL("http://localhost:8080/swagger/doc.json"),
		ginSwagger.DefaultModelsExpandDepth(-1))

	r.Run()
}
```
