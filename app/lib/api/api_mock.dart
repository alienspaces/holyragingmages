// Application packages
import 'package:holyragingmages/api/api.dart';

class ApiMock implements Api {
  // API token
  String _apiToken;

  String get apiToken {
    return _apiToken;
  }

  set apiToken(String token) {
    this._apiToken = token;
  }

  // Internal constructor
  ApiMock();

  // postAuth creates a new entity
  Future<List<dynamic>> postAuth(Map<String, dynamic> data) async {
    return Future.delayed(Duration(microseconds: 1), () {
      List<dynamic> responseData = [
        {
          "blah": 1,
        },
      ];
      return responseData;
    });
  }

  // getEntities returns a list of entities
  Future<List<dynamic>> getEntities(String accountId) async {
    return Future.delayed(Duration(microseconds: 1), () {
      List<dynamic> responseData = [
        {
          "blah": 1,
        },
      ];
      return responseData;
    });
  }

  // postEntity creates a new entity
  Future<List<dynamic>> postEntity(String accountId, Map<String, dynamic> data) async {
    return Future.delayed(Duration(microseconds: 1), () {
      List<dynamic> responseData = [
        {
          "blah": 1,
        },
      ];
      return responseData;
    });
  }

  // putEntity updates an existing entity
  Future<List<dynamic>> putEntity(String id, Map<String, dynamic> data) async {
    return Future.delayed(Duration(microseconds: 1), () {
      List<dynamic> responseData = [
        {
          "blah": 1,
        },
      ];
      return responseData;
    });
  }
}
