import 'package:flutter/material.dart';
import 'package:logging/logging.dart';
import 'package:provider/provider.dart';

// Application packages
import 'package:holyragingmages/api/api.dart';
import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/widgets/mage_create_attributes.dart';
import 'package:holyragingmages/widgets/mage_create_avatar.dart';
import 'package:holyragingmages/widgets/mage_create_button.dart';
import 'package:holyragingmages/widgets/mage_create_name.dart';

class MageCreateScreen extends StatefulWidget {
  final Api api;

  MageCreateScreen({Key key, this.api}) : super(key: key);

  @override
  _MageCreateScreenState createState() => _MageCreateScreenState();
}

class _MageCreateScreenState extends State<MageCreateScreen> {
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
    final log = Logger('MageCreateScreen - build');

    log.info("Building");

    return Scaffold(
      appBar: AppBar(
        title: Text('Create Mage'),
      ),
      body: Container(
        child: Column(
          children: <Widget>[
            MageCreateAvatarWidget(),
            MageCreateNameWidget(),
            MageCreateAttributesWidget(),
          ],
        ),
      ),
      floatingActionButton: MageCreateButtonWidget(),
      resizeToAvoidBottomInset: false,
    );
  }
}
