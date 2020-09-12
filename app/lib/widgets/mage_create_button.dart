import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:logging/logging.dart';

import '../models/models.dart';

class MageCreateButtonWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageCreateScreen - build');

    log.info("Building");

    // Mage model
    var mageModel = Provider.of<MageModel>(context);
    var mageListModel = Provider.of<MageListModel>(context);

    bool _createEnabled() {
      if (mageModel.name == null || mageModel.name.length == 0) {
        log.info('Mage name is null or empty, create disabled');
        return false;
      }
      if (mageModel.attributePoints != 0) {
        log.info(
            'Mage points ${mageModel.attributePoints} are unspent, create disabled');
        return false;
      }
      return true;
    }

    void _addMage() {
      mageListModel.addMage(mageModel);
      Navigator.pop(context);
    }

    return FloatingActionButton(
      onPressed: _createEnabled() ? _addMage : null,
      disabledElevation: 0.0,
    );
  }
}
