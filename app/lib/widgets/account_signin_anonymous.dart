import 'package:flutter/material.dart';
import 'package:logging/logging.dart';

class AccountSignInAnonymousWidget extends StatelessWidget {
  final Function signInCallback;

  AccountSignInAnonymousWidget({this.signInCallback});

  @override
  Widget build(BuildContext context) {
    // Logger
    final log = Logger('AccountSignInGoogleWidget - build');

    log.info("Building");

    return OutlineButton(
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(40)),
      borderSide: BorderSide(color: Colors.grey),
      child: Padding(
        padding: const EdgeInsets.fromLTRB(0, 10, 0, 10),
        child: Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            Padding(
              padding: const EdgeInsets.only(left: 10),
              child: const Text('Skip'),
            ),
          ],
        ),
      ),
      onPressed: signInCallback,
    );
  }
}
