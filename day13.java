import java.util.*;
import java.util.regex.*;

import static java.lang.Long.parseLong;
import static java.lang.System.in;
import static java.lang.System.out;

public class day13 {
    public static void main(String[] args) {
        var scanner = new Scanner(in);
        var numbers = Pattern.compile(".*?(\\d+).*?(\\d+)");
        var totals = new long[]{0, 0};
        while (scanner.hasNextLine()) {
            var line = scanner.nextLine();
            if (line.isBlank()) { continue; }
            var m = numbers.matcher(line); m.matches();
            long ax = parseLong(m.group(1)), ay = parseLong(m.group(2));
            m = numbers.matcher(scanner.nextLine()); m.matches();
            long bx = parseLong(m.group(1)), by = parseLong(m.group(2));
            m = numbers.matcher(scanner.nextLine()); m.matches();
            long px = parseLong(m.group(1)), py = parseLong(m.group(2));
            totals[0] += spent(ax, ay, bx, by, px, py);
            totals[1] += spent(ax, ay, bx, by, px + 10000000000000L, py + 10000000000000L);
        }
        out.println(totals[0] + " " + totals[1]);
    }

    static long spent(long ax, long ay, long bx, long by, long px, long py) {
        var btimes = (py * ax - px * ay) / (ax * by - bx * ay);
        var atimes = (px - btimes * bx) / ax;
        if (atimes * ax + btimes * bx != px || atimes * ay + btimes * by != py) {
            return 0;
        }
        return atimes * 3 + btimes;
    }
}
