import 'package:test/test.dart';
import 'package:logging/logging.dart';

import 'package:holyragingmages/models/mage.dart';

void main() {
  // Logging
  Logger.root.level = Level.INFO;
  Logger.root.onRecord.listen((record) {
    print(
        '${record.level.name}: ${record.time}: ${record.loggerName}: ${record.message}');
  });

  test('New mage defaults', () {
    var mage = new MageModel();

    expect(mage.points, 10, reason: 'Points equals expected value');
    expect(mage.strength, 10, reason: 'Strength equals expected value');
    expect(mage.dexterity, 10, reason: 'Dexterity equals expected value');
    expect(mage.intelligence, 10, reason: 'Intelligence equals expected value');
  });

  test('Adjusting strength', () {
    var mage = new MageModel();

    mage.strength = 12;

    expect(mage.strength, 12, reason: 'Strength equals expected value');
    expect(mage.points, 8, reason: 'Points equals expected value');

    mage.strength = 10;
    expect(mage.strength, 10, reason: 'Strength equals expected value');
    expect(mage.points, 10, reason: 'Points equals expected value');
  });

  test('Adjusting dexterity', () {
    var mage = new MageModel();

    mage.dexterity = 12;

    expect(mage.dexterity, 12, reason: 'Dexterity equals expected value');
    expect(mage.points, 8, reason: 'Points equals expected value');

    mage.dexterity = 10;
    expect(mage.dexterity, 10, reason: 'Dexterity equals expected value');
    expect(mage.points, 10, reason: 'Points equals expected value');
  });

  test('Adjusting intelligence', () {
    var mage = new MageModel();

    mage.intelligence = 12;

    expect(mage.intelligence, 12, reason: 'Intelligence equals expected value');
    expect(mage.points, 8, reason: 'Points equals expected value');

    mage.intelligence = 10;
    expect(mage.intelligence, 10, reason: 'Intelligence equals expected value');
    expect(mage.points, 10, reason: 'Points equals expected value');
  });

  test('New mage from JSON', () {
    Map<String, dynamic> mageJson = {
      "id": "9f6f269b-b025-4817-8a25-f014e79db609",
      "strength": 12,
      "dexterity": 12,
      "intelligence": 12,
    };

    var mage = MageModel.fromJson(mageJson);

    expect(mage.points, 4, reason: 'Points equals expected value');
    expect(mage.strength, 12, reason: 'Strength equals expected value');
    expect(mage.dexterity, 12, reason: 'Dexterity equals expected value');
    expect(mage.intelligence, 12, reason: 'Intelligence equals expected value');
  });
}
