const std = @import("std");

const puzzleInput = @embedFile("input.txt");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    std.debug.print("part1: {d}\n", .{try part1(puzzleInput, allocator)});
    std.debug.print("part2: {d}\n", .{try part2(puzzleInput, allocator)});
}

const Lines = struct {
    rows: []const []const u8,
    width: usize,
};

fn loadLines(input: []const u8, allocator: std.mem.Allocator) !Lines {
    var list = std.ArrayList([]const u8).empty;
    defer list.deinit(allocator);

    var it = std.mem.splitScalar(u8, input, '\n');
    while (it.next()) |raw_line| {
        var line = raw_line;
        if (line.len > 0 and line[line.len - 1] == '\r') line = line[0 .. line.len - 1];
        if (line.len == 0) continue;
        try list.append(allocator, line);
    }

    var width: usize = 0;
    for (list.items) |line| width = @max(width, line.len);

    return .{ .rows = try list.toOwnedSlice(allocator), .width = width };
}

fn charAt(lines: Lines, row: usize, col: usize) u8 {
    const line = lines.rows[row];
    return if (col < line.len) line[col] else ' ';
}

fn columnHasInk(lines: Lines, col: usize) bool {
    for (lines.rows) |line| {
        if (col < line.len and line[col] != ' ') return true;
    }
    return false;
}

fn nextSegment(lines: Lines, start_col: usize) ?struct { start: usize, end: usize } {
    var col = start_col;
    while (col < lines.width and !columnHasInk(lines, col)) col += 1;
    if (col >= lines.width) return null;
    const start = col;
    while (col < lines.width and columnHasInk(lines, col)) col += 1;
    return .{ .start = start, .end = col };
}

fn parseRowValue(lines: Lines, row: usize, start: usize, end: usize) !i128 {
    var value: i128 = 0;
    var found = false;
    var pos = start;
    while (pos < end) : (pos += 1) {
        const ch = charAt(lines, row, pos);
        if (ch >= '0' and ch <= '9') {
            found = true;
            value = value * 10 + @as(i128, @intCast(ch - '0'));
        }
    }
    if (!found) return error.InvalidProblem;
    return value;
}

fn parseColumnValue(lines: Lines, col: usize) !i128 {
    var value: i128 = 0;
    var found = false;
    var row: usize = 0;
    while (row + 1 < lines.rows.len) : (row += 1) {
        const ch = charAt(lines, row, col);
        if (ch >= '0' and ch <= '9') {
            found = true;
            value = value * 10 + @as(i128, @intCast(ch - '0'));
        }
    }
    if (!found) return error.InvalidProblem;
    return value;
}

fn solve(lines: Lines, right_to_left: bool) !i64 {
    if (lines.rows.len == 0) return 0;

    var total: i128 = 0;
    var col: usize = 0;
    while (nextSegment(lines, col)) |seg| {
        col = seg.end;

        var op: u8 = 0;
        var idx = seg.start;
        while (idx < seg.end) : (idx += 1) {
            const ch = charAt(lines, lines.rows.len - 1, idx);
            if (ch != ' ') {
                op = ch;
                break;
            }
        }
        if (op == 0) return error.InvalidProblem;

        var result: i128 = if (op == '+') 0 else 1;

        if (!right_to_left) {
            for (lines.rows[0 .. lines.rows.len - 1], 0..) |_, row| {
                const value = try parseRowValue(lines, row, seg.start, seg.end);
                switch (op) {
                    '+' => result += value,
                    '*' => result *= value,
                    else => return error.UnknownOperation,
                }
            }
        } else {
            var c = seg.end;
            while (c > seg.start) {
                c -= 1;
                const value = try parseColumnValue(lines, c);
                switch (op) {
                    '+' => result += value,
                    '*' => result *= value,
                    else => return error.UnknownOperation,
                }
            }
        }

        total += result;
    }

    return @as(i64, @intCast(total));
}

pub fn part1(input: []const u8, allocator: std.mem.Allocator) !i64 {
    const lines = try loadLines(input, allocator);
    defer allocator.free(lines.rows);

    return solve(lines, false);
}

pub fn part2(input: []const u8, allocator: std.mem.Allocator) !i64 {
    const lines = try loadLines(input, allocator);
    defer allocator.free(lines.rows);

    return solve(lines, true);
}

const example_input =
    \\123 328  51 64 
    \\ 45 64  387 23 
    \\  6 98  215 314
    \\*   +   *   +  
;

test "part1" {
    const result = try part1(example_input, std.testing.allocator);
    try std.testing.expectEqual(4277556, result);
}

test "part2" {
    const result = try part2(example_input, std.testing.allocator);
    try std.testing.expectEqual(3263827, result);
}
