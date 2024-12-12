import java.io.*;
import java.util.*;

import static java.lang.System.in;
import static java.lang.System.out;

public class day06 {
    public static void main(String[] args) {
        var map = new BufferedReader(new InputStreamReader(in)).lines().map(String::toCharArray).toList().toArray(new char[][]{});
        var guard = new Guard(map);
        var result = guard.walk();
        out.print(result.visited);
        var loops = 0;
        for (int row = 0; row < map.length; row++) {
            for (int col = 0; col < map[row].length; col++) {
                if (map[row][col] == '.') {
                    map[row][col] = '#';
                    var r = guard.walk();
                    if (r.isLoop()) {
                        loops++;
                    }
                    map[row][col] = '.';
                }
            }
        }
        out.println(" " + loops);
    }

    static class Guard {
        final char[][] map;
        Pos p;

        Guard(char[][] map) {
            this.map = map;
            for (int row = 0; row < map.length; row++) {
                for (int col = 0; col < map[row].length; col++) {
                    if (map[row][col] == '^') {
                        p = new Pos(row, col);
                        break;
                    }
                }
            }
        }

        WalkResult walk() {
            var pos = new Pos(p.row, p.col);
            Set<Pos> visited = new HashSet<>();
            visited.add(pos);
            Set<DirectionalKey> visitedDir = new HashSet<>();
            var dir = Direction.UP;
            visitedDir.add(new DirectionalKey(pos, dir));
            while (true) {
                var next = pos.move(dir);
                if (!inside(next)) {
                    return new WalkResult(false, visited.size());
                }
                if (visitedDir.contains(new DirectionalKey(next, dir))) {
                    return new WalkResult(true, visited.size());
                }
                if (map[next.row][next.col] == '#') {
                    dir = dir.turnRight();
                } else {
                    pos = next;
                }
                visited.add(pos);
                visitedDir.add(new DirectionalKey(pos, dir));
            }
        }

        boolean inside(Pos p) {
            return p.row >= 0 && p.row < map.length && p.col >= 0 && p.col < map[p.row].length;
        }
    }

    enum Direction {
        UP(-1, 0), RIGHT(0, 1), DOWN(1, 0), LEFT(0, -1);

        final int dr, dc;

        Direction(int dr, int dc) {
            this.dr = dr;
            this.dc = dc;
        }

        Direction turnRight() {
            return switch (this) {
                case UP -> RIGHT;
                case RIGHT -> DOWN;
                case DOWN -> LEFT;
                case LEFT -> UP;
            };
        }
    }

    record Pos (int row, int col) {
        Pos move(Direction dir) {
            return new Pos(row + dir.dr, col + dir.dc);
        }
    }

    record DirectionalKey (Pos p, Direction d) {}
    record WalkResult (boolean isLoop, int visited) {}
}
