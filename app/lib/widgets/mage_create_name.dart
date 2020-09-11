import 'package:flutter/material.dart';
import 'package:logging/logging.dart';

class MageCreateNameWidget extends StatefulWidget {
  final String value;

  MageCreateNameWidget({
    Key key,
    this.value,
  }) : super(key: key);

  @override
  MageCreateNameWidgetState createState() => new MageCreateNameWidgetState();
}

class MageCreateNameWidgetState extends State<MageCreateNameWidget> {
  TextEditingController _controller;

  void initState() {
    super.initState();
    _controller = TextEditingController();
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
    );
  }
}
