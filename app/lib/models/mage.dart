import 'package:logging/logging.dart';
import 'package:flutter/foundation.dart';

// Application packages
import 'package:holyragingmages/api/api.dart';

// Constants
const int initialAttributeValue = 10;
const int initialAttributePoints = 40;

/// Mage encapsulates a mages data and methods
class Mage extends ChangeNotifier {
  // Backend API
  final Api _api = new Api();

  // Properties
  String id;
  String accountId;
  String _name;
  int _strength;
  int _dexterity;
  int _intelligence;
  int _attributePoints;
  int experiencePoints;
  int coins;

  Mage() {
    this.initModel();
  }

  factory Mage.fromJson(Map<String, dynamic> json) {
    // Logger
    final log = Logger('Mage - fromJson');

    var mage = new Mage();

    log.info('Creating mage from $json');

    mage.updateFromJson(json);

    return mage;
  }

  void updateFromJson(Map<String, dynamic> json) {
    this.id = json['id'];
    this.accountId = json['account_id'];
    this.name = json['name'];

    // Attribute points are "at least" the sum of the current attributes. Anything beyond
    // that are available to distribute.
    this.attributePoints =
        json['attribute_points'] != null ? json['attribute_points'] : initialAttributePoints;

    this.strength = json['strength'];
    this.dexterity = json['dexterity'];
    this.intelligence = json['intelligence'];

    this.experiencePoints = json['experiencePoints'] != null ? json['experiencePoints'] : 0;
    this.coins = json['coins'];
  }

  Map<String, dynamic> toJson() {
    // Logger
    final log = Logger('Mage - toJson');

    Map<String, dynamic> json = {};

    json["data"] = {
      "id": this.id,
      "account_id": this.accountId,
      "name": this.name,
      "strength": this.strength,
      "dexterity": this.dexterity,
      "intelligence": this.intelligence,
      "attribute_points": this._attributePoints,
      "experience_points": this.experiencePoints,
      "coins": this.coins,
    };

    log.info('Returning json $json');

    return json;
  }

  String get name {
    return this._name;
  }

  set name(String value) {
    this._name = value;
    notifyListeners();
  }

  int get strength {
    return this._strength;
  }

  set strength(int value) {
    // Logger
    final log = Logger('Mage - strength');

    if (this._attributePoints == null) {
      throw 'Mage points must be set before adjusting attributes';
    }

    var difference = this._strength != null ? this._strength - value : 0 - value;

    var available = this._attributePoints - (this._strength + this._dexterity + this._intelligence);

    log.info(
        'Adjust value $value current ${this._strength} difference $difference available $available');

    if (available + difference >= 0) {
      log.info('Setting strength to $value');
      this._strength = value;
      notifyListeners();
      return;
    }

    log.info('Leaving strength as ${this._strength}');
    return;
  }

  int get dexterity {
    return this._dexterity;
  }

  set dexterity(int value) {
    // Logger
    final log = Logger('Mage - dexterity');

    if (this._attributePoints == null) {
      throw 'Mage points must be set before adjusting attributes';
    }

    var difference = this._dexterity - value;
    var available = this._attributePoints - (this._strength + this._dexterity + this._intelligence);

    log.info(
        'Adjust value $value current ${this._strength} difference $difference available $available');

    if (available + difference >= 0) {
      log.info('Setting dexterity to $value');
      this._dexterity = value;
      notifyListeners();
      return;
    }

    log.info('Leaving dexterity as ${this._dexterity}');
    return;
  }

  int get intelligence {
    return this._intelligence;
  }

  set intelligence(int value) {
    // Logger
    final log = Logger('Mage - intelligence');

    if (this._attributePoints == null) {
      throw 'Mage points must be set before adjusting attributes';
    }

    var difference = this._intelligence - value;
    var available = this._attributePoints - (this._strength + this._dexterity + this._intelligence);

    log.info(
        'Adjust value $value current ${this._strength} difference $difference available $available');

    if (available + difference >= 0) {
      log.info('Setting intelligence to $value');
      this._intelligence = value;
      notifyListeners();
      return;
    }

    log.info('Leaving intelligence as ${this._intelligence}');
    return;
  }

  int get attributePoints {
    var available = this._attributePoints - (this._strength + this._dexterity + this._intelligence);
    return available;
  }

  set attributePoints(int value) {
    this._attributePoints = value;
  }

  void initModel() {
    // When not given an ID we can assume this is a newly created mage
    if (this.id == null) {
      this._attributePoints = initialAttributePoints;
      this._strength = initialAttributeValue;
      this._dexterity = initialAttributeValue;
      this._intelligence = initialAttributeValue;
      this.experiencePoints = 0;
      this.coins = 0;
    }
  }

  // Save this mage to the server
  void save() async {
    // Logger
    final log = Logger('Mage - addMage');

    List<dynamic> magesData;
    try {
      magesData = await this._api.postEntity(this.accountId, this.toJson());
    } catch (e) {
      log.warning('Failed adding mage $e');
      return;
    }

    for (Map<String, dynamic> mageData in magesData) {
      log.info('Post has mage data $mageData');
      this.updateFromJson(mageData);
    }

    // Notify listeners
    notifyListeners();

    return;
  }
}
