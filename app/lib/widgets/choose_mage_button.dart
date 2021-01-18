import 'package:flutter/material.dart';
import 'package:logging/logging.dart';
// import 'package:provider/provider.dart';

// Application packages
// import 'package:holyragingmages/models/models.dart';

class ChooseMageButtonWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('ChooseMageButtonWidget - build');

    log.info("Building");

    return Container(
      child: FlatButton(
        onPressed: null,
        child: Text('Next'),
      ),
    );
  }
}
