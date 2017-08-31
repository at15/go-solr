.PHONY: solr local-install install uninstall reinstall stop start create config

solr:
	./solr/bin/solr start -f
solr-v:
	./solr/bin/solr start -f -v
local-install:
	./local_install.sh
install: local-install start create
uninstall:
	rm -rf solr
reinstall: uninstall install
start:
	./solr/bin/solr start
stop:
	./solr/bin/solr stop
config:
	rm -rf ./solr/server/solr/configsets/film
	rm -rf ./solr/server/solr/configsets/job
	cp -r ./solr/server/solr/configsets/data_driven_schema_configs ./solr/server/solr/configsets/film
	cp -r ./solr/server/solr/configsets/data_driven_schema_configs ./solr/server/solr/configsets/job
	cp ../example/film/managed-schema.xml ./solr/server/solr/configsets/film/conf/managed-schema
	cp ../example/job/managed-schema.xml ./solr/server/solr/configsets/job/conf/managed-schema
create: config
	solrgo core create demo
	solrgo core create film --configSet film
	solrgo core create job --configSet job

