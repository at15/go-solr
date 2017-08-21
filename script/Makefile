.PHONY: solr install uninstall reinstall stop

solr:
	./solr/bin/solr start -f
solr-v:
	./solr/bin/solr start -f -v
install:
	./local_install.sh
	./solr/bin/solr start
	solrgo core create demo
	./solr/bin/solr stop
uninstall:
	rm -rf solr
reinstall: uninstall install
stop:
	./solr/bin/solr stop
