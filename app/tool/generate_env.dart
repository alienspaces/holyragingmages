import 'dart:convert';
import 'dart:io';

// Generates lib/env.dart from current environment
// USAGE: dart tool/generate_env.dart
Future<void> main() async {
  final config = {
    'apiHost': Platform.environment['APP_API_HOST'],
  };

  final filename = 'lib/env.dart';
  await File(filename)
      .writeAsString('final environment = ${json.encode(config)};');
}
