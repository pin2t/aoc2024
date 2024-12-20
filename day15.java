import java.util.*;

import static java.lang.System.in;
import static java.lang.System.out;

public class day15 {
    final char[][] map;
    final char[][] wideMap;
    final List<Direction> moves = new ArrayList<>();

    public static void main(String[] args) {
        new day15().main();
    }

    day15() {
        var scanner = new Scanner(in);
        var lines = new ArrayList<String>();
        while (scanner.hasNextLine()) {
            var line = scanner.nextLine();
            if (line.isBlank()) break;
            lines.add(line);
        }
        map = new char[lines.size()][];
        wideMap = new char[lines.size()][];
        for (int i = 0; i < lines.size(); i++) {
            var line = lines.get(i);
            map[i] = line.toCharArray();
            wideMap[i] = new char[2 * line.length()];
            for (int j = 0; j < line.length(); j++) {
                switch (line.charAt(j)) {
                case '.': wideMap[i][2 * j] = wideMap[i][2 * j + 1] = '.'; break;
                case '#': wideMap[i][2 * j] = wideMap[i][2 * j + 1] = '#'; break;
                case '@': wideMap[i][2 * j] = '@'; break;
                case 'O': wideMap[i][2 * j] = '['; wideMap[i][2 * j + 1] = ']'; break;
                }
            }
        }
        while (scanner.hasNextLine()) {
            var line = scanner.nextLine();
            for (int i = 0; i < line.length(); i++) {
                moves.add(Direction.parse(line.charAt(i)));
            }
        }
    }

    void main() {
        Pos robot = robot(map);
        for (var m : moves) {
            var next = robot.move(m);
            if (map[next.row()][next.col()] == 'O' && canMove(next, m))
                move(next, m);
            if (map[next.row()][next.col()] == '.' || map[next.row()][next.col()] == '@')
                robot = next;
        }
        out.print(coords(map));
        robot = robot(wideMap);
        for (var m : moves) {
            var next = robot.move(m);
            if (wideMap[next.row()][next.col()] == '#') continue;
            switch (m) {
            case UP:
            case DOWN:
                if (wideMap[next.row()][next.col()] == ']' && canMoveWide(next.move(Direction.LEFT), m))
                    moveWide(next.move(Direction.LEFT), m);
                if (wideMap[next.row()][next.col()] == '[' && canMoveWide(next, m))
                    moveWide(next, m);
                break;
            case LEFT:
                if (wideMap[next.row()][next.col()] == ']' && canMoveWide(next.move(Direction.LEFT), m))
                    moveWide(next.move(Direction.LEFT), m);
                break;
            case RIGHT:
                if (wideMap[next.row()][next.col()] == '[' && canMoveWide(next, m))
                    moveWide(next, m);
                break;
            }
            if (wideMap[next.row()][next.col()] == '.') {
                wideMap[next.row()][next.col()] = '@';
                wideMap[robot.row()][robot.col()] = '.';
                robot = next;
            }
        }
        out.println(" " + coords(wideMap));
    }

    boolean canMove(Pos p, Direction dir) {
        var next = p.move(dir);
        if (map[next.row()][next.col()] == 'O')
            return canMove(next, dir);
        return map[next.row()][next.col()] != '#';
    }

    void move(Pos p, Direction dir) {
        var next = p.move(dir);
        if (map[next.row()][next.col()] == 'O')
            move(next, dir);
        map[p.row()][p.col()] = '.';
        map[next.row()][next.col()] = 'O';
    }

    boolean canMoveWide(Pos p, Direction dir) {
        if (wideMap[p.row()][p.col()] == '#') { return false; }
        var next = p.move(dir);
        if (wideMap[next.row()][next.col()] == '[' || wideMap[next.row()][next.col()] == ']') {
            if (dir == Direction.UP || dir == Direction.DOWN)
                return canMoveWide(next, dir) && canMoveWide(next.move(Direction.RIGHT), dir);
            else
                return canMoveWide(next.move(dir), dir);
        }
        return wideMap[next.row()][next.col()] != '#';
    }

    void moveWide(Pos pLeft, Direction dir) {
        var next = pLeft.move(dir);
        switch (dir) {
        case UP:
        case DOWN:
            if (wideMap[next.row()][next.col()] == ']')
                moveWide(next.move(Direction.LEFT), dir);
            if (wideMap[next.row()][next.col()] == '[')
                moveWide(next, dir);
            if (wideMap[next.row()][next.col() + 1] == '[')
                moveWide(next.move(Direction.RIGHT), dir);
            break;
        case LEFT:
        case RIGHT:
            next = next.move(dir);
            if (wideMap[next.row()][next.col()] == '[')
                moveWide(next, dir);
            break;
        }
        wideMap[next.row()][next.col()] = '[';
        wideMap[next.row()][next.col() + 1] = ']';
        wideMap[pLeft.row()][pLeft.col()] = '.';
        wideMap[pLeft.row()][pLeft.col() + 1] = '.';
    }

    int coords(char[][] map) {
        var result = 0;
        for (int r = 0; r < map.length; r++)
            for (int c = 0; c < map[r].length; c++)
                if (map[r][c] == '[' || map[r][c] == 'O')
                    result += 100 * r + c;
        return result;
    }

    Pos robot(char[][] map) {
        for (int r = 0; r < map.length; r++)
            for (int c = 0; c < map[r].length; c++)
                if (map[r][c] == '@') return new Pos(r, c);
        throw new IllegalStateException("robot not found");
    }
}

record Pos(int row, int col) {
    Pos move(Direction d) { return new Pos(row + d.row, col + d.col); }
}

enum Direction {
    UP(-1, 0), RIGHT(0, 1), DOWN(1, 0), LEFT(0, -1);
    final int row, col;

    Direction(int row, int col) { this.row = row; this.col = col; }

    static Direction parse(char c) {
        return switch (c) {
        case '^' -> UP;
        case '>' -> RIGHT;
        case 'v' -> DOWN;
        case '<' -> LEFT;
        default -> throw new IllegalArgumentException("invalid direction" + c);
        };
    }

    @Override
    public String toString() {
        return switch (this) {
        case UP -> "^";
        case RIGHT -> ">";
        case DOWN -> "v";
        case LEFT -> "<";
        };
    }
}

