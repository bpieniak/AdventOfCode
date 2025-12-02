const std = @import("std");

pub fn build(b: *std.Build) void {
    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{});

    const day = b.option(u8, "day", "Which day to build/run/test") orelse 1;

    var name_buf: [16]u8 = undefined;
    const exe_name = std.fmt.bufPrint(&name_buf, "aoc_day{d:0>2}", .{day}) catch unreachable;

    var path_buf: [64]u8 = undefined;
    const root_path = std.fmt.bufPrint(&path_buf, "src/day{d:0>2}/main.zig", .{day}) catch unreachable;

    const root_module = b.createModule(.{
        .root_source_file = b.path(root_path),
        .target = target,
        .optimize = optimize,
    });

    const exe = b.addExecutable(.{
        .name = exe_name,
        .root_module = root_module,
    });
    b.installArtifact(exe);

    const run_cmd = b.addRunArtifact(exe);
    run_cmd.step.dependOn(b.getInstallStep());
    if (b.args) |args| run_cmd.addArgs(args);
    const run_step = b.step("run", "Run selected day");
    run_step.dependOn(&run_cmd.step);

    const exe_tests = b.addTest(.{
        .root_module = root_module,
    });
    const run_exe_tests = b.addRunArtifact(exe_tests);
    const test_step = b.step("test", "Run tests for selected day");
    test_step.dependOn(&run_exe_tests.step);
}
