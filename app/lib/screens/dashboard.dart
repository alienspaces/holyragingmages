import 'package:flutter/material.dart';
import 'package:google_sign_in/google_sign_in.dart';

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
    if (_currentUser != null) {
      return Scaffold(
        appBar: AppBar(
          title: Text('Mages'),
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

      // return Column(
      //   mainAxisAlignment: MainAxisAlignment.spaceAround,
      //   children: <Widget>[
      //     ListTile(
      //       leading: GoogleUserCircleAvatar(
      //         identity: _currentUser,
      //       ),
      //       title: Text(_currentUser.displayName ?? ''),
      //       subtitle: Text(_currentUser.email ?? ''),
      //     ),
      //     const Text("Signed in successfully."),
      //     Text(_contactText ?? ''),
      //     RaisedButton(
      //       child: const Text('SIGN OUT'),
      //       onPressed: _handleSignOut,
      //     ),
      //     RaisedButton(
      //       child: const Text('REFRESH'),
      //       onPressed: _handleGetContact,
      //     ),
      //   ],
      // );
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

  // @override
  // Widget build(BuildContext context) {
  //   return Scaffold(
  //     appBar: AppBar(
  //       title: Text('Mages'),
  //     ),
  //     body: Container(
  //       child: Center(
  //         child: MageListWidget(),
  //       ),
  //     ),
  //     floatingActionButton: FloatingActionButton(
  //       onPressed: () {
  //         Navigator.pushNamed(context, '/mage_create');
  //       },
  //     ),
  //   );
  // }
}
