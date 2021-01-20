import 'dart:async';
import 'package:flutter/material.dart';
import 'package:logging/logging.dart';
import 'package:carousel_slider/carousel_slider.dart';

// Application packages
import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/widgets/mage_animated.dart';
import 'package:holyragingmages/widgets/mage_card_basic.dart';

class ChooseMageListWidget extends StatefulWidget {
  final Function({Mage mage}) chooseMageCallback;
  final List<Mage> starterMageList;

  ChooseMageListWidget({Key key, this.starterMageList, this.chooseMageCallback}) : super(key: key);

  @override
  _ChooseMageListWidgetState createState() => _ChooseMageListWidgetState();
}

class _ChooseMageListWidgetState extends State<ChooseMageListWidget> {
  MageAction mageAction = MageAction.idle;
  Map<String, MageAction> mageActionMap = {};

  @override
  void initState() {
    // Logger
    final log = Logger('ChooseMageListWidget - initState');

    log.info('Initialising');

    for (var mage in widget.starterMageList) {
      mageActionMap[mage.name] = MageAction.idle;
    }

    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('ChooseMageListWidget - build');

    log.info("Building");

    // TODO:
    // - Maintain mageAction per start mage so all
    // cards are not animating at once..
    // - Change 'casting' to attacking? Make sure all mages
    // have multiple sets of images
    // - Provide animation start and end callbacks so
    // a caller widget can update its own state when
    // an animation starts and finishes.

    void onPageChangedHandler(int pageIdx, CarouselPageChangedReason reason) {
      // Logger
      final log = Logger('ChooseMageListWidget - onPageChangedHandler');

      log.info("Page idx $pageIdx reason $reason");
    }

    void chooseMage(int idx) {
      // Change mage to casting action
      setState(() {
        mageActionMap[widget.starterMageList[idx].name] = MageAction.casting;
      });

      Timer(Duration(seconds: 3), () {
        setState(() {
          mageActionMap[widget.starterMageList[idx].name] = MageAction.idle;
        });

        widget.chooseMageCallback(mage: widget.starterMageList[idx]);
      });
    }

    // Build mage
    Widget buildMageCard(int idx) {
      log.info(
          'Building mage card with mage name >${widget.starterMageList[idx].name}< action >$mageAction<');
      return MageCardBasic(
        mage: widget.starterMageList[idx],
        mageAction: mageActionMap[widget.starterMageList[idx].name],
      );
    }

    return CarouselSlider.builder(
      itemCount: widget.starterMageList.length,
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
        color: Colors.grey[400],
        child: Column(
          children: <Widget>[
            buildMageCard(idx),
            Expanded(
              child: Container(
                alignment: Alignment.center,
                child: ElevatedButton(
                  onPressed: () => chooseMage(idx),
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
