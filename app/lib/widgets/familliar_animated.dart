import 'dart:async';
import 'package:flutter/material.dart';
import 'package:logging/logging.dart';

enum FamilliarAction { idle, attack }

const Map<FamilliarAction, String> actionImageMap = {
  FamilliarAction.idle: 'idle',
  FamilliarAction.attack: 'attack',
};

class FamilliarAnimatedWidget extends StatefulWidget {
  final String familliarAvatar;
  final FamilliarAction familliarAction;
  final int imageCount;

  FamilliarAnimatedWidget({
    Key key,
    this.familliarAvatar,
    this.familliarAction,
    this.imageCount,
  }) : super(key: key);

  @override
  FamilliarAnimatedWidgetState createState() => new FamilliarAnimatedWidgetState();
}

class FamilliarAnimatedWidgetState extends State<FamilliarAnimatedWidget> {
  // Map of mage action to images
  Map<FamilliarAction, List<Image>> actionImageList = {};
  // Current mage action to animate
  FamilliarAction familliarAction;
  // Current index of image to display
  int currentIdx = 0;
  // Timer used to manage animation
  Timer timer;

  @override
  void initState() {
    // Logger
    final log = Logger('FamilliarAnimatedWidget - initState');

    log.info('Initialising - with mage action ${widget.familliarAction}');

    familliarAction = widget.familliarAction;

    loadImages();

    super.initState();
  }

  void loadImages() {
    // Logger
    final log = Logger('FamilliarAnimatedWidget - loadImages');

    log.info('Loading images...');

    // Initialise action image list
    actionImageList = {};

    for (var familliarAction in [FamilliarAction.idle, FamilliarAction.attack]) {
      // Initialise action image list
      actionImageList[familliarAction] = [];

      String imageName = actionImageMap[familliarAction];
      String imagePath = 'assets/images/familliars/${widget.familliarAvatar}/$imageName/$imageName';

      for (int idx = 0; idx <= widget.imageCount; idx++) {
        String assetName = "${imagePath}_${idx.toString().padLeft(3, '0')}.png";
        log.info('Adding image assetName $assetName');
        Key imageKey = Key('${widget.familliarAvatar}-$imageName-$idx');
        Image image = Image(key: imageKey, image: AssetImage(assetName));
        log.info('Added ${image.toString()}');
        actionImageList[familliarAction].add(image);
      }
    }
  }

  void cacheImages() {
    // Pre-cache images
    for (var familliarAction in [FamilliarAction.idle, FamilliarAction.attack]) {
      for (var idx = 0; idx <= widget.imageCount; idx++) {
        precacheImage(actionImageList[familliarAction][idx].image, context);
      }
    }
  }

  @override
  void didChangeDependencies() {
    // Logger
    final log = Logger('FamilliarAnimatedWidget - didChangeDependencies');

    log.info('Caching mage ${actionImageMap[familliarAction]}');

    // Pre-cache images
    cacheImages();

    // Change image periodically
    if (timer == null && mounted) {
      timer = Timer.periodic(Duration(milliseconds: 100), (timer) {
        setState(() {
          currentIdx++;
          if (currentIdx == widget.imageCount) {
            currentIdx = 0;
          }
        });
      });
    }

    super.didChangeDependencies();
  }

  @override
  void dispose() {
    // Cancel timer
    if (timer != null) {
      timer.cancel();
    }

    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('FamilliarAnimatedWidget - build');

    if (familliarAction != widget.familliarAction) {
      log.info("Familliar action changed - ${actionImageMap[widget.familliarAction]}");
      setState(() {
        familliarAction = widget.familliarAction;
      });
    }

    return Container(
      child: actionImageList[familliarAction][currentIdx],
    );
  }
}
