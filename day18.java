import java.io.*;
import java.util.*;

import static java.lang.Integer.parseInt;
import static java.lang.System.in;
import static java.lang.System.out;
import static java.util.Comparator.comparingInt;

public class day18 {
    public static void main(String[] args) {
        var bytes = new ArrayList<Pos>();
        new BufferedReader(new InputStreamReader(in)).lines().forEach(l -> {
            bytes.add(Pos.parse(l));
        });
        out.print(new Path(bytes.subList(0, 1024), 70).steps());
        var l = 0; var r = bytes.size() - 1;
        while (r > l) {
            var m = (r + l) / 2;
            if (new Path(bytes.subList(0, m), 70).isExists()) {
                l = m + 1;
            } else {
                r = m - 1;
            }
        }
        out.println(" " + bytes.get(r - 1).x() + "," + bytes.get(r - 1).y());
    }
}

class Path {
    int steps = -1;

    Path(List<Pos> corrupted, int size) {
        record State(Pos p, int steps) {}
        var queue = new PriorityQueue<State>(comparingInt(State::steps));
        queue.offer(new State(new Pos(0, 0), 0));
        var end = new Pos(size, size);
        var processed = new HashSet<Pos>();
        while (!queue.isEmpty()) {
            var s = queue.poll();
            if (s.p.equals(end)) { steps = s.steps; break; }
            if (!processed.add(s.p)) continue;
            for (var d : Direction.values()) {
                var next = s.p.move(d);
                if (next.x() >= 0 && next.x() <= size && next.y() >= 0 && next.y() <= size && !corrupted.contains(next)) {
                    queue.offer(new State(next, s.steps + 1));
                }
            }
        }
    }

    boolean isExists() { return steps > 0; }

    int steps() { return steps; }
}

record Pos(int x, int y) {
    static Pos parse(String s) {
        var pair = s.split(",");
        return new Pos(parseInt(pair[0]), parseInt(pair[1]));
    }
    Pos move(Direction d) { return new Pos(x + d.x, y + d.y); }
}

enum Direction {
    UP(-1, 0), RIGHT(0, 1), DOWN(1, 0), LEFT(0, -1);
    final int x, y;

    Direction(int x, int y) { this.x = x; this.y = y; }
}
