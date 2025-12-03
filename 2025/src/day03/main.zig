const std = @import("std");

const puzzleInput = @embedFile("input.txt");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    std.debug.print("part1: {d}\n", .{try part1(puzzleInput, allocator)});
    std.debug.print("part2: {d}\n", .{try part2(puzzleInput, allocator)});
}

pub fn part1(input: []const u8, allocator: std.mem.Allocator) !i64 {
    var sum: i64 = 0;
    var it = std.mem.splitAny(u8, input, "\r\n");
    while (it.next()) |x| {
        if (x.len == 0) continue;
        sum += try FindMaxKJoltage(x, 2, allocator);
    }
    return sum;
}

pub fn part2(input: []const u8, allocator: std.mem.Allocator) !i64 {
    var sum: i64 = 0;
    var it = std.mem.splitAny(u8, input, "\r\n");
    while (it.next()) |x| {
        if (x.len == 0) continue;
        sum += try FindMaxKJoltage(x, 12, allocator);
    }
    return sum;
}

pub fn FindMaxKJoltage(bank: []const u8, k: usize, allocator: std.mem.Allocator) !i64 {
    if (k == 0) {
        return -1;
    }

    var digits_u8 = try allocator.alloc(u8, bank.len);
    defer allocator.free(digits_u8);

    var count: usize = 0;
    for (bank) |char| {
        if (!std.ascii.isDigit(char)) continue;
        // convert an ASCII char to integer
        digits_u8[count] = char - '0';
        count += 1;
    }

    if (count < k) {
        return -1;
    }

    var keep = try allocator.alloc(u8, k);
    defer allocator.free(keep);

    var top: usize = 0;
    var to_drop = count - k;
    for (digits_u8[0..count]) |d| {
        while (top > 0 and to_drop > 0 and keep[top - 1] < d) {
            top -= 1;
            to_drop -= 1;
        }
        if (top < k) {
            keep[top] = d;
            top += 1;
        } else if (to_drop > 0) {
            // Already have k digits; drop the current one if we still must drop.
            to_drop -= 1;
        }
    }

    var max_joltage: i64 = 0;
    for (keep[0..k]) |digit| {
        max_joltage = max_joltage * 10 + digit;
    }

    return max_joltage;
}

const example_input =
    \\987654321111111
    \\811111111111119
    \\234234234234278
    \\818181911112111
;

test "part1" {
    const result = try part1(example_input, std.testing.allocator);
    try std.testing.expectEqual(357, result);
}

test "part2" {
    const result = try part2(example_input, std.testing.allocator);
    try std.testing.expectEqual(3121910778619, result);
}
