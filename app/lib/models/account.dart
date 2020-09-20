// import 'dart:collection';
import 'package:logging/logging.dart';
import 'package:flutter/foundation.dart';
import 'package:google_sign_in/google_sign_in.dart';

// import '../api/api.dart';

class AccountModel extends ChangeNotifier {
  // Properties
  String id;
  String email;
  String name;
  String provider;
  String providerAccountId;
  String providerToken;

  // Provider specific account types
  GoogleSignInAccount _googleAccount;

  GoogleSignIn _googleSignIn = GoogleSignIn(
    scopes: <String>[
      'email',
      'profile',
      'openid',
    ],
  );

  AccountModel() {
    this.initModel();
  }

  Future<void> handleGoogleSignIn() async {
    try {
      await _googleSignIn.signIn();
    } catch (error) {
      print(error);
    }
  }

// Generic sign out handles whichever provider initially used
  Future<void> handleSignOut() async {
    // Logger
    final log = Logger('AccountModel - handleSignOut');

    // Signed in with Google
    if (this._googleAccount != null) {
      log.info('Signing out');
      _googleSignIn.disconnect().then((value) {
        log.info('Signed out $value');

        this._googleAccount = null;
        this.provider = null;
        this.email = null;
        this.name = null;
        this.providerAccountId = null;
        this.providerToken = null;

        // Notify
        notifyListeners();
      });
    }
    return null;
  }

  void initModel() {
    // Logger
    final log = Logger('AccountModel - onCurrentUserChanged');

    _googleSignIn.onCurrentUserChanged
        .listen((GoogleSignInAccount account) async {
      log.info('Signed in $account');
      this._googleAccount = account;
      this.provider = 'google';
      this.email = account.email;
      this.name = account.displayName;
      this.providerAccountId = account.id;

      // Access token
      GoogleSignInAuthentication auth = await account.authentication;

      this.providerToken = auth.accessToken;

      log.info('Signed in provider ${this.provider}');
      log.info('Signed in providerAccountId ${this.providerAccountId}');
      log.info('Signed in providerToken ${this.providerToken}');

      // Notify
      notifyListeners();
    });
    _googleSignIn.signInSilently(suppressErrors: false);
  }
}
