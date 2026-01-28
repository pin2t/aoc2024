import java.io.*;
import java.util.*;

import static java.lang.Math.abs;
import static java.lang.System.in;
import static java.lang.System.out;
import static java.util.Comparator.comparingInt;

public class day20 {
    public static void main(String[] args) {
        var track = new BufferedReader(new InputStreamReader(in))
            .lines().map(String::toCharArray).toList().toArray(new char[][]{});
        Pos start = null, end = null;
        for (int row = 0; row < track.length; row++) {
            for (int col = 0; col < track[row].length; col++) {
                switch (track[row][col]) {
                case 'S': start = new Pos(row, col); break;
                case 'E': end = new Pos(row, col); break;
                }
            }
        }
        record State(Pos p, int ps, State prev) {}
        var queue = new PriorityQueue<State>(comparingInt(State::ps));
        var processed = new HashSet<Pos>();
        var path = new ArrayList<Pos>();
        queue.offer(new State(start, 0, null));
        while (!queue.isEmpty()) {
            var s = queue.poll();
            if (s.p.equals(end)) {
                while (s != null) {
                    path.add(s.p);
                    s = s.prev;
                }
                break;
            }
            if (!processed.add(s.p)) continue;
            for (var d : Direction.values()) {
                var next = s.p.move(d);
                if (track[next.row()][next.col()] != '#') {
                    queue.offer(new State(next, s.ps + 1, s));
                }
            }
        }
        var n2 = 0; var n20 = 0;
        for (int i = 0; i < path.size() - 100; i++) {
            for (int j = i + 100; j < path.size(); j++) {
                var len = abs(path.get(i).row() - path.get(j).row()) + abs(path.get(i).col() - path.get(j).col());
                if (len <= 20 && j - i - len >= 100) n20++;
                if (len <= 2 && j - i - len >= 100) n2++;
            }
        }
        out.println(n2 + " " + n20);
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
}

