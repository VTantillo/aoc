local left = {}
local right = {}

for line in io.lines("../../../puzzles/2024/01/input.txt") do
	local line_nums = {}
	for val in string.gmatch(line, "%w+") do
		table.insert(line_nums, tonumber(val))
	end

	print(line_nums[1], line_nums[2])

	table.insert(left, line_nums[1])
	table.insert(right, line_nums[2])
end

table.sort(left)
table.sort(right)

for k, v in pairs(left) do
	print(v)
end
