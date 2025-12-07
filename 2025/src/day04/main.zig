const std = @import("std");

const puzzleInput = @embedFile("input.txt");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    std.debug.print("part1: {d}\n", .{try part1(puzzleInput, allocator)});
    std.debug.print("part2: {d}\n", .{try part2(puzzleInput, allocator)});
}

const Grid = struct {
    rows: [][]u8,
    allocator: std.mem.Allocator,

    pub fn init(input: []const u8, allocator: std.mem.Allocator) !Grid {
        const rows = try buildGrid(input, allocator);
        return .{ .rows = rows, .allocator = allocator };
    }

    pub fn deinit(self: *Grid) void {
        for (self.rows) |row| {
            self.allocator.free(row);
        }
        self.allocator.free(self.rows);
    }

    pub fn dump(self: *Grid) void {
        for (self.rows) |row| {
            std.debug.print("{s}\n", .{row});
        }
    }
};

pub fn part1(input: []const u8, allocator: std.mem.Allocator) !i64 {
    var sum: i64 = 0;

    var grid = try Grid.init(input, allocator);
    defer grid.deinit();

    for (grid.rows, 0..) |row, y| {
        for (row, 0..) |_, x| {
            if (grid.rows[y][x] != '@') {
                continue;
            }

            if (countRolls(grid.rows, x, y) < 4) {
                sum += 1;
            }
        }
    }

    return sum;
}

pub fn part2(input: []const u8, allocator: std.mem.Allocator) !i64 {
    var sum: i64 = 0;

    var grid = try Grid.init(input, allocator);
    defer grid.deinit();

    while (true) {
        const removedRolls = removeRolls(grid.rows);
        sum += removedRolls;

        if (removedRolls == 0) {
            break;
        }
    }

    return sum;
}

pub fn removeRolls(diagram: [][]u8) i64 {
    var removed: i64 = 0;
    for (diagram, 0..) |row, y| {
        for (row, 0..) |_, x| {
            if (diagram[y][x] != '@') {
                continue;
            }

            if (countRolls(diagram, x, y) < 4) {
                removed += 1;
                diagram[y][x] = '.';
            }
        }
    }

    return removed;
}

pub fn countRolls(diagram: [][]u8, x: usize, y: usize) u8 {
    var rolls: u8 = 0;

    const y_max = diagram.len;
    const x_max = diagram[y].len;
    const y_i = @as(isize, @intCast(y));
    const x_i = @as(isize, @intCast(x));

    const dirs = [3]isize{ -1, 0, 1 };
    for (dirs) |x_dir| {
        for (dirs) |y_dir| {
            if (x_dir == 0 and y_dir == 0) {
                continue;
            }

            const check_x = x_i + x_dir;
            const check_y = y_i + y_dir;

            if (check_x < 0 or check_y < 0) continue;

            const ux = @as(usize, @intCast(check_x));
            const uy = @as(usize, @intCast(check_y));

            if (ux >= x_max or uy >= y_max) {
                continue;
            }

            if (diagram[uy][ux] == '@') {
                rolls += 1;
            }
        }
    }

    return rolls;
}

pub fn buildGrid(input: []const u8, alloc: std.mem.Allocator) ![][]u8 {
    var lines = std.ArrayList([]u8).empty;
    defer lines.deinit(alloc);

    var it = std.mem.splitScalar(u8, input, '\n');
    while (it.next()) |line| {
        if (line.len == 0) continue;
        try lines.append(alloc, try alloc.dupe(u8, line));
    }

    return try lines.toOwnedSlice(alloc);
}

const example_input =
    \\..@@.@@@@.
    \\@@@.@.@.@@
    \\@@@@@.@.@@
    \\@.@@@@..@.
    \\@@.@@@@.@@
    \\.@@@@@@@.@
    \\.@.@.@.@@@
    \\@.@@@.@@@@
    \\.@@@@@@@@.
    \\@.@.@@@.@.
;

test "part1" {
    const result = try part1(example_input, std.testing.allocator);
    try std.testing.expectEqual(13, result);
}

test "part2" {
    const result = try part2(example_input, std.testing.allocator);
    try std.testing.expectEqual(43, result);
}
