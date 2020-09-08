import 'package:flutter/material.dart';
import 'package:logging/logging.dart';

import 'mage_create_attribute.dart';

class MageCreateWidget extends StatefulWidget {
  @override
  MageCreateWidgetState createState() {
    return MageCreateWidgetState();
  }
}

class MageCreateWidgetState extends State<MageCreateWidget> {
  final _mageNameController = TextEditingController();
  int strength;
  int dexterity;
  int intelligence;
  int points;

  @override
  void initState() {
    super.initState();
    strength = 10;
    dexterity = 10;
    intelligence = 10;
    points = 10;
  }

  void _incrementStrength() {
    setState(() {
      if (points > 0) {
        strength++;
        points--;
      }
    });
  }

  void _decrementStrength() {
    setState(() {
      if (strength > 10) {
        strength--;
        points++;
      }
    });
  }

  void _incrementDexterity() {
    setState(() {
      if (points > 0) {
        dexterity++;
        points--;
      }
    });
  }

  void _decrementDexterity() {
    setState(() {
      if (dexterity > 10) {
        dexterity--;
        points++;
      }
    });
  }

  void _incrementIntelligence() {
    setState(() {
      if (points > 0) {
        intelligence++;
        points--;
      }
    });
  }

  void _decrementIntelligence() {
    setState(() {
      if (intelligence > 10) {
        intelligence--;
        points++;
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageCreateWidget - build');

    log.info("Building");

    return Column(
      children: <Widget>[
        CircleAvatar(
          maxRadius: 60.0,
          backgroundImage: AssetImage("assets/avatars/2.jpg"),
        ),
        TextField(
          decoration: InputDecoration(
            filled: true,
            labelText: "Name",
          ),
          controller: _mageNameController,
        ),
        Row(children: <Widget>[
          Expanded(
            flex: 5,
            child: Text('Points Remaining'),
          ),
          Expanded(
            flex: 5,
            child: Text(this.points.toString()),
          ),
        ]),
        // Strength
        MageCreateAttribute(
          name: 'Strength',
          value: this.strength,
          incrementValue: _incrementStrength,
          decrementValue: _decrementStrength,
          incrementEnabled: this.points > 0,
          decrementEnabled: this.strength > 10,
        ),
        // Dexterity
        MageCreateAttribute(
          name: 'Dexterity',
          value: this.dexterity,
          incrementValue: _incrementDexterity,
          decrementValue: _decrementDexterity,
          incrementEnabled: this.points > 0,
          decrementEnabled: this.dexterity > 10,
        ),
        // Intelligence
        MageCreateAttribute(
          name: 'Intelligence',
          value: this.intelligence,
          incrementValue: _incrementIntelligence,
          decrementValue: _decrementIntelligence,
          incrementEnabled: this.points > 0,
          decrementEnabled: this.intelligence > 10,
        ),
        // TODO: Save button and save new mage
      ],
    );
  }
}
