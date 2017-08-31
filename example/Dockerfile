FROM solr:6.6.0

USER solr

# TODO: COPY does not seems to follow USER, so we copy to a tmp location and use cat to create the file with right perm
COPY ./film/managed-schema.xml /tmp/managed-schema.film
COPY ./job/managed-schema.xml /tmp/managed-schema.job

# duplicate the data driven configset
RUN cp -r /opt/solr/server/solr/configsets/data_driven_schema_configs /opt/solr/server/solr/configsets/film && \
    cp -r /opt/solr/server/solr/configsets/data_driven_schema_configs /opt/solr/server/solr/configsets/job && \
    cat /tmp/managed-schema.film > /opt/solr/server/solr/configsets/film/conf/managed-schema && \
    cat  /tmp/managed-schema.job > /opt/solr/server/solr/configsets/job/conf/managed-schema