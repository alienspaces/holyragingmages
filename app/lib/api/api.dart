import 'package:dio/dio.dart';

import 'mage.dart';

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
    Dio dio = new Dio();
    response = await dio.get("/api/mages", queryParameters: {});
    // Mage http request and response handling
    MageHandler mageResponse = new MageHandler();
    return mageResponse.parseResponse(response);
  }
}
