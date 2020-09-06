import 'package:flutter/material.dart';

import '../widgets/mage_list.dart';
import '../widgets/mage_create_button.dart';

class DashboardScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Dashboard'),
      ),
      body: Container(
        child: Center(
          child: MageListWidget(),
        ),
      ),
      floatingActionButton: MageCreateButtonWidget(),
    );
  }
}
