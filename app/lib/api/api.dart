import 'package:dio/dio.dart';

class Api {
  Dio dio;

  Api() {
    // Dio http client
    BaseOptions options = new BaseOptions(
      baseUrl: "http://10.1.1.22:8082",
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
          "/mage/api/mages",
          queryParameters: {},
          options: Options(),
        );

    if (response.data == null) {
      return null;
    }

    return response.data["data"];
  }
}
