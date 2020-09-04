import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../models/models.dart';

class MageListWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    // mage list
    var mageList = Provider.of<MageListModel>(context);
    return GridView.count(
      crossAxisCount: 2,
      childAspectRatio: 0.90,
      children: List.generate(
        mageList.mages.length,
        (index) => ListTile(
          title: Text(
            mageList.mages[index].name,
          ),
        ),
      ),
    );
  }
}
