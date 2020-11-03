import 'dart:collection';
import 'package:logging/logging.dart';
import 'package:flutter/foundation.dart';

// Application packages
import 'package:holyragingmages/api/api.dart';
import 'package:holyragingmages/models/mage.dart';
import 'package:holyragingmages/fault.dart';

enum ModelState { initial, processing, done }

// MageCollection contains a collection of mages
class MageCollection extends ChangeNotifier {
  // Singleton instance
  static MageCollection _instance;

  // Model state
  ModelState _state = ModelState.initial;

  // Model errors
  Fault _fault;
  Fault get fault => _fault;
  void _setFault(Fault fault) {
    _fault = fault;
    notifyListeners();
  }

  // Mage list
  final List<Mage> _mages = [];

  // Backend API
  final Api _api = new Api();

  // Singleton
  factory MageCollection() {
    if (_instance == null) {
      _instance = MageCollection._internal();
    }
    return _instance;
  }

  MageCollection._internal() {
    this._mages.clear();
  }

  // Mage list getter
  UnmodifiableListView<Mage> get mages => UnmodifiableListView(_mages);

  // Model state getter
  ModelState get state => _state;

  void _setState(ModelState state) {
    _state = state;
    notifyListeners();
  }

  // Count mages
  int count() {
    return this._mages.length;
  }

  // Load mages
  void load(String accountId) async {
    // Processing
    _setState(ModelState.processing);

    // Call on API to fetch mages
    List<dynamic> magesData;
    try {
      magesData = await this._api.getEntities(accountId);
    } catch (e) {
      _setFault(Fault(e.toString()));
    }

    // Clear mages first
    this.clear();

    for (Map<String, dynamic> mageData in magesData) {
      var mage = Mage.fromJson(mageData);
      this._mages.add(mage);
    }

    // Processing
    _setState(ModelState.done);

    // Notify listeners
    notifyListeners();
  }

  // Clear mages
  void clear() {
    // Logger
    final log = Logger('Mage - clearMages');

    log.info('Clearing mages');

    this._mages.clear();

    log.info('Mages cleared ${this.count()}');

    notifyListeners();
  }
}
