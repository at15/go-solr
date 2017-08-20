.PHONY: solr install reinstall

solr:
	./solr/bin/solr start -f
solr-v:
	./solr/bin/solr start -f -v
install:
	./local_install.sh
reinstall:
	rm -rf solr
	./local_install.sh