const std = @import("std");

const Day01 = @import("src/01.zig");
const Day02 = @import("src/02.zig");

pub fn main() !void {
    const stdout = std.io.getStdOut().writer();

    var single_threaded_arena = std.heap.ArenaAllocator.init(std.heap.page_allocator);
    defer single_threaded_arena.deinit();
    var thread_safe_arena = std.heap.ThreadSafeAllocator{ .child_allocator = single_threaded_arena.allocator() };
    const allocator = thread_safe_arena.allocator();

    const args = try std.process.argsAlloc(allocator);

    for (args) |arg| {
        try stdout.print("arg: {s}\n", .{arg});
    }

    if (args.len > 1) {
        runAll();
    }

    const res = try Day01.solve(allocator, Day01.input);

    try stdout.print(switch (@TypeOf(res.part1)) {
        []const u8 => "part1: {s}",
        else => "part1: {any}",
    } ++ "\n", .{res.part1});

    try stdout.print(switch (@TypeOf(res.part2)) {
        []const u8 => "part2: {s}",
        else => "part2: {any}",
    } ++ "\n", .{res.part2});
}

fn runAll(allocator: std.mem.Allocator) !void {
    const res01 = try Day01.solve(allocator, Day01.input);

    try printRes(res01.part1);
    try printRes(res01.part2);
}

fn printRes(res: anytype) !void {
    const stdout = std.io.getStdOut().writer();

    try stdout.print(switch (@TypeOf(res)) {
        []const u8 => "part2: {s}",
        else => "part2: {any}",
    } ++ "\n", .{res});
}
