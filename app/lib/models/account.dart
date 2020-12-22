// import 'dart:collection';
// import 'dart:ffi';

import 'package:logging/logging.dart';
import 'package:flutter/foundation.dart';
import 'package:google_sign_in/google_sign_in.dart';

// Application packages
import 'package:holyragingmages/api/api.dart';

class Account extends ChangeNotifier {
  // Api
  final Api api;

  // Provider
  String provider;
  String providerAccountId;
  String providerToken;

  // Account
  String id;
  String email;
  String name;

  GoogleSignIn _googleSignIn = GoogleSignIn(
    scopes: <String>[
      'email',
      'profile',
      'openid',
    ],
  );

  // Constructor
  Account({Key key, this.api}) {
    // Logger
    final log = Logger('Account - constructor');
    log.info('Constructing new account');
  }

  Future<void> handleGoogleSignIn() async {
    // Logger
    final log = Logger('Account - handleGoogleSignIn');
    return Future(() async {
      log.info('Signing in');
      GoogleSignInAccount account = await _googleSignIn.signIn();
      return await signInAccount(account);
    });
  }

  // Generic sign out handles whichever provider initially used
  Future<void> handleGoogleSignOut() async {
    // Logger
    final log = Logger('Account - handleSignOut');

    return Future(() async {
      log.info('Signing out');
      GoogleSignInAccount account = await _googleSignIn.disconnect();
      return await signOutAccount(account);
    });
  }

  Future<void> signInAccount(GoogleSignInAccount account) async {
    // Logger
    final log = Logger('Account - signInAccount');

    return Future(() async {
      log.info('Signing into account');
      // Provider
      this.provider = 'google';
      this.providerToken = null;
      this.providerAccountId = account.id;

      // Provider token
      GoogleSignInAuthentication auth = await account.authentication;

      this.providerToken = auth.accessToken;

      log.info('Signed in provider ${this.provider}');
      log.info('Signed in providerAccountId ${this.providerAccountId}');
      log.info('Signed in providerToken ${this.providerToken}');

      // JWT for additional API calls
      Map<String, dynamic> data = {
        "data": {
          "provider": "google",
          "provider_account_id": this.providerAccountId,
          "provider_token": this.providerToken,
        },
      };

      List<dynamic> accountsData = await this.api.postAuth(data);
      log.info('Post returned ${accountsData.length} length');
      for (Map<String, dynamic> accountData in accountsData) {
        log.info('Post has account data $accountData');
        // Account
        this.id = accountData['account_id'];
        this.name = accountData['account_name'];
        this.email = accountData['account_email'];
        // Set API token to use from now on
        this.api.apiToken = accountData['token'];
      }

      log.info('Signed in');
    });
  }

  Future<void> signOutAccount(GoogleSignInAccount account) async {
    // Logger
    final log = Logger('Account - signOutAccount');

    return Future(() {
      log.info('Signing out of account');
      this.provider = null;
      this.providerToken = null;
      this.providerAccountId = null;

      this.id = null;
      this.email = null;
      this.name = null;

      log.info('Signed out');
    });
  }
}
