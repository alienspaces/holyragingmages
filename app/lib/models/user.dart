import 'package:meta/meta.dart';

/// A User encapsulates all User specific data
class UserModel {
  final String id;
  String name;

  UserModel({
    @required this.id,
    @required this.name,
  });
}
