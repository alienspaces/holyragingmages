import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:logging/logging.dart';

// Application packages
import 'package:holyragingmages/api/api.dart';
import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/widgets/mage_list.dart';

enum MageListScreenState { ready, processing }

class MageListScreen extends StatefulWidget {
  final Api api;

  MageListScreen({Key key, this.api}) : super(key: key);

  @override
  _MageListScreenState createState() => _MageListScreenState();
}

class _MageListScreenState extends State<MageListScreen> {
  // Screen state
  MageListScreenState state = MageListScreenState.ready;

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
    log.fine('Account provider id ${accountModel.providerAccountId ?? ''}');
    log.fine('Account provider token ${accountModel.providerToken ?? ''}');

    // Mage list model
    var mageListModel = Provider.of<MageCollection>(context, listen: false);

    accountModel.handleGoogleSignOut().then((_) {
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
      return Scaffold(
        body: Container(
          alignment: Alignment.center,
          child: CircularProgressIndicator(),
        ),
      );
    }

    // Mage list model
    var mageListModel = Provider.of<MageCollection>(context);

    return Scaffold(
      appBar: AppBar(
        title: Text('Mages'),
        actions: <Widget>[
          RaisedButton(
            child: const Text('SIGN OUT'),
            onPressed: handleSignOut,
          ),
        ],
      ),
      body: Container(
        child: Center(
          child: MageListWidget(),
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: mageListModel.count() >= 4
            ? null
            : () {
                Navigator.pushNamed(context, '/mage_create');
              },
        child: Icon(Icons.add),
      ),
    );
  }
}
