package com.sohamkamani.storage;

import java.io.IOException;
import java.nio.file.Paths;
import com.google.cloud.storage.Blob;
import com.google.cloud.storage.BlobId;
import com.google.cloud.storage.BlobInfo;
import com.google.cloud.storage.Storage;
import com.google.cloud.storage.StorageOptions;
import com.google.cloud.storage.Storage.BlobListOption;


public class App {

  private static String projectId = "sohamkamani-demo";
  private static String bucketName = "sohamkamani-demo-bucket";
  private static String objectName = "sample.txt";

  public static void main(String[] args) throws Exception {
    uploadFile();

    // readFile();
    listFiles();
    // downloadFile();
    // deleteFile();
    // deleteFile();
  }

  // upload file to GCS
  public static void uploadFile() throws IOException {
    // Create a new GCS client
    Storage storage = StorageOptions.newBuilder().setProjectId(projectId).build().getService();
    // The blob ID indentifies the newly created blob, which consists of a bucket name and an object
    // name
    BlobId blobId = BlobId.of(bucketName, "my_folder/" + objectName);
    BlobInfo blobInfo = BlobInfo.newBuilder(blobId).build();

    // The filepath on our local machine that we want to upload
    String filePath = "C://var/sample.txt";

    // upload the file and print the status
    storage.createFrom(blobInfo, Paths.get(filePath));
    System.out
        .println("File " + filePath + " uploaded to bucket " + bucketName + " as " + objectName);
  }

  // download file from GCS
  public static void downloadFile() throws IOException {
    // we'll download the same file to another file path
    String filePath = "/tmp/sample_downloaded.txt";

    // Create a new GCS client and get the blob object from the blob ID
    Storage storage = StorageOptions.newBuilder().setProjectId(projectId).build().getService();
    BlobId blobId = BlobId.of(bucketName, objectName);
    Blob blob = storage.get(blobId);

    // download the file and print the status
    blob.downloadTo(Paths.get(filePath));
    System.out.println("File " + objectName + " downloaded to " + filePath);
  }

  // read contents of the file
  public static void readFile() throws IOException {
    // Create a new GCS client and get the blob object from the blob ID
    Storage storage = StorageOptions.newBuilder().setProjectId(projectId).build().getService();
    BlobId blobId = BlobId.of(bucketName, objectName);
    Blob blob = storage.get(blobId);

    // read the contents of the file and print it
    String contents = new String(blob.getContent());
    System.out.println("Contents of file " + objectName + ": " + contents);
  }

  // delete an existing file from GCS
  public static void deleteFile() throws IOException {
    // Create a new GCS client and get the blob object from the blob ID
    Storage storage = StorageOptions.newBuilder().setProjectId(projectId).build().getService();
    BlobId blobId = BlobId.of(bucketName, objectName);
    Blob blob = storage.get(blobId);

    if (blob == null) {
      System.out.println("File " + objectName + " does not exist in bucket " + bucketName);
      return;
    }

    // delete the file and print the status
    blob.delete();
    System.out.println("File " + objectName + " deleted from bucket " + bucketName);
  }

  // list all files in a folder or bucket
  public static void listFiles() throws IOException {
    // Create a new GCS client and get the blob object from the blob ID
    Storage storage = StorageOptions.newBuilder().setProjectId(projectId).build().getService();

    System.out.println("Files in bucket " + bucketName + ":");
    // list all the blobs in the bucket
    for (Blob blob : storage
        .list(bucketName, BlobListOption.currentDirectory(), BlobListOption.prefix("my_folder/"))
        .iterateAll()) {
      System.out.println(blob.getName());
    }
  }
}
