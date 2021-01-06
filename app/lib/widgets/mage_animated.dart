import 'dart:async';
import 'package:flutter/material.dart';
import 'package:logging/logging.dart';

class MageAnimatedWidget extends StatefulWidget {
  final String imagePath;
  final int imageCount;

  MageAnimatedWidget({
    Key key,
    this.imagePath,
    this.imageCount,
  }) : super(key: key);

  @override
  MageAnimatedWidgetState createState() => new MageAnimatedWidgetState();
}

class MageAnimatedWidgetState extends State<MageAnimatedWidget> {
  List<Image> imageList = [];
  int currentIdx = 0;
  Timer timer;

  @override
  void initState() {
    // Logger
    final log = Logger('MageAnimatedWidget - initState');

    if (imageList.length == 0) {
      for (int idx = 0; idx <= widget.imageCount; idx++) {
        String imagePath = "${widget.imagePath}${idx.toString().padLeft(3, '0')}.png";
        log.finer('Adding image path $imagePath');
        Image image = Image(image: AssetImage(imagePath));
        log.finer('Added ${image.toString()}');
        imageList.add(image);
      }
    }

    super.initState();
  }

  @override
  void didChangeDependencies() {
    // Pre-cache images
    for (var idx = 0; idx <= widget.imageCount; idx++) {
      precacheImage(imageList[idx].image, context);
    }

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

    log.finer("Building");

    return Container(
      child: imageList[currentIdx],
    );
  }
}
