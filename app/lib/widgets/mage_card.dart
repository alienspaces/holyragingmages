import 'package:flutter/material.dart';
import 'package:logging/logging.dart';

import '../models/models.dart';

class MageCard extends StatelessWidget {
  final MageModel mage;

  MageCard(this.mage);

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
              child: CircleAvatar(
                maxRadius: 50.0,
                backgroundImage: AssetImage("assets/avatars/2.jpg"),
              ),
            ),
            // Familliar avatar
            Container(
              alignment: Alignment.bottomRight,
              padding: EdgeInsets.only(top: 60),
              child: CircleAvatar(
                maxRadius: 20.0,
                backgroundImage: AssetImage("assets/avatars/2.jpg"),
              ),
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
        // Mage attributes
        Row(
          children: <Widget>[
            Expanded(
              flex: 7,
              child: Text("Strength"),
            ),
            Expanded(
              flex: 3,
              child: Align(
                alignment: Alignment.centerRight,
                child: Text(this.mage.strength.toString()),
              ),
            ),
          ],
        ),
        Row(
          children: <Widget>[
            Expanded(
              flex: 7,
              child: Text("Dexterity"),
            ),
            Expanded(
              flex: 3,
              child: Align(
                alignment: Alignment.centerRight,
                child: Text(this.mage.dexterity.toString()),
              ),
            ),
          ],
        ),
        Row(
          children: <Widget>[
            Expanded(
              flex: 7,
              child: Text("Intelligence"),
            ),
            Expanded(
              flex: 3,
              child: Align(
                alignment: Alignment.centerRight,
                child: Text(this.mage.intelligence.toString()),
              ),
            ),
          ],
        ),
        // Mage coins
        Row(
          children: <Widget>[
            Expanded(
              flex: 7,
              child: Text("Coins"),
            ),
            Expanded(
              flex: 3,
              child: Align(
                alignment: Alignment.centerRight,
                child: Text(this.mage.coins.toString()),
              ),
            ),
          ],
        ),
        // Mage experience
        Row(
          children: <Widget>[
            Expanded(
              flex: 7,
              child: Text("Experience"),
            ),
            Expanded(
              flex: 3,
              child: Align(
                alignment: Alignment.centerRight,
                child: Text(this.mage.experiencePoints.toString()),
              ),
            ),
          ],
        ),
        Container(
          alignment: Alignment.center,
          child: FlatButton(
            onPressed: () {
              Navigator.pushNamed(context, '/mage_play');
            },
            color: Colors.green,
            disabledColor: Colors.grey,
            child: Text("Play"),
          ),
        ),
      ],
    );
  }
}
