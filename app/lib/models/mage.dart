// import 'package:logging/logging.dart';
import 'package:flutter/foundation.dart';

// Application packages
import 'package:holyragingmages/api/api.dart';
import 'package:holyragingmages/models/entity.dart';

/// Mage encapsulates a mages data and methods
class Mage extends Entity {
  Mage({Key key, Api api}) : super(api: api);

  // From JSON
  Mage.fromJson(Api api, Map<String, dynamic> json) : super.fromJson(api, json);
}
