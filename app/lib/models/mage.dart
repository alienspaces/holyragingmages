import 'dart:collection';

import 'package:flutter/cupertino.dart';

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
  final List<MageModel> _mages = [
    MageModel(name: "Bruce"),
    MageModel(name: "Margeret"),
  ];

  UnmodifiableListView<MageModel> get mages => UnmodifiableListView(_mages);

  /// Adds [mage] to cart.
  void add(MageModel mage) {
    _mages.add(mage);
    // This call tells the widgets that are listening to this model to rebuild.
    notifyListeners();
  }

  /// Removes all items from the cart.
  void removeAll() {
    _mages.clear();
    // This call tells the widgets that are listening to this model to rebuild.
    notifyListeners();
  }
}
