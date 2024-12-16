import java.util.*;
import java.util.regex.*;

import static java.lang.Long.parseLong;
import static java.lang.System.in;
import static java.lang.System.out;

public class day11 {
    public static void main(String[] args) {
        var input = new Scanner(in).nextLine();
        var stones = new Stones(input);
        for (int i = 1; i <= 25; i++) stones.change();
        out.print(stones.count());
        stones = new Stones(input);
        for (int i = 1; i <= 75; i++) stones.change();
        out.println(" " + stones.count());
    }
}

class Stones {
    static final Pattern re = Pattern.compile("\\d+");
    Map<Long, Long> numbers = new HashMap<>();

    Stones(String input) {
        for (var m : new Scanner(input).findAll(re).toList()) {
            numbers.merge(parseLong(m.group()), 1L, Long::sum);
        }
    }

    void change() {
        var changed = new HashMap<Long, Long>();
        for (var e : numbers.entrySet()) {
            var n = e.getKey();
            var count = e.getValue();
            if (n == 0) {
                changed.merge(1L, count, Long::sum);
            } else {
                var s = n.toString();
                if (s.length() % 2 == 0) {
                    changed.merge(parseLong(s.substring(0, s.length() / 2)), count, Long::sum);
                    changed.merge(parseLong(s.substring(s.length() / 2)), count, Long::sum);
                } else {
                    changed.merge(n * 2024, count, Long::sum);
                }
            }
        }
        numbers = changed;
    }

    long count() {
        return numbers.values().stream().mapToLong(Long::valueOf).sum();
    }
}