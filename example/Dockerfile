FROM solr:6.6.0

# duplicate the data driven configset
RUN cp -r /opt/solr/server/solr/configsets/data_driven_schema_configs /opt/solr/server/solr/configsets/film && \
    cp -r /opt/solr/server/solr/configsets/data_driven_schema_configs /opt/solr/server/solr/configsets/job
# overwrite the managed schema
COPY ./film/managed-schema.xml /opt/solr/server/solr/configsets/film/conf/managed-schema
COPY ./job/managed-schema.xml /opt/solr/server/solr/configsets/job/conf/managed-schema