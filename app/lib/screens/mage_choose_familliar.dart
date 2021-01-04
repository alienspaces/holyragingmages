import 'package:flutter/material.dart';
import 'package:logging/logging.dart';
import 'package:provider/provider.dart';

// Application packages
import 'package:holyragingmages/api/api.dart';
import 'package:holyragingmages/models/models.dart';

class MageChooseFamilliarScreen extends StatefulWidget {
  final Api api;

  MageChooseFamilliarScreen({Key key, this.api}) : super(key: key);

  @override
  _MageChooseFamilliarScreenState createState() => _MageChooseFamilliarScreenState();
}

class _MageChooseFamilliarScreenState extends State<MageChooseFamilliarScreen> {
  @override
  void initState() {
    // Account model
    var accountModel = Provider.of<Account>(context, listen: false);

    // Mage model
    var mageModel = Provider.of<Mage>(context, listen: false);

    mageModel.clear();
    mageModel.accountId = accountModel.id;

    super.initState();
  }

  @override
  void dispose() {
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageChooseFamilliarScreen - build');

    log.info("Building");

    // Styling
    EdgeInsetsGeometry padding = EdgeInsets.fromLTRB(15, 10, 15, 10);

    return Scaffold(
      appBar: AppBar(
        title: Text('Choose Familliar'),
      ),
      body: Container(
        padding: padding,
        child: Column(
          children: <Widget>[],
        ),
      ),
      resizeToAvoidBottomInset: false,
    );
  }
}
