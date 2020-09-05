import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:logging/logging.dart';

import '../models/models.dart';
import 'mage_card.dart';

class MageListWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageListWidget - build');

    log.info("Building");

    // Mage list model
    var mageListModel = Provider.of<MageListModel>(context);

    // List of mages
    var mages = mageListModel.mages;

    // No mages yet
    if (mages.length == 0) {
      log.info("Fetching mages");
      mageListModel.refreshMages();
      return Text("No mages yet");
    }

    // Have mages
    return GridView.count(
      crossAxisCount: 2,
      childAspectRatio: 0.90,
      children: List.generate(
        mages.length,
        (index) => ListTile(
          title: MageCard(mages[index]),
        ),
      ),
    );
  }
}
