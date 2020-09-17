import 'package:flutter/material.dart';
import 'package:google_sign_in/google_sign_in.dart';
import 'package:logging/logging.dart';

import '../widgets/mage_list.dart';

GoogleSignIn _googleSignIn = GoogleSignIn(
  scopes: <String>[
    'email',
    'profile',
    'openid',
  ],
);

class DashboardScreen extends StatefulWidget {
  @override
  State createState() => DashboardScreenState();
}

class DashboardScreenState extends State<DashboardScreen> {
  GoogleSignInAccount _currentUser;

  @override
  void initState() {
    super.initState();
    _googleSignIn.onCurrentUserChanged.listen((GoogleSignInAccount account) {
      setState(() {
        _currentUser = account;
      });
    });
    _googleSignIn.signInSilently(suppressErrors: false);
  }

  Future<void> _handleSignIn() async {
    try {
      await _googleSignIn.signIn();
    } catch (error) {
      print(error);
    }
  }

  Future<void> _handleSignOut() => _googleSignIn.disconnect();

  Widget _buildBody() {
    // Logger
    final log = Logger('DashboardScreen - _buildBody');

    log.info("Building");

    if (_currentUser != null) {
      // Current user
      log.info('Current user ID ${_currentUser.id ?? ''}');
      log.info('Current user displayName ${_currentUser.displayName ?? ''}');
      log.info('Current user email ${_currentUser.email ?? ''}');

      return Scaffold(
        appBar: AppBar(
          title: Text('Mages'),
          actions: <Widget>[
            RaisedButton(
              child: const Text('SIGN OUT'),
              onPressed: _handleSignOut,
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
                onPressed: _handleSignIn,
              ),
            ],
          ),
        ),
      );
    }
  }

  @override
  Widget build(BuildContext context) {
    return _buildBody();
  }
}
