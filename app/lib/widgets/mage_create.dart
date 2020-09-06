import 'package:flutter/material.dart';

class MageCreateWidget extends StatefulWidget {
  @override
  MageCreateWidgetState createState() {
    return MageCreateWidgetState();
  }
}

class MageCreateWidgetState extends State<MageCreateWidget> {
  //
  @override
  Widget build(BuildContext context) {
    //
    return Column(
      children: <Widget>[
        CircleAvatar(
          maxRadius: 60.0,
          backgroundImage: AssetImage("assets/avatars/2.jpg"),
        ),
        TextField(
          decoration: InputDecoration(
            labelText: "Name",
          ),
        ),
        // Strength
        Row(
          children: <Widget>[
            Expanded(
              flex: 4,
              child: Text("Strength"),
            ),
            Expanded(
              flex: 2,
              child: Align(
                alignment: Alignment.centerLeft,
                child:
                    FlatButton(onPressed: null, child: Icon(Icons.arrow_back)),
              ),
            ),
            Expanded(
              flex: 2,
              child: Align(
                alignment: Alignment.center,
                child: Text("10"),
              ),
            ),
            Expanded(
              flex: 2,
              child: Align(
                alignment: Alignment.centerLeft,
                child: FlatButton(
                    onPressed: null, child: Icon(Icons.arrow_forward)),
              ),
            ),
          ],
        ),
        // Dexterity
        Row(
          children: <Widget>[
            Expanded(
              flex: 4,
              child: Text("Dexterity"),
            ),
            Expanded(
              flex: 2,
              child: Align(
                alignment: Alignment.centerLeft,
                child:
                    FlatButton(onPressed: null, child: Icon(Icons.arrow_back)),
              ),
            ),
            Expanded(
              flex: 2,
              child: Align(
                alignment: Alignment.center,
                child: Text("10"),
              ),
            ),
            Expanded(
              flex: 2,
              child: Align(
                alignment: Alignment.centerLeft,
                child: FlatButton(
                    onPressed: null, child: Icon(Icons.arrow_forward)),
              ),
            ),
          ],
        ),
        // Intelligence
        Row(
          children: <Widget>[
            Expanded(
              flex: 4,
              child: Text("Intelligence"),
            ),
            Expanded(
              flex: 2,
              child: Align(
                alignment: Alignment.centerLeft,
                child:
                    FlatButton(onPressed: null, child: Icon(Icons.arrow_back)),
              ),
            ),
            Expanded(
              flex: 2,
              child: Align(
                alignment: Alignment.center,
                child: Text("10"),
              ),
            ),
            Expanded(
              flex: 2,
              child: Align(
                alignment: Alignment.centerLeft,
                child: FlatButton(
                    onPressed: null, child: Icon(Icons.arrow_forward)),
              ),
            ),
          ],
        ),
      ],
    );
  }
}
