import 'dart:async';
import 'package:flutter/material.dart';
import 'package:logging/logging.dart';
import 'package:carousel_slider/carousel_slider.dart';

// Application packages
import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/widgets/familliar_animated.dart';
import 'package:holyragingmages/widgets/familliar_card_basic.dart';

class ChooseFamilliarListWidget extends StatefulWidget {
  final Function({Familliar familliar}) chooseFamilliarCallback;
  final List<Familliar> starterFamilliarList;

  ChooseFamilliarListWidget({Key key, this.starterFamilliarList, this.chooseFamilliarCallback})
      : super(key: key);

  @override
  _ChooseFamilliarListWidgetState createState() => _ChooseFamilliarListWidgetState();
}

class _ChooseFamilliarListWidgetState extends State<ChooseFamilliarListWidget> {
  FamilliarAction familliarAction = FamilliarAction.idle;
  Map<String, FamilliarAction> familliarActionMap = {};

  @override
  void initState() {
    // Logger
    final log = Logger('ChooseFamilliarListWidget - initState');

    log.info('Initialising');

    for (var familliar in widget.starterFamilliarList) {
      familliarActionMap[familliar.name] = FamilliarAction.idle;
    }

    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('ChooseFamilliarListWidget - build');

    log.info("Building");

    // TODO:
    // - Maintain familliarAction per start familliar so all
    // cards are not animating at once..
    // - Change 'casting' to attacking? Make sure all familliars
    // have multiple sets of images
    // - Provide animation start and end callbacks so
    // a caller widget can update its own state when
    // an animation starts and finishes.

    void onPageChangedHandler(int pageIdx, CarouselPageChangedReason reason) {
      // Logger
      final log = Logger('ChooseFamilliarListWidget - onPageChangedHandler');

      log.info("Page idx $pageIdx reason $reason");
    }

    void chooseFamilliar(int idx) {
      // Change familliar to casting action
      setState(() {
        familliarActionMap[widget.starterFamilliarList[idx].name] = FamilliarAction.attack;
      });

      Timer(Duration(milliseconds: 1300), () {
        setState(() {
          familliarActionMap[widget.starterFamilliarList[idx].name] = FamilliarAction.idle;
        });

        widget.chooseFamilliarCallback(familliar: widget.starterFamilliarList[idx]);
      });
    }

    // Build familliar
    Widget buildFamilliarCard(int idx) {
      log.info(
          'Building familliar card with familliar name >${widget.starterFamilliarList[idx].name}< action >$familliarAction<');
      return FamilliarCardBasic(
        familliar: widget.starterFamilliarList[idx],
        familliarAction: familliarActionMap[widget.starterFamilliarList[idx].name],
      );
    }

    return CarouselSlider.builder(
      itemCount: widget.starterFamilliarList.length,
      options: CarouselOptions(
        height: 400,
        aspectRatio: 16 / 9,
        // viewportFraction: 0.8,
        viewportFraction: .7,
        initialPage: 0,
        enableInfiniteScroll: true,
        enlargeCenterPage: true,
        scrollDirection: Axis.horizontal,
        onPageChanged: onPageChangedHandler,
      ),
      itemBuilder: (BuildContext context, int idx) => Container(
        color: Colors.grey[400],
        child: Column(
          children: <Widget>[
            buildFamilliarCard(idx),
            Expanded(
              child: Container(
                alignment: Alignment.center,
                child: ElevatedButton(
                  onPressed: () => chooseFamilliar(idx),
                  style: ElevatedButton.styleFrom(
                    primary: Colors.yellow[600],
                    onPrimary: Colors.black,
                  ),
                  child: Text('Choose'),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
