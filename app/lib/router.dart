import 'package:flutter/material.dart';
import 'package:logging/logging.dart';

// Application packages
import 'package:holyragingmages/screens/screens.dart';

class RouteGenerator {
  static Route<dynamic> generateRoute(RouteSettings settings) {
    // Logger
    final log = Logger('RouteGenerator - generateRoute');

    // Arguments
    final args = settings.arguments;

    log.info('Routing with args $args');

    switch (settings.name) {
      case '/':
        return MaterialPageRoute(builder: (context) => MageListScreen(), maintainState: false);
      case '/mage_create':
        return MaterialPageRoute(builder: (context) => MageCreateScreen());
      case '/mage_play':
        return MaterialPageRoute(builder: (context) => MagePlayScreen());
      default:
        return _errorRoute();
    }
  }

  static Route<dynamic> _errorRoute() {
    return MaterialPageRoute(builder: (context) {
      return Scaffold(
        appBar: AppBar(
          title: Text('Error'),
        ),
        body: Center(
          child: Text('Page not found'),
        ),
      );
    });
  }
}
