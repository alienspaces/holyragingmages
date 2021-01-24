import 'dart:collection';
import 'package:logging/logging.dart';
import 'package:flutter/foundation.dart';

// Application packages
import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/api/api.dart';
import 'package:holyragingmages/models/familliar.dart';
import 'package:holyragingmages/fault.dart';

// StarterFamilliarCollection contains a collection of mages
class StarterFamilliarCollection extends ChangeNotifier {
  // Api
  final Api api;

  // Familliars
  final List<Familliar> _familliars = [];
  UnmodifiableListView<Familliar> get familliars => UnmodifiableListView(_familliars);

  // Faults
  Fault fault;

  // State
  ModelState state = ModelState.initial;

  // Constructor
  StarterFamilliarCollection({Key key, this.api}) {
    this._familliars.clear();
  }

  // Count familliar
  int count() {
    return this._familliars.length;
  }

  bool canLoad() {
    // Logger
    final log = Logger('StarterFamilliarCollection - canLoad');
    if (state == ModelState.processing) {
      log.info('State is $state, cannot load');
      return false;
    }
    log.info('Can load');
    return true;
  }

  // Load familliar
  Future<void> load() async {
    // Logger
    final log = Logger('StarterFamilliarCollection - load');

    log.info('Loading familliars');

    state = ModelState.processing;

    // Get entities
    this.api.getEntities(null, type: EntityTypeStarterFamilliar).then((List<dynamic> entitiesData) {
      log.info('Adding familliars');

      // Clear familliar
      this._familliars.clear();

      // Add familliar
      for (Map<String, dynamic> entityData in entitiesData) {
        var entity = Familliar.fromJson(api, entityData);
        this._familliars.add(entity);
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

  // Clear familliar
  void clear() {
    // Logger
    final log = Logger('StarterFamilliarCollection - clearFamilliars');

    log.info('Clearing familliar');

    this._familliars.clear();

    log.info('Familliars cleared ${this.count()}');

    notifyListeners();
  }
}
