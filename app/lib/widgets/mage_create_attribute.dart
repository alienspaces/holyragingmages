import 'package:flutter/material.dart';
import 'package:logging/logging.dart';

typedef void IncrementValueCallback();
typedef void DecrementValueCallback();

class MageCreateAttribute extends StatefulWidget {
  final String name;
  final int value;
  // final IncrementValueCallback incrementValue;
  // final DecrementValueCallback decrementValue;
  final VoidCallback incrementValue;
  final VoidCallback decrementValue;

  MageCreateAttribute({
    Key key,
    this.name,
    this.value,
    this.incrementValue,
    this.decrementValue,
  }) : super(key: key);

  @override
  MageCreateAttributeState createState() => new MageCreateAttributeState();
}

class MageCreateAttributeState extends State<MageCreateAttribute> {
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
              onPressed: widget.decrementValue,
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
              onPressed: widget.incrementValue,
              child: Icon(Icons.arrow_forward),
            ),
          ),
        ),
      ],
    );
  }
}
