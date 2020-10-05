import 'package:flutter/material.dart';
// import 'package:google_sign_in/google_sign_in.dart';
import 'package:provider/provider.dart';
import 'package:logging/logging.dart';

import '../models/models.dart';
import '../widgets/mage_list.dart';

class DashboardScreen extends StatelessWidget {
  // Display the sign in page if we do not have an account
  Widget _buildBody(BuildContext context) {
    // Logger
    final log = Logger('DashboardScreen - _buildBody');

    log.info("Building");

    // Account model
    var accountModel = Provider.of<AccountModel>(context);

    // Mage list model
    var mageListModel = Provider.of<MageListModel>(context);

    if (accountModel.id != null) {
      // Current user
      log.info('Current user id ${accountModel.id ?? ''}');
      log.info('Current user name ${accountModel.name ?? ''}');
      log.info('Current user email ${accountModel.email ?? ''}');
      log.info(
          'Current user provider id ${accountModel.providerAccountId ?? ''}');
      log.info(
          'Current user provider token ${accountModel.providerToken ?? ''}');

      return Scaffold(
        appBar: AppBar(
          title: Text('Mages'),
          actions: <Widget>[
            RaisedButton(
              child: const Text('SIGN OUT'),
              onPressed: () {
                accountModel.handleSignOut();
                mageListModel.removeMages();
              },
            ),
          ],
        ),
        body: Container(
          child: Center(
            child: MageListWidget(),
          ),
        ),
        floatingActionButton: FloatingActionButton(
          onPressed: () {
            Navigator.pushNamed(context, '/mage_create');
          },
        ),
      );
    } else {
      // Account model
      log.info('Current user id ${accountModel.id ?? ''}');
      log.info('Current user name ${accountModel.name ?? ''}');
      log.info('Current user email ${accountModel.email ?? ''}');
      log.info(
          'Current user provider id ${accountModel.providerAccountId ?? ''}');
      log.info(
          'Current user provider token ${accountModel.providerToken ?? ''}');

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
                onPressed: accountModel.handleGoogleSignIn,
              ),
            ],
          ),
        ),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return _buildBody(context);
  }
}
