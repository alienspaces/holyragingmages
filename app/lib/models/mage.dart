import 'dart:collection';
import 'package:logging/logging.dart';
import 'package:flutter/foundation.dart';

import '../api/api.dart';

// Constants
const int initialAttributeValue = 10;
const int initialAttributePoints = 40;

/// MageModel encapsulates a mages data and methods
class MageModel extends ChangeNotifier {
  // Properties
  String id;
  String _name;
  int _strength;
  int _dexterity;
  int _intelligence;
  int _attributePoints;
  int experiencePoints;
  int coins;

  MageModel() {
    this.initModel();
  }

  factory MageModel.fromJson(Map<String, dynamic> json) {
    // Logger
    final log = Logger('MageModel - fromJson');

    var mage = new MageModel();

    log.info('Creating mage from $json');

    mage.id = json['id'];
    mage.name = json['name'];
    // Points is "at least" the sum of the current attributes. Anything beyond
    // that are available to distribute.
    mage.attributePoints = json['attribute_points'] != null
        ? json['attribute_points']
        : initialAttributePoints;
    mage.strength = json['strength'];
    mage.dexterity = json['dexterity'];
    mage.intelligence = json['intelligence'];

    mage.experiencePoints =
        json['experiencePoints'] != null ? json['experiencePoints'] : 0;
    mage.coins = json['coins'];

    return mage;
  }

  Map<String, dynamic> toJson() {
    // Logger
    final log = Logger('MageModel - toJson');

    Map<String, dynamic> json = {};

    json["data"] = {
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
    final log = Logger('MageModel - strength');

    if (this._attributePoints == null) {
      throw 'Mage points must be set before adjusting attributes';
    }

    var difference =
        this._strength != null ? this._strength - value : 0 - value;

    var available = this._attributePoints -
        (this._strength + this._dexterity + this._intelligence);

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
    final log = Logger('MageModel - dexterity');

    if (this._attributePoints == null) {
      throw 'Mage points must be set before adjusting attributes';
    }

    var difference = this._dexterity - value;
    var available = this._attributePoints -
        (this._strength + this._dexterity + this._intelligence);

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
    final log = Logger('MageModel - intelligence');

    if (this._attributePoints == null) {
      throw 'Mage points must be set before adjusting attributes';
    }

    var difference = this._intelligence - value;
    var available = this._attributePoints -
        (this._strength + this._dexterity + this._intelligence);

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
    var available = this._attributePoints -
        (this._strength + this._dexterity + this._intelligence);
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
}

// MageListModel contains a collection of MagesModels and provides access
// to server API's for managing mages
class MageListModel extends ChangeNotifier {
  final List<MageModel> _mages = [];
  // api
  final Api api = new Api();

  UnmodifiableListView<MageModel> get mages => UnmodifiableListView(_mages);

  void addMage(MageModel mage) {
    // Logger
    final log = Logger('MageModel - addMage');

    // Validate required
    if (mage.name == null) {
      throw 'Mage name must be set before adding a mage';
    }

    Future<List<dynamic>> magesFuture = this.api.postEntity(mage.toJson());
    magesFuture.then((magesData) {
      log.info('Post returned ${magesData.length} length');
      for (Map<String, dynamic> mageData in magesData) {
        log.info('Post has mage data $mageData');
        var mage = MageModel.fromJson(mageData);
        _mages.add(mage);
      }
      // Notify listeners
      notifyListeners();
    });
  }

  void refreshEntities() {
    // Call on API to fetch mages
    Future<List<dynamic>> magesFuture = this.api.getEntities();
    magesFuture.then((magesData) {
      for (Map<String, dynamic> mageData in magesData) {
        var mage = MageModel.fromJson(mageData);
        _mages.add(mage);
      }
      // Notify listeners
      notifyListeners();
    });
  }

  void removeMages() {
    _mages.clear();
    // This call tells the widgets that are listening to this model to rebuild.
    notifyListeners();
  }
}
