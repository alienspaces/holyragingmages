import 'package:flutter/material.dart';
import 'package:holyragingmages/env.dart';

class DashboardScreen extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    final String apiUrl = environment['apiUrl'];
    return Scaffold(
      appBar: AppBar(
        title: Text('Dashboard $apiUrl'),
      ),
      body: Container(),
    );
  }
}
