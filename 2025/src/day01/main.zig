const std = @import("std");
const expect = std.testing.expect;

const puzzleInput = @embedFile("input.txt");

pub fn main() !void {
    std.debug.print("day1: {d}\n", .{try part1(puzzleInput)});
}

pub fn part1(input: []const u8) !i64 {
    var location: i32 = 50;
    var password: i32 = 0;

    var it = std.mem.tokenizeScalar(u8, input, '\n');
    while (it.next()) |token| {
        const integer = try std.fmt.parseInt(i32, token[1..], 10);

        switch (token[0]) {
            'L' => {
                location -= integer;
            },
            'R' => {
                location += integer;
            },
            else => {},
        }

        location = @mod(location, 100);

        if (location == 0) {
            password += 1;
        }
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
    try expect(result == 3);
}
