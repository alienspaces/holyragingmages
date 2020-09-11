import 'package:flutter/material.dart';
import 'package:logging/logging.dart';
import 'package:provider/provider.dart';

import '../models/models.dart';

import 'mage_create_attribute.dart';

class MageCreateWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageCreateWidget - build');

    log.info("Building");

    // Mage models
    var mageModel = Provider.of<MageModel>(context);

    // TODO: Add saving mage and adding to mage list
    // var mageListModel = Provider.of<MageListModel>(context);

    final _mageNameController = TextEditingController();

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
            child: Text(mageModel.points.toString()),
          ),
        ]),
        // Strength
        MageCreateAttribute(
          name: 'Strength',
          value: mageModel.strength,
          incrementValue: _incrementStrength,
          decrementValue: _decrementStrength,
          incrementEnabled: mageModel.points > 0,
          decrementEnabled: mageModel.strength > 10,
        ),
        // Dexterity
        MageCreateAttribute(
          name: 'Dexterity',
          value: mageModel.dexterity,
          incrementValue: _incrementDexterity,
          decrementValue: _decrementDexterity,
          incrementEnabled: mageModel.points > 0,
          decrementEnabled: mageModel.dexterity > 10,
        ),
        // Intelligence
        MageCreateAttribute(
          name: 'Intelligence',
          value: mageModel.intelligence,
          incrementValue: _incrementIntelligence,
          decrementValue: _decrementIntelligence,
          incrementEnabled: mageModel.points > 0,
          decrementEnabled: mageModel.intelligence > 10,
        ),
      ],
    );
  }
}
