import 'package:flutter/material.dart';
import 'package:logging/logging.dart';

// Application packages
import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/widgets/familliar_animated.dart';

class FamilliarCardBasic extends StatelessWidget {
  final Familliar familliar;
  final FamilliarAction familliarAction;

  FamilliarCardBasic({this.familliar, this.familliarAction = FamilliarAction.idle});

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('FamilliarCardBasic - build');

    log.info("Building - with familliar action >$familliarAction<");

    double calculateFillWidth(BuildContext context, int attributeValue) {
      // Logger
      final log = Logger('ChooseFamilliarListWidget - calculateFillWidth');

      double parentWidth = MediaQuery.of(context).size.width;

      log.finer('Parent width         $parentWidth');
      log.finer('Attribute value      $attributeValue');

      int attributePercentage = ((attributeValue / 20) * 100).toInt();
      log.finer('Attribute percentage $attributePercentage');

      double childWidth = ((attributePercentage / 130) * parentWidth);
      log.finer('Child width          $childWidth');

      return childWidth;
    }

    Color calculateFillColour(BuildContext context, int attributeValue) {
      // Logger
      final log = Logger('ChooseFamilliarListWidget - calculateFillWidth');

      double parentWidth = MediaQuery.of(context).size.width;

      log.finer('Parent width         $parentWidth');
      log.finer('Attribute value      $attributeValue');

      int attributePercentage = ((attributeValue / 20) * 100).toInt();
      log.finer('Attribute percentage $attributePercentage');

      int shadeOffset = ((attributePercentage / 100) * 255).toInt();
      log.finer('Shade offset         $shadeOffset');

      return Color.fromARGB(
          255, (200 - shadeOffset / 2).toInt(), (200 - shadeOffset / 4).toInt(), 255 - shadeOffset);
    }

    return Container(
      child: Column(children: <Widget>[
        // Name
        Container(
          padding: EdgeInsets.fromLTRB(10, 10, 10, 10),
          child: Text('${familliar.name}'),
        ),
        // Avatar
        Container(
          padding: EdgeInsets.fromLTRB(10, 10, 10, 10),
          height: 150,
          child: FamilliarAnimatedWidget(
            familliarAvatar: familliar.avatar,
            familliarAction: familliarAction,
            imageCount: 11,
          ),
        ),
        // Strength
        Container(
          child: Row(
            children: <Widget>[
              Expanded(
                flex: 6,
                child: Container(
                  alignment: Alignment.centerLeft,
                  margin: EdgeInsets.fromLTRB(10, 0, 10, 2),
                  child: Container(
                    padding: EdgeInsets.all(3),
                    width: calculateFillWidth(context, familliar.strength),
                    color: calculateFillColour(context, familliar.strength),
                    child: Text('Strength'),
                  ),
                ),
              ),
              Expanded(
                flex: 1,
                child: Container(
                  padding: EdgeInsets.fromLTRB(10, 0, 10, 0),
                  child: Text('${familliar.strength}'),
                ),
              ),
            ],
          ),
        ),
        // Dexterity
        Container(
          child: Row(
            children: <Widget>[
              Expanded(
                flex: 6,
                child: Container(
                  alignment: Alignment.centerLeft,
                  margin: EdgeInsets.fromLTRB(10, 0, 10, 2),
                  child: Container(
                    padding: EdgeInsets.all(3),
                    width: calculateFillWidth(context, familliar.dexterity),
                    color: calculateFillColour(context, familliar.dexterity),
                    child: Text('Dexterity'),
                  ),
                ),
              ),
              Expanded(
                flex: 1,
                child: Container(
                  padding: EdgeInsets.fromLTRB(10, 0, 10, 0),
                  child: Text('${familliar.dexterity}'),
                ),
              ),
            ],
          ),
        ),
        // Intelligence
        Container(
          child: Row(
            children: <Widget>[
              Expanded(
                flex: 6,
                child: Container(
                  alignment: Alignment.centerLeft,
                  margin: EdgeInsets.fromLTRB(10, 0, 10, 2),
                  child: Container(
                    padding: EdgeInsets.all(3),
                    width: calculateFillWidth(context, familliar.intelligence),
                    color: calculateFillColour(context, familliar.intelligence),
                    child: Text('Intelligence'),
                  ),
                ),
              ),
              Expanded(
                flex: 1,
                child: Container(
                  padding: EdgeInsets.fromLTRB(10, 0, 10, 0),
                  child: Text('${familliar.intelligence}'),
                ),
              ),
            ],
          ),
        ),
      ]),
    );
  }
}
