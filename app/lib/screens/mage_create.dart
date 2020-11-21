import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:logging/logging.dart';

// Application packages
import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/widgets/mage_create.dart';
import 'package:holyragingmages/widgets/mage_create_button.dart';

class MageCreateScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageCreateScreen - build');

    log.info("Building");

    return WillPopScope(
      onWillPop: () {
        log.info('Popping');
        return Future.value(true);
      },
      child: ChangeNotifierProvider<Mage>(
        create: (context) => Mage(),
        child: Scaffold(
          appBar: AppBar(
            title: Text('Create Mage'),
          ),
          body: Container(
            child: Center(
              child: MageCreateWidget(),
            ),
          ),
          floatingActionButton: MageCreateButtonWidget(),
          resizeToAvoidBottomInset: false,
        ),
      ),
    );
  }
}
