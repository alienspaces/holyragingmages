# Holy Raging Mages

> This is simply insane! ~ alienspaces

## Overview

Source code for the game [Holy Raging Mages](https://holyragingmages.com).

You are a mage with:

* Three spells
* Three items
* A familliar

You are in this situation where you have to keep fighting monsters and other mages to survive! It is terrible really..

You can:

* Learn new spells (they don't come free)
* Replace spells with other spells (because you can only have three)
* Collect or buy new items (collect off the dead or buy with your money $$)
* Replace items with other items (because you can only have three)
* Replace your familliar with another familliar (sometimes they die or are just not very good as a pet)
* Improve your skill in weapons and spells over time simply by using them successfully.
* Get more money $$
* Fight monsters for average items and little money
* Fight other mages for better items and better money
* Decide how you want to fight (this is kind of important)

### Mages and Familliars

* All mages and familliars will have the following attributes that determine how effective they are in combat with different weapons, armour and spells:
  * Strength
  * Dexterity
  * Intelligence
* Money is earned for being victorious in battle and can be spent purchasing new weapons or armour, training in new spells or adopting a new familliar.
* All familliars
  * Have a single offensive, defensive or supportive spell
  * Cannot change the spells or items they come with
  * Cannot level up

### Spells and Items

* Items and spells will be categorised
  * Offensive
  * Defensive
  * Regenerative
* Items and spells can have combination effects when used successively
* Items can have limited doses or charges (potions and wands)
* Spells can have cool down timers (must wait a certain number of turns before you can use the spell again)

### Attributes

A mages attributes determine how affective a mage is in battle.

#### Strength

##### Weapon Damage

For every point of strength a mage has above a weapons strength requirement the mages effectiveness when using that weapon improves.

For every point of strength a mage has below a weapons strength requirement the mages effectiveness when using that weapon declines.

##### Armour Weight

For every point of strength the mage has below an armours strength requirement the mage will lose the equivalent number of points in dexterity.

| Armour Type | Protection | Strength Requirement |
|-------------|------------|----------------------|
| Heavy  | 4 | 16 ([^1]15) |
| Medium | 3 | 14 ([^2]13) |
| Light  | 2 | 12 ([^3]11) |

[^1]: When heavy armour proficiency is chosen

[^2]: When medium armour proficiency is chosen

[^3]: When light armour proficiency is chosen

#### Dexterity

Dexterity determines a mages accuracy when attacking with a weapon or spell and also determines a mages ability to dodge or evade and attack.

##### Dodge and Evading

| Dexterity | Dodge / Evade Adjustment |
|-----------|--------------------------|
| <= 3      | -4  |
| 4/5       | -3  |
| 6/7       | -2  |
| 8/9       | -1  |
| 10        | 0   |
| 11/12     | +1  |
| 13/14     | +2  |
| 15/16     | +3  |
| >= 17     | +4  |

##### Accuracy

For every point of dexterity a mage has above any weapon or spells requirement the mages accuracy when using that weapon or spell improves.

For every point of dexterity a mage has below any weapon or spells requirement the mages accuracy when using that weapon or spell declines.

#### Intelligence

Intelligence determines how effective a mage is at casting spells.

##### Spell Damage and Duration

For every point of intelligence the mage has above a spells intelligence requirement the mages effectiveness when casting that spell improves.

For every point of intelligence the mage has below a spells intelligence requirement the mages effectiveness when casting that spell declines.

### Proficiencies

A mage may choose to be proficient in two areas only.

A proficiency in a weapon, armour or spell type lowers the ability requirements of all weapons, armour or spells of that type for the mage by a point. The result is that a mage will be more accurate and cause more damage when using a weapon or spell of that type or potentially allow a mage to wear heavier armour with no penalties.

* Weapon
  * Swords
  * Axes
  * Clubs
  * Bows
* Armour
  * Heavy
  * Medium
  * Light
* Spell
  * Fire
  * Ice
  * Earth

### Gameplay

Players do not control turn by turn what actions their mage or familliar do.

Instead a player will be able to decide how their mage and familliar fight by adjusting how offensive, defensive or supportive their actions should be.

This setting can be adjusted for the mage and the familliar independently.

* When playing offensively spells and items categorised as offensive will be used more often
* When playing defensively spells and items categorised as defensive will be used more often
* When playing supportively spells and items categorised as supportive will be used more often

## License

COPYRIGHT 2021 ALIENSPACES alienspaces@gmail.com
