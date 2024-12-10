import java.io.*;
import java.util.*;

import static java.lang.System.in;
import static java.lang.System.out;

public class day04 {
    public static void main(String[] args) {
        var world = new World(new BufferedReader(new InputStreamReader(in)).lines().toList());
        out.println(world.xmasses() + " " + world.crosses());
    }

    static class World {
        static final char[] XMAS = "XMAS".toCharArray();
        static final char[] MAS = "MAS".toCharArray();
        static final List<Direction> DIRS = List.of(new Direction(1, 0), new Direction(1, 1), new Direction(0, 1), new Direction(-1, 1),
            new Direction(-1, 0), new Direction(-1, -1), new Direction(0, -1), new Direction(1, -1));
        
        final List<String> world;

        World(List<String> lines) { this.world = lines; }

        boolean match(char[] word, int r, int c, int dr, int dc) {
            for (int i = 0; i < word.length; i++) {
                var rr = r + i * dr;
                var cc = c + i * dc;
                if (rr < 0 || rr >= world.size() || cc < 0 || cc >= world.get(r).length() || world.get(rr).charAt(cc) != word[i]) {
                    return false;
                }
            }
            return true;
        }

        int xmasses() {
            var xmases = 0;
            for (int r = 0; r < world.size(); r++) {
                for (int c = 0; c < world.get(r).length(); c++) {
                    for (var d : DIRS) {
                        if (match(XMAS, r, c, d.dr, d.dc)) xmases++;
                    }
                }
            }
            return xmases;
        }

        int crosses() {
            var crosses = 0;
            for (int r = 0; r < world.size(); r++) {
                for (int c = 0; c < world.get(r).length(); c++) {
                    if ((match(MAS, r - 1, c - 1, 1, 1) || match(MAS, r + 1, c + 1, -1, -1)) &&
                        (match(MAS, r - 1, c + 1, 1, -1) || match(MAS, r + 1, c - 1, -1, 1))) {
                        crosses++;
                    }
                }
            }
            return crosses;
        }
    }

    record Direction(int dr, int dc) {};
}
