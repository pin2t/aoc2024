import java.io.*;
import java.util.*;
import java.util.regex.*;

import static java.lang.Integer.parseInt;
import static java.lang.System.in;
import static java.lang.System.out;
import static java.util.regex.Pattern.compile;

public class day14 {
    public static void main(String[] args) {
        var robots = new ArrayList<Robot>();
        new BufferedReader(new InputStreamReader(in)).lines().forEach(line -> {
            robots.add(new Robot(line));
        });
        for (int second = 1; ; second++) {
            int q1 = 0, q2 = 0, q3 = 0, q4 = 0;
            for (var r : robots) {
                r.move();
                switch (r.quadrant()) {
                case 1: q1++; break;
                case 2: q2++; break;
                case 3: q3++; break;
                case 4: q4++; break;
                }
            }
            if (second == 100) out.print(q1 * q2 * q3 * q4);
            if (q1 > robots.size() / 2 || q2 > robots.size() / 2 || q3 > robots.size() / 2 || q4 > robots.size() / 2) {
                out.println(" " + second);
                break;
            }
        }
    }
}

class Robot {
    static Pattern numbers = compile(".*?(-?\\d+),(\\d+).*?(-?\\d+),(-?\\d+)");
    static int sizex = 101;
    static int sizey = 103;
    int x, y, vx, vy;

    public Robot(String s) {
        var m = numbers.matcher(s);
        m.matches();
        this.x = parseInt(m.group(1));
        this.y = parseInt(m.group(2));
        this.vx = parseInt(m.group(3));
        this.vy = parseInt(m.group(4));
    }

    void move() {
        x += vx; if (x >= sizex) x -= sizex; else if (x < 0) x += sizex;
        y += vy; if (y >= sizey) y -= sizey; else if (y < 0) y += sizey;
    }

    int quadrant() {
        if (x < sizex / 2) {
            if (y < sizey / 2)
                return 1;
            else if (y > sizey / 2)
                return 3;
        } else if (x > sizex / 2) {
            if (y < sizey / 2)
                return 2;
            else if (y > sizey / 2)
                return 4;
        }
        return 0;
    }
}
