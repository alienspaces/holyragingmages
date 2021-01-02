import 'dart:async';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:logging/logging.dart';

// Application packages
import 'package:holyragingmages/api/api.dart';
import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/screens/processing.dart';
import 'package:holyragingmages/widgets/account_signout.dart';
import 'package:holyragingmages/widgets/mage_list.dart';

enum MageListScreenState { ready, processing }

// Refresh token duration
const Duration refreshTokenDuration = Duration(minutes: 5);

class MageListScreen extends StatefulWidget {
  final Api api;

  MageListScreen({Key key, this.api}) : super(key: key);

  @override
  _MageListScreenState createState() => _MageListScreenState();
}

class _MageListScreenState extends State<MageListScreen> {
  // Screen state
  MageListScreenState state = MageListScreenState.ready;
  Timer timer;

  @override
  void initState() {
    // Logger
    final log = Logger('MageListScreen - initState');

    // TODO: Implement in account model..

    // Refresh API token
    timer = new Timer.periodic(refreshTokenDuration, (Timer timer) {
      log.info('Refreshing token ${widget.api.apiToken}');

      // Auth data
      Map<String, dynamic> data = {
        "data": {
          "token": widget.api.apiToken,
        },
      };

      widget.api.refreshAuth(data).then((accountsData) {
        log.info('Post returned ${accountsData.length} length');
        for (Map<String, dynamic> accountData in accountsData) {
          log.info('Post has account data $accountData');

          // Set API token to use from now on
          widget.api.apiToken = accountData['token'];
        }
      });
    });

    super.initState();
  }

  @override
  void dispose() {
    timer.cancel();
    super.dispose();
  }

  // Handle sign out then reroute
  void handleSignOut() {
    // Logger
    final log = Logger('MageListScreen - handleSignOut');

    // State - processing
    setState(() {
      state = MageListScreenState.processing;
    });

    // Account model
    var accountModel = Provider.of<Account>(context, listen: false);

    // Account
    log.fine('Account id ${accountModel.id ?? ''}');
    log.fine('Account name ${accountModel.name ?? ''}');
    log.fine('Account email ${accountModel.email ?? ''}');

    // Mage list model
    var mageListModel = Provider.of<MageCollection>(context, listen: false);

    accountModel.handleSignOut().then((_) {
      log.info('Account signed out, routing..');

      // Clear mages
      mageListModel.clear();

      // Reroute
      Navigator.of(context).pushReplacementNamed('/');

      // State - ready
      setState(() {
        state = MageListScreenState.ready;
      });
    });
  }

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageListScreen - build');

    log.info("Building");

    // Processing
    if (state == MageListScreenState.processing) {
      return ProcessingScreen();
    }

    return Scaffold(
      appBar: AppBar(
        title: Text('Mages'),
        actions: <Widget>[
          AccountSignOutWidget(
            signOutCallback: handleSignOut,
          ),
        ],
      ),
      body: Container(
        child: Center(
          child: MageListWidget(),
        ),
      ),
    );
  }
}
