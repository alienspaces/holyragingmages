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
        case '/mage_choose_character':
          log.fine('Returning mage choose character screen');
          return MageChooseCharacterScreen(api: api);
        case '/mage_choose_familliar':
          log.fine('Returning mage choose familliar screen');
          return MageChooseFamilliarScreen(api: api);
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

    Widget resolveTransition(BuildContext context, RouteSettings settings,
        Animation<double> animation, Animation<double> anotherAnimation, Widget child) {
      // Fade transition
      Widget fadeTransition(Widget child) {
        animation = CurvedAnimation(curve: Curves.easeInOut, parent: animation);
        return FadeTransition(
          opacity: animation,
          child: child,
        );
      }

      // Slide transition
      Widget slideTransition(Widget child) {
        animation = CurvedAnimation(curve: Curves.easeInOut, parent: animation);
        return SlideTransition(
          position: Tween(begin: Offset(-1.0, 0.0), end: Offset(0.0, 0.0)).animate(animation),
          child: child,
        );
      }

      switch (settings.name) {
        case '/':
          log.fine('Returning account login screen');
          return fadeTransition(child);
        case '/mage_list':
          log.fine('Returning mage list screen');
          return fadeTransition(child);
        case '/mage_create':
          log.fine('Returning mage create screen');
          return slideTransition(child);
        case '/mage_choose_character':
          log.fine('Returning mage create screen');
          return slideTransition(child);
        case '/mage_choose_familliar':
          log.fine('Returning mage create screen');
          return slideTransition(child);
        case '/mage_play':
          log.fine('Returning mage play screen');
          return slideTransition(child);
        case '/mage_train':
          log.fine('Returning mage train screen');
          return slideTransition(child);
        default:
          return slideTransition(child);
      }
    }

    return PageRouteBuilder(
      pageBuilder: (context, animation, anotherAnimation) {
        return resolveRoute(context, settings);
      },
      transitionDuration: Duration(milliseconds: 500),
      transitionsBuilder: (context, animation, anotherAnimation, child) {
        return resolveTransition(context, settings, animation, anotherAnimation, child);
      },
    );
  }
}
