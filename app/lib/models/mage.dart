import 'dart:collection';

import 'package:flutter/foundation.dart';

import '../api/api.dart';

/// A Mage encapsulates all mage specific data
class MageModel extends ChangeNotifier {
  // Server sourced properties
  String id;
  String name;
  int strength;
  int dexterity;
  int intelligence;
  int experience;
  int coin;

  // Runtime properties
  bool currentMage;

  MageModel({
    this.id,
    this.name,
    this.strength,
    this.dexterity,
    this.intelligence,
    this.experience,
    this.coin,
  });
}

class MageListModel extends ChangeNotifier {
  final List<MageModel> _mages = [];
  final Api api = new Api();

  UnmodifiableListView<MageModel> get mages => UnmodifiableListView(_mages);

  /// Adds [mage] to list
  void add(MageModel mage) {
    // Call API to save new mage
    _mages.add(mage);
    // Notify listeners
    notifyListeners();
  }

  /// Get all mages
  List<MageModel> refreshMages() {
    // Call on API to fetch mages
    Future<List<MageData>> magesFuture = this.api.getMages();
    magesFuture.then((magesData) {
      for (MageData mageData in magesData) {
        var mage = new MageModel(
          id: mageData.id,
          name: mageData.name,
          strength: mageData.strength,
          dexterity: mageData.dexterity,
          intelligence: mageData.intelligence,
          experience: mageData.experience,
          coin: mageData.coin,
        );
        _mages.add(mage);
      }
      // Notify listeners
      notifyListeners();

      return _mages;
    });

    return null;
  }

  /// Removes all mages from list.
  void removeAll() {
    _mages.clear();
    // This call tells the widgets that are listening to this model to rebuild.
    notifyListeners();
  }
}
