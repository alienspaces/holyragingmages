import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:logging/logging.dart';

// Application packages
import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/widgets/mage_card.dart';

class MageListWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageListWidget - build');

    log.info("Building");

    // Account model
    var accountModel = Provider.of<Account>(context);

    // Mage list model
    var mageCollectionModel = Provider.of<MageCollection>(context);

    if (accountModel.id == null) {
      log.info("Account is null");
      return Text("Not signed in");
    }

    // List of mages
    var mages = mageCollectionModel.mages;

    // No mages yet
    if (mages.length == 0) {
      log.info("Fetching mages");
      mageCollectionModel.load(accountModel.id);
      return Text("No mages yet");
    }

    // Calculate aspect ratio so we can have 4 mages
    var size = MediaQuery.of(context).size;
    final double itemHeight = (size.height - kToolbarHeight - 50) / 2;
    final double itemWidth = size.width / 2;

    // Have mages
    return GridView.count(
      crossAxisCount: 2,
      crossAxisSpacing: 4,
      mainAxisSpacing: 4,
      childAspectRatio: (itemWidth / itemHeight),
      children: List.generate(
        mages.length,
        (index) => Container(
          padding: EdgeInsets.all(10.0),
          child: MageCard(mages[index]),
        ),
      ),
    );
  }
}
