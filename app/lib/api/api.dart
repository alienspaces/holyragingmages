import 'package:dio/dio.dart';

import '../env.dart';

class Api {
  Dio dio;
  final String apiUrl = environment['apiUrl'];

  Api() {
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

  // getEntities returns a list of entities
  Future<List<dynamic>> getEntities() async {
    Response response = await this.dio.get(
      "/entity/api/entities",
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
          "/entity/api/entities",
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
          "/entity/api/entities/$id",
          data: data,
        );

    if (response.data == null) {
      return null;
    }

    return response.data["data"];
  }
}
