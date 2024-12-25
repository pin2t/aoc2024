import java.util.*;

import static java.lang.System.in;
import static java.lang.System.out;

public class day19 {
    static String[] towels;
    public static void main(String[] args) {
        var scanner = new Scanner(in);
        towels = scanner.nextLine().split(", ");
        scanner.nextLine();
        var result1 = 0L; var result2 = 0L;
        while (scanner.hasNextLine()) {
            var p = possibles(scanner.nextLine());
            if (p > 0) result1++;
            result2 += p;
        }
        out.printf("%d %d%n", result1, result2);
    }

    static final Map<String, Long> cache = new HashMap<>();

    static long possibles(String design) {
        if (design.isEmpty()) return 1L;
        if (cache.containsKey(design)) return cache.get(design);
        var result = 0L;
        for (var t : towels) {
            if (t.length() <= design.length() && design.startsWith(t)) {
                result += possibles(design.substring(t.length()));
            }
        }
        cache.put(design, result);
        return result;
    }
}
