import 'package:flutter/material.dart';
import 'package:logging/logging.dart';

class MageCreateAttributeWidget extends StatefulWidget {
  final String name;
  final int value;
  final VoidCallback incrementValue;
  final VoidCallback decrementValue;
  final bool incrementEnabled;
  final bool decrementEnabled;

  MageCreateAttributeWidget({
    Key key,
    this.name,
    this.value,
    this.incrementValue,
    this.decrementValue,
    this.incrementEnabled,
    this.decrementEnabled,
  }) : super(key: key);

  @override
  MageCreateAttributeWidgetState createState() =>
      new MageCreateAttributeWidgetState();
}

class MageCreateAttributeWidgetState extends State<MageCreateAttributeWidget> {
  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageCreateAttributeWidget - build');
    log.info("Building");

    return Row(
      children: <Widget>[
        Expanded(
          flex: 4,
          child: Text(widget.name),
        ),
        Expanded(
          flex: 2,
          child: Align(
            alignment: Alignment.centerLeft,
            child: FlatButton(
              onPressed: widget.decrementEnabled ? widget.decrementValue : null,
              child: Icon(Icons.arrow_back),
            ),
          ),
        ),
        Expanded(
          flex: 2,
          child: Align(
            alignment: Alignment.center,
            child: Text(widget.value.toString()),
          ),
        ),
        Expanded(
          flex: 2,
          child: Align(
            alignment: Alignment.centerLeft,
            child: FlatButton(
              onPressed: widget.incrementEnabled ? widget.incrementValue : null,
              child: Icon(Icons.arrow_forward),
            ),
          ),
        ),
      ],
    );
  }
}
