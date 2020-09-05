import 'package:flutter/cupertino.dart';

import '../models/models.dart';

class MageCard extends StatelessWidget {
  final MageModel mage;

  MageCard(this.mage);

  @override
  Widget build(BuildContext context) {
    return Column(
      children: <Widget>[
        Row(
          children: <Widget>[
            Expanded(
              flex: 5,
              child: Text("Name"),
            ),
            Expanded(
              flex: 5,
              child: Text(this.mage.name),
            ),
          ],
        ),
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
      ],
    );
  }
}
