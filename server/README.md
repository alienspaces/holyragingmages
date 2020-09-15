# Holy Raging Mages - Server

Server application for the [Holy Raging Mages](https://holyragingmages.com]) game.

## Start

Start all services

```bash
./script/start
```

[http://localhost:8082/entity/api](http://localhost:8082/entity/api)

[http://localhost:8082/spell/api](http://localhost:8082/spell/api)

[http://localhost:8082/item/api](http://localhost:8082/item/api)

## Stop

Stop all services

```bash
./script/stop
```

## Technical

Server

* Separate domain oriented services
  * Go
  * Postgres
  * K8s
* Account creation with OAuth/third party providers only

```plantuml
left to right direction
[Public API] ..> [Mage]
[Public API] ..> [Item]
[Public API] ..> [Spell]
[Public API] ..> [Fight]
[Public API] ..> [Tactic]
[Fight] ..> [Mage]
[Fight] ..> [Spell]
[Fight] ..> [Item]
[Fight] ..> [Tactic]
```

## License

COPYRIGHT 2020 ALIENSPACES alienspaces@gmail.com
