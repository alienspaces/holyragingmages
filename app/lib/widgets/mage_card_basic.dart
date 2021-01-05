import 'package:flutter/material.dart';
import 'package:logging/logging.dart';

// Application packages
import 'package:holyragingmages/models/models.dart';

class MageCardBasic extends StatelessWidget {
  final Mage mage;

  MageCardBasic(this.mage);

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageListWidget - build');

    log.info("Building");

    return Column(
      mainAxisSize: MainAxisSize.min,
      children: <Widget>[
        Stack(
          children: <Widget>[
            // Mage avatar
            Container(
              alignment: Alignment.topCenter,
              // child: CircleAvatar(
              //   maxRadius: 50.0,
              //   backgroundImage: AssetImage("assets/avatars/2.jpg"),
              // ),
            ),
          ],
        ),
        // Mage name
        Container(
          padding: EdgeInsets.all(3.0),
          child: Align(
            alignment: Alignment.bottomCenter,
            child: Text(this.mage.name),
          ),
        ),
      ],
    );
  }
}
