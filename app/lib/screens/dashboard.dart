import 'package:flutter/material.dart';
// import 'package:google_sign_in/google_sign_in.dart';
import 'package:provider/provider.dart';
import 'package:logging/logging.dart';

import '../models/models.dart';
import '../widgets/mage_list.dart';

class DashboardScreen extends StatelessWidget {
  Widget _buildBody(BuildContext context) {
    // Logger
    final log = Logger('DashboardScreen - _buildBody');

    log.info("Building");

    // Mage list model
    var accountModel = Provider.of<AccountModel>(context);

    if (accountModel.providerAccountId != null) {
      // Current user
      log.info('Current user ID ${accountModel.id ?? ''}');
      log.info('Current user displayName ${accountModel.name ?? ''}');
      log.info('Current user email ${accountModel.email ?? ''}');

      return Scaffold(
        appBar: AppBar(
          title: Text('Mages'),
          actions: <Widget>[
            RaisedButton(
              child: const Text('SIGN OUT'),
              onPressed: accountModel.handleSignOut,
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
