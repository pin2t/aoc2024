import java.io.*;
import java.util.*;
import java.util.concurrent.atomic.*;

import static java.lang.System.in;
import static java.lang.System.out;

public class day10 {
    public static void main(String[] args) {
        new day10().main();
    }

    day10() {
        var lines = new BufferedReader(new InputStreamReader(in)).lines().toList();
        map = new int[lines.size()][];
        for (int i = 0; i < lines.size(); i++) {
            var l = lines.get(i);
            map[i] = new int[l.length()];
            for (int j = 0; j < l.length(); j++) {
                map[i][j] = l.charAt(j) - '0';
            }
        }
    }

    void main() {
        for (int r = 0; r < map.length; r++) {
            for (int c = 0; c < map[r].length; c++) {
                if (map[r][c] != 0) { continue; }
                visited.clear();
                search(new Pos(r, c));
            }
        }
        out.println(scores.get() + " " + ratings.get());
    }

    int[][] map;
    final Set<Pos> visited = new HashSet<>();
    final AtomicInteger scores = new AtomicInteger();
    final AtomicInteger ratings = new AtomicInteger();

    void search(Pos p) {
        if (map[p.row()][p.col()] == 9) {
            if (!visited.contains(p)) scores.incrementAndGet();
            ratings.incrementAndGet();
            visited.add(p);
        }
        for (var d : Direction.values()) {
            var next = p.move(d);
            if (next.row() >= 0 && next.row() < map.length && next.col() >= 0 && next.col() < map[next.row()].length &&
                map[next.row()][next.col()] - map[p.row()][p.col()] == 1) {
                search(next);
            }
        }
    }
}

enum Direction {
    UP(-1, 0), RIGHT(0, 1), DOWN(1, 0), LEFT(0, -1);
    final int dr, dc;

    Direction(int dr, int dc) { this.dr = dr; this.dc = dc; }
}

record Pos (int row, int col) {
    Pos move(Direction dir) {
        return new Pos(row + dir.dr, col + dir.dc);
    }
}
