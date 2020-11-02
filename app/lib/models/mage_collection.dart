import 'dart:collection';
import 'package:logging/logging.dart';
import 'package:flutter/foundation.dart';

import '../api/api.dart';
import 'mage.dart';

// MageCollection contains a collection of MagesModels and provides access
// to server API's for managing mages
class MageCollection extends ChangeNotifier {
  // Singleton instance
  static MageCollection _instance;

  // _mages - Internal list of mages
  final List<Mage> _mages = [];

  // api - Backend interface
  final Api api = new Api();

  // mages - Returns a list of mages
  UnmodifiableListView<Mage> get mages => UnmodifiableListView(_mages);

  // count - Returns the count of mages
  int count() {
    return _mages.length;
  }

  // Singleton
  factory MageCollection() {
    if (_instance == null) {
      _instance = MageCollection._internal();
    }
    return _instance;
  }

  MageCollection._internal() {
    _mages.clear();
  }

  // addMage - Adds a new mage to the list
  Future<Mage> addMage(String accountId, Mage mage) async {
    // Logger
    final log = Logger('Mage - addMage');

    // Maximum allowed
    if (this.count() >= 4) {
      log.warning('Cannot add mage, mage list length ${this.count()}');
      return null;
    }

    // Validate required
    if (mage.name == null) {
      throw 'Mage name must be set before adding a mage';
    }

    List<dynamic> magesData;
    try {
      magesData = await this.api.postEntity(accountId, mage.toJson());
    } catch (e) {
      log.warning('Failed adding mage $e');
      return null;
    }

    for (Map<String, dynamic> mageData in magesData) {
      log.info('Post has mage data $mageData');
      mage = Mage.fromJson(mageData);
      this._mages.add(mage);
    }

    // Notify listeners
    notifyListeners();

    return mage;
  }

  void refreshMages(String accountId) {
    // Call on API to fetch mages
    Future<List<dynamic>> magesFuture = this.api.getEntities(accountId);
    magesFuture.then((magesData) {
      for (Map<String, dynamic> mageData in magesData) {
        var mage = Mage.fromJson(mageData);
        this._mages.add(mage);
      }
      // Notify listeners
      notifyListeners();
    });
  }

  void clearMages() {
    // Logger
    final log = Logger('Mage - clearMages');

    log.info('Clearing mages');

    this._mages.clear();

    log.info('Mages cleared ${this.count()}');

    notifyListeners();
  }
}
