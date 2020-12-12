import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:logging/logging.dart';

// Application packages
import 'package:holyragingmages/models/models.dart';

class MageCreateButtonWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageCreateButtonWidget - build');

    log.info("Building");

    // Account model
    var accountModel = Provider.of<Account>(context);

    // Mage model
    var mageModel = Provider.of<Mage>(context);
    var mageCollectionModel = Provider.of<MageCollection>(context);

    bool _createEnabled() {
      if (mageModel.name == null || mageModel.name.length == 0) {
        log.info('Mage name is null or empty, create disabled');
        return false;
      }
      if (mageModel.availableAttributePoints() != 0) {
        log.info('Mage points ${mageModel.attributePoints} are unspent, create disabled');
        return false;
      }
      return true;
    }

    void _saveMage() {
      // Logger
      final log = Logger('MageCreateButtonWidget - _saveMage');

      log.info('Saving mage');
      mageModel.save().then((t) {
        log.info('Mage saved, reloading mage collection');
        mageCollectionModel.load(accountModel.id).then((t) {
          log.info('Mage collection reloaded, popping navigation');
          Navigator.pop(context);
        });
      });
    }

    return FloatingActionButton(
      onPressed: _createEnabled() ? _saveMage : null,
      disabledElevation: 0.0,
    );
  }
}
