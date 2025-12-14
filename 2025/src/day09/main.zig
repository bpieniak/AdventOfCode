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
    const tiles = try read_tiles(input, allocator);
    defer allocator.free(tiles);

    var max_area: u64 = 0;

    for (0..tiles.len) |i| {
        for (i..tiles.len) |j| {
            const curr_area = tiles[i].area(tiles[j]);

            max_area = @max(max_area, curr_area);
        }
    }

    return max_area;
}

pub fn part2(input: []const u8, allocator: std.mem.Allocator) !usize {
    const reds = try read_tiles(input, allocator);
    defer allocator.free(reds);

    var edges = try buildEdges(reds, allocator);
    defer edges.deinit(allocator);

    var max_area: usize = 0;

    for (reds, 0..) |a, i| {
        for (reds[i + 1 ..]) |b| {
            const minx = @min(a.x, b.x);
            const maxx = @max(a.x, b.x);
            const miny = @min(a.y, b.y);
            const maxy = @max(a.y, b.y);

            if (!rectangleInside(edges, minx, miny, maxx, maxy)) continue;

            const area = @as(usize, @intCast((maxx - minx + 1) * (maxy - miny + 1)));
            max_area = @max(max_area, area);
        }
    }

    return max_area;
}

const tile = struct {
    x: i64,
    y: i64,

    fn area(self: tile, other: tile) usize {
        const w = @abs(other.x - self.x) + 1;
        const h = @abs(other.y - self.y) + 1;
        return w * h;
    }
};

fn read_tiles(input: []const u8, allocator: std.mem.Allocator) ![]tile {
    var tiles = std.ArrayList(tile).empty;
    errdefer tiles.deinit(allocator);

    var it = std.mem.splitScalar(u8, input, '\n');
    while (it.next()) |line| {
        var lineIt = std.mem.splitScalar(u8, line, ',');

        const x = try std.fmt.parseInt(i64, lineIt.next() orelse return error.BadLine, 10);
        const y = try std.fmt.parseInt(i64, lineIt.next() orelse return error.BadLine, 10);

        try tiles.append(allocator, .{
            .x = x,
            .y = y,
        });
    }

    return tiles.toOwnedSlice(allocator);
}

const EdgeSet = struct {
    vertical: []VertEdge,
    horizontal: []HorzEdge,
    min: tile,
    max: tile,

    fn deinit(self: *EdgeSet, allocator: std.mem.Allocator) void {
        allocator.free(self.vertical);
        allocator.free(self.horizontal);
    }
};

const VertEdge = struct { x: i64, y0: i64, y1: i64 };
const HorzEdge = struct { y: i64, x0: i64, x1: i64 };

fn buildEdges(reds: []const tile, allocator: std.mem.Allocator) !EdgeSet {
    var min = reds[0];
    var max = reds[0];

    var verts = std.ArrayList(VertEdge).empty;
    errdefer verts.deinit(allocator);
    var hors = std.ArrayList(HorzEdge).empty;
    errdefer hors.deinit(allocator);

    for (reds, 0..) |r, i| {
        const nxt = reds[(i + 1) % reds.len];
        if (r.x == nxt.x) {
            const lo = @min(r.y, nxt.y);
            const hi = @max(r.y, nxt.y);
            try verts.append(allocator, VertEdge{ .x = r.x, .y0 = lo, .y1 = hi });
        } else {
            const lo = @min(r.x, nxt.x);
            const hi = @max(r.x, nxt.x);
            try hors.append(allocator, HorzEdge{ .y = r.y, .x0 = lo, .x1 = hi });
        }

        min.x = @min(min.x, r.x);
        max.x = @max(max.x, r.x);
        min.y = @min(min.y, r.y);
        max.y = @max(max.y, r.y);
    }

    const vertSlice = try verts.toOwnedSlice(allocator);
    const horSlice = try hors.toOwnedSlice(allocator);

    std.sort.block(VertEdge, vertSlice, {}, struct {
        fn lessThan(_: void, a: VertEdge, b: VertEdge) bool {
            return a.x < b.x;
        }
    }.lessThan);
    std.sort.block(HorzEdge, horSlice, {}, struct {
        fn lessThan(_: void, a: HorzEdge, b: HorzEdge) bool {
            return a.y < b.y;
        }
    }.lessThan);

    return EdgeSet{ .vertical = vertSlice, .horizontal = horSlice, .min = min, .max = max };
}

