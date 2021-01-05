import 'package:flame/sprite.dart';

// Application packages
import 'package:holyragingmages/screens/processing.dart';
import 'package:holyragingmages/components/mage.dart';

class MageDruid extends Mage {
  MageDruid({ProcessingGame game, double startX, startY, endX, endY})
      : super(game: game, startX: startX, startY: startY, endX: endX, endY: endY) {
    mageSprite = Sprite('druid/casting/Casting_000.png');
  }
}
