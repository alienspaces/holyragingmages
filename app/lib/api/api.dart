// Application exports
export 'package:holyragingmages/api/api_impl.dart';
export 'package:holyragingmages/api/api_mock.dart';

abstract class Api {
  set apiToken(String token);
  Future<List<dynamic>> postAuth(Map<String, dynamic> data);
  Future<List<dynamic>> getEntities(String accountId);
  Future<List<dynamic>> postEntity(String accountId, Map<String, dynamic> data);
  Future<List<dynamic>> putEntity(String id, Map<String, dynamic> data);
}
