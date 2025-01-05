const std = @import("std");

pub const input = @embedFile("./inputs/02.txt");

const Direction = enum { forward, up, down };

const Instruction = struct { direction: Direction, val: usize };

pub fn solve(allocator: std.mem.Allocator, inputData: []const u8) !struct { part1: usize, part2: usize } {
    const instructions = try parseInput(allocator, inputData);
    defer instructions.deinit();

    const res1 = try part1(instructions);
    const res2 = try part2(instructions);

    return .{ .part1 = res1, .part2 = res2 };
}

fn part1(instructions: std.ArrayList(Instruction)) !usize {
    var h_pos: usize = 0;
    var depth: usize = 0;

    for (instructions.items) |ins| {
        switch (ins.direction) {
            .forward => h_pos += ins.val,
            .up => depth -= ins.val,
            .down => depth += ins.val,
        }
    }

    return h_pos * depth;
}

fn part2(instructions: std.ArrayList(Instruction)) !usize {
    var h_pos: usize = 0;
    var depth: usize = 0;
    var aim: usize = 0;

    for (instructions.items) |ins| {
        switch (ins.direction) {
            .forward => {
                h_pos += ins.val;
                depth += aim * ins.val;
            },
            .up => aim -= ins.val,
            .down => aim += ins.val,
        }
    }

    return h_pos * depth;
}

fn parseInput(allocator: std.mem.Allocator, inputData: []const u8) !std.ArrayList(Instruction) {
    var res = std.ArrayList(Instruction).init(allocator);

    const trimmed = std.mem.trim(u8, inputData, &std.ascii.whitespace);

    var linesIter = std.mem.splitSequence(u8, trimmed, "\n");

    while (linesIter.next()) |line| {
        var spl = std.mem.splitScalar(u8, line, ' ');

        const dirStr = spl.next() orelse "";

        const dir: Direction = if (std.mem.eql(u8, dirStr, "forward"))
            .forward
        else if (std.mem.eql(u8, dirStr, "down"))
            .down
        else
            .up;

        const val = try std.fmt.parseInt(usize, spl.next() orelse "", 10);

        const ins: Instruction = .{ .direction = dir, .val = val };

        try res.append(ins);
    }

    return res;
}

test "solve for test input" {
    const test_input =
        \\forward 5
        \\down 5
        \\forward 8
        \\up 3
        \\down 8
        \\forward 2
    ;

    const res = try solve(std.testing.allocator, test_input);

    const res1 = res.part1;
    const res2 = res.part2;

    const expected_part1 = 150;
    const expected_part2 = 900;

    try std.testing.expectEqual(expected_part1, res1);
    try std.testing.expectEqual(expected_part2, res2);
}

test "solve for input" {
    const res = try solve(std.testing.allocator, input);

    const res1 = res.part1;
    const res2 = res.part2;

    const expected_part1 = 2036120;
    const expected_part2 = 2015547716;

    try std.testing.expectEqual(expected_part1, res1);
    try std.testing.expectEqual(expected_part2, res2);
}
