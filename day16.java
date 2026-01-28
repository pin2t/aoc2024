import java.io.*;
import java.util.*;

import static java.lang.Math.min;
import static java.lang.System.in;
import static java.lang.System.out;
import static java.util.Collections.singleton;

public class day16 {
    public static void main(String[] args) {
        var maze = new BufferedReader(new InputStreamReader(in)).lines().map(String::toCharArray).toList().toArray(new char[][]{});
        Pos start = null, end = null;
        for (int r = 0; r < maze.length; r++) {
            for (int c = 0; c < maze[r].length; c++) {
                if (maze[r][c] == 'S') start = new Pos(r, c);
                if (maze[r][c] == 'E') end = new Pos(r, c);
            }
        }
        record State(Pos p, Direction d, int score, Set<Pos> path) {}
        record DirectionalPos(Pos p, Direction d) {}
        var queue = new PriorityQueue<State>(Comparator.comparingInt(State::score));
        var processed = new HashMap<DirectionalPos, Integer>();
        var seats = new HashSet<Pos>();
        var minscore = 1000000000;
        queue.add(new State(start, Direction.RIGHT, 0, singleton(start)));
        while (!queue.isEmpty()) {
            var s = queue.poll();
            if (s.score > minscore) { continue; }
            var key = new DirectionalPos(s.p, s.d);
            if (processed.containsKey(key) && processed.get(key) < s.score()) { continue; }
            processed.put(key, s.score());
            if (s.p.equals(end)) {
                minscore = min(minscore, s.score());
                seats.addAll(s.path);
            }
            var next = s.p.move(s.d);
            if (maze[next.row()][next.col()] != '#') {
                var path = new HashSet<>(s.path);
                path.add(next);
                queue.add(new State(next, s.d, s.score() + 1, path));
            }
            queue.add(new State(s.p, s.d.toLeft(), s.score() + 1000, s.path));
            queue.add(new State(s.p, s.d.toRight(), s.score() + 1000, s.path));
        }
        out.println(minscore + " " + seats.size());
    }

    record Pos(int row, int col) {
        Pos move(Direction d) { return new Pos(row + d.row, col + d.col); }
    }

    enum Direction {
        UP(-1, 0), RIGHT(0, 1), DOWN(1, 0), LEFT(0, -1);
        final int row, col;

        Direction(int row, int col) { this.row = row; this.col = col; }

        Direction toRight() {
            return switch (this) {
                case UP -> RIGHT;
                case RIGHT -> DOWN;
                case DOWN -> LEFT;
                case LEFT -> UP;
            };
        }

        Direction toLeft() {
            return switch (this) {
                case UP -> LEFT;
                case RIGHT -> UP;
                case DOWN -> RIGHT;
                case LEFT -> DOWN;
            };
        }
    }
}