fn pointOnEdge(p: tile, e: VertEdge) bool {
    if (p.x != e.x) return false;
    return p.y >= e.y0 and p.y <= e.y1;
}
fn pointOnEdgeH(p: tile, e: HorzEdge) bool {
    if (p.y != e.y) return false;
    return p.x >= e.x0 and p.x <= e.x1;
}

fn pointInside(edges: EdgeSet, p: tile) bool {
    for (edges.vertical) |v| {
        if (pointOnEdge(p, v)) return true;
    }
    for (edges.horizontal) |h| {
        if (pointOnEdgeH(p, h)) return true;
    }

    var crossings: usize = 0;
    for (edges.vertical) |v| {
        if (p.y < v.y0 or p.y >= v.y1) continue;
        if (v.x > p.x) crossings += 1;
    }

    return (crossings % 2) == 1;
}

fn rangeOverlap(a0: i64, a1: i64, b0: i64, b1: i64) bool {
    const lo = @max(a0, b0);
    const hi = @min(a1, b1);
    return lo < hi;
}

fn lowerBoundV(arr: []const VertEdge, x: i64) usize {
    var lo: usize = 0;
    var hi: usize = arr.len;
    while (lo < hi) {
        const mid = (lo + hi) / 2;
        if (arr[mid].x < x) {
            lo = mid + 1;
        } else {
            hi = mid;
        }
    }
    return lo;
}
fn lowerBoundH(arr: []const HorzEdge, y: i64) usize {
    var lo: usize = 0;
    var hi: usize = arr.len;
    while (lo < hi) {
        const mid = (lo + hi) / 2;
        if (arr[mid].y < y) {
            lo = mid + 1;
        } else {
            hi = mid;
        }
    }
    return lo;
}

fn hasCrossing(edges: EdgeSet, minx: i64, miny: i64, maxx: i64, maxy: i64) bool {
    const vx_lo = lowerBoundV(edges.vertical, minx + 1);
    const vx_hi = lowerBoundV(edges.vertical, maxx);
    var i = vx_lo;
    while (i < vx_hi) : (i += 1) {
        const v = edges.vertical[i];
        if (v.x <= minx or v.x >= maxx) continue;
        if (rangeOverlap(v.y0, v.y1, miny, maxy)) return true;
    }

    const hy_lo = lowerBoundH(edges.horizontal, miny + 1);
    const hy_hi = lowerBoundH(edges.horizontal, maxy);
    var j = hy_lo;
    while (j < hy_hi) : (j += 1) {
        const h = edges.horizontal[j];
        if (h.y <= miny or h.y >= maxy) continue;
        if (rangeOverlap(h.x0, h.x1, minx, maxx)) return true;
    }

    return false;
}

fn rectangleInside(edges: EdgeSet, minx: i64, miny: i64, maxx: i64, maxy: i64) bool {
    if (minx < edges.min.x or maxx > edges.max.x or miny < edges.min.y or maxy > edges.max.y) return false;

    if (!pointInside(edges, .{ .x = minx, .y = miny })) return false;
    if (!pointInside(edges, .{ .x = minx, .y = maxy })) return false;
    if (!pointInside(edges, .{ .x = maxx, .y = miny })) return false;
    if (!pointInside(edges, .{ .x = maxx, .y = maxy })) return false;

    if (hasCrossing(edges, minx, miny, maxx, maxy)) return false;

    return true;
}

const example_input =
    \\7,1
    \\11,1
    \\11,7
    \\9,7
    \\9,5
    \\2,5
    \\2,3
    \\7,3
;

test "part1" {
    const result = try part1(example_input, std.testing.allocator);
    try std.testing.expectEqual(50, result);
}

test "part2" {
    const result = try part2(example_input, std.testing.allocator);
    try std.testing.expectEqual(24, result);
}
