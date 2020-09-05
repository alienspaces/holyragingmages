import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:logging/logging.dart';

import '../models/models.dart';

class MageListWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageListWidget - build');

    log.info("Building");

    // mage list
    var mageListModel = Provider.of<MageListModel>(context);
    var mages = mageListModel.mages;

    if (mages.length == 0) {
      log.info("Fetching mages");
      mageListModel.refreshMages();
      return Text("No mages yet");
    }

    return GridView.count(
      crossAxisCount: 2,
      childAspectRatio: 0.90,
      children: List.generate(
        mages.length,
        (index) => ListTile(
          title: Text(
            mages[index].name,
          ),
        ),
      ),
    );
  }
}
