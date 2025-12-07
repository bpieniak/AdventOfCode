const std = @import("std");

const puzzleInput = @embedFile("input.txt");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    std.debug.print("part1: {d}\n", .{try part1(puzzleInput, allocator)});
    std.debug.print("part2: {d}\n", .{try part2(puzzleInput, allocator)});
}

pub fn part1(input: []const u8, allocator: std.mem.Allocator) !usize {
    var slits_num: usize = 0;

    var it = std.mem.splitScalar(u8, input, '\n');

    const first_line = it.next() orelse return error.Oops;

    const width = first_line.len;
    const start_point = std.mem.indexOfScalar(u8, first_line, 'S') orelse return error.Oops;

    var track_rays: []u8 = try allocator.alloc(u8, width);
    defer allocator.free(track_rays);

    var track_rays_next: []u8 = try allocator.alloc(u8, width);
    defer allocator.free(track_rays_next);

    @memset(track_rays, '.');
    @memset(track_rays_next, '.');

    track_rays[start_point] = '|';

    while (it.next()) |line| {
        @memset(track_rays_next, '.');

        for (line, 0..) |char, i| {
            if (char == '^' and track_rays[i] == '|') {
                // split
                if (i > 0) track_rays_next[i - 1] = '|';
                if (i + 1 < width) track_rays_next[i + 1] = '|';
                slits_num += 1;
            } else if (track_rays[i] == '|') {
                // pass
                track_rays_next[i] = '|';
            }
        }

        @memcpy(track_rays, track_rays_next);
    }

    return slits_num;
}

pub fn part2(input: []const u8, allocator: std.mem.Allocator) !usize {
    var it = std.mem.splitScalar(u8, input, '\n');

    const first_line = it.next() orelse return error.Oops;

    const width = first_line.len;
    const start_point = std.mem.indexOfScalar(u8, first_line, 'S') orelse return error.Oops;

    var track_rays: []usize = try allocator.alloc(usize, width);
    defer allocator.free(track_rays);

    var track_rays_next: []usize = try allocator.alloc(usize, width);
    defer allocator.free(track_rays_next);

    @memset(track_rays, 0);
    @memset(track_rays_next, 0);

    // track how many timelines (rays) occupy each column as we march downward
    track_rays[start_point] = 1;

    while (it.next()) |line| {
        @memset(track_rays_next, 0);

        for (line, 0..) |char, i| {
            const rays_here = track_rays[i];
            if (rays_here == 0) continue;

            if (char == '^') {
                if (i > 0) track_rays_next[i - 1] += rays_here;
                if (i + 1 < width) track_rays_next[i + 1] += rays_here;
            } else {
                track_rays_next[i] += rays_here;
            }
        }

        @memcpy(track_rays, track_rays_next);
    }

    var timelines: usize = 0;
    for (track_rays) |count| timelines += count;
    return timelines;
}

const example_input =
    \\.......S.......
    \\...............
    \\.......^.......
    \\...............
    \\......^.^......
    \\...............
    \\.....^.^.^.....
    \\...............
    \\....^.^...^....
    \\...............
    \\...^.^...^.^...
    \\...............
    \\..^...^.....^..
    \\...............
    \\.^.^.^.^.^...^.
    \\...............
;

test "part1" {
    const result = try part1(example_input, std.testing.allocator);
    try std.testing.expectEqual(21, result);
}

test "part2" {
    const result = try part2(example_input, std.testing.allocator);
    try std.testing.expectEqual(40, result);
}
