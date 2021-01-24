import 'dart:async';
import 'package:flutter/material.dart';
import 'package:logging/logging.dart';
import 'package:provider/provider.dart';

// Application packages
import 'package:holyragingmages/api/api.dart';
import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/widgets/choose_mage_list.dart';

class ChooseMageScreen extends StatefulWidget {
  final Api api;

  ChooseMageScreen({Key key, this.api}) : super(key: key);

  @override
  _ChooseMageScreenState createState() => _ChooseMageScreenState();
}

class _ChooseMageScreenState extends State<ChooseMageScreen> {
  // Loading state
  ModelState _loadingState = ModelState.initial;

  @override
  void initState() {
    // Logger
    final log = Logger('ChooseMageScreen - initState');

    log.info("Initialising");

    // Account model
    var accountModel = Provider.of<Account>(context, listen: false);

    // Mage model
    var mageModel = Provider.of<Mage>(context, listen: false);

    mageModel.clear();
    mageModel.accountId = accountModel.id;

    // Starter mage collection model
    var mageStarterCollectionModel = Provider.of<StarterMageCollection>(context, listen: false);

    if (mageStarterCollectionModel.canLoad()) {
      log.info("Fetching starter mages");
      setState(() {
        _loadingState = ModelState.processing;
      });

      // Load starter mages
      mageStarterCollectionModel.load().then((FutureOr<void> v) {
        setState(() {
          _loadingState = ModelState.done;
        });
      });
    }

    super.initState();
  }

  @override
  void dispose() {
    super.dispose();
  }

  void chooseMage({Mage mage}) {
    // Logger
    final log = Logger('ChooseMageScreen - chooseMage');

    log.info('Choose mage name ${mage.id} ${mage.name}');

    // Mage model
    var mageModel = Provider.of<Mage>(context, listen: false);

    // Copy starter mage into mage
    mageModel.copyFrom(mage);

    log.info('Navigating to choose familliar');

    // Navigate to choosing familliar
    Navigator.of(context).pushNamed('/choose-familliar');
  }

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('ChooseMageScreen - build');

    log.info("Building");

    // Starter mage collection model
    var mageStarterCollectionModel = Provider.of<StarterMageCollection>(context, listen: true);

    if (_loadingState == ModelState.processing || mageStarterCollectionModel.mages.isEmpty) {
      log.info("Processing");
      return Container(
        child: Text('......'),
      );
    }

    // Styling
    EdgeInsetsGeometry padding = EdgeInsets.fromLTRB(15, 10, 15, 10);

    return Scaffold(
      appBar: AppBar(
        title: Text('Choose Your Mage'),
      ),
      body: Container(
        padding: padding,
        child: Column(
          children: <Widget>[
            Container(
              child: Expanded(
                child: ChooseMageListWidget(
                  starterMageList: mageStarterCollectionModel.mages,
                  chooseMageCallback: chooseMage,
                ),
              ),
            ),
          ],
        ),
      ),
      resizeToAvoidBottomInset: false,
    );
  }
}
