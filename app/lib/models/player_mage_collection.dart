import 'dart:collection';
import 'package:logging/logging.dart';
import 'package:flutter/foundation.dart';

// Application packages
import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/api/api.dart';
import 'package:holyragingmages/models/mage.dart';
import 'package:holyragingmages/fault.dart';

// PlayerMageCollection contains a collection of mages
class PlayerMageCollection extends ChangeNotifier {
  // Api
  final Api api;

  // Mages
  final List<Mage> _mages = [];
  UnmodifiableListView<Mage> get mages => UnmodifiableListView(_mages);

  // Faults
  Fault fault;

  // State
  ModelState state = ModelState.initial;

  // Constructor
  PlayerMageCollection({Key key, this.api}) {
    this._mages.clear();
  }

  // Count mages
  int count() {
    return this._mages.length;
  }

  bool canLoad() {
    // Logger
    final log = Logger('PlayerMageCollection - canLoad');
    if (state == ModelState.processing) {
      log.info('State is $state, cannot load');
      return false;
    }
    log.info('Can load');
    return true;
  }

  // Load mages
  Future<void> load(String accountId) async {
    // Logger
    final log = Logger('PlayerMageCollection - load');

    log.info('Loading mages');

    state = ModelState.processing;

    // Get entities
    this.api.getEntities(accountId).then((List<dynamic> entitiesData) {
      log.info('Adding mages');

      // Clear mages
      this._mages.clear();

      // Add mages
      for (Map<String, dynamic> entityData in entitiesData) {
        var entity = Mage.fromJson(api, entityData);
        this._mages.add(entity);
      }

      // Done
      state = ModelState.done;

      // Notify listeners
      notifyListeners();
    }).catchError((e) {
      // Fault
      fault = Fault(e.toString());

      // Done
      state = ModelState.done;

      // Notify listeners
      notifyListeners();
    });
  }

  // Clear mages
  void clear() {
    // Logger
    final log = Logger('PlayerMageCollection - clearMages');

    log.info('Clearing mages');

    this._mages.clear();

    log.info('Mages cleared ${this.count()}');

    notifyListeners();
  }
}
