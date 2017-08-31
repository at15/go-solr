# Film

This is an example from Solr distribution `<sorl>/example/films`.
We already updated the schema, most fields all inferred by Solr.

````xml
<field name="name" type="text_general" multiValued="false" stored="true" />
<field name="initial_release_date" type="tdate" stored="true"/>
```` 

````bash
solrgo core index film films.json
````

TODO

````log
INFO[0000] Core job already exists pkg=gosolr 
cd example/film; solrgo core index film films.json
FATA[0000] can't index films.json: solr: can't update document http://solr:8983/solr/film/update?commit=true&wt=json: solr: 500: Error persisting managed schema /opt/solr/server/solr/configsets/film/conf/managed-schema pkg=gosolr 
````