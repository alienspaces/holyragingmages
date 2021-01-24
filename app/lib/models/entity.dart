import 'package:logging/logging.dart';
import 'package:flutter/foundation.dart';

// Application packages
import 'package:holyragingmages/api/api.dart';

// Constants
const int initialAttributeValue = 10;
const int initialAttributePoints = 40;

/// Entity encapsulates an entities data and methods
class Entity extends ChangeNotifier {
  // Api
  final Api api;

  // Properties
  String id;
  String accountId;
  String name;
  String avatar;
  int _strength;
  int _dexterity;
  int _intelligence;
  int attributePoints;
  int experiencePoints;
  int coins;

  // Constructor
  Entity({Key key, this.api}) {
    // When not given an ID we can assume this is a newly created entity
    if (this.id == null) {
      this._strength = initialAttributeValue;
      this._dexterity = initialAttributeValue;
      this._intelligence = initialAttributeValue;
      this.attributePoints = initialAttributePoints;
      this.experiencePoints = 0;
      this.coins = 0;
    }
  }

  Entity.fromJson(this.api, Map<String, dynamic> json) {
    // Logger
    final log = Logger('Entity - fromJson');

    // this.api = api;

    log.info('Creating entity from $json');

    this.updateFromJson(json);

    log.info('- id ${this.id}');
    log.info('- accountId ${this.accountId}');
    log.info('- name ${this.name}');
    log.info('- avatar ${this.avatar}');
    log.info('- attributePoints ${this.attributePoints}');
    log.info('- strength ${this.strength}');
    log.info('- dexterity ${this.dexterity}');
    log.info('- intelligence ${this.intelligence}');
    log.info('- experiencePoints ${this.experiencePoints}');
    log.info('- coins ${this.coins}');
  }

  void updateFromJson(Map<String, dynamic> json) {
    this.id = json['id'];
    this.accountId = json['account_id'];
    this.name = json['name'];
    this.avatar = json['avatar'];

    // Attribute points are "at least" the sum of the current attributes. Anything beyond
    // that are available to distribute.
    this.attributePoints =
        json['attribute_points'] != null ? json['attribute_points'] : initialAttributePoints;

    this._strength = json['strength'];
    this._dexterity = json['dexterity'];
    this._intelligence = json['intelligence'];

    this.experiencePoints = json['experiencePoints'] != null ? json['experiencePoints'] : 0;
    this.coins = json['coins'];
  }

  Map<String, dynamic> toJson() {
    // Logger
    final log = Logger('Entity - toJson');

    Map<String, dynamic> data = {};
    if (this.id != null) {
      data["id"] = this.id;
    }
    if (this.accountId != null) {
      data["account_id"] = this.accountId;
    }
    if (this.name != null) {
      data["name"] = this.name;
    }
    if (this.avatar != null) {
      data["avatar"] = this.avatar;
    }
    if (this.strength != null) {
      data["strength"] = this.strength;
    }
    if (this.dexterity != null) {
      data["dexterity"] = this.dexterity;
    }
    if (this.intelligence != null) {
      data["intelligence"] = this.intelligence;
    }
    if (this.attributePoints != null) {
      data["attribute_points"] = this.attributePoints;
    }
    if (this.experiencePoints != null) {
      data["experience_points"] = this.experiencePoints;
    }
    if (this.coins != null) {
      data["coins"] = this.coins;
    }

    Map<String, dynamic> json = {"data": data};

    log.info('Returning json $json');

    return json;
  }

  int get strength {
    return this._strength;
  }

  set strength(int value) {
    // Logger
    final log = Logger('Entity - strength');

    if (this.attributePoints == null) {
      throw 'Entity points must be set before adjusting attributes';
    }

    var difference = this._strength != null ? this._strength - value : 0 - value;

    var available = this.attributePoints - (this._strength + this._dexterity + this._intelligence);

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
    final log = Logger('Entity - dexterity');

    if (this.attributePoints == null) {
      throw 'Entity points must be set before adjusting attributes';
    }

    var difference = this._dexterity - value;
    var available = this.attributePoints - (this._strength + this._dexterity + this._intelligence);

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
    final log = Logger('Entity - intelligence');

    if (this.attributePoints == null) {
      throw 'Entity points must be set before adjusting attributes';
    }

    var difference = this._intelligence - value;
    var available = this.attributePoints - (this._strength + this._dexterity + this._intelligence);

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

  int get availableAttributePoints {
    var available = this.attributePoints - (this._strength + this._dexterity + this._intelligence);
    return available;
  }

  // Clear all attributes from this entity
  void clear() {
    this.id = null;
    this.accountId = null;
    this.name = null;
    this.avatar = null;
    this._strength = initialAttributeValue;
    this._dexterity = initialAttributeValue;
    this._intelligence = initialAttributeValue;
    this.attributePoints = initialAttributePoints;
    this.experiencePoints = null;
    this.coins = null;
  }

  // Copy properties from another entity
  void copyFrom(Entity entity) {
    this.name = entity.name;
    this.avatar = entity.avatar;
    this._strength = entity.strength;
    this._dexterity = entity.dexterity;
    this._intelligence = entity.intelligence;
    this.attributePoints = entity.attributePoints;
    this.experiencePoints = entity.experiencePoints;
    this.coins = entity.coins;
  }

  // Save this entity to the server
  Future<void> save() async {
    // Logger
    final log = Logger('Entity - save');

    Map<String, dynamic> saveEntity = this.toJson();

    log.info('Saving entity account ID ${this.accountId} Entity $saveEntity');

    List<dynamic> entitysData;
    try {
      entitysData = await api.postEntity(
        this.accountId,
        saveEntity,
      );
    } catch (e) {
      log.warning('Failed adding entity $e');
      return;
    }

    for (Map<String, dynamic> entityData in entitysData) {
      log.info('Post has entity data $entityData');
      this.updateFromJson(entityData);
    }

    // Notify listeners
    notifyListeners();

    return;
  }
}
