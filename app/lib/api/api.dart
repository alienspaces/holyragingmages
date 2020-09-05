import 'package:dio/dio.dart';

import 'mage.dart';
export 'mage.dart';

class BaseHandler {}

class BaseData {
  String id;
  String createdAt;
  String updatedAt;
}

class Api {
  // getMages returns a list of mages
  Future<List<MageData>> getMages() async {
    Response response;
    // Dio http client
    BaseOptions options = new BaseOptions(
      baseUrl: "http://10.1.1.22:8082",
    );
    Dio dio = new Dio(options);
    // dio.interceptors.add(LogInterceptor(responseBody: false));
    response = await dio.get("/mage/api/mages",
        queryParameters: {}, options: Options());
    // Mage http request and response handling
    MageHandler mageResponse = new MageHandler();
    return mageResponse.parseResponse(response);
  }
}
