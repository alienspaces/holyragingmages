import 'package:flutter/material.dart';
import 'package:logging/logging.dart';

// Application packages

class ProcessingScreen extends StatelessWidget {
  ProcessingScreen({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('ProcessingScreen - build');

    log.info("Building");

    return Scaffold(
      body: Container(
        alignment: Alignment.center,
        color: Colors.black,
        child: CircularProgressIndicator(),
      ),
    );
  }
}
