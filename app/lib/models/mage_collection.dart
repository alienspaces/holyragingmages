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
  // Mages
  final List<Mage> _mages = [];
  UnmodifiableListView<Mage> get mages => UnmodifiableListView(_mages);

  // API
  final Api _api = new Api();

  // Faults
  Fault fault;

  // State
  ModelState state = ModelState.initial;

  // Constructor
  MageCollection() {
    this._mages.clear();
  }

  // Count mages
  int count() {
    return this._mages.length;
  }

  bool canLoad() {
    if (state == ModelState.processing) {
      return false;
    }
    return true;
  }

  // Load mages
  Future<void> load(String accountId) async {
    // Logger
    final log = Logger('Mage - load');

    log.info('Loading mages');

    state = ModelState.processing;

    // Get entities
    this._api.getEntities(accountId).then((List<dynamic> entitiesData) {
      log.info('Adding mages');

      // Clear mages
      this._mages.clear();

      // Add mages
      for (Map<String, dynamic> entityData in entitiesData) {
        var entity = Mage.fromJson(entityData);
        this._mages.add(entity);
      }

      // Done
      state = ModelState.done;

      // Notify listeners
      notifyListeners();
    }).catchError((e) {
      // Fault
      fault = Fault(e.toString());

      // Notify listeners
      notifyListeners();
    });
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
