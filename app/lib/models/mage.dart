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

  factory MageModel.fromJson(Map<String, dynamic> json) {
    return MageModel(
      id: json["id"],
      name: json["name"],
      strength: json["strength"],
      dexterity: json["dexterity"],
      intelligence: json["intelligence"],
      experience: json["experience"],
      coin: json["coin"],
    );
  }
}

class MageListModel extends ChangeNotifier {
  final List<MageModel> _mages = [];
  final Api api = new Api();

  UnmodifiableListView<MageModel> get mages => UnmodifiableListView(_mages);

  /// Adds [mage] to list
  void addMage(MageModel mage) {
    // Call API to save new mage
    _mages.add(mage);
    // Notify listeners
    notifyListeners();
  }

  /// Get all mages
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
