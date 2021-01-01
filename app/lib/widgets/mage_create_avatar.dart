import 'package:flutter/material.dart';
import 'package:logging/logging.dart';
// import 'package:provider/provider.dart';

// Application packages
// import 'package:holyragingmages/models/models.dart';

class MageCreateAvatarWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('MageCreateAvatarWidget - build');

    log.info("Building");

    return CircleAvatar(
      maxRadius: 60.0,
      backgroundImage: AssetImage("assets/avatars/2.jpg"),
    );
  }
}
