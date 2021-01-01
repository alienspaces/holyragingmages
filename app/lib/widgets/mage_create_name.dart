import 'package:flutter/material.dart';
import 'package:logging/logging.dart';
import 'package:provider/provider.dart';

// Application packages
import 'package:holyragingmages/models/models.dart';

class MageCreateNameWidget extends StatefulWidget {
  MageCreateNameWidget({
    Key key,
  }) : super(key: key);

  @override
  MageCreateNameWidgetState createState() => new MageCreateNameWidgetState();
}

class MageCreateNameWidgetState extends State<MageCreateNameWidget> {
  TextEditingController _controller;

  void initState() {
    super.initState();
    _controller = TextEditingController();
    _controller.addListener(() {
      // Mage models
      var mageModel = Provider.of<Mage>(context, listen: false);
      final text = _controller.text;
      mageModel.name = text;
    });
  }

  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageCreateAttributeWidget - build');
    log.info("Building");

    return TextField(
      decoration: InputDecoration(
        filled: true,
        labelText: "Name",
      ),
      controller: _controller,
      autofocus: true,
    );
  }
}
