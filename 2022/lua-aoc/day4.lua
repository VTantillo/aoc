local day_four = function(input)
	local overlap = 0

	local elves = {}
	for k, v in pairs(input) do
		elves[k] = {}
		for s in string.gmatch(v, "%d*") do
			table.insert(elves[k], tonumber(s))
		end
	end

	for k, v in pairs(elves) do
		local is_overlapping = false
		if v[1] >= v[3] and v[2] <= v[4] then
			is_overlapping = true
		elseif v[1] <= v[3] and v[2] >= v[4] then
			is_overlapping = true
		elseif v[1] >= v[3] and v[1] <= v[4] then
			is_overlapping = true
		elseif v[1] <= v[3] and v[3] <= v[2] then
			is_overlapping = true
		end

		if is_overlapping then
			overlap = overlap + 1
		else
			print(k, "1-start(1):", v[1], "1-end(2):", v[2])
			print(k, "2-start(3):", v[3], "2-end(4):", v[4])
			print("----")
		end
	end

	return overlap
end

local example = [[
2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
]]

local lines = {}
for line in io.lines("../input/day-4.txt") do
	table.insert(lines, line)
end

local ex_lines = {}
for line in string.gmatch(example, "[^\r\n]+") do
	table.insert(ex_lines, line)
end

local result = day_four(lines)

local answer = 2
print(result)
