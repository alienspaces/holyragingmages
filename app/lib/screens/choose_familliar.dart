import 'dart:async';
import 'package:flutter/material.dart';
import 'package:logging/logging.dart';
import 'package:provider/provider.dart';

// Application packages
import 'package:holyragingmages/api/api.dart';
import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/widgets/choose_familliar_list.dart';

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

    // Familliar model
    var familliarModel = Provider.of<Familliar>(context, listen: false);

    familliarModel.clear();
    familliarModel.accountId = accountModel.id;

    // Starter familliar collection model
    var familliarStarterCollectionModel =
        Provider.of<StarterFamilliarCollection>(context, listen: false);

    if (familliarStarterCollectionModel.canLoad()) {
      log.info("Fetching starter familliars");
      setState(() {
        _loadingState = ModelState.processing;
      });

      // Load starter familliars
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

  void chooseFamilliar({Familliar familliar}) {
    // Logger
    final log = Logger('ChooseFamilliarScreen - chooseFamilliar');

    log.info('Choose familliar name ${familliar.id} ${familliar.name}');

    // Familliar model
    var familliarModel = Provider.of<Familliar>(context, listen: false);

    // Copy starter familliar into familliar
    familliarModel.copyFrom(familliar);

    // Navigate to choosing familliar
    Navigator.of(context).pushNamed('/choose-familliar');
  }

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('ChooseFamilliarScreen - build');

    log.info("Building");

    // Starter familliar collection model
    var familliarStarterCollectionModel =
        Provider.of<StarterFamilliarCollection>(context, listen: true);

    if (_loadingState == ModelState.processing ||
        familliarStarterCollectionModel.familliars.isEmpty) {
      log.info("Processing");
      return Container(
        child: Text('......'),
      );
    }

    // Styling
    EdgeInsetsGeometry padding = EdgeInsets.fromLTRB(15, 10, 15, 10);

    return Scaffold(
      appBar: AppBar(
        title: Text('Choose Your Familliar'),
      ),
      body: Container(
        padding: padding,
        child: Column(
          children: <Widget>[
            Container(
              child: Expanded(
                child: ChooseFamilliarListWidget(
                  starterFamilliarList: familliarStarterCollectionModel.familliars,
                  chooseFamilliarCallback: chooseFamilliar,
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
