const std = @import("std");

const puzzleInput = @embedFile("input.txt");

pub fn main() !void {
    std.debug.print("part1: {d}\n", .{try part1(puzzleInput)});
    std.debug.print("part2: {d}\n", .{try part2(puzzleInput)});
}

pub fn part1(input: []const u8) !i64 {
    var sum: i64 = 0;

    var it = std.mem.splitAny(u8, input, ",");
    while (it.next()) |x| {
        var id_iterator = std.mem.splitAny(u8, x, "-");

        const from_str = id_iterator.next() orelse return error.InvalidRange;
        const to_str = id_iterator.next() orelse return error.InvalidRange;

        var from = try std.fmt.parseInt(i64, from_str, 10);
        const to = try std.fmt.parseInt(i64, to_str, 10);

        while (from <= to) : (from += 1) {
            if (isInvalidId(from)) {
                sum += from;
            }
        }
    }
    return sum;
}

fn isInvalidId(id: i64) bool {
    var buf: [20]u8 = undefined;
    const id_str = std.fmt.bufPrint(&buf, "{d}", .{id}) catch unreachable;

    if (id_str.len % 2 != 0) {
        return false;
    }

    return std.mem.eql(u8, id_str[0 .. id_str.len / 2], id_str[id_str.len / 2 ..]);
}

pub fn part2(input: []const u8) !i64 {
    var sum: i64 = 0;

    var it = std.mem.splitAny(u8, input, ",");
    while (it.next()) |x| {
        var id_iterator = std.mem.splitAny(u8, x, "-");

        const from_str = id_iterator.next() orelse return error.InvalidRange;
        const to_str = id_iterator.next() orelse return error.InvalidRange;

        var from = try std.fmt.parseInt(i64, from_str, 10);
        const to = try std.fmt.parseInt(i64, to_str, 10);

        while (from <= to) : (from += 1) {
            if (isInvalidId2(from)) {
                sum += from;
            }
        }
    }
    return sum;
}

fn isInvalidId2(id: i64) bool {
    var buf: [32]u8 = undefined;
    const id_str = std.fmt.bufPrint(&buf, "{d}", .{id}) catch unreachable;

    const N = id_str.len;

    var L: usize = 1;
    while (L <= N / 2) : (L += 1) {
        if (N % L == 0) {
            const pattern = id_str[0..L];

            const repeats = N / L;
            var i: usize = 1;

            var matches = true;
            while (i < repeats) : (i += 1) {
                const segment_start = i * L;
                const segment = id_str[segment_start .. segment_start + L];

                if (!std.mem.eql(u8, pattern, segment)) {
                    matches = false;
                    break;
                }
            }

            if (matches) {
                return true;
            }
        }
    }

    return false;
}

const example_input =
    \\11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124
;

test "part1" {
    const result = try part1(example_input);
    try std.testing.expectEqual(1227775554, result);
}

test "part2" {
    const result = try part2(example_input, std.testing.allocator);
    try std.testing.expectEqual(4174379265, result);
}
