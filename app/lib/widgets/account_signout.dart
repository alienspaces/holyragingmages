import 'package:flutter/material.dart';
import 'package:logging/logging.dart';

class AccountSignOutWidget extends StatelessWidget {
  final Function signOutCallback;

  AccountSignOutWidget({this.signOutCallback});

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('SignOutWidget - build');

    log.info("Building");

    return OutlineButton(
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(40)),
      borderSide: BorderSide(color: Colors.grey),
      child: Padding(
        padding: const EdgeInsets.only(left: 10),
        child: const Text('Sign Out'),
      ),
      onPressed: signOutCallback,
    );
  }
}
