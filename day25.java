import java.io.*;
import java.util.*;
import java.util.concurrent.atomic.*;

import static java.lang.System.in;
import static java.lang.System.out;

public class day25 {
    public static void main(String[] args) {
        var locks = new HashSet<List<String>>();
        var keys = new HashSet<List<String>>();
        AtomicReference<List<String>> block = new AtomicReference<>(new ArrayList<>());
        new BufferedReader(new InputStreamReader(in)).lines().forEach(line -> {
            if (line.isBlank()) {
                if (block.get().get(0).chars().allMatch(c -> c == '#')) {
                    locks.add(block.get());
                } else {
                    keys.add(block.get());
                }
                block.set(new ArrayList<>());
            } else {
                block.get().add(line);
            }
        });
        if (block.get().get(0).chars().allMatch(c -> c == '#')) {
            locks.add(block.get());
        } else {
            keys.add(block.get());
        }
        var comb = 0;
        for (var l : locks) {
            for (var k : keys) {
                var fit = true;
                for (int i = 0; i < l.size() && fit; i++) {
                    for (int j = 0; j < l.get(i).length() && fit; j++) {
                        fit = l.get(i).charAt(j) != '#' || k.get(i).charAt(j) != '#';
                    }
                }
                if (fit) comb++;
            }
        }
        out.println(comb);
    }
}
