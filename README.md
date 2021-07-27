# web-service-gin

Sample implementation of the REST framework Gin.

##  Files and folders

 1. go.mod - consists of the module name and the list of packages used
 2. main.go - entry point. Added the routes. At moment validation of token is not in use
 3. docs/swagger.yml - Yaml file with paths defined. It can be accessed at http://localhost:8080/docs
 4. api/album,user - The route handler functions are added in api folder under respective sub folder.
 5. Dockerfile  - Commands	to run the app on docker
 6. makefile -  Commands to start/build/run on docker

## References

 1. https://tutorialedge.net/golang
 2. https://blog.logrocket.com/
 3. https://levelup.gitconnected.com/
 
## Unit - Test
To execute the unit test cases, following steps need to be executed
1. make run-test

## Test Coverage
Test coverage is a term that describes how much of a package's code is exercised by running the package's tests. If executing the test suite causes 80% of the package's source statements to be run, we say that the test coverage is 80%.

1.  go test ./... -v -coverprofile coverage/cover.out ./
2. go tool cover -html=coverage/cover.out -o coverage/cover.html
3. open coverage/cover.html