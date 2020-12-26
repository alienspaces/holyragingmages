import 'dart:ui';
import 'package:flame/sprite.dart';
import 'package:logging/logging.dart';

// Application packages
import 'package:holyragingmages/screens/processing.dart';

class Mage {
  final ProcessingGame game;
  Sprite mageSprite;
  Rect mageRect;
  final double startX, startY, endX, endY;
  double _currX, _currY;
  bool flipHorizontally = false;
  bool flipVertically = false;

  Mage({this.game, this.startX, this.startY, this.endX, this.endY}) {
    // Logger
    final log = Logger('Mage - constructor');

    _currX = startX;
    _currY = startY;

    log.fine("Constructing currX >$_currX< endX >$endX< currY >$_currY< endY >$endY<");
  }

  void render(Canvas c) {
    c.save();
    mageRect = Rect.fromLTWH(_currX, _currY, game.mageSize, game.mageSize);
    if (flipHorizontally) {
      c.translate(mageRect.center.dx, mageRect.center.dy);
      c.scale(-1, 1);
      c.translate(-mageRect.center.dx, -mageRect.center.dy);
    }
    mageSprite.renderRect(c, mageRect.inflate(1));
    c.restore();
  }

  void update(double t) {
    // Logger
    final log = Logger('Mage - update');

    log.fine("Updating currX >$_currX< endX >${this.endX}<");
    log.fine("Updating currY >$_currY< endY >${this.endY}<");

    if (_currX < endX) {
      _currX++;
    } else if (_currX > endX) {
      _currX--;
    }
    if (_currY < endY) {
      _currY++;
    } else if (_currY > endY) {
      _currY--;
    }
  }
}
