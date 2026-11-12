/*
 * SPACE-AICOS is pleased to support the open source community by making
 * SPACE-AICOS available.
 * Copyright (C) 2026 China Academy of Information and Communications Technology (CAICT).
 * All rights reserved.
 * Licensed under the MIT License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on
 * an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */

function getNextSequence(name) {
    var ret = db.cc_idgenerator.findAndModify(
        {
            query: { _id: name },
            update: { $inc: { SequenceID: 1 } },
            new: true
        }
    );
    return ret.SequenceID;
}

db.cc_ObjAttDes.find().forEach(function (myDoc) {
    var nid = getNextSequence("cc_ObjAttDes")
    print("nid: " + nid);
    db.cc_ObjAttDes.update(
        {
            "bk_obj_id": myDoc.bk_obj_id,
            "bk_property_id": myDoc.bk_property_id,
            "bk_supplier_account": myDoc.bk_supplier_account,
        },
        {
            "$set": {
                "id": nid,
            }
        })
});

