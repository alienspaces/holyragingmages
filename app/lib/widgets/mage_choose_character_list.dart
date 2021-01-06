import 'package:flutter/material.dart';
import 'package:logging/logging.dart';
import 'package:carousel_slider/carousel_slider.dart';
// import 'package:provider/provider.dart';

// Application packages
// import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/widgets/mage_animated.dart';

// TODO: Store starter character on server with items, spells
class MageTemplate {
  String description;
  int strength;
  int dexterity;
  int intelligence;
  String imagePath;
  int imageCount;
  ImageProvider avatarImage;

  MageTemplate({
    this.description,
    this.strength,
    this.dexterity,
    this.intelligence,
    this.imagePath,
    this.imageCount,
    this.avatarImage,
  });
}

List<MageTemplate> mageTemplates = [
  MageTemplate(
    description: 'Dark Armoured',
    strength: 18,
    dexterity: 10,
    intelligence: 10,
    imagePath: 'assets/images/dark-knight/idle/Idle_',
    imageCount: 11,
    avatarImage: AssetImage("assets/avatars/1.jpg"),
  ),
  MageTemplate(
    description: 'Druid',
    strength: 14,
    dexterity: 14,
    intelligence: 10,
    imagePath: 'assets/images/druid/idle/Idle_',
    imageCount: 11,
    avatarImage: AssetImage("assets/avatars/2.jpg"),
  ),
  MageTemplate(
    description: 'Fairy',
    strength: 10,
    dexterity: 14,
    intelligence: 14,
    imagePath: 'assets/images/fairy/idle/Idle_',
    imageCount: 11,
    avatarImage: AssetImage("assets/avatars/3.jpg"),
  ),
  MageTemplate(
    description: 'Necromancer',
    strength: 14,
    dexterity: 10,
    intelligence: 14,
    imagePath: 'assets/images/necromancer/idle/Idle_',
    imageCount: 11,
    avatarImage: AssetImage("assets/avatars/3.jpg"),
  ),
];

class MageChooseCharacterListWidget extends StatelessWidget {
  // Calculate attribute background fill width
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

  // Calculate attribute background fill color
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
            child: Text('${mageTemplates[idx].description}'),
          ),
          // Avatar
          Container(
            padding: EdgeInsets.fromLTRB(10, 10, 10, 10),
            height: 150,
            child: MageAnimatedWidget(
              imageCount: mageTemplates[idx].imageCount,
              imagePath: mageTemplates[idx].imagePath,
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
                      width: calculateFillWidth(context, mageTemplates[idx].strength),
                      color: calculateFillColour(context, mageTemplates[idx].strength),
                      child: Text('Strength'),
                    ),
                  ),
                ),
                Expanded(
                  flex: 1,
                  child: Container(
                    padding: EdgeInsets.fromLTRB(10, 0, 10, 0),
                    child: Text('${mageTemplates[idx].strength}'),
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
                      width: calculateFillWidth(context, mageTemplates[idx].dexterity),
                      color: calculateFillColour(context, mageTemplates[idx].dexterity),
                      child: Text('Dexterity'),
                    ),
                  ),
                ),
                Expanded(
                  flex: 1,
                  child: Container(
                    padding: EdgeInsets.fromLTRB(10, 0, 10, 0),
                    child: Text('${mageTemplates[idx].dexterity}'),
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
                      width: calculateFillWidth(context, mageTemplates[idx].intelligence),
                      color: calculateFillColour(context, mageTemplates[idx].intelligence),
                      child: Text('Intelligence'),
                    ),
                  ),
                ),
                Expanded(
                  flex: 1,
                  child: Container(
                    padding: EdgeInsets.fromLTRB(10, 0, 10, 0),
                    child: Text('${mageTemplates[idx].intelligence}'),
                  ),
                ),
              ],
            ),
          ),
        ]),
      );
    }

    return CarouselSlider.builder(
      itemCount: mageTemplates.length,
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
