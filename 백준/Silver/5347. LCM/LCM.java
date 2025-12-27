import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class Main{
    public static void main(String[] args) throws IOException{
        BufferedReader reader = new BufferedReader((new InputStreamReader(System.in)));
        StringBuilder builder = new StringBuilder();

        int n = Integer.parseInt(reader.readLine());
        for (int i = 0; i<n; i++){
            String[] ab = reader.readLine().split(" ");
            int a = Integer.parseInt(ab[0]);
            int b = Integer.parseInt(ab[1]);
            builder.append(lcm(a, b)).append("\n");
        }
        System.out.println(builder);
        reader.close();
    }

    public static int gcb(int a, int b){
        while (a>0){
            int r = b%a;
            b = a;
            a = r;
        }
        return b;
    }

    public static long lcm(int a, int b){
        return (long) a/gcb(a, b)*b;
    }
}