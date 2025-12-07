const std = @import("std");

const puzzleInput = @embedFile("input.txt");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    std.debug.print("part1: {d}\n", .{try part1(puzzleInput, allocator)});
    std.debug.print("part2: {d}\n", .{try part2(puzzleInput, allocator)});
}

const ID_range = struct {
    from: usize,
    to: usize,

    pub fn init(line: []const u8) !ID_range {
        const dash_idx = std.mem.indexOfScalar(u8, line, '-') orelse return error.InvalidRange;
        const from = try std.fmt.parseInt(usize, line[0..dash_idx], 10);
        const to = try std.fmt.parseInt(usize, line[dash_idx + 1 ..], 10);
        return .{ .from = from, .to = to };
    }
};

pub fn part1(input: []const u8, allocator: std.mem.Allocator) !i64 {
    var ranges = std.ArrayList(ID_range).empty;
    defer ranges.deinit(allocator);

    var fresh_count: i64 = 0;
    var parsing_ranges = true;

    var it = std.mem.splitAny(u8, input, "\r\n");
    while (it.next()) |line| {
        if (line.len == 0) {
            parsing_ranges = false;
            continue;
        }

        const trimmed = std.mem.trim(u8, line, " ");

        if (parsing_ranges) {
            try ranges.append(allocator, try ID_range.init(trimmed));
            continue;
        }

        const id = try std.fmt.parseInt(usize, trimmed, 10);
        for (ranges.items) |range| {
            if (id >= range.from and id <= range.to) {
                fresh_count += 1;
                break;
            }
        }
    }

    return fresh_count;
}

pub fn part2(input: []const u8, allocator: std.mem.Allocator) !i64 {
    var ranges = std.ArrayList(ID_range).empty;
    defer ranges.deinit(allocator);

    var it = std.mem.splitAny(u8, input, "\r\n");
    while (it.next()) |line| {
        if (line.len == 0) {
            break;
        }

        const trimmed = std.mem.trim(u8, line, " ");

        const range = try ID_range.init(trimmed);

        try ranges.append(allocator, range);
    }

    if (ranges.items.len == 0) return 0;

    std.mem.sort(ID_range, ranges.items, {}, struct {
        fn lessThan(_: void, a: ID_range, b: ID_range) bool {
            if (a.from == b.from) return a.to < b.to;
            return a.from < b.from;
        }
    }.lessThan);

    var total: usize = 0;
    var current = ranges.items[0];
    for (ranges.items[1..]) |r| {
        if (r.from <= current.to + 1) {
            if (r.to > current.to) current.to = r.to;
            continue;
        }

        total += current.to - current.from + 1;
        current = r;
    }
    total += current.to - current.from + 1;

    return @as(i64, @intCast(total));
}

const example_input =
    \\3-5
    \\10-14
    \\16-20
    \\12-18
    \\
    \\1
    \\5
    \\8
    \\11
    \\17
    \\32
;

test "part1" {
    const result = try part1(example_input, std.testing.allocator);
    try std.testing.expectEqual(3, result);
}

test "part2" {
    const result = try part2(example_input, std.testing.allocator);
    try std.testing.expectEqual(14, result);
}
