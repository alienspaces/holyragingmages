import 'package:flutter/material.dart';
import 'package:logging/logging.dart';
import 'package:carousel_slider/carousel_slider.dart';

// Application packages
import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/widgets/mage_animated.dart';

class MageChooseCharacterListWidget extends StatelessWidget {
  final List<Mage> starterMageList;

  MageChooseCharacterListWidget({Key key, this.starterMageList}) : super(key: key);

  double calculateFillWidth(BuildContext context, int attributeValue) {
    // Logger
    final log = Logger('MageChooseCharacterListWidget - calculateFillWidth');

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
    final log = Logger('MageChooseCharacterListWidget - calculateFillWidth');

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

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageChooseCharacterListWidget - build');

    log.info("Building");

    void onPageChangedHandler(int pageIdx, CarouselPageChangedReason reason) {
      // Logger
      final log = Logger('MageChooseCharacterListWidget - onScrolledHandler');

      log.info("Page idx $pageIdx reason $reason");
    }

    // Build mage
    Widget buildMageCard(int idx) {
      return Container(
        // width: 450,
        color: Colors.grey[400],
        child: Column(children: <Widget>[
          // Description
          Container(
            padding: EdgeInsets.fromLTRB(10, 10, 10, 10),
            child: Text('${starterMageList[idx].name}'),
          ),
          // Avatar
          Container(
            padding: EdgeInsets.fromLTRB(10, 10, 10, 10),
            height: 150,
            child: MageAnimatedWidget(
              mageAvatar: starterMageList[idx].avatar,
              mageAction: 'idle',
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
                      width: calculateFillWidth(context, starterMageList[idx].strength),
                      color: calculateFillColour(context, starterMageList[idx].strength),
                      child: Text('Strength'),
                    ),
                  ),
                ),
                Expanded(
                  flex: 1,
                  child: Container(
                    padding: EdgeInsets.fromLTRB(10, 0, 10, 0),
                    child: Text('${starterMageList[idx].strength}'),
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
                      width: calculateFillWidth(context, starterMageList[idx].dexterity),
                      color: calculateFillColour(context, starterMageList[idx].dexterity),
                      child: Text('Dexterity'),
                    ),
                  ),
                ),
                Expanded(
                  flex: 1,
                  child: Container(
                    padding: EdgeInsets.fromLTRB(10, 0, 10, 0),
                    child: Text('${starterMageList[idx].dexterity}'),
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
                      width: calculateFillWidth(context, starterMageList[idx].intelligence),
                      color: calculateFillColour(context, starterMageList[idx].intelligence),
                      child: Text('Intelligence'),
                    ),
                  ),
                ),
                Expanded(
                  flex: 1,
                  child: Container(
                    padding: EdgeInsets.fromLTRB(10, 0, 10, 0),
                    child: Text('${starterMageList[idx].intelligence}'),
                  ),
                ),
              ],
            ),
          ),
        ]),
      );
    }

    return CarouselSlider.builder(
      itemCount: starterMageList.length,
      options: CarouselOptions(
        height: 400,
        aspectRatio: 16 / 9,
        viewportFraction: 0.8,
        initialPage: 0,
        enableInfiniteScroll: true,
        enlargeCenterPage: true,
        scrollDirection: Axis.horizontal,
        onPageChanged: onPageChangedHandler,
      ),
      itemBuilder: (BuildContext context, int idx) => Container(
        child: buildMageCard(idx),
      ),
    );
  }
}
