import 'dart:collection';
import 'package:logging/logging.dart';

import 'package:flutter/foundation.dart';

import '../api/api.dart';

// Constants
const int initialAttributeValue = 10;
const int initialAttributePoints = 40;

/// A Mage encapsulates all mage specific data
class MageModel extends ChangeNotifier {
  // Server sourced properties
  String id;
  String _name;
  int _strength;
  int _dexterity;
  int _intelligence;
  int _points;
  int experience;
  int coin;

  // Runtime properties
  bool currentMage;

  MageModel() {
    this.initDefaults();
  }

  factory MageModel.fromJson(Map<String, dynamic> json) {
    var mage = new MageModel();

    // Logger
    final log = Logger('MageModel - fromJson');

    log.info('Creating mage from $json');

    mage.id = json['id'];
    mage.name = json['name'];
    // Points is "at least" the sum of the current attributes. Anything beyond
    // that are available to distribute.
    mage.points =
        json['points'] != null ? json['points'] : initialAttributePoints;
    mage.strength = json['strength'];
    mage.dexterity = json['dexterity'];
    mage.intelligence = json['intelligence'];

    mage.experience = json['experience'];
    mage.coin = json['coin'];

    return mage;
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

    if (this._points == null) {
      throw 'Mage points must be set before adjusting attributes';
    }

    var difference =
        this._strength != null ? this._strength - value : 0 - value;

    var available =
        this._points - (this._strength + this._dexterity + this._intelligence);

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

    if (this._points == null) {
      throw 'Mage points must be set before adjusting attributes';
    }

    var difference = this._dexterity - value;
    var available =
        this._points - (this._strength + this._dexterity + this._intelligence);

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

    if (this._points == null) {
      throw 'Mage points must be set before adjusting attributes';
    }

    var difference = this._intelligence - value;
    var available =
        this._points - (this._strength + this._dexterity + this._intelligence);

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

  int get points {
    var available =
        this._points - (this._strength + this._dexterity + this._intelligence);
    return available;
  }

  set points(int value) {
    this._points = value;
  }

  void initDefaults() {
    // When not given an ID we can assume this is a newly created mage
    if (this.id == null) {
      this._points = initialAttributePoints;
      this._strength = initialAttributeValue;
      this._dexterity = initialAttributeValue;
      this._intelligence = initialAttributeValue;
      this.experience = 0;
      this.coin = 0;
    }
  }
}

class MageListModel extends ChangeNotifier {
  final List<MageModel> _mages = [];
  final Api api = new Api();

  UnmodifiableListView<MageModel> get mages => UnmodifiableListView(_mages);

  /// Creates a [mage] adding it to the existing mages list
  void addMage(MageModel mage) {
    // Validate required
    if (mage.name == null) {
      throw 'Mage name must be set before adding a mage';
    }
    _mages.add(mage);
    // Notify listeners
    notifyListeners();
  }

  /// Refresh all mages
  void refreshMages() {
    // Call on API to fetch mages
    Future<List<dynamic>> magesFuture = this.api.getMages();
    magesFuture.then((magesData) {
      for (Map<String, dynamic> mageData in magesData) {
        var mage = MageModel.fromJson(mageData);
        _mages.add(mage);
      }
      // Notify listeners
      notifyListeners();
    });
  }

  /// Removes all mages from list.
  void removeMages() {
    _mages.clear();
    // This call tells the widgets that are listening to this model to rebuild.
    notifyListeners();
  }
}
