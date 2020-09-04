import 'package:dio/dio.dart';
import 'package:logging/logging.dart';

import 'api.dart';

class MageHandler extends BaseHandler {
  List<MageData> parseResponse(Response response) {
    // Logger
    final log = Logger('parseResponse');

    log.info('Parsing response $response');

    List<MageData> mageData = [];
    return mageData;
  }
}

class MageData extends BaseData {
  String name;
  int strength;
  int dexterity;
  int intelligence;
  int experience;
  int coin;
}
