# Core

> In Solr, the term core is used to refer to a single index and associated transaction log and configuration files (including the solrconfig.xml and Schema files, among others). Your Solr installation can have multiple cores if needed, which allows you to index data with different structures in the same server, and maintain more control over how your data is presented to different audiences. In SolrCloud mode you will be more familiar with the term collection. Behind the scenes a collection consists of one or more cores

> Multiple cores in Solr are analogous to a database having multiple tables

- http://lucene.apache.org/solr/guide/6_6/coreadmin-api.html
- http://lucene.apache.org/solr/guide/6_6/solr-cores-and-solr-xml.html

## Create

- you can use GET, no need to use POST, but both should work, I guess solr doesn't filter http methods
- [ ] TODO: wrong install dir, config set etc.
- [ ] TODO: collection, when we start using solr cloud

200, success

````json
{"responseHeader":{"status":0,"QTime":239},"core":"films2"}
````

500, duplicated cores

````json
{"responseHeader":{"status":500,"QTime":1},"error":{"metadata":["error-class","org.apache.solr.common.SolrException","root-error-class","org.apache.solr.common.SolrException"],"msg":"Core with name 'films2' already exists.","trace":"org.apache.solr.common.SolrException: Core with name 'films2' already exists.\n\tat org.apache.solr.core.CoreContainer.create(CoreContainer.java:840)\n\tat org.apache.solr.handler.admin.CoreAdminOperation.lambda$static$0(CoreAdminOperation.java:91)\n\tat org.apache.solr.handler.admin.CoreAdminOperation.execute(CoreAdminOperation.java:384)\n\tat org.apache.solr.handler.admin.CoreAdminHandler$CallInfo.call(CoreAdminHandler.java:388)\n\tat org.apache.solr.handler.admin.CoreAdminHandler.handleRequestBody(CoreAdminHandler.java:174)\n\tat org.apache.solr.handler.RequestHandlerBase.handleRequest(RequestHandlerBase.java:173)\n\tat org.apache.solr.servlet.HttpSolrCall.handleAdmin(HttpSolrCall.java:748)\n\tat org.apache.solr.servlet.HttpSolrCall.handleAdminRequest(HttpSolrCall.java:729)\n\tat org.apache.solr.servlet.HttpSolrCall.call(HttpSolrCall.java:510)\n\tat org.apache.solr.servlet.SolrDispatchFilter.doFilter(SolrDispatchFilter.java:361)\n\tat org.apache.solr.servlet.SolrDispatchFilter.doFilter(SolrDispatchFilter.java:305)\n\tat org.eclipse.jetty.servlet.ServletHandler$CachedChain.doFilter(ServletHandler.java:1691)\n\tat org.eclipse.jetty.servlet.ServletHandler.doHandle(ServletHandler.java:582)\n\tat org.eclipse.jetty.server.handler.ScopedHandler.handle(ScopedHandler.java:143)\n\tat org.eclipse.jetty.security.SecurityHandler.handle(SecurityHandler.java:548)\n\tat org.eclipse.jetty.server.session.SessionHandler.doHandle(SessionHandler.java:226)\n\tat org.eclipse.jetty.server.handler.ContextHandler.doHandle(ContextHandler.java:1180)\n\tat org.eclipse.jetty.servlet.ServletHandler.doScope(ServletHandler.java:512)\n\tat org.eclipse.jetty.server.session.SessionHandler.doScope(SessionHandler.java:185)\n\tat org.eclipse.jetty.server.handler.ContextHandler.doScope(ContextHandler.java:1112)\n\tat org.eclipse.jetty.server.handler.ScopedHandler.handle(ScopedHandler.java:141)\n\tat org.eclipse.jetty.server.handler.ContextHandlerCollection.handle(ContextHandlerCollection.java:213)\n\tat org.eclipse.jetty.server.handler.HandlerCollection.handle(HandlerCollection.java:119)\n\tat org.eclipse.jetty.server.handler.HandlerWrapper.handle(HandlerWrapper.java:134)\n\tat org.eclipse.jetty.rewrite.handler.RewriteHandler.handle(RewriteHandler.java:335)\n\tat org.eclipse.jetty.server.handler.HandlerWrapper.handle(HandlerWrapper.java:134)\n\tat org.eclipse.jetty.server.Server.handle(Server.java:534)\n\tat org.eclipse.jetty.server.HttpChannel.handle(HttpChannel.java:320)\n\tat org.eclipse.jetty.server.HttpConnection.onFillable(HttpConnection.java:251)\n\tat org.eclipse.jetty.io.AbstractConnection$ReadCallback.succeeded(AbstractConnection.java:273)\n\tat org.eclipse.jetty.io.FillInterest.fillable(FillInterest.java:95)\n\tat org.eclipse.jetty.io.SelectChannelEndPoint$2.run(SelectChannelEndPoint.java:93)\n\tat org.eclipse.jetty.util.thread.strategy.ExecuteProduceConsume.executeProduceConsume(ExecuteProduceConsume.java:303)\n\tat org.eclipse.jetty.util.thread.strategy.ExecuteProduceConsume.produceConsume(ExecuteProduceConsume.java:148)\n\tat org.eclipse.jetty.util.thread.strategy.ExecuteProduceConsume.run(ExecuteProduceConsume.java:136)\n\tat org.eclipse.jetty.util.thread.QueuedThreadPool.runJob(QueuedThreadPool.java:671)\n\tat org.eclipse.jetty.util.thread.QueuedThreadPool$2.run(QueuedThreadPool.java:589)\n\tat java.lang.Thread.run(Thread.java:745)\n","code":500}}
````

## Commit

> In SQL databases, you perform commit after updates. Similarly, Solr also requires a commit, and the changes are searchable only after the commit has been triggered on that core


