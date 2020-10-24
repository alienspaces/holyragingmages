import 'package:flutter/material.dart';
import 'package:logging/logging.dart';
import 'package:provider/provider.dart';

import 'package:holyragingmages/screens/screens.dart';
import 'package:holyragingmages/models/models.dart';

void main() {
  // Logging
  Logger.root.level = Level.INFO;
  Logger.root.onRecord.listen((record) {
    print(
        '${record.level.name}: ${record.time}: ${record.loggerName}: ${record.message}');
  });

  runApp(HolyRagingMages());
}

class HolyRagingMages extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MultiProvider(
      // Global providers
      providers: [
        ChangeNotifierProvider(create: (context) => AccountModel()),
        ChangeNotifierProvider(create: (context) => MageListModel()),
      ],
      child: MaterialApp(
        initialRoute: '/',
        routes: {
          '/': (context) => DashboardScreen(),
          '/mage_create': (context) => MageCreateScreen(),
          '/mage_play': (context) => MagePlayScreen(),
        },
      ),
    );
  }
}
