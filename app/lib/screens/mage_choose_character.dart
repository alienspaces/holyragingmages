import 'dart:async';
import 'package:flutter/material.dart';
import 'package:logging/logging.dart';
import 'package:provider/provider.dart';

// Application packages
import 'package:holyragingmages/api/api.dart';
import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/widgets/mage_choose_character_list.dart';
import 'package:holyragingmages/widgets/mage_choose_character_button.dart';

class MageChooseCharacterScreen extends StatefulWidget {
  final Api api;

  MageChooseCharacterScreen({Key key, this.api}) : super(key: key);

  @override
  _MageChooseCharacterScreenState createState() => _MageChooseCharacterScreenState();
}

class _MageChooseCharacterScreenState extends State<MageChooseCharacterScreen> {
  // Loading state
  ModelState _loadingState = ModelState.initial;

  @override
  void initState() {
    // Logger
    final log = Logger('MageChooseCharacterScreen - initState');

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

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageChooseCharacterScreen - build');

    log.info("Building");

    if (_loadingState == ModelState.processing) {
      log.info("Processing");
      return Container(
        child: Text('......'),
      );
    }

    // Starter mage collection model
    var mageStarterCollectionModel = Provider.of<StarterMageCollection>(context, listen: true);

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
                child: MageChooseCharacterListWidget(
                  starterMageList: mageStarterCollectionModel.mages,
                ),
              ),
            ),
            Container(
              child: MageChooseCharacterButtonWidget(),
            ),
          ],
        ),
      ),
      resizeToAvoidBottomInset: false,
    );
  }
}
