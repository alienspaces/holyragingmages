import 'package:flutter/material.dart';
import 'package:logging/logging.dart';
import 'package:provider/provider.dart';

// Application packages
import 'package:holyragingmages/api/api.dart';
import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/screens/screens.dart';

class RouteGenerator {
  final Api api;

  RouteGenerator({Key key, this.api});

  Route<dynamic> generateRoute(RouteSettings settings) {
    // Logger
    final log = Logger('RouteGenerator - generateRoute');

    // Arguments
    final args = settings.arguments;

    log.info('Routing args >$args<');
    log.info('Routing path >${settings.name}<');

    Widget resolveRoute(BuildContext context, RouteSettings settings) {
      // Account model
      var accountModel = Provider.of<Account>(context, listen: false);
      if (accountModel.id == null) {
        log.info('Account ID is null, returning login screen');
        return AccountLoginScreen(api: api);
      }

      switch (settings.name) {
        case '/':
          log.fine('Returning account login screen');
          return AccountLoginScreen(api: api);
        case '/mage_list':
          log.fine('Returning mage list screen');
          return MageListScreen(api: api);
        case '/mage_create':
          log.fine('Returning mage create screen');
          return MageCreateScreen(api: api);
        case '/mage_play':
          log.fine('Returning mage play screen');
          return MagePlayScreen(api: api);
        case '/mage_train':
          log.fine('Returning mage train screen');
          return MageTrainScreen(api: api);
        default:
          return Scaffold(
            body: Center(
              child: CircularProgressIndicator(),
            ),
          );
      }
    }

    return MaterialPageRoute(
      builder: (context) {
        return resolveRoute(context, settings);
      },
    );
  }
}
