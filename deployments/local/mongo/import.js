mongo = new Mongo();
db = mongo.getDB("Instance");
data = cat("/data/data.json")
insert = db.instances.insertMany(JSON.parse(data))
print("Inseridos "+insert.insertedIds.length+" Registros no database Instance")