import 'package:flutter/material.dart';
import 'package:logging/logging.dart';

// Application packages
import 'package:holyragingmages/api/api.dart';

class AccountProcessingScreen extends StatelessWidget {
  final Api api;

  AccountProcessingScreen({Key key, this.api}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('AccountProcessingScreen - build');

    log.info("Building");

    return Scaffold(
      body: Container(
        alignment: Alignment.center,
        child: CircularProgressIndicator(),
      ),
    );
  }
}
