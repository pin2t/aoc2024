import java.util.*;

import static java.lang.Integer.parseInt;
import static java.lang.System.in;
import static java.lang.System.out;

public class day05 {
    public static void main(String[] args) {
        var scanner = new Scanner(in);
        var rules = new ArrayList<Pair>();
        while (scanner.hasNext()) {
            var line = scanner.nextLine();
            if (line.isEmpty()) break;
            var s = line.split("\\|");
            rules.add(new Pair(parseInt(s[0]), parseInt(s[1])));
        }
        var mid1 = 0; var mid2 = 0;
        while (scanner.hasNext()) {
            var ns = new Numbers(scanner.nextLine(), rules);
            if (ns.isSorted())
                mid1 += ns.middle();
            else {
                ns.sort();
                mid2 += ns.middle();
            }
        }
        out.println(mid1 + " " + mid2);
    }

    record Pair(int a, int b) {}

    static class Numbers {
        final List<Integer> numbers;
        final Comparator<Integer> cmp;

        Numbers(String line, List<Pair> sortingRules) {
            this.numbers = new ArrayList<>(Arrays.stream(line.split(",")).map(Integer::parseInt).toList());
            this.cmp = (n1, n2) -> {
                for (var p : sortingRules) {
                    if (n1.equals(p.a) && n2.equals(p.b)) return -1;
                    if (n1.equals(p.b) && n2.equals(p.a)) return 1;
                }
                return 0;
            };
        }

        int middle() { return numbers.get(numbers.size() / 2); }

        boolean isSorted() {
            for (int i = 0; i < numbers.size() - 1; i++)
                for (int j = i + 1; j < numbers.size(); j++)
                    if (cmp.compare(numbers.get(i), numbers.get(j)) > 0) return false;
            return true;
        }

        void sort() { numbers.sort(cmp); }
    }
}
