import 'package:dio/dio.dart';

// Application packages
import 'package:holyragingmages/env.dart';
import 'package:holyragingmages/api/api.dart';

class ApiImpl implements Api {
  // HTTP package
  Dio dio;
  // API host
  static String apiHost = environment['apiHost'];
  // API token
  String _apiToken;

  String get apiToken {
    return _apiToken;
  }

  set apiToken(String token) {
    this._apiToken = token;

    // Add token to all DIO requests
    this.dio.options.headers["Authorization"] = "Bearer " + this.apiToken;
  }

  // Internal constructor
  ApiImpl() {
    // Dio http client
    BaseOptions options = new BaseOptions(
      baseUrl: apiHost,
      headers: {
        "content-type": "application/json; charset=utf-8",
      },
    );

    dio = new Dio(options);

    // Logging
    dio.interceptors.add(LogInterceptor(responseBody: false));
  }

  // postAuth creates a new entity
  Future<List<dynamic>> postAuth(Map<String, dynamic> data) async {
    Response response = await this.dio.post(
          "/api/auth",
          data: data,
        );

    if (response.data == null) {
      return null;
    }

    return response.data["data"];
  }

  // getEntities returns a list of entities
  Future<List<dynamic>> getEntities(String accountId) async {
    Response response = await this.dio.get(
      "/api/accounts/" + accountId + "/entities",
      queryParameters: {},
    );

    if (response.data == null) {
      return null;
    }

    return response.data["data"];
  }

  // postEntity creates a new entity
  Future<List<dynamic>> postEntity(String accountId, Map<String, dynamic> data) async {
    Response response = await this.dio.post(
          "/api/accounts/" + accountId + "/entities",
          data: data,
        );

    if (response.data == null) {
      return null;
    }

    return response.data["data"];
  }

  // putEntity updates an existing entity
  Future<List<dynamic>> putEntity(String id, Map<String, dynamic> data) async {
    Response response = await this.dio.put(
          "/api/entities/$id",
          data: data,
        );

    if (response.data == null) {
      return null;
    }

    return response.data["data"];
  }
}
