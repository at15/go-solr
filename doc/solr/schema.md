# Schema

- http://lucene.apache.org/solr/guide/6_6/schema-api.html
- https://lucene.apache.org/solr/guide/6_6/field-type-definitions-and-properties.html
- https://lucene.apache.org/solr/guide/6_6/field-types-included-with-solr.html

## Get

- GET localhost:8983/solr/demo/schema?wt=json

see [managed-schema.json](managed-schema.json), includes field, field types, dynamic fields

## Add

````json
{
   "add-field":{
     "name":"sell-by",
     "type":"tdate",
     "stored":true 
   }
}
````

````json
{
    "responseHeader": {
        "status": 0,
        "QTime": 31
    }
}
````

## Delete

````json
{
  "delete-field" : { "name":"sell-by" }
}
````

````json
{
    "responseHeader": {
        "status": 0,
        "QTime": 16
    }
}
````

- [ ] TODO: even when there is error, the response code is also 200 ...

````json
{
    "responseHeader": {
        "status": 0,
        "QTime": 0
    },
    "errors": [
        {
            "delete-field": {
                "name": "sell-by"
            },
            "errorMessages": [
                "The field 'sell-by' is not present in this schema, and so cannot be deleted.\n"
            ]
        }
    ]
}
````