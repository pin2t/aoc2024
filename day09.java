import java.util.*;

import static java.lang.System.in;
import static java.lang.System.out;

public class day09 {
    public static void main(String[] args) {
        var scanner = new Scanner(in);
        var map = scanner.nextLine();
        var disk = new Disk(map);
        disk.compactBLocks();
        out.print(disk.checksum());
        disk = new Disk(map);
        disk.compactFiles();
        out.println(" " + disk.checksum());
    }
}

class Disk {
    static final int EMPTY = -1;
    final int[] blocks;

    Disk(String map) {
        int id = 0;
        var bl =  new ArrayList<Integer>();
        for (int i = 0; i < map.length(); i++) {
            var length = map.charAt(i) - '0';
            if (i % 2 == 0) {
                for (int j = 0; j < length; j++) bl.add(id);
                id++;
            } else {
                for (int j = 0; j < length; j++) bl.add(-1);
            }
        }
        this.blocks = new int[bl.size()];
        for (int i = 0; i < bl.size(); i++) this.blocks[i] = bl.get(i);
    }

    void compactBLocks() {
        var empty = 0;
        var block = blocks.length - 1;
        while (empty < block) {
            for (;block >= 0 && blocks[block] == EMPTY; block--);
            for (;empty < block && blocks[empty] != EMPTY; empty++);
            if (empty < block) {
                blocks[empty] = blocks[block];
                blocks[block] = EMPTY;
                block--;
            }
        }
    }

    void compactFiles() {
        var emptyFirst = 0; var emptyLast = 0;
        var fileFirst = blocks.length - 1; var fileLast = blocks.length - 1;
        while (emptyLast < fileFirst) {
            for (;fileLast >= 0 && blocks[fileLast] == EMPTY; fileLast--);
            for (fileFirst = fileLast; fileFirst >= 0 && blocks[fileFirst] == blocks[fileLast]; fileFirst--);
            fileFirst++;
            do {
                for (; emptyFirst < fileFirst && blocks[emptyFirst] != EMPTY; emptyFirst++);
                for (emptyLast = emptyFirst; emptyLast < fileFirst && blocks[emptyLast] == EMPTY; emptyLast++);
                emptyLast--;
                if (emptyLast < fileFirst && emptyLast - emptyFirst >= fileLast - fileFirst) {
                    for (int i = 0; i <= fileLast - fileFirst; i++) {
                        blocks[emptyFirst + i] =  blocks[fileFirst + i];
                        blocks[fileFirst + i] = EMPTY;
                    }
                    break;
                }
                emptyFirst = emptyLast + 1;
            } while (emptyFirst < fileFirst);
            fileLast = fileFirst - 1;
            emptyFirst = emptyLast = 0;
        }
    }

    long checksum() {
        long checksum = 0L;
        for (int i = 0; i < blocks.length; i++) {
            if (blocks[i] != EMPTY)
                checksum += (long) i * blocks[i];
        }
        return checksum;
    }
}
