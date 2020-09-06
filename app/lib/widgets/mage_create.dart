import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../models/models.dart';

class MageCreateWidget extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    var mageList = Provider.of<MageListModel>(context);
    return Container(
      color: Colors.green,
      child: IconButton(
        icon: Icon(Icons.add_a_photo),
        onPressed: () {
          _addMage(mageList);
        },
      ),
    );
  }

  void _addMage(MageListModel mageList) {
    MageModel mage = new MageModel(
      id: null,
      name: "Henry",
      strength: null,
      dexterity: null,
      intelligence: null,
    );
    mageList.addMage(mage);
  }
}
