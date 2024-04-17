import com.google.cloud.storage.Storage;
import com.google.cloud.storage.StorageOptions;
import com.google.cloud.storage.Blob;
import java.nio.file.Files;
import java.nio.file.Paths;

public class App {
    public static void main(String[] args) {
        Storage storage = StorageOptions.getDefaultInstance().getService();
        Blob blob = storage.get("my-bucket", "my-object");
        byte[] content = blob.getContent();
        Files.write(Paths.get("local-file.txt"), content);
        System.out.println("File contents: " + new String(content));
    }
}