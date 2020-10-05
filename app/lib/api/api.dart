import 'package:dio/dio.dart';

import '../env.dart';

class Api {
  // Singleton instance
  static Api _instance;
  // HTTP package
  Dio dio;
  // API host
  static String apiUrl = environment['apiUrl'];
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

  // Singleton
  factory Api() {
    if (_instance == null) {
      _instance = Api._internal();
    }
    return _instance;
  }

  // Internal constructor
  Api._internal() {
    // Dio http client
    BaseOptions options = new BaseOptions(
      baseUrl: apiUrl,
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
  Future<List<dynamic>> getEntities() async {
    Response response = await this.dio.get(
      "/api/entities",
      queryParameters: {},
    );

    if (response.data == null) {
      return null;
    }

    return response.data["data"];
  }

  // postEntity creates a new entity
  Future<List<dynamic>> postEntity(Map<String, dynamic> data) async {
    Response response = await this.dio.post(
          "/api/entities",
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
