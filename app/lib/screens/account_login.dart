import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:logging/logging.dart';

// Application packages
import 'package:holyragingmages/api/api.dart';
import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/screens/processing.dart';
import 'package:holyragingmages/widgets/account_signin_google.dart';
import 'package:holyragingmages/widgets/account_signin_anonymous.dart';

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

  @override
  void initState() {
    super.initState();
  }

  // Handle anonymous sign in then reroute
  void handleAnonymousSignIn() {
    // Logger
    final log = Logger('AccountLoginScreen - handleAnonymousSignIn');

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

    accountModel.handleAnonymousSignIn().then((_) {
      log.info('Account signed in, routing..');

      // Reroute
      Navigator.of(context).pushReplacementNamed('/mage_list');

      // State - ready
      setState(() {
        state = AccountLoginScreenState.ready;
      });
    });
  }

  // Handle Google sign in
  void handleGoogleSignIn() {
    // Logger
    final log = Logger('AccountLoginScreen - handleGoogleSignIn');

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
        margin: const EdgeInsets.fromLTRB(20, 0, 20, 0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.spaceAround,
          children: <Widget>[
            Container(
              child: Column(
                mainAxisAlignment: MainAxisAlignment.spaceAround,
                children: <Widget>[
                  // Google sign in
                  Container(
                    height: 45.0,
                    margin: EdgeInsets.fromLTRB(0, 10, 0, 10),
                    child: AccountSignInGoogleWidget(signInCallback: handleGoogleSignIn),
                  ),
                  // Anonymous sign in
                  Container(
                    height: 45.0,
                    margin: EdgeInsets.fromLTRB(0, 10, 0, 10),
                    child: AccountSignInAnonymousWidget(signInCallback: handleAnonymousSignIn),
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }
}
