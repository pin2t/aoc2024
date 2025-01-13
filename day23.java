import java.io.*;
import java.util.*;

import static java.lang.System.in;
import static java.lang.System.out;

public class day23 {
    public static void main(String[] args) {
        Map<String, Set<String>> connections = new HashMap<>();
        new BufferedReader(new InputStreamReader(in)).lines().forEach(line -> {
            var c = line.split("-");
            connections.computeIfAbsent(c[0], _ -> new HashSet<>()).add(c[1]);
            connections.computeIfAbsent(c[1], _ -> new HashSet<>()).add(c[0]);
        });
        var tconnections = 0;
        var clist = connections.keySet().stream().toList();
        for (int i = 0; i < clist.size() - 2; i++) {
            for (int j = i + 1; j < clist.size() - 1; j++) {
                for (int k = j + 1; k < clist.size(); k++) {
                    if (connections.get(clist.get(i)).contains(clist.get(j)) &&
                        connections.get(clist.get(j)).contains(clist.get(k)) &&
                        connections.get(clist.get(i)).contains(clist.get(k)) &&
                        (clist.get(i).startsWith("t") || clist.get(j).startsWith("t") || clist.get(k).startsWith("t"))) {
                        tconnections++;
                    }
                }
            }
        }
        Set<String> largest = new HashSet<>();
        for (var c : connections.keySet()) {
            for (var other : connections.get(c)) {
                Set<String> lan = new HashSet<>();
                lan.add(c);
                lan.add(other);
                for (var second : connections.get(c)) {
                    if (!lan.contains(second) && connections.get(second).containsAll(lan)) {
                        lan.add(second);
                    }
                }
                if (lan.size() > largest.size()) largest = lan;
            }
        }
        out.println(tconnections + " " + String.join(",", largest.stream().sorted().toList()));
    }
}
