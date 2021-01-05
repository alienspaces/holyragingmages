import 'dart:ui';
import 'package:flame/flame.dart';
import 'package:flame/game.dart';
import 'package:flutter/material.dart';
import 'package:logging/logging.dart';

// Application packages
import 'package:holyragingmages/components/mage_druid.dart';
import 'package:holyragingmages/components/mage_necromancer.dart';

class ProcessingGame extends Game {
  Size screenSize;
  double mageSize;
  MageDruid leftMage;
  MageNecromancer rightMage;

  ProcessingGame() {
    initialise();
  }

  void initialise() async {
    // Logger
    final log = Logger('ProcessingGame - initialise');

    log.fine("Initialising");

    await Flame.images.loadAll(<String>[
      'druid/casting/Casting_000.png',
      'necromancer/casting/Casting_000.png',
    ]);
    resize(await Flame.util.initialDimensions());
  }

  void resize(Size size) {
    screenSize = size;
    mageSize = screenSize.width / 3;
    super.resize(size);
  }

  void render(Canvas canvas) {
    // Logger
    final log = Logger('ProcessingGame - render');

    log.fine("Rendering");

    // Background
    Rect bgRect = Rect.fromLTWH(0, 0, screenSize.width, screenSize.height);
    Paint bgPaint = Paint();
    bgPaint.color = Color(0xff000000);
    canvas.drawRect(bgRect, bgPaint);

    // Mages
    if (leftMage == null) {
      leftMage = MageDruid(
        game: this,
        startX: 10,
        startY: (screenSize.height / 2) - 50,
        endX: (screenSize.width / 2) - 20,
        endY: (screenSize.height / 2),
      );
    }
    leftMage.render(canvas);

    if (rightMage == null) {
      rightMage = MageNecromancer(
        game: this,
        startX: screenSize.width - (screenSize.width / 3) - 10,
        startY: (screenSize.height / 2) - 50,
        endX: (screenSize.width / 2) + 10,
        endY: (screenSize.height / 2),
      );
    }
    rightMage.render(canvas);
  }

  void update(double t) {
    // Logger
    final log = Logger('ProcessingGame - update');

    log.fine("Updating - $t");

    // Mages
    if (leftMage != null) {
      leftMage.update(t);
    }
    if (rightMage != null) {
      rightMage.update(t);
    }
  }
}

class ProcessingScreen extends StatelessWidget {
  ProcessingScreen({Key key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('ProcessingScreen - build');

    log.fine("Building");

    ProcessingGame game = ProcessingGame();

    return Scaffold(
      body: Container(
        alignment: Alignment.center,
        color: Colors.black,
        child: game.widget,
      ),
    );
  }
}
