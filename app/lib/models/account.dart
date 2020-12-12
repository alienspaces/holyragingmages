// import 'dart:collection';
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

  // Provider specific account types
  GoogleSignInAccount _googleAccount;

  GoogleSignIn _googleSignIn = GoogleSignIn(
    scopes: <String>[
      'email',
      'profile',
      'openid',
    ],
  );

  // Constructor
  Account({Key key, this.api}) {
    this.initModel();
  }

  Future<void> handleGoogleSignIn() async {
    // Logger
    final log = Logger('Account - handleGoogleSignIn');
    try {
      log.info('Signing in');
      await _googleSignIn.signIn();
    } catch (error) {
      log.warning('Error signing in $error');
    }
  }

  // Generic sign out handles whichever provider initially used
  Future<void> handleSignOut() async {
    // Logger
    final log = Logger('Account - handleSignOut');

    // Signed in with Google
    if (this._googleAccount != null) {
      log.info('Signing out');
      _googleSignIn.disconnect().then((value) {
        log.info('Signed out $value');

        this._googleAccount = null;
        this.provider = null;
        this.providerAccountId = null;
        this.providerToken = null;
        this.id = null;
        this.name = null;
        this.email = null;

        // Notify
        notifyListeners();
      });
    }
    return null;
  }

  void verifyAccount(GoogleSignInAccount account) async {
    // Logger
    final log = Logger('Account - verifyAccount');
    // Account has changed however could mean the user has logged out
    log.info('Account changed $account');

    if (account != null) {
      // Provider
      this._googleAccount = account;
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
    } else {
      this._googleAccount = null;
      this.provider = null;
      this.providerToken = null;
      this.providerAccountId = null;

      this.id = null;
      this.email = null;
      this.name = null;
    }

    // Notify
    notifyListeners();

    return;
  }

  void initModel() {
    // Logger
    // final log = Logger('Account - initModel');

    _googleSignIn.onCurrentUserChanged.listen(this.verifyAccount);

    // try {
    //   log.info('Signing in silently');
    //   _googleSignIn.signInSilently(suppressErrors: true);
    // } catch (error) {
    //   log.warning('Error signing in silently $error');
    // }
  }
}
