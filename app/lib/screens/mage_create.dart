import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:logging/logging.dart';

import '../models/models.dart';
import '../widgets/mage_create.dart';
import '../widgets/mage_create_button.dart';

class MageCreateScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageCreateScreen - build');

    log.info("Building");

    return ChangeNotifierProvider<Mage>(
      create: (context) => Mage(),
      child: Scaffold(
        appBar: AppBar(
          title: Text('Mage Create'),
        ),
        body: Container(
          child: Center(
            child: MageCreateWidget(),
          ),
        ),
        floatingActionButton: MageCreateButtonWidget(),
        resizeToAvoidBottomInset: false,
      ),
    );
  }
}
