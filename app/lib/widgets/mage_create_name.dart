import 'package:flutter/material.dart';
import 'package:logging/logging.dart';

typedef UpdateValueCallback = void Function(String value);

class MageCreateNameWidget extends StatefulWidget {
  final String value;
  final UpdateValueCallback updateValue;

  MageCreateNameWidget({
    Key key,
    this.value,
    this.updateValue,
  }) : super(key: key);

  @override
  MageCreateNameWidgetState createState() => new MageCreateNameWidgetState();
}

class MageCreateNameWidgetState extends State<MageCreateNameWidget> {
  TextEditingController _controller;

  void initState() {
    super.initState();
    _controller = TextEditingController(text: widget.value);
    _controller.addListener(() {
      final text = _controller.text;
      widget.updateValue(text);
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
