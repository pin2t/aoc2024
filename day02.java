import java.io.*;
import java.util.*;

import static java.lang.System.in;
import static java.lang.System.out;

public class day02 {
    public static void main(String[] args) {
        var safes = new int[]{0, 0};
        new BufferedReader(new InputStreamReader(in)).lines().forEach(line -> {
            var report = new Report(line);
            if (report.safe()) {
                safes[0]++;
                safes[1]++;
            } else if (report.canBeSafe()) {
                safes[1]++;
            }
        });
        out.println(safes[0] + " " + safes[1]);
    }

    static class Report extends ArrayList<Integer> {
        Report(String input) {
            super(Arrays.stream(input.split("\\s")).map(Integer::valueOf).toList());
        }

        Report(List<Integer> source, int except) {
            super(source);
            remove(except);
        }

        boolean safe() {
            boolean increasing = getFirst().compareTo(getLast()) < 0;
            for (int i = 1; i < size(); i++) {
                if (increasing && (get(i) - get(i - 1) <= 0 || get(i) - get(i - 1) > 3) ||
                    !increasing && (get(i) - get(i - 1) >= 0 || get(i - 1) - get(i) > 3)) {
                    return false;
                }
            }
            return true;
        }

       boolean canBeSafe() {
            for (int i = 0; i < size(); i++)
                if (new Report(this, i).safe()) return true;
            return false;
        }
    }
}
