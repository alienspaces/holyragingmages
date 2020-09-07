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

  @override
  void initState() {
    super.initState();
    strength = 10;
    dexterity = 10;
    intelligence = 10;
  }

  void _incrementStrength() {
    setState(() {
      strength++;
    });
  }

  void _decrementStrength() {
    setState(() {
      strength--;
    });
  }

  void _incrementDexterity() {
    setState(() {
      dexterity++;
    });
  }

  void _decrementDexterity() {
    setState(() {
      dexterity--;
    });
  }

  void _incrementIntelligence() {
    setState(() {
      intelligence++;
    });
  }

  void _decrementIntelligence() {
    setState(() {
      intelligence--;
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
        // TODO: Attribute points remaining
        // Strength
        MageCreateAttribute(
          name: 'Strength',
          value: this.strength,
          incrementValue: _incrementStrength,
          decrementValue: _decrementStrength,
        ),
        // Dexterity
        MageCreateAttribute(
          name: 'Dexterity',
          value: this.dexterity,
          incrementValue: _incrementDexterity,
          decrementValue: _decrementDexterity,
        ),
        // Intelligence
        MageCreateAttribute(
          name: 'Intelligence',
          value: this.intelligence,
          incrementValue: _incrementIntelligence,
          decrementValue: _decrementIntelligence,
        ),
      ],
    );
  }
}
