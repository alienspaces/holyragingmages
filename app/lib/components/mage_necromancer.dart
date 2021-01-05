import 'package:flame/sprite.dart';

// Application packages
import 'package:holyragingmages/screens/processing.dart';
import 'package:holyragingmages/components/mage.dart';

class MageNecromancer extends Mage {
  MageNecromancer({ProcessingGame game, double startX, startY, endX, endY})
      : super(game: game, startX: startX, startY: startY, endX: endX, endY: endY) {
    mageSprite = Sprite('necromancer/casting/Casting_000.png');
    flipHorizontally = true;
  }
}
