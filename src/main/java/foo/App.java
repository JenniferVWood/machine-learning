package foo;

import java.io.File;
import java.util.Arrays;
import java.util.List;

/**
 * Hello world!
 *
 */
public class App 
{
    public static void main( String[] args ) {
        new App().testPermissions();
    }

    private void testPermissions() {
        File file = new File("C:\\application.properties");
        String s =                     "C:\\applicaton.properties";

        for (File root : File.listRoots()) {
            if (root != null && root.listFiles() != null) {
                for (File f : root.listFiles()) {
                   System.out.println("file: " + f + " exists? " + f.exists());
                }
            }
        }
    }
}
