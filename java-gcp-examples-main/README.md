# Java Google Cloud Examples

This is the example repo for my blog post on [Java Google Cloud Examples](java/google-cloud)

To run the examples in this repository:

1. Make sure you have Java, JDK, and Maven installed
2. Clone the repo
3. Run `mvn -DMAIN_CLASS=<class-to-run>  clean compile assembly:single` - this should give you a JAR file in the `target` folder
4. Run `java -jar target/java-gcp-examples-1.0-SNAPSHOT-jar-with-dependencies.jar` to run the JAR file

For example, if you want to run the example in [src/main/java/com/sohamkamani/cloudsql](./src/main/java/com/sohamkamani/cloudsql), you should run:

```bash
mvn -DMAIN_CLASS=com.sohamkamani.cloudsql.App clean compile assembly:single \ 
&& java -jar target/java-gcp-examples-1.0-SNAPSHOT-jar-with-dependencies.jar
```