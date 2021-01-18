import 'package:flutter/material.dart';
import 'package:logging/logging.dart';
import 'package:carousel_slider/carousel_slider.dart';

// Application packages
import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/widgets/familliar_animated.dart';

class ChooseFamilliarListWidget extends StatelessWidget {
  final Function chooseFamilliarCallback;
  final List<Mage> starterFamilliarList;

  ChooseFamilliarListWidget({Key key, this.starterFamilliarList, this.chooseFamilliarCallback})
      : super(key: key);

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

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('ChooseFamilliarListWidget - build');

    log.info("Building");

    void onPageChangedHandler(int pageIdx, CarouselPageChangedReason reason) {
      // Logger
      final log = Logger('ChooseFamilliarListWidget - onScrolledHandler');

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
            child: Text('${starterFamilliarList[idx].name}'),
          ),
          // Avatar
          Container(
            padding: EdgeInsets.fromLTRB(10, 10, 10, 10),
            height: 150,
            child: FamilliarAnimatedWidget(
              familliarAvatar: starterFamilliarList[idx].avatar,
              familliarAction: 'idle',
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
                      width: calculateFillWidth(context, starterFamilliarList[idx].strength),
                      color: calculateFillColour(context, starterFamilliarList[idx].strength),
                      child: Text('Strength'),
                    ),
                  ),
                ),
                Expanded(
                  flex: 1,
                  child: Container(
                    padding: EdgeInsets.fromLTRB(10, 0, 10, 0),
                    child: Text('${starterFamilliarList[idx].strength}'),
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
                      width: calculateFillWidth(context, starterFamilliarList[idx].dexterity),
                      color: calculateFillColour(context, starterFamilliarList[idx].dexterity),
                      child: Text('Dexterity'),
                    ),
                  ),
                ),
                Expanded(
                  flex: 1,
                  child: Container(
                    padding: EdgeInsets.fromLTRB(10, 0, 10, 0),
                    child: Text('${starterFamilliarList[idx].dexterity}'),
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
                      width: calculateFillWidth(context, starterFamilliarList[idx].intelligence),
                      color: calculateFillColour(context, starterFamilliarList[idx].intelligence),
                      child: Text('Intelligence'),
                    ),
                  ),
                ),
                Expanded(
                  flex: 1,
                  child: Container(
                    padding: EdgeInsets.fromLTRB(10, 0, 10, 0),
                    child: Text('${starterFamilliarList[idx].intelligence}'),
                  ),
                ),
              ],
            ),
          ),
        ]),
      );
    }

    return CarouselSlider.builder(
      itemCount: starterFamilliarList.length,
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
