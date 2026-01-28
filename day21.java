import java.io.*;
import java.util.*;

import static java.lang.Integer.parseInt;
import static java.lang.Math.abs;
import static java.lang.Math.min;
import static java.lang.System.in;
import static java.lang.System.out;

public class day21 {
    public static void main(String[] args) {
        var results = new long[]{0, 0};
        for (var code : new BufferedReader(new InputStreamReader(in)).lines().toList()) {
            var prefix = parseInt(code.substring(0, 3));
            results[0] += new Sequence(code, 2).length() * prefix;
            results[1] += new Sequence(code, 25).length() * prefix;
        }
        out.println(results[0] + " " + results[1]);
    }
}

class Sequence {
    static final Map<Character, Pos> numPad = new HashMap<>() {{
        put('7', new Pos(0, 0));
        put('8', new Pos(0, 1));
        put('9', new Pos(0, 2));
        put('4', new Pos(1, 0));
        put('5', new Pos(1, 1));
        put('6', new Pos(1, 2));
        put('1', new Pos(2, 0));
        put('2', new Pos(2, 1));
        put('3', new Pos(2, 2));
        put('0', new Pos(3, 1));
        put('A', new Pos(3, 2));
    }};
    static final Map<Character, Pos> dirPad = new HashMap<>() {{
        put('^', new Pos(0, 1));
        put('A', new Pos(0, 2));
        put('<', new Pos(1, 0));
        put('v', new Pos(1, 1));
        put('>', new Pos(1, 2));
    }};

    int limit;
    String code;

    record Move(Pos start, Pos end) {}
    record CacheKey(String code, int indirection) {}
    Map<Move, List<String>> moves = new HashMap<>();
    Map<CacheKey, Long> cache = new HashMap<>();

    Sequence(String code, int limit) {
        this.code = code;
        this.limit = limit;
    }

    long length() {
        return length(code, 0, numPad, new Pos(3, 2));
    }

    private long length(String code, int indirection, Map<Character, Pos> pad, Pos from) {
        var ck = new CacheKey(code, indirection);
        var l = cache.get(ck);
        if (l != null) return l;
        var len = 0L;
        for (int i = 0; i < code.length(); i++) {
            var to = pad.get(code.charAt(i));
            List<String> sequences = new ArrayList<>();
            if (indirection > 0) {
                var sk = new Move(from, to);
                sequences = moves.get(sk);
                if (sequences == null) {
                    sequences = new ArrayList<>();
                    generate(sequences, "", from, to, pad);
                    moves.put(sk, sequences);
                }
            } else {
                generate(sequences, "", from, to, pad);
            }
            var m = 1000000000000L;
            for (var s : sequences) {
                m = min(m, indirection == limit ? s.length() : length(s, indirection + 1, dirPad, new Pos(0, 2)));
            }
            len += m;
            from = to;
        }
        cache.put(ck, len);
        return len;
    }

    private void generate(List<String> sequences, String current, Pos from, Pos to, Map<Character, Pos> pad) {
        if (from.equals(to)) {
            sequences.add(current + "A");
            return;
        }
        for (var d : Direction.values()) {
            var next = from.move(d);
            if (next.distance(to) < from.distance(to) && pad.containsValue(next)) {
                generate(sequences, current + d.symbol(), next, to, pad);
            }
        }
    }
    enum Direction {
        UP(-1, 0), RIGHT(0, 1), DOWN(1, 0), LEFT(0, -1);
        final int dr, dc;

        Direction(int dr, int dc) { this.dr = dr; this.dc = dc; }

        char symbol() {
            return switch(this) {
                case UP -> '^';
                case RIGHT -> '>';
                case DOWN -> 'v';
                case LEFT -> '<';
            };
        }
    }

    record Pos (int row, int col) {
        Pos move(Direction dir) {
            return new Pos(row + dir.dr, col + dir.dc);
        }

        int distance(Pos to) {
            return abs(this.row - to.row) + abs(this.col - to.col);
        }
    }
}