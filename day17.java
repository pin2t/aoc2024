import java.util.*;
import java.util.regex.*;

import static java.lang.Long.parseLong;
import static java.lang.System.in;
import static java.lang.System.out;

public class day17 {
    static final Pattern NUMBER_PTN = Pattern.compile("\\d+");
    public static void main(String[] args) {
        var scanner = new Scanner(in);
        var a = parseLong(new Scanner(scanner.nextLine()).findInLine(NUMBER_PTN));
        var b = parseLong(new Scanner(scanner.nextLine()).findInLine(NUMBER_PTN));
        var c = parseLong(new Scanner(scanner.nextLine()).findInLine(NUMBER_PTN));
        scanner.nextLine();
        var program = new Program(scanner.nextLine());
        out.println(program.run(a, b, c) + " " + program.replicate());
    }
}

class Program {
    final int[] program;

    Program(String s) {
        program = new Scanner(s).findAll(day17.NUMBER_PTN).map(r -> r.group()).mapToInt(Integer::parseInt).toArray();
    }

    String run(long a, long b, long c) {
        var output = new StringBuilder();
        for (int ip = 0; ip < program.length; ip += 2) {
            var combo = switch (program[ip + 1]) {
                case 0, 1, 2, 3 -> program[ip + 1];
                case 4 -> a;
                case 5 -> b;
                case 6 -> c;
                default -> 0L;
            };
            switch (program[ip]) {
            case 0: a = (long) (a / Math.pow(2, (double)combo)); break;
            case 1: b = b ^ program[ip + 1]; break;
            case 2: b = combo % 8; break;
            case 3: if (a > 0) { ip = program[ip + 1]; ip -= 2; } break;
            case 4: b = b ^ c; break;
            case 5:
                if (!output.isEmpty()) { output.append(","); }
                output.append(combo % 8);
                break;
            case 6: b = (long) (a / Math.pow(2, (double)combo)); break;
            case 7: c = (long) (a / Math.pow(2, (double)combo)); break;
            }
        }
        return output.toString();
    }

    long replicate() {
        return runTarget(Arrays.stream(program).boxed().toList(), 0);
    }

    long runTarget(List<Integer> target, long inita) {
        if (target.isEmpty()) { return inita; }
        for (int d = 0; d < 8; d++) {
            var a = (inita << 3) | d;
            long b = 0, c = 0;
            var out = 0L;
            for (int ip = 0; ip < program.length - 2; ip += 2) {
                var combo = switch (program[ip + 1]) {
                    case 0, 1, 2, 3 -> program[ip + 1];
                    case 4 -> a;
                    case 5 -> b;
                    case 6 -> c;
                    default -> 0L;
                };
                switch (program[ip]) {
                case 1: b = b ^ program[ip + 1]; break;
                case 2: b = combo % 8; break;
                case 4: b = b ^ c; break;
                case 5: out = combo % 8; break;
                case 6: b = (long) (a / Math.pow(2, (double)combo)); break;
                case 7: c = (long) (a / Math.pow(2, (double)combo)); break;
                }
            }
            if (out == (long)target.getLast()) {
                var newTarget = new ArrayList<>(target);
                newTarget.removeLast();
                var prev = runTarget(newTarget, a);
                if (prev != -1) return prev;
            }
        }
        return -1;
    }
}
