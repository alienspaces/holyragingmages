import 'dart:async';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:logging/logging.dart';

// Application packages
import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/widgets/mage_card_basic.dart';
import 'package:holyragingmages/widgets/mage_card_new.dart';

class MageListWidget extends StatefulWidget {
  @override
  _MageListWidgetState createState() => _MageListWidgetState();
}

class _MageListWidgetState extends State<MageListWidget> {
  // Loading state
  ModelState _loadingState = ModelState.initial;

  @override
  void initState() {
    // Logger
    final log = Logger('MageListWidget - initState');

    log.info("Initialising");

    // Account model
    var accountModel = Provider.of<Account>(context, listen: false);

    // Mage collection model
    var mageCollectionModel = Provider.of<PlayerMageCollection>(context, listen: false);

    if (mageCollectionModel.canLoad()) {
      log.info("Fetching mages");
      setState(() {
        _loadingState = ModelState.processing;
      });
      mageCollectionModel.load(accountModel.id).then((FutureOr<void> v) {
        setState(() {
          _loadingState = ModelState.done;
        });
      });
    }

    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageListWidget - build');

    log.info("Building");

    // Mage list model
    var mageCollectionModel = Provider.of<PlayerMageCollection>(context);

    // List of mages
    var mages = mageCollectionModel.mages;

    if (_loadingState == ModelState.processing) {
      log.info("Processing");
      return Container(
        child: Text('......'),
      );
    }

    // Calculate aspect ratio so we can have 4 mages
    var size = MediaQuery.of(context).size;
    final double itemHeight = (size.height - kToolbarHeight - 50) / 3;
    final double itemWidth = size.width / 2;

    Widget generateCard(int index) {
      // Margin
      EdgeInsetsGeometry cardMargin = EdgeInsets.all(10.0);
      // Padding
      EdgeInsetsGeometry cardPadding = EdgeInsets.all(10.0);

      if (index < mages.length) {
        return Container(
          margin: cardMargin,
          padding: cardPadding,
          alignment: Alignment.center,
          child: MageCardBasic(mages[index]),
        );
      }

      return Container(
        color: Colors.blue[300],
        margin: cardMargin,
        padding: cardPadding,
        alignment: Alignment.center,
        child: MageCardNew(),
      );
    }

    // Have mages
    return GridView.count(
      crossAxisCount: 2,
      crossAxisSpacing: 4,
      mainAxisSpacing: 4,
      childAspectRatio: (itemWidth / itemHeight),
      children: List.generate(
        6,
        (index) => generateCard(index),
      ),
    );
  }
}
