# Holy Raging Mages

## TODO

### Now

* Common HTTP server testing packaged with service module (if applicable)
* Test data builder packaged with service module
  * Complete template Get handler function
  * Add GET and PUT HTTP server tests

### Next

* Common config constants so modules aren't free wheeling config "strings" all over the place
* Remove need for repositories to have their own database *Tx
  * Why did I write this ^^
* Script `test-service $1`
  * Establish environment and run specific service tests
* Script `generate`
  * Generate a new service from the template service
* Service to service communications (gRPC)
* Mage service
* Other services

## License

COPYRIGHT 2020 ALIENSPACES alienspaces@gmail.com
