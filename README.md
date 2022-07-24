# GO API

- For this API Ive used mongodb, this is a test api that implements updates, lists and removes of mongodb files within a collection
- This Project can be build by first running go mod init
- After that you can run go build -o binary_name main
- This API uses the official mongodb driver since mgo has been out of maintenance for some time now

# Explanation of the Files
- main.go -> main.go only contains the initialization of the http server and the instantiation of a router
- routes.go -> routes.go contains the static routes for the API with its structs as well as a function that creates the routes based on an array
- movie.go -> movie.go only contains the creation of the necessary structs for the type movie
- actions.go -> actions.go is the file that contains all the logic for the API and the file that makes the interaction with mongodb possible

# Usage

The usage of this API in my case was just for testing and fun, there will probably be an update where I containerize the application as well as I will run snyk to check if the API/container image has some known vulnerabilities, I will also push a folder with the necessary tests to check the functionality of the application and a Jenkins file, in my case to update/remove content Ive used Postman which was extremely helpful to do some manual testing

# Tools I've used

- vscode with the ansible extension
- Postman
- mongodb
- Studio 3T for mongodb (this last one is to be able to manage the collections in a visual way)