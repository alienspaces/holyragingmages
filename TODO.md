# Holy Raging Mages

## TODO

### Now

* Modify entities API paths to include `:account_id`
* When `:account_id` present in path enforce path `:account_id` on all entity resources

### Next

* Attribute points
  * All `attribute_points` should be fully allocated across attributes on mage creation
  * Sum of attributes points should never exceed `attribute_points` when updating attributes
* Mage items
  * Default items when a mage is created
  * Enhance mage API to equipped items
  * Update mage card UI to show equipped items
  * Provide ability to click on item in UI to show details
* Mage spells
  * Default spells when a mage is created
  * Enhance mage API to include equipped spells
  * Update mage card UI to show equipped spells
  * Provide ability to click on spell in UI to show details

### Future

* Server - make upgrading Go version a bit easier
* Server - make handler unit tests cleaning up POST data correctly

## License

COPYRIGHT 2020 ALIENSPACES alienspaces@gmail.com
