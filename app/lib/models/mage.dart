import 'dart:collection';

import 'package:flutter/cupertino.dart';
import 'package:meta/meta.dart';

/// A Mage encapsulates all mage specific data
class MageModel {
  final String id;
  String name;
  int strength;
  int dexterity;
  int intelligence;
  int experience;
  int coin;

  MageModel({
    @required this.id,
    @required this.name,
    @required this.strength,
    @required this.dexterity,
    @required this.intelligence,
    this.experience,
    this.coin,
  });
}

class MageListModel extends ChangeNotifier {
  final List<MageModel> _mages = [];

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
