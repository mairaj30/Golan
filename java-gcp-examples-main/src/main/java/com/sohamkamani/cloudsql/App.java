package com.sohamkamani.cloudsql;

import java.sql.ResultSet;
import java.sql.SQLException;
import javax.sql.DataSource;

public class App {
  public static void main(String[] args) throws SQLException {
    DataSource dataSource = CloudSqlConnectionPoolFactory.createConnectionPool();

    ResultSet rs =
        dataSource.getConnection().prepareStatement("select * from birds").executeQuery();

    while (rs.next()) {
      System.out
          .println("name: " + rs.getString("bird") + " description:" + rs.getString("description"));
    }
  }
}
