import 'package:flutter/material.dart';
import 'package:provider/provider.dart';
import 'package:logging/logging.dart';

// Application packages
import 'package:holyragingmages/api/api.dart';
import 'package:holyragingmages/models/models.dart';

class MagePlayScreen extends StatelessWidget {
  final Api api;

  MagePlayScreen({Key key, this.api}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageCreateScreen - build');

    log.info("Building");

    return ChangeNotifierProvider<Mage>(
      create: (context) => Mage(api: api),
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
