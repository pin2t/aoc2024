import static java.lang.IO.*;
import static java.lang.String.format;

enum Op { AND, OR, XOR }

static Map<String, Gate> gates = new HashMap<>();
static Map<String, Boolean> values = new HashMap<>();

record Gate(Op op, String in1, String in2) {
    boolean value() {
        switch (op) {
            case AND: return values.get(in1) && values.get(in2);
            case OR:  return values.get(in1) || values.get(in2);
            case XOR: return values.get(in1) ^ values.get(in2);
        }
        return false;
    }

    private boolean call(String name, int bit, BiFunction<Gate, Integer, Boolean> gateFunc) {
        if (!gates.containsKey(name)) {
            return false;
        }
        return gateFunc.apply(gates.get(name), bit);
    }

    boolean isAddition(int bit) {
        if (bit == 0) {
            return op == Op.XOR && (in1.equals("x00") && in2.equals("y00") || in1.equals("y00") && in2.equals("x00"));
        }
        return op == Op.XOR && (call(in1, bit, Gate::isIntermediate) && call(in2, bit, Gate::isCarry) ||
                call(in1, bit, Gate::isCarry) && call(in2, bit, Gate::isIntermediate));
    }

    boolean isIntermediate(int bit) {
        return op == Op.XOR && (in1.equals(format("x%02d", bit)) || in2.equals(format("y%02d", bit)) ||
                in1.equals(format("y%02d", bit)) || in2.equals(format("x%02d", bit)));
    }

    boolean isCarry(int bit) {
        if (bit == 1) {
            return op == Op.AND && (in1.equals("x00") && in2.equals("y00") || in1.equals("y00") && in2.equals("x00"));
        }
        return op == Op.OR && (call(in1, bit - 1, Gate::isDirectCarry) && call(in2, bit - 1, Gate::isRecarry) ||
                call(in1, bit - 1, Gate::isRecarry) && call(in2, bit - 1, Gate::isDirectCarry));
    }

    boolean isDirectCarry(int bit) {
        return op == Op.AND && (in1.equals(format("x%02d", bit)) || in2.equals(format("y%02d", bit)) ||
                in1.equals(format("y%02d", bit)) || in2.equals(format("x%02d", bit)));
    }

    boolean isRecarry(int bit) {
        return op == Op.AND && (call(in1, bit, Gate::isIntermediate) && call(in2, bit, Gate::isCarry) ||
                call(in1, bit, Gate::isCarry) && call(in2, bit, Gate::isIntermediate));
    }
}

int progress() {
    int bit = 0;
    while (gates.containsKey(format("z%02d", bit)) && gates.get(format("z%02d", bit)).isAddition(bit)) {
        bit++;
    }
    return bit;
}

void main() {
    var line = readln();
    while (!line.isBlank()) {
        var s = line.split(":\\s");
        values.put(s[0], "1".equals(s[1]));
        line = readln();
    }
    line = readln();
    while (line != null) {
        var s = line.split("\\s");
        gates.put(s[4], new Gate(Op.valueOf(s[1]), s[0], s[2]));
        line = readln();
    }
    var zs = gates.keySet().stream().filter(k -> k.charAt(0) == 'z').count();
    while (values.keySet().stream().filter(k -> k.charAt(0) == 'z').count() < zs) {
        for (var g : gates.entrySet()) {
            if (values.containsKey(g.getValue().in1) && values.containsKey(g.getValue().in2)) {
                values.put(g.getKey(), g.getValue().value());
            }
        }
    }
    var zv = 0L;
    for (int i = 0; i < 64; i++) {
        zv |= (values.getOrDefault(format("z%02d", i), false) ? 1L : 0L) << i;
    }
    var swaps = new ArrayList<String>();
    while (swaps.size() < 8) {
        var base = progress();
        outer:
        for (var ea : gates.entrySet()) {
            for (var eb : gates.entrySet()) {
                if (ea.getKey().equals(eb.getKey())) { continue; }
                var a = ea.getValue();
                var b = eb.getValue();
                gates.put(ea.getKey(), b);
                gates.put(eb.getKey(), a);
                if (progress() > base) {
                    swaps.add(ea.getKey());
                    swaps.add(eb.getKey());
                    break outer;
                }
                gates.put(ea.getKey(), a);
                gates.put(eb.getKey(), b);
            }
        }
    }
    swaps.sort(Comparator.naturalOrder());
    println(zv + " " + String.join(",", swaps));
}
