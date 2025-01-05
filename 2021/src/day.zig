const std = @import("std");

const Result = union(enum) { usize: usize, string: []const u8 };

// Day interface - is this how to create interfaces in zig?? WTF
// https://www.openmymind.net/Zig-Interfaces/
fn Day(comptime InputT: type) type {
    return struct {
        ptr: *anyopaque,
        parseFn: *const fn (ptr: *anyopaque) InputT,
        part1Fn: *const fn (ptr: *anyopaque, input: InputT) Result,
        part2Fn: *const fn (ptr: *anyopaque, input: InputT) Result,

        fn init(ptr: anytype) Day(InputT) {
            const T = @TypeOf(ptr);
            const ptr_info = @typeInfo(T);

            const gen = struct {
                pub fn parse(pointer: *anyopaque) InputT {
                    const self: T = @ptrCast(@alignCast(pointer));
                    return ptr_info.Pointer.child.parse(self);
                }
                pub fn part1(pointer: *anyopaque, input: InputT) Result {
                    const self: T = @ptrCast(@alignCast(pointer));
                    return ptr_info.Pointer.child.part1(self, input);
                }
                pub fn part2(pointer: *anyopaque, input: InputT) Result {
                    const self: T = @ptrCast(@alignCast(pointer));
                    return ptr_info.Pointer.child.part2(self, input);
                }
            };

            return .{
                .ptr = ptr,
                .parseFn = gen.parse,
                .part1Fn = gen.part1,
                .part2Fn = gen.part2,
            };
        }

        pub fn parse(self: Day(InputT)) InputT {
            return self.parseFn(self.ptr);
        }
        pub fn part1(self: Day(InputT), input: InputT) Result {
            return self.part1Fn(self.ptr, input);
        }
        pub fn part2(self: Day(InputT), input: InputT) Result {
            return self.part2Fn(self.ptr, input);
        }
    };
}

const Day01 = struct {
    inputData: []const u8,
    allocator: std.mem.Allocator,

    pub fn day(self: *Day01) Day(usize) {
        return Day(usize).init(self);
    }

    pub fn parse(self: *Day01) usize {
        std.debug.print("day01 parse. InputData: {s}\n", .{self.inputData});

        return 123;
    }
    pub fn part1(self: *Day01, input: usize) Result {
        _ = input;
        _ = self;

        std.debug.print("day01 part1\n", .{});

        return Result{ .usize = 456 };
    }
    pub fn part2(self: *Day01, input: usize) Result {
        _ = input;
        _ = self;

        std.debug.print("day01 part2\n", .{});

        return Result{ .usize = 789 };
    }
};

const Day02 = struct {
    inputData: []const u8,
    allocator: std.mem.Allocator,

    pub fn day(self: *Day02) Day([]const u8) {
        return Day([]const u8).init(self);
    }

    pub fn parse(self: *Day02) []const u8 {
        std.debug.print("day02 parse. InputData: {s}\n", .{self.inputData});

        return "abc";
    }
    pub fn part1(self: *Day02, input: []const u8) Result {
        _ = input;
        _ = self;

        std.debug.print("day02 part1\n", .{});

        return Result{ .string = "def" };
    }
    pub fn part2(self: *Day02, input: []const u8) Result {
        _ = input;
        _ = self;

        std.debug.print("day02 part2\n", .{});

        return Result{ .string = "ghi" };
    }
};

pub fn main() void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    const allocator = gpa.allocator();

    var day01 = Day01{
        .inputData = "foo",
        .allocator = allocator,
    };
    var day02 = Day02{
        .inputData = "bar",
        .allocator = allocator,
    };

    const day01Day = day01.day();
    const day02Day = day02.day();

    const input01 = day01Day.parse();
    const res01_1 = day01Day.part1(input01);
    const res01_2 = day01Day.part2(input01);

    std.debug.print("input01: {d}\n", .{input01});
    printResult(res01_1);
    printResult(res01_2);

    const input02 = day02Day.parse();
    const res02_1 = day02Day.part1(input02);
    const res02_2 = day02Day.part2(input02);

    std.debug.print("input02: {s}\n", .{input02});
    printResult(res02_1);
    printResult(res02_2);
}

fn printResult(res: Result) void {
    switch (res) {
        .usize => std.debug.print("result is {any}\n", .{res.usize}),
        .string => std.debug.print("result is {s}\n", .{res.string}),
    }
}
