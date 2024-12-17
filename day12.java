import java.io.*;
import java.util.*;

import static java.lang.System.in;
import static java.lang.System.out;
import static java.util.Collections.emptySet;

public class day12 {
    public static void main(String[] args) {
        var map = new BufferedReader(new InputStreamReader(in)).lines().map(String::toCharArray).toList().toArray(new char[][]{});
        var processed = new HashSet<Pos>();
        long sum1 = 0, sum2 = 0;
        for (int row = 0; row < map.length; row++) {
            for (int col = 0; col < map[row].length; col++) {
                var p = new Pos(row, col);
                if (processed.contains(p)) { continue; }
                var region = new Region(map, p, processed);
                sum1 += region.area() * region.perimeter();
                sum2 += region.area() * region.sides();
            }
        }
        out.println(sum1 + " " + sum2);
    }
}

class Region {
    final Set<Pos> region =new HashSet<>();
    final char[][] map;
    long perimeter;
    char type;

    Region(char[][] map, Pos p, Set<Pos> processed) {
        this.map = map;
        this.type = map[p.row()][p.col()];
        expand(type, p, processed);
    }

    long area() { return region.size(); }

    long perimeter() { return perimeter; }

    long sides() {
        var sides = 0;
        var outDirections = new HashMap<Pos, Set<Direction>>();
        for (int row = 0; row < map.length; row++) {
            for (int col = 0; col < map[row].length; col++) {
                var pp = new Pos(row, col);
                if (!region.contains(pp)) { continue; }
                for (var d : Direction.values()) {
                    var np = pp.move(d);
                    if (inside(np) && map[np.row()][np.col()] == type) { continue; }
                    var left = pp.move(d.turnLeft());
                    var right = pp.move(d.turnRight());
                    if (!outDirections.getOrDefault(left, emptySet()).contains(d) &&
                            !outDirections.getOrDefault(right, emptySet()).contains(d)) {
                        sides++;
                    }
                    if (!outDirections.containsKey(pp)) {
                        outDirections.put(pp, new HashSet<>());
                    }
                    outDirections.get(pp).add(d);
                }
            }
        }
        return sides;
    }

    void expand(char type, Pos p, Set<Pos> processed) {
        if (!inside(p) || map[p.row()][p.col()] != type) {
            perimeter++;
            return;
        }
        if (!processed.add(p)) { return; }
        region.add(p);
        for (var d : Direction.values()) {
            expand(type, p.move(d), processed);
        }
    }

    boolean inside(Pos p) {
        return p.row() >= 0 && p.row() < map.length && p.col() >= 0 && p.col() < map[p.row()].length;
    }
}

record Pos(int row, int col) {
    Pos move(Direction d) { return new Pos(row + d.dr, col + d.dc); }
}

enum Direction {
    UP(-1, 0), RIGHT(0, 1), DOWN(1, 0), LEFT(0, -1);
    final int dr, dc;

    Direction(int dr, int dc) { this.dr = dr; this.dc = dc; }

    Direction turnRight() {
        return switch (this) {
            case UP -> RIGHT;
            case RIGHT -> DOWN;
            case DOWN -> LEFT;
            case LEFT -> UP;
        };
    }

    Direction turnLeft() {
        return switch (this) {
            case UP -> LEFT;
            case RIGHT -> UP;
            case DOWN -> RIGHT;
            case LEFT -> DOWN;
        };
    }
}
