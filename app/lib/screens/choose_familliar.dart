import 'dart:async';
import 'package:flutter/material.dart';
import 'package:logging/logging.dart';
import 'package:provider/provider.dart';

// Application packages
import 'package:holyragingmages/api/api.dart';
import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/widgets/choose_familliar_list.dart';
import 'package:holyragingmages/widgets/choose_familliar_button.dart';

class ChooseFamilliarScreen extends StatefulWidget {
  final Api api;

  ChooseFamilliarScreen({Key key, this.api}) : super(key: key);

  @override
  _ChooseFamilliarScreenState createState() => _ChooseFamilliarScreenState();
}

class _ChooseFamilliarScreenState extends State<ChooseFamilliarScreen> {
  // Loading state
  ModelState _loadingState = ModelState.initial;

  @override
  void initState() {
    // Logger
    final log = Logger('ChooseFamilliarScreen - initState');

    log.info("Initialising");

    // Account model
    var accountModel = Provider.of<Account>(context, listen: false);

    // Mage model
    var mageModel = Provider.of<Mage>(context, listen: false);

    mageModel.clear();
    mageModel.accountId = accountModel.id;

    // Starter mage collection model
    var familliarStarterCollectionModel =
        Provider.of<StarterFamilliarCollection>(context, listen: false);

    if (familliarStarterCollectionModel.canLoad()) {
      log.info("Fetching starter mages");
      setState(() {
        _loadingState = ModelState.processing;
      });
      familliarStarterCollectionModel.load().then((FutureOr<void> v) {
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
    final log = Logger('ChooseFamilliarScreen - build');

    log.info("Building");

    if (_loadingState == ModelState.processing) {
      log.info("Processing");
      return Container(
        child: Text('......'),
      );
    }

    // Starter mage collection model
    var familliarStarterCollectionModel = Provider.of<StarterMageCollection>(context, listen: true);

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
                child: ChooseFamilliarListWidget(
                  starterFamilliarList: familliarStarterCollectionModel.mages,
                ),
              ),
            ),
            Container(
              child: ChooseFamilliarButtonWidget(),
            ),
          ],
        ),
      ),
      resizeToAvoidBottomInset: false,
    );
  }
}
