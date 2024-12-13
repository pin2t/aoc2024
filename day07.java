import java.io.*;
import java.util.*;
import java.util.function.*;
import java.util.regex.*;

import static java.lang.System.in;
import static java.lang.System.out;

public class day07 {
    public static void main(String[] args) {
        var re = Pattern.compile("\\d+");
        var results = new long[]{0, 0};
        new BufferedReader(new InputStreamReader(in)).lines().forEach(line -> {
            var ns = new Scanner(line).findAll(re).mapToLong(m -> Long.parseLong(m.group())).toArray();
            if (match(ns, 1, 0, List.of((c, n) -> c + n, (c, n) -> c * n))) {
                results[0] += ns[0];
            }
            if (match(ns, 1, 0, List.of((c, n) -> c + n, (c, n) -> c * n, (c, n) ->
                Long.parseLong(c.toString() + n.toString())
            ))) {
                results[1] += ns[0];
            }
        });
        out.println(results[0] + " " + results[1]);
    }

    static boolean match(long[] ns, int i, long c, List<BiFunction<Long, Long, Long>> ops) {
        if (i == ns.length - 1) {
            return ns[0] == ops.get(0).apply(c, ns[i]) ||
                   ns[0] == ops.get(1).apply(c, ns[i]) ||
                   (ops.size() > 2 && ns[0] == ops.get(2).apply(c, ns[i]));
        }
        return match(ns, i + 1, ops.get(0).apply(c, ns[i]), ops) ||
               match(ns, i + 1, ops.get(1).apply(c, ns[i]), ops) ||
               (ops.size() > 2 && match(ns, i + 1, ops.get(2).apply(c, ns[i]), ops));
    }
}
