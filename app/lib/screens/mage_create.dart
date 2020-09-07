import 'package:flutter/material.dart';

import '../widgets/mage_create.dart';

class MageCreateScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Mage Create'),
      ),
      body: Container(
        child: Center(
          child: MageCreateWidget(),
        ),
      ),
      resizeToAvoidBottomInset: false,
    );
  }
}
