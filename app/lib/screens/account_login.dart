import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:logging/logging.dart';

// Application packages
import 'package:holyragingmages/api/api.dart';
import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/screens/processing.dart';

enum AccountLoginScreenState { ready, processing }

class AccountLoginScreen extends StatefulWidget {
  final Api api;

  AccountLoginScreen({Key key, this.api}) : super(key: key);

  @override
  _AccountLoginScreenState createState() => _AccountLoginScreenState();
}

class _AccountLoginScreenState extends State<AccountLoginScreen> {
  // Screen state
  AccountLoginScreenState state = AccountLoginScreenState.ready;

  // Handle sign in then reroute
  void handleSignIn() {
    // Logger
    final log = Logger('AccountLoginScreen - handleSignIn');

    // State - processing
    setState(() {
      state = AccountLoginScreenState.processing;
    });

    // Account model
    var accountModel = Provider.of<Account>(context, listen: false);

    // Account
    log.fine('Account id ${accountModel.id ?? ''}');
    log.fine('Account name ${accountModel.name ?? ''}');
    log.fine('Account email ${accountModel.email ?? ''}');
    log.fine('Account provider id ${accountModel.providerAccountId ?? ''}');
    log.fine('Account provider token ${accountModel.providerToken ?? ''}');

    accountModel.handleGoogleSignIn().then((_) {
      log.info('Account signed in, routing..');

      // Reroute
      Navigator.of(context).pushReplacementNamed('/mage_list');

      // State - ready
      setState(() {
        state = AccountLoginScreenState.ready;
      });
    });
  }

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('AccountLoginScreen - build');

    log.info("Building");

    // Processing
    if (state == AccountLoginScreenState.processing) {
      return ProcessingScreen();
    }

    // Ready
    return Scaffold(
      appBar: AppBar(
        title: Text('Sign In'),
      ),
      body: Container(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.spaceAround,
          children: <Widget>[
            const Text("You are not currently signed in."),
            RaisedButton(
              child: const Text('SIGN IN'),
              onPressed: handleSignIn,
            ),
          ],
        ),
      ),
    );
  }
}
