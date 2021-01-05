import 'package:flutter/material.dart';
import 'package:logging/logging.dart';
import 'package:carousel_slider/carousel_slider.dart';
// import 'package:provider/provider.dart';

// Application packages
// import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/widgets/mage_animated.dart';

class CharacterChoice {
  String description;
  int strength;
  int dexterity;
  int intelligence;
  String imagePath;
  int imageCount;
  ImageProvider avatarImage;

  CharacterChoice({
    this.description,
    this.strength,
    this.dexterity,
    this.intelligence,
    this.imagePath,
    this.imageCount,
    this.avatarImage,
  });
}

List<CharacterChoice> characterChoices = [
  CharacterChoice(
    description: 'Dark Armoured',
    strength: 18,
    dexterity: 10,
    intelligence: 10,
    imagePath: 'assets/images/dark-knight/idle/Idle_',
    imageCount: 11,
    avatarImage: AssetImage("assets/avatars/1.jpg"),
  ),
  CharacterChoice(
    description: 'Druid',
    strength: 14,
    dexterity: 14,
    intelligence: 10,
    imagePath: 'assets/images/druid/idle/Idle_',
    imageCount: 11,
    avatarImage: AssetImage("assets/avatars/2.jpg"),
  ),
  CharacterChoice(
    description: 'Fairy',
    strength: 10,
    dexterity: 14,
    intelligence: 14,
    imagePath: 'assets/images/fairy/idle/Idle_',
    imageCount: 11,
    avatarImage: AssetImage("assets/avatars/3.jpg"),
  ),
  CharacterChoice(
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

    log.warning('Parent width         $parentWidth');
    log.warning('Attribute value      $attributeValue');

    int attributePercentage = ((attributeValue / 20) * 100).toInt();
    log.warning('Attribute percentage $attributePercentage');

    double childWidth = ((attributePercentage / 130) * parentWidth);
    log.warning('Child width          $childWidth');

    return childWidth;
  }

  // Calculate attribute background fill color
  Color calculateFillColour(BuildContext context, int attributeValue) {
    // Logger
    final log = Logger('MageChooseCharacterListWidget - calculateFillWidth');

    double parentWidth = MediaQuery.of(context).size.width;

    log.warning('Parent width         $parentWidth');
    log.warning('Attribute value      $attributeValue');

    int attributePercentage = ((attributeValue / 20) * 100).toInt();
    log.warning('Attribute percentage $attributePercentage');

    int shadeOffset = ((attributePercentage / 100) * 255).toInt();
    log.warning('Shade offset         $shadeOffset');

    return Color.fromARGB(
        255, (200 - shadeOffset / 2).toInt(), (200 - shadeOffset / 4).toInt(), 255 - shadeOffset);
  }

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageChooseCharacterListWidget - build');

    log.info("Building");

    // Build mage
    Widget buildMageCard(int idx) {
      return Container(
        // width: 450,
        color: Colors.grey[400],
        child: Column(children: <Widget>[
          // Description
          Container(
            padding: EdgeInsets.fromLTRB(10, 10, 10, 10),
            child: Text('${characterChoices[idx].description}'),
          ),
          // Avatar
          Container(
            padding: EdgeInsets.fromLTRB(10, 10, 10, 10),
            height: 150,
            child: MageAnimatedWidget(
              imageCount: characterChoices[idx].imageCount,
              imagePath: characterChoices[idx].imagePath,
            ),
            // child: CircleAvatar(
            //   maxRadius: 60.0,
            //   backgroundImage: characterChoices[idx].avatarImage,
            // ),
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
                      width: calculateFillWidth(context, characterChoices[idx].strength),
                      color: calculateFillColour(context, characterChoices[idx].strength),
                      child: Text('Strength'),
                    ),
                  ),
                ),
                Expanded(
                  flex: 1,
                  child: Container(
                    padding: EdgeInsets.fromLTRB(10, 0, 10, 0),
                    child: Text('${characterChoices[idx].strength}'),
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
                      width: calculateFillWidth(context, characterChoices[idx].dexterity),
                      color: calculateFillColour(context, characterChoices[idx].dexterity),
                      child: Text('Dexterity'),
                    ),
                  ),
                ),
                Expanded(
                  flex: 1,
                  child: Container(
                    padding: EdgeInsets.fromLTRB(10, 0, 10, 0),
                    child: Text('${characterChoices[idx].dexterity}'),
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
                      width: calculateFillWidth(context, characterChoices[idx].intelligence),
                      color: calculateFillColour(context, characterChoices[idx].intelligence),
                      child: Text('Intelligence'),
                    ),
                  ),
                ),
                Expanded(
                  flex: 1,
                  child: Container(
                    padding: EdgeInsets.fromLTRB(10, 0, 10, 0),
                    child: Text('${characterChoices[idx].intelligence}'),
                  ),
                ),
              ],
            ),
          ),
        ]),
      );
    }

    return CarouselSlider.builder(
      itemCount: characterChoices.length,
      options: CarouselOptions(
        height: 400,
        aspectRatio: 16 / 9,
        viewportFraction: 0.8,
        initialPage: 0,
        enableInfiniteScroll: true,
        enlargeCenterPage: true,
        scrollDirection: Axis.horizontal,
      ),
      itemBuilder: (BuildContext context, int idx) => Container(
        child: buildMageCard(idx),
      ),
    );
  }
}
