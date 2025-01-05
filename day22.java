import java.io.*;
import java.util.*;

import static java.lang.System.in;
import static java.lang.System.out;

public class day22 {
    public static void main(String[] args) {
        long[] sum = new long[]{0L};
        var prices = new HashMap<Long, Integer>();
        new BufferedReader(new InputStreamReader(in)).lines().map(Long::parseLong).forEach(n -> {
            var processed = new HashSet<Integer>();
            long nn = n; var c1 = 0; var c2 = 0; var c3 = 0;
            for (int i = 0; i < 2000; i++) {
                var nx = ((nn * 64) ^ nn) % 16777216;
                nx = ((nx / 32) ^ nx) % 16777216;
                nx = ((nx * 2048) ^ nx) % 16777216;
                if (i > 2) {
                    var key = c1 * 1000000 + c2 * 10000 + c3 * 100 + (int) (nx % 10 - nn % 10);
                    if (processed.add(key)) {
                        prices.merge((long) key, (int) nx % 10, Integer::sum);
                    }
                }
                c1 = c2; c2 = c3; c3 = (int) (nx % 10 - nn % 10);
                nn = nx;
            }
            sum[0] += nn;
        });
        out.println(sum[0] + " " + prices.values().stream().max(Integer::compareTo).get());
    }
}
