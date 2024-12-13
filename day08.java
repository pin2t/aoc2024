import java.util.*;

import static java.lang.System.in;
import static java.lang.System.out;
import static java.util.Collections.singletonList;

public class day08 {
    public static void main(String[] args) {
        var antennas = new HashMap<Character, List<Pos>>();
        var row = new int[]{0};
        var cols = 0;
        var scanner = new Scanner(in);
        while (scanner.hasNextLine()) {
            var line = scanner.nextLine();
            for (int col = 0; col < line.length(); col++) {
                if (line.charAt(col) == '.') { continue; }
                int fcol = col;
                antennas.compute(line.charAt(col), (a, p) -> {
                    if (p == null) { p = new ArrayList<>(); }
                    p.add(new Pos(row[0], fcol));
                    return p;
                });
            }
            cols = line.length();
            row[0]++;
        }
        out.printf("%d %d%n",
            new Antinodes(row[0], cols, antennas).get(1),
            new Antinodes(row[0], cols, antennas).all()
        );
    }

    static class Antinodes {
        final int rows, cols;
        final HashMap<Character, List<Pos>> antennas;
        final Set<Pos> antinodes = new HashSet<>();

        Antinodes(int rows, int cols, HashMap<Character, List<Pos>> antennas) {
            this.rows = rows;
            this.cols = cols;
            this.antennas = antennas;
        }

        int get(int distance) {
            var prev = antinodes.size();
            for (var an : antennas.values()) {
                for (int i = 0; i < an.size() - 1; i++) {
                    for (int j = i + 1; j < an.size(); j++) {
                        var a1 = new Pos(
                            an.get(i).row - distance * (an.get(j).row - an.get(i).row),
                            an.get(i).col - distance * (an.get(j).col - an.get(i).col)
                        );
                        var a2 = new Pos(
                            an.get(j).row - distance * (an.get(i).row - an.get(j).row),
                            an.get(j).col - distance * (an.get(i).col - an.get(j).col)
                        );
                        if (inside(a1)) antinodes.add(a1);
                        if (inside(a2)) antinodes.add(a2);
                    }
                }
            }
            return antinodes.size() - prev;
        }

        int all() {
            var result = 0;
            for (int d = 0; d < rows; d++) {
                var step = get(d);
                if (step == 0) break;
                result += step;
            }
            return result;
        }

        boolean inside(Pos p) {
            return p.row >= 0 && p.row < rows && p.col >= 0 && p.col < cols;
        }
    }

    record Pos (int row, int col) {}
}
