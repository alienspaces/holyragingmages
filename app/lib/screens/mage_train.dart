import 'package:flutter/material.dart';
import 'package:logging/logging.dart';

// Application packages
import 'package:holyragingmages/api/api.dart';

class MageTrainScreen extends StatelessWidget {
  final Api api;

  MageTrainScreen({Key key, this.api}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageCreateScreen - build');

    log.info("Building");

    return Scaffold(
      appBar: AppBar(
        title: Text('Create Mage'),
      ),
      body: Container(),
    );
  }
}
