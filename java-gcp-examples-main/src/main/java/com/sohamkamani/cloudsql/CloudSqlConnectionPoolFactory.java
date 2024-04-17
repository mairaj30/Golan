package com.sohamkamani.cloudsql;

import javax.sql.DataSource;
import com.zaxxer.hikari.HikariConfig;
import com.zaxxer.hikari.HikariDataSource;

public class CloudSqlConnectionPoolFactory {
  private static final String INSTANCE_CONNECTION_NAME =
      "sohamkamani-website:us-central1:my-sample-database";
  private static final String DB_USER = "postgres";
  private static final String DB_PASS = "o.^5PU~RU,ai-*j&";
  private static final String DB_NAME = "bird_encyclopedia";

  public static DataSource createConnectionPool() {
    HikariConfig config = new HikariConfig();
    config.setJdbcUrl(String.format("jdbc:postgresql:///%s", DB_NAME));
    config.setUsername(DB_USER);
    config.setPassword(DB_PASS);
    config.addDataSourceProperty("socketFactory", "com.google.cloud.sql.postgres.SocketFactory");
    config.addDataSourceProperty("cloudSqlInstance", INSTANCE_CONNECTION_NAME);

    return new HikariDataSource(config);
  }
}
