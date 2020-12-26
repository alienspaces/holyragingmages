import 'package:flutter/material.dart';
import 'package:logging/logging.dart';

// Application packages

class MageCardNew extends StatelessWidget {
  MageCardNew();

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageCardNew - build');

    log.info("Building");

    return InkWell(
      onTap: () {
        Navigator.pushNamed(context, '/mage_create');
      },
      child: Container(
        alignment: Alignment.center,
        constraints: BoxConstraints.expand(),
        child: Text('Create'),
      ),
    );
  }
}
