import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:logging/logging.dart';

import '../models/models.dart';

class MagePlayScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageCreateScreen - build');

    log.info("Building");

    return ChangeNotifierProvider<Mage>(
      create: (context) => Mage(),
      child: Scaffold(
        appBar: AppBar(
          title: Text('Mage Play'),
        ),
        body: Container(),
        resizeToAvoidBottomInset: false,
      ),
    );
  }
}
