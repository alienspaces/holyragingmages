import 'dart:async';
import 'package:flutter/material.dart';
import 'package:logging/logging.dart';

enum MageAction { idle, casting }

const Map<MageAction, String> actionImageMap = {
  MageAction.idle: 'idle',
  MageAction.casting: 'casting',
};

class MageAnimatedWidget extends StatefulWidget {
  final String mageAvatar;
  final MageAction mageAction;
  final int imageCount;

  MageAnimatedWidget({
    Key key,
    this.mageAvatar,
    this.mageAction,
    this.imageCount,
  }) : super(key: key);

  @override
  MageAnimatedWidgetState createState() => new MageAnimatedWidgetState();
}

class MageAnimatedWidgetState extends State<MageAnimatedWidget> {
  // Map of mage action to images
  Map<MageAction, List<Image>> actionImageList = {};
  // Current mage action to animate
  MageAction mageAction;
  // Current index of image to display
  int currentIdx = 0;
  // Timer used to manage animation
  Timer timer;

  @override
  void initState() {
    // Logger
    final log = Logger('MageAnimatedWidget - initState');

    log.info('Initialising - with mage action ${widget.mageAction}');

    mageAction = widget.mageAction;

    loadImages();

    super.initState();
  }

  void loadImages() {
    // Logger
    final log = Logger('MageAnimatedWidget - loadImages');

    log.info('Loading images...');

    // Initialise action image list
    actionImageList = {};

    for (var mageAction in [MageAction.idle, MageAction.casting]) {
      // Initialise action image list
      actionImageList[mageAction] = [];

      String imageName = actionImageMap[mageAction];
      String imagePath = 'assets/images/mages/${widget.mageAvatar}/$imageName/$imageName';

      for (int idx = 0; idx <= widget.imageCount; idx++) {
        String assetName = "${imagePath}_${idx.toString().padLeft(3, '0')}.png";
        log.info('Adding image assetName $assetName');
        Image image = Image(image: AssetImage(assetName));
        log.info('Added ${image.toString()}');
        actionImageList[mageAction].add(image);
      }
    }
  }

  void cacheImages() {
    // Pre-cache images
    for (var mageAction in [MageAction.idle, MageAction.casting]) {
      for (var idx = 0; idx <= widget.imageCount; idx++) {
        precacheImage(actionImageList[mageAction][idx].image, context);
      }
    }
  }

  @override
  void didChangeDependencies() {
    // Logger
    final log = Logger('MageAnimatedWidget - didChangeDependencies');

    log.info('Caching mage ${actionImageMap[mageAction]}');

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
    final log = Logger('MageAnimatedWidget - build');

    if (mageAction != widget.mageAction) {
      log.info("Mage action changed - ${actionImageMap[widget.mageAction]}");
      setState(() {
        mageAction = widget.mageAction;
      });
    }

    return Container(
      child: actionImageList[mageAction][currentIdx],
    );
  }
}
