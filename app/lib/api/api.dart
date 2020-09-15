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

  // getMages returns a list of mages
  Future<List<dynamic>> getMages() async {
    Response response = await this.dio.get(
      "/entity/api/entities",
      queryParameters: {},
    );

    if (response.data == null) {
      return null;
    }

    return response.data["data"];
  }

  // postMage creates a new mage
  Future<List<dynamic>> postMage(Map<String, dynamic> data) async {
    Response response = await this.dio.post(
          "/entity/api/entities",
          data: data,
        );

    if (response.data == null) {
      return null;
    }

    return response.data["data"];
  }

  // putMage creates a new mage
  Future<List<dynamic>> putMage(String id, Map<String, dynamic> data) async {
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
