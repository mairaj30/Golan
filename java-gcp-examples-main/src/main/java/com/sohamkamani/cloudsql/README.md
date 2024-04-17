This folder contains the example for my blog post on [connecting to CloudSQL in Java](https://sohamkamani.com/java/cloudsql/)

To run this example execute this command from the main directory:

```bash
mvn -DMAIN_CLASS=com.sohamkamani.cloudsql.App clean compile assembly:single \ 
&& java -jar target/java-gcp-examples-1.0-SNAPSHOT-jar-with-dependencies.jar
```