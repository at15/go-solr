FROM solr:6.6.0

USER solr

COPY ./film/managed-schema.xml /tmp/managed-schema.film
COPY ./job/managed-schema.xml /tmp/managed-schema.job

# duplicate the data driven configset
RUN cp -r /opt/solr/server/solr/configsets/data_driven_schema_configs /opt/solr/server/solr/configsets/film && \
    cp -r /opt/solr/server/solr/configsets/data_driven_schema_configs /opt/solr/server/solr/configsets/job && \
    cat /tmp/managed-schema.film > /opt/solr/server/solr/configsets/film/conf/managed-schema && \
    cat  /tmp/managed-schema.job > /opt/solr/server/solr/configsets/job/conf/managed-schema
# https://github.com/at15/go-solr/issues/19, it seems USER does not work
#USER solr
## overwrite the managed schema
#COPY ./film/managed-schema.xml /opt/solr/server/solr/configsets/film/conf/managed-schema
#COPY ./job/managed-schema.xml /opt/solr/server/solr/configsets/job/conf/managed-schema
# FIXME: USER seems does not work for copy, so we have to manually chown, which creates another layer
# chown: changing ownership of '/opt/solr/server/solr/configsets/film/conf/managed-schema': Operation not permitted
#RUN chown solr:solr /opt/solr/server/solr/configsets/film/conf/managed-schema && \
#    chown solr:solr /opt/solr/server/solr/configsets/job/conf/managed-schema