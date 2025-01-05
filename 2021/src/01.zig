const std = @import("std");

pub const input = @embedFile("./inputs/01.txt");

pub fn solve(allocator: std.mem.Allocator, inputData: []const u8) !struct { part1: usize, part2: usize } {
    const depths = try parseInput(allocator, inputData);
    defer depths.deinit();

    const res1 = try part1(depths);
    const res2 = try part2(depths);

    return .{ .part1 = res1, .part2 = res2 };
}

pub fn part1(depths: std.ArrayList(usize)) !usize {
    var res: usize = 0;

    for (depths.items[1..], 1..) |depth, i| {
        const prev = depths.items[i - 1];

        if (depth > prev) {
            res += 1;
        }
    }

    return res;
}

pub fn part2(depths: std.ArrayList(usize)) !usize {
    var res: usize = 0;

    var prev_window = depths.items[0] + depths.items[1] + depths.items[2];

    for (3..depths.items.len) |i| {
        const window = depths.items[i] + depths.items[i - 1] + depths.items[i - 2];

        if (window > prev_window) {
            res += 1;
        }

        prev_window = window;
    }

    return res;
}

pub fn parseInput(allocator: std.mem.Allocator, inputData: []const u8) !std.ArrayList(usize) {
    var res = std.ArrayList(usize).init(allocator);

    const trimmed = std.mem.trim(u8, inputData, &std.ascii.whitespace);

    var linesIter = std.mem.splitSequence(u8, trimmed, "\n");

    while (linesIter.next()) |line| {
        const n = try std.fmt.parseInt(usize, line, 10);

        try res.append(n);
    }

    return res;
}

test "solve for test input" {
    const test_input =
        \\199
        \\200
        \\208
        \\210
        \\200
        \\207
        \\240
        \\269
        \\260
        \\263
    ;

    const res = try solve(std.testing.allocator, test_input);

    const res1 = res.part1;
    const res2 = res.part2;

    const expected_part1 = 7;
    const expected_part2 = 5;

    try std.testing.expectEqual(expected_part1, res1);
    try std.testing.expectEqual(expected_part2, res2);
}

test "solve for input" {
    const res = try solve(std.testing.allocator, input);

    const res1 = res.part1;
    const res2 = res.part2;

    const expected_part1 = 1602;
    const expected_part2 = 1633;

    try std.testing.expectEqual(expected_part1, res1);
    try std.testing.expectEqual(expected_part2, res2);
}
