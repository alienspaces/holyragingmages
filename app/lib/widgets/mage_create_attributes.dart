import 'package:flutter/material.dart';
import 'package:logging/logging.dart';
import 'package:provider/provider.dart';

// Application packages
import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/widgets/mage_create_attribute.dart';

class MageCreateAttributesWidget extends StatefulWidget {
  @override
  _MageCreateAttributesWidgetState createState() => _MageCreateAttributesWidgetState();
}

class _MageCreateAttributesWidgetState extends State<MageCreateAttributesWidget> {
  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageCreateWidget - build');

    log.info("Building");

    // Mage models
    var mageModel = Provider.of<Mage>(context);

    void unfocus() {
      log.warning('Unfocus');
      FocusScopeNode currentFocus = FocusScope.of(context);
      if (!currentFocus.hasPrimaryFocus) {
        log.warning('Unfocus - not focused');
        currentFocus.unfocus();
      }
    }

    void _incrementStrength() {
      unfocus();
      mageModel.strength++;
    }

    void _decrementStrength() {
      unfocus();
      mageModel.strength--;
    }

    void _incrementDexterity() {
      unfocus();
      mageModel.dexterity++;
    }

    void _decrementDexterity() {
      unfocus();
      mageModel.dexterity--;
    }

    void _incrementIntelligence() {
      unfocus();
      mageModel.intelligence++;
    }

    void _decrementIntelligence() {
      unfocus();
      mageModel.intelligence--;
    }

    // Styling
    EdgeInsetsGeometry padding = EdgeInsets.fromLTRB(0, 5, 0, 5);

    return Column(
      children: <Widget>[
        Container(
          padding: padding,
          alignment: Alignment.center,
          child: Text('Points Remaining: ${mageModel.availableAttributePoints}'),
        ),
        // Strength
        Container(
          padding: padding,
          child: MageCreateAttributeWidget(
            name: 'Strength',
            value: mageModel.strength,
            incrementValue: _incrementStrength,
            decrementValue: _decrementStrength,
            incrementEnabled: mageModel.availableAttributePoints > 0,
            decrementEnabled: mageModel.strength > initialAttributeValue,
          ),
        ),
        // Dexterity
        Container(
          padding: padding,
          child: MageCreateAttributeWidget(
            name: 'Dexterity',
            value: mageModel.dexterity,
            incrementValue: _incrementDexterity,
            decrementValue: _decrementDexterity,
            incrementEnabled: mageModel.availableAttributePoints > 0,
            decrementEnabled: mageModel.dexterity > initialAttributeValue,
          ),
        ),
        // Intelligence
        Container(
          padding: padding,
          child: MageCreateAttributeWidget(
            name: 'Intelligence',
            value: mageModel.intelligence,
            incrementValue: _incrementIntelligence,
            decrementValue: _decrementIntelligence,
            incrementEnabled: mageModel.availableAttributePoints > 0,
            decrementEnabled: mageModel.intelligence > initialAttributeValue,
          ),
        ),
      ],
    );
  }
}
