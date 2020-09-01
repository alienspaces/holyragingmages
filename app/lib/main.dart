import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'package:holyragingmages/screens/screens.dart';
import 'package:holyragingmages/models/models.dart';

// Stub...
void main() {
  runApp(HolyRagingMages());
}

class HolyRagingMages extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      providers: [
        Provider(create: (context) => MageListModel()),
      ],
      child: MaterialApp(
        home: DashboardScreen(),
      ),
    );
  }
}
