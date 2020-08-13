import 'dart:convert';
import 'dart:io';

// Generates lib/env.dart from current environment
// USAGE: dart tool/generate_env.dart
Future<void> main() async {
  final config = {
    'apiUrl': Platform.environment['APP_API_URL'],
  };

  final filename = 'lib/env.dart';
  await File(filename)
      .writeAsString('final environment = ${json.encode(config)};');
}
