const std = @import("std");
const expect = std.testing.expect;

const puzzleInput = @embedFile("input.txt");

pub fn main() !void {
    std.debug.print("part1: {d}\n", .{try part1(puzzleInput)});
    std.debug.print("part2: {d}\n", .{try part2(puzzleInput)});
}

pub fn part1(input: []const u8) !i64 {
    var location: i32 = 50;
    var password: i32 = 0;

    var it = std.mem.tokenizeScalar(u8, input, '\n');
    while (it.next()) |token| {
        const steps = try std.fmt.parseInt(i32, token[1..], 10);

        switch (token[0]) {
            'L' => {
                location = @mod(location - steps, 100);
            },
            'R' => {
                location = @mod(location + steps, 100);
            },
            else => {},
        }

        if (location == 0) {
            password += 1;
        }
    }

    return password;
}

pub fn part2(input: []const u8) !i64 {
    var location: i32 = 50;
    var password: i32 = 0;

    var it = std.mem.tokenizeScalar(u8, input, '\n');
    while (it.next()) |token| {
        const steps = try std.fmt.parseInt(i32, token[1..], 10);
        var hits: i32 = 0;

        switch (token[0]) {
            'L' => {
                if (location == 0) {
                    hits = @divFloor(steps, 100);
                } else if (steps >= location) {
                    hits = @divFloor(steps - location, 100) + 1;
                }

                location = @mod(location - steps, 100);
            },
            'R' => {
                if (location == 0) {
                    hits = @divFloor(steps, 100);
                } else {
                    const first_zero = 100 - location;
                    if (steps >= first_zero) {
                        hits = @divFloor(steps - first_zero, 100) + 1;
                    }
                }

                location = @mod(location + steps, 100);
            },
            else => {},
        }

        password += hits;
    }

    return password;
}

const example_input =
    \\L68
    \\L30
    \\R48
    \\L5
    \\R60
    \\L55
    \\L1
    \\L99
    \\R14
    \\L82
;

test "part1" {
    const result = try part1(example_input);
    try std.testing.expectEqual(result, 3);
}

test "part2" {
    const result = try part2(example_input);
    try std.testing.expectEqual(6, result);
}
