# Holy Raging Mages

This is a game.

## Overview

You are a mage with:

* Three spells only
* Three items only
* A familliar (that is Mage talk for a pet)
* Strength, dexterity and intelligence (those are "stats")
* Money $$

You are in this situation where you have to keep fighting monsters and other mages to survive! It is terrible really..

You can:

* Learn new spells (they don't come free)
* Replace spells with other spells (because you can only have three)
* Collect or buy new items (collect off the dead or buy with your money $$)
* Replace items with other items (because you can only have three)
* Replace your familliar with another familliar (sometimes they die or are just not very good as a pet)
* Improve your strength, dexterity and intelligence (lets you learn cooler spells and use more powerful items)
* Get more money $$
* Fight monsters for average items and little money
* Fight other mages for better items and better money
* Decide how you want to fight (this is kind of important)

### Mages and Familliars

* Strength
  * Health (for staying alive)
  * Items may have strength pre-requisites
* Dexterity
  * Dodge (for avoiding being hit)
  * Items may have dexterity pre-requisites
* Intelligence
  * Power (for casting spells)
  * Items may have intelligence pre-requisites
* Experience Points (earned for winning, when you level up you can add points to strength, dexterity or intelligence)
* Money (earned for winning)

### Spells and Items

* Items and spells can have combination effects (use this type of spell and then this spell results in this added effect)
* Items can have limited doses or charges (potions and wands)
* Spells can have cool down timers (must wait a certain number of turns before you can use the spell again)

## Start

Start all services

```bash
./script/start
```

[http://localhost:8081/](http://localhost:8081/)

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

Client

* Terminal client
  * Go
* Android application
  * Flutter

## License

COPYRIGHT 2020 ALIENSPACES alienspaces@gmail.com
