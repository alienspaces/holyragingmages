import 'package:logging/logging.dart';

import 'api.dart';

class MageHandler extends BaseHandler {
  List<MageData> convert(Map<String, dynamic> data) {
    // Logger
    final log = Logger('parseResponse');

    log.info('Parsing response $data');

    List<MageData> magesData = [];
    if (data["data"] != null) {
      (data["data"] as List).forEach((mage) {
        log.info('Adding mage $mage');

        MageData mageData = new MageData(
          id: mage["id"],
          name: mage["name"],
          strength: mage["strength"],
          dexterity: mage["dexterity"],
          intelligence: mage["intelligence"],
        );
        magesData.add(mageData);
      });
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

  MageData({
    this.id,
    this.name,
    this.strength,
    this.dexterity,
    this.intelligence,
  });
}
