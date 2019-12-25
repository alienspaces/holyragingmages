# Holy Raging Mages

This is a game.

## Overview

You are a mage with:

* Five spells only
* Five items only
* A familliar (that is Mage talk for a pet)
* Strength, dexterity and intelligence (those are "stats")
* Money $$

You are in this situation where you have to keep fighting monsters and other mages to survive! It is terrible really..

You can:

* Learn new spells (they don't come free)
* Replace spells with other spells (because you can only have five)
* Collect or buy new items (collect off the dead or buy with your money $$)
* Replace items with other items (because you can only have five)
* Replace your familliar with another familliar (sometimes they die or are just not very good as a pet)
* Improve your strength, dexterity and intelligence (lets you learn cooler spells and use more powerful items)
* Get more money $$
* Fight monsters for average items and little money
* Fight other mages for better items and better money
* Decide how you want to fight (this is kind of important)

## Technical

Server

* Separate domain oriented services written in Go
* Deployed to GCE/GKE
* Account creation with OAuth/third party providers only

```plantuml
left to right direction
[Public API] ..> [Entity]
[Public API] ..> [Item]
[Public API] ..> [Spell]
[Entity] ..> [Item]
[Entity] ..> [Spell]
```

Client

* Terminal client written in Go
* Android application built with Flutter

## License

COPYRIGHT 2020 ALIENSPACES alienspaces@gmail.com
