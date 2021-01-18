import 'package:flutter/material.dart';
import 'package:logging/logging.dart';
import 'package:carousel_slider/carousel_slider.dart';

// Application packages
import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/widgets/mage_card_basic.dart';

class ChooseMageListWidget extends StatelessWidget {
  final Function({Mage mage}) chooseMageCallback;
  final List<Mage> starterMageList;

  ChooseMageListWidget({Key key, this.starterMageList, this.chooseMageCallback}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('ChooseMageListWidget - build');

    log.info("Building");

    void onPageChangedHandler(int pageIdx, CarouselPageChangedReason reason) {
      // Logger
      final log = Logger('ChooseMageListWidget - onScrolledHandler');

      log.info("Page idx $pageIdx reason $reason");
    }

    // Build mage
    Widget buildMageCard(int idx) {
      return MageCardBasic(mage: starterMageList[idx]);
    }

    return CarouselSlider.builder(
      itemCount: starterMageList.length,
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
            Container(
              child: FlatButton(
                onPressed: () => chooseMageCallback(mage: starterMageList[idx]),
                color: Colors.orange,
                child: Text('Choose'),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
