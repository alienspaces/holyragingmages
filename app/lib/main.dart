import 'package:flutter/material.dart';
import 'package:logging/logging.dart';
import 'package:provider/provider.dart';
import 'package:flutter/services.dart';

// Application packages
import 'package:holyragingmages/router.dart';
import 'package:holyragingmages/models/models.dart';
import 'package:holyragingmages/api/api.dart';

void main() {
  // Logging
  Logger.root.level = Level.INFO;
  Logger.root.onRecord.listen((record) {
    print('${record.level.name}: ${record.time}: ${record.loggerName}: ${record.message}');
  });

  runApp(HolyRagingMages());
}

// Api
Api api = ApiImpl();

// Routing
RouteGenerator router = RouteGenerator(api: api);

class HolyRagingMages extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    // Portrait mode only
    SystemChrome.setPreferredOrientations([
      DeviceOrientation.portraitUp,
      DeviceOrientation.portraitDown,
    ]);

    return MultiProvider(
      // Global providers
      providers: [
        ChangeNotifierProvider(create: (context) => Account(api: api)),
        ChangeNotifierProvider(create: (context) => Mage(api: api)),
        ChangeNotifierProvider(create: (context) => Familliar(api: api)),
        ChangeNotifierProvider(create: (context) => PlayerMageCollection(api: api)),
        ChangeNotifierProvider(create: (context) => StarterMageCollection(api: api)),
        ChangeNotifierProvider(create: (context) => StarterFamilliarCollection(api: api)),
      ],
      child: MaterialApp(
        initialRoute: '/',
        onGenerateRoute: router.generateRoute,
      ),
    );
  }
}
