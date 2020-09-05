import 'package:dio/dio.dart';
import 'package:logging/logging.dart';

import 'api.dart';

class MageHandler extends BaseHandler {
  List<MageData> parseResponse(Response response) {
    // Logger
    final log = Logger('parseResponse');

    log.info('Parsing response $response');

    List<MageData> magesData = [];
    if (response.data != null) {
      Map<String, dynamic> data = response.data;
      if (data["data"] != null) {
        (data["data"] as List).forEach((mage) {
          log.info('Adding mage $mage');

          MageData mageData = new MageData(
            id: mage["id"],
            name: mage["name"],
          );
          magesData.add(mageData);
        });
      }
    }

    return magesData;
  }
}

class MageData extends BaseData {
  String id;
  String name;
  int strength;
  int dexterity;
  int intelligence;
  int experience;
  int coin;

  MageData({this.id, this.name});
}
