const std = @import("std");

const puzzleInput = @embedFile("input.txt");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    std.debug.print("part1: {d}\n", .{try part1(puzzleInput, allocator)});
    std.debug.print("part2: {d}\n", .{try part2(puzzleInput, allocator)});
}

const juntion_box = struct {
    x: usize,
    y: usize,
    z: usize,
};

const distance = struct {
    box1: usize,
    box2: usize,
    distance: u64,
};

fn less(_: void, a: distance, b: distance) bool {
    return a.distance < b.distance;
}

pub fn part1(input: []const u8, allocator: std.mem.Allocator) !usize {
    return solve1(input, allocator, 1000);
}

fn solve1(input: []const u8, allocator: std.mem.Allocator, connection_limit: usize) !usize {
    var boxes = try parseBoxes(input, allocator);
    defer allocator.free(boxes);

    var distances = std.ArrayList(distance).empty;
    defer distances.deinit(allocator);

    for (0..boxes.len) |i| {
        for (i + 1..boxes.len) |j| {
            // d = (xi - xj)^2 + (yi - yj)^2 + (zi - zj)^2
            // The actual distance is sqrt of this, but sorting of sqrt will give same order so we can skip it.
            const dx = @as(i64, @intCast(boxes[i].x)) - @as(i64, @intCast(boxes[j].x));
            const dy = @as(i64, @intCast(boxes[i].y)) - @as(i64, @intCast(boxes[j].y));
            const dz = @as(i64, @intCast(boxes[i].z)) - @as(i64, @intCast(boxes[j].z));

            const d = @as(u64, @intCast(dx * dx + dy * dy + dz * dz));

            try distances.append(allocator, .{ .box1 = i, .box2 = j, .distance = d });
        }
    }

    std.sort.block(distance, distances.items, {}, less);

    var parents = try allocator.alloc(usize, boxes.len);
    defer allocator.free(parents);
    var sizes = try allocator.alloc(usize, boxes.len);
    defer allocator.free(sizes);

    for (parents, 0..) |*p, idx| {
        p.* = idx;
        sizes[idx] = 1;
    }

    var edges_used: usize = 0;
    const max_edges = @min(connection_limit, distances.items.len);

    for (distances.items) |edge| {
        if (edges_used >= max_edges) break;
        _ = unite(&parents, &sizes, edge.box1, edge.box2);
        edges_used += 1;
    }

    var component_sizes = std.ArrayList(usize).empty;
    defer component_sizes.deinit(allocator);

    for (parents, 0..) |_, idx| {
        if (parents[idx] == idx) {
            try component_sizes.append(allocator, sizes[idx]);
        }
    }

    const desc = struct {
        fn less(_: void, a: usize, b: usize) bool {
            return a > b;
        }
    };
    std.sort.block(usize, component_sizes.items, {}, desc.less);

    if (component_sizes.items.len < 3) return error.NotEnoughComponents;

    return component_sizes.items[0] * component_sizes.items[1] * component_sizes.items[2];
}

pub fn part2(input: []const u8, allocator: std.mem.Allocator) !usize {
    var boxes = try parseBoxes(input, allocator);
    defer allocator.free(boxes);

    var distances = std.ArrayList(distance).empty;
    defer distances.deinit(allocator);

    for (0..boxes.len) |i| {
        for (i + 1..boxes.len) |j| {
            const dx = @as(i64, @intCast(boxes[i].x)) - @as(i64, @intCast(boxes[j].x));
            const dy = @as(i64, @intCast(boxes[i].y)) - @as(i64, @intCast(boxes[j].y));
            const dz = @as(i64, @intCast(boxes[i].z)) - @as(i64, @intCast(boxes[j].z));

            const d = @as(u64, @intCast(dx * dx + dy * dy + dz * dz));

            try distances.append(allocator, .{ .box1 = i, .box2 = j, .distance = d });
        }
    }

    std.sort.block(distance, distances.items, {}, less);

    var parents = try allocator.alloc(usize, boxes.len);
    defer allocator.free(parents);
    var sizes = try allocator.alloc(usize, boxes.len);
    defer allocator.free(sizes);

    for (parents, 0..) |*p, idx| {
        p.* = idx;
        sizes[idx] = 1;
    }

    var components = boxes.len;

    for (distances.items) |edge| {
        if (unite(&parents, &sizes, edge.box1, edge.box2)) {
            components -= 1;
            const product = boxes[edge.box1].x * boxes[edge.box2].x;
            if (components == 1) return product;
        }
    }

    return error.NotConnected;
}

fn parseBoxes(input: []const u8, allocator: std.mem.Allocator) ![]juntion_box {
    var boxes = std.ArrayList(juntion_box).empty;
    errdefer boxes.deinit(allocator);

    var lines = std.mem.splitScalar(u8, input, '\n');
    while (lines.next()) |line| {
        if (line.len == 0) continue;

        var parts = std.mem.splitScalar(u8, line, ',');
        const x = try std.fmt.parseInt(usize, parts.next() orelse return error.BadLine, 10);
        const y = try std.fmt.parseInt(usize, parts.next() orelse return error.BadLine, 10);
        const z = try std.fmt.parseInt(usize, parts.next() orelse return error.BadLine, 10);

        try boxes.append(allocator, .{ .x = x, .y = y, .z = z });
    }
    return boxes.toOwnedSlice(allocator);
}

fn find(parents: *[]usize, x: usize) usize {
    var root = x;
    while (parents.*[root] != root) {
        root = parents.*[root];
    }
    var node = x;
    while (parents.*[node] != root) {
        const next = parents.*[node];
        parents.*[node] = root;
        node = next;
    }
    return root;
}

fn unite(parents: *[]usize, sizes: *[]usize, a: usize, b: usize) bool {
    var ra = find(parents, a);
    var rb = find(parents, b);
    if (ra == rb) return false;

    if (sizes.*[ra] < sizes.*[rb]) {
        const tmp = ra;
        ra = rb;
        rb = tmp;
    }

    parents.*[rb] = ra;
    sizes.*[ra] += sizes.*[rb];
    return true;
}

const example_input =
    \\162,817,812
    \\57,618,57
    \\906,360,560
    \\592,479,940
    \\352,342,300
    \\466,668,158
    \\542,29,236
    \\431,825,988
    \\739,650,466
    \\52,470,668
    \\216,146,977
    \\819,987,18
    \\117,168,530
    \\805,96,715
    \\346,949,466
    \\970,615,88
    \\941,993,340
    \\862,61,35
    \\984,92,344
    \\425,690,689
;

test "part1" {
    const result = try solve1(example_input, std.testing.allocator, 10);
    try std.testing.expectEqual(40, result);
}

test "part2" {
    const result = try part2(example_input, std.testing.allocator);
    try std.testing.expectEqual(25272, result);
}
