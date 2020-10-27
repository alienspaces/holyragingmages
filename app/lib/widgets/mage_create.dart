import 'package:flutter/material.dart';
import 'package:logging/logging.dart';
import 'package:provider/provider.dart';

import '../models/models.dart';

import 'mage_create_attribute.dart';
import 'mage_create_name.dart';

class MageCreateWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageCreateWidget - build');

    log.info("Building");

    // Mage models
    var mageModel = Provider.of<Mage>(context);

    void _incrementStrength() {
      mageModel.strength++;
    }

    void _decrementStrength() {
      mageModel.strength--;
    }

    void _incrementDexterity() {
      mageModel.dexterity++;
    }

    void _decrementDexterity() {
      mageModel.dexterity--;
    }

    void _incrementIntelligence() {
      mageModel.intelligence++;
    }

    void _decrementIntelligence() {
      mageModel.intelligence--;
    }

    void _updateName(String value) {
      mageModel.name = value;
    }

    return Column(
      children: <Widget>[
        CircleAvatar(
          maxRadius: 60.0,
          backgroundImage: AssetImage("assets/avatars/2.jpg"),
        ),
        MageCreateNameWidget(
          value: mageModel.name,
          updateValue: _updateName,
        ),
        Row(children: <Widget>[
          Expanded(
            flex: 5,
            child: Text('Points Remaining'),
          ),
          Expanded(
            flex: 5,
            child: Text(mageModel.attributePoints.toString()),
          ),
        ]),
        // Strength
        MageCreateAttributeWidget(
          name: 'Strength',
          value: mageModel.strength,
          incrementValue: _incrementStrength,
          decrementValue: _decrementStrength,
          incrementEnabled: mageModel.attributePoints > 0,
          decrementEnabled: mageModel.strength > 10,
        ),
        // Dexterity
        MageCreateAttributeWidget(
          name: 'Dexterity',
          value: mageModel.dexterity,
          incrementValue: _incrementDexterity,
          decrementValue: _decrementDexterity,
          incrementEnabled: mageModel.attributePoints > 0,
          decrementEnabled: mageModel.dexterity > 10,
        ),
        // Intelligence
        MageCreateAttributeWidget(
          name: 'Intelligence',
          value: mageModel.intelligence,
          incrementValue: _incrementIntelligence,
          decrementValue: _decrementIntelligence,
          incrementEnabled: mageModel.attributePoints > 0,
          decrementEnabled: mageModel.intelligence > 10,
        ),
      ],
    );
  }
}
