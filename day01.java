import java.util.*;

import static java.lang.Integer.parseInt;
import static java.lang.Math.abs;
import static java.lang.System.in;
import static java.lang.System.out;

public class day01 {
    public static void main(String[] args) {
        var left = new ArrayList<Integer>();
        var right = new ArrayList<Integer>();
        new Scanner(in).findAll("(\\d+)\\s+(\\d+)").forEach(pair -> {
            left.add(parseInt(pair.group(1)));
            right.add(parseInt(pair.group(2)));
        });
        left.sort(Integer::compareTo);
        right.sort(Integer::compareTo);
        int total = 0, similarity = 0;
        for (int i = 0; i < left.size(); i++) {
            int l = left.get(i);
            total += abs(l - right.get(i));
            similarity += (int) (l * right.stream().filter(n -> n.equals(l)).count());
        }
        out.println(total + " " + similarity);
    }
}
