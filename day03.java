import java.io.*;
import java.util.*;
import java.util.regex.*;

import static java.lang.Integer.parseInt;
import static java.lang.System.in;
import static java.lang.System.out;

public class day03 {
    public static void main(String[] args) {
        var cmd = Pattern.compile("mul\\((\\d+),(\\d+)\\)|do\\(\\)|don't\\(\\)");
        var muls = new int[]{0, 0};
        var _do = new boolean[]{true};
        new BufferedReader(new InputStreamReader(in)).lines().forEach(line -> {
            new Scanner(line).findAll(cmd).forEach(match -> {
                if (match.group(0).startsWith("mul")) {
                    var n1 = parseInt(match.group(1));
                    var n2 = parseInt(match.group(2));
                    muls[0] += n1 * n2;
                    if (_do[0]) { muls[1] += n1 * n2; }
                } else {
                    _do[0] = "do()".equals(match.group(0));
                }
            });
        });
        out.printf("%d %d\n", muls[0], muls[1]);
    }
}
