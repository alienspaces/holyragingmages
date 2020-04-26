# Holy Raging Mages

## TODO

### Now

* Support passing modeller to runner init
  * Service server will replace it with every request
  * Service daemon could reuse it
  * CLI application can use it

### Next

* Run all with Docker Compose
* Common HTTP server testing packaged with service module (if applicable)
* Config
  * Common constants so modules aren't free wheeling configuration key "strings" all over the place
  * Database stored configuration
  * Run-time restart trigger
* Service to service communications (gRPC)

## License

COPYRIGHT 2020 ALIENSPACES alienspaces@gmail.com
