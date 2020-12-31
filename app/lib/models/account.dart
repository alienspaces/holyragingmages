import 'dart:io';
import 'package:device_info_plus/device_info_plus.dart';
import 'package:flutter/foundation.dart';
import 'package:google_sign_in/google_sign_in.dart';
import 'package:logging/logging.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:uuid/uuid.dart';

// Application packages
import 'package:holyragingmages/api/api.dart';

// Account Provider Type
enum AccountProviderType { anonymous, google }

// Account Provider
class AccountProvider {
  AccountProviderType type;
  String accountId;
  String token;

  AccountProvider({this.type, this.accountId, this.token});
}

class Account extends ChangeNotifier {
  // Api
  final Api api;

  // Device
  bool isAndroid;
  bool isIOS;
  bool isWeb;

  // Account
  String id;
  String email;
  String name;

  // Provider
  AccountProviderType accountProviderType;

  // Google Sign In
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

    initialise();
  }

  void initialise() {
    // Logger
    final log = Logger('Account - initialise');

    DeviceInfoPlugin deviceInfo = DeviceInfoPlugin();

    if (kIsWeb) {
      deviceInfo.webBrowserInfo.then((WebBrowserInfo webBrowserInfo) {
        log.info('Running on web: ${webBrowserInfo.userAgent}');
        isWeb = true;
      });
    } else if (Platform.isAndroid) {
      deviceInfo.androidInfo.then((AndroidDeviceInfo androidInfo) {
        log.info('Running on Android: ${androidInfo.model}');
        isAndroid = true;
      });
    } else if (Platform.isIOS) {
      deviceInfo.iosInfo.then((IosDeviceInfo iosInfo) {
        log.info('Running on IOS: ${iosInfo.utsname.machine}');
        isIOS = true;
      });
    } else {
      log.warning('Unsupported device');
    }
  }

  Future<void> handleAnonymousSignIn() async {
    // Logger
    final log = Logger('Account - handleAnonymousSignIn');

    SharedPreferences prefs = await SharedPreferences.getInstance();

    // Locally stored account ID
    String accountId = (prefs.getString('accountId') ?? Uuid().v4());
    await prefs.setString('accountId', accountId);

    // Provider
    accountProviderType = AccountProviderType.anonymous;

    AccountProvider provider = AccountProvider(
      type: accountProviderType,
      accountId: accountId,
      token: "", // Empty token for anonymous
    );

    return Future(() async {
      log.info('Signing in');
      return await signInAccount(provider);
    });
  }

  Future<void> handleGoogleSignIn() async {
    // Logger
    final log = Logger('Account - handleGoogleSignIn');

    // Provider
    accountProviderType = AccountProviderType.google;

    AccountProvider provider = AccountProvider(
      type: accountProviderType,
    );

    return Future(() async {
      log.info('Signing in');

      int errCount = 0;
      GoogleSignInAccount account;

      while (errCount < 6) {
        try {
          account = await _googleSignIn.signIn();
        } catch (e) {
          log.warning('Error count >$errCount<: Google sign in error ${e.toString()}');
          errCount++;
          await Future.delayed(Duration(seconds: errCount), null);
          continue;
        }
        break;
      }

      // Provider token
      GoogleSignInAuthentication auth = await account.authentication;

      provider.accountId = account.id;
      provider.token = auth.accessToken;

      return await signInAccount(provider);
    });
  }

  // Generic sign out handles whichever provider initially used
  Future<void> handleSignOut() async {
    // Logger
    final log = Logger('Account - handleSignOut');

    return Future(() async {
      log.info('Signing out');

      if (accountProviderType == AccountProviderType.google) {
        GoogleSignInAccount account = await _googleSignIn.disconnect();
        log.info('Signed out of Google account ${account.toString()}');
      }

      return await signOutAccount();
    });
  }

  Future<void> signInAccount(AccountProvider provider) async {
    // Logger
    final log = Logger('Account - signInAccount');

    return Future(() async {
      log.info('Signing into account');

      log.info('Signed in provider ${provider.type.toString()}');
      log.info('Signed in providerAccountId ${provider.accountId}');
      log.info('Signed in providerToken ${provider.token}');

      // Auth data
      Map<String, dynamic> data = {
        "data": {
          "provider": provider.type.toString().substring(provider.type.toString().indexOf('.') + 1),
          "provider_account_id": provider.accountId,
          "provider_token": provider.token,
        },
      };

      // Auth post
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

  Future<void> signOutAccount() async {
    // Logger
    final log = Logger('Account - signOutAccount');

    return Future(() {
      log.info('Signing out of account');

      this.id = null;
      this.email = null;
      this.name = null;

      log.info('Signed out');
    });
  }
}
