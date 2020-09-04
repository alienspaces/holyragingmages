import 'package:flutter/material.dart';

import '../widgets/mage_list.dart';
import '../widgets/mage_create.dart';
import '../env.dart';

class DashboardScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final String apiUrl = environment['apiUrl'];
    return Scaffold(
      appBar: AppBar(
        title: Text('Dashboard $apiUrl'),
      ),
      body: Container(
        child: Center(
          child: MageListWidget(),
        ),
      ),
      floatingActionButton: MageCreateWidget(),
    );
  }
}
