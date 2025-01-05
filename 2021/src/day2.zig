const std = @import("std");

const Result = union(enum) { usize: usize, string: []const u8 };

const Day01 = struct {
    inputData: []const u8,
    allocator: std.mem.Allocator,

    pub fn parse(self: Day01) usize {
        std.debug.print("day01 parse. InputData: {s}\n", .{self.inputData});

        return 123;
    }
    pub fn part1(self: Day01, input: usize) Result {
        _ = input;
        _ = self;

        std.debug.print("day01 part1\n", .{});

        return Result{ .usize = 456 };
    }
    pub fn part2(self: Day01, input: usize) Result {
        _ = input;
        _ = self;

        std.debug.print("day01 part2\n", .{});

        return Result{ .usize = 789 };
    }
};

const Day02 = struct {
    inputData: []const u8,
    allocator: std.mem.Allocator,

    pub fn parse(self: Day02) []const u8 {
        std.debug.print("day02 parse. InputData: {s}\n", .{self.inputData});

        return "abc";
    }
    pub fn part1(self: Day02, input: []const u8) Result {
        _ = input;
        _ = self;

        std.debug.print("day02 part1\n", .{});

        return Result{ .string = "def" };
    }
    pub fn part2(self: Day02, input: []const u8) Result {
        _ = input;
        _ = self;

        std.debug.print("day02 part2\n", .{});

        return Result{ .string = "ghi" };
    }
};

pub fn main() void {
    // var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    // const allocator = gpa.allocator();
    //
    // const day01 = Day01{
    //     .inputData = "foo",
    //     .allocator = allocator,
    // };
    // const day02 = Day02{
    //     .inputData = "bar",
    //     .allocator = allocator,
    // };
    //
    // runDay(day01);
    // runDay(day02);

    runner();
}

fn runDay(day: anytype) void {
    const input = day.parse();
    const res1 = day.part1(input);
    const res2 = day.part2(input);

    printResult(res1);
    printResult(res2);
}

fn printResult(res: Result) void {
    switch (res) {
        .usize => std.debug.print("result is {any}\n", .{res.usize}),
        .string => std.debug.print("result is {s}\n", .{res.string}),
    }
}

fn runner() void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();

    const d = "01";

    const inputPath = std.fs.path.join(allocator, &[_][]const u8{ "inputs", std.fmt.allocPrint(
        allocator,
        "{s}.txt",
        .{d},
    ) catch unreachable }) catch unreachable;

    std.debug.print("input path: {s}", .{inputPath});

    const Day = getDayStruct(d);

    const dd = Day{
        .inputData = @embedFile(inputPath),
        .allocator = allocator,
    };

    runDay(dd);
}

fn getDayStruct(type_name: []const u8) type {
    return if (std.mem.eql(u8, type_name, "01")) Day01 else if (std.mem.eql(u8, type_name, "02")) Day02 else Day01;
}
