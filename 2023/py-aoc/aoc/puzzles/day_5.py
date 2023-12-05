from pydantic import BaseModel


class RangeMap(BaseModel):
    dst_start: int
    src_start: int
    range_len: int


class SeedRange(BaseModel):
    start: int
    range_len: int


def part_1(input: list[str]) -> int:
    print("begin stuff: making the range maps")
    seed_to_soil = get_range_map(input, "seed-to-soil")
    seed_to_soil.sort(key=lambda x: x.src_start)

    print("seed to soil")

    soil_to_fertilizer = get_range_map(input, "soil-to-fertilizer")
    soil_to_fertilizer.sort(key=lambda x: x.src_start)

    print("soil to fert")
    fertilizer_to_water = get_range_map(input, "fertilizer-to-water")
    fertilizer_to_water.sort(key=lambda x: x.src_start)

    print("fert to water")
    water_to_light = get_range_map(input, "water-to-light")
    water_to_light.sort(key=lambda x: x.src_start)

    print("water to light")
    light_to_temp = get_range_map(input, "light-to-temperature")
    light_to_temp.sort(key=lambda x: x.src_start)

    print("light to temp")
    temp_to_humidity = get_range_map(input, "temperature-to-humidity")
    temp_to_humidity.sort(key=lambda x: x.src_start)

    print("temp to humid")
    humidity_to_location = get_range_map(input, "humidity-to-location")
    humidity_to_location.sort(key=lambda x: x.src_start)

    print("humid to location")

    seeds = get_seeds(input)

    print("begin checking seeds", len(seeds))

    low_location: int | None = None
    for sr in seeds:
        print(f"checking {sr.range_len} starting at {sr.start}")
        for seed in range(sr.start, sr.range_len + sr.start):
            soil = get_dst(seed, seed_to_soil)
            fertilizer = get_dst(soil, soil_to_fertilizer)
            water = get_dst(fertilizer, fertilizer_to_water)
            light = get_dst(water, water_to_light)
            temp = get_dst(light, light_to_temp)
            humidity = get_dst(temp, temp_to_humidity)
            location = get_dst(humidity, humidity_to_location)

            if low_location is None:
                print("setting low", location)
                low_location = location

            if low_location is not None and location < low_location:
                low_location = location
                print(f"new low location: {low_location}")
        print(f"done checking, current low {low_location}")

    print("done checking seeds", low_location)

    return low_location if low_location is not None else 0


def get_dst(src: int, range_map: list[RangeMap]) -> int:
    for m in range_map:
        if src >= m.src_start and src < m.src_start + m.range_len:
            diff = m.src_start - m.dst_start
            return src - diff

    return src


def get_seeds(input: list[str]) -> list[SeedRange]:
    seed_list: list[int] = []
    for line in input:
        if "seeds:" in line:
            for seed in line.split(" "):
                if seed.strip().isdigit():
                    seed_list.append(int(seed))

    seeds: list[SeedRange] = []
    for s in range(0, len(seed_list), 2):
        seeds.append(SeedRange(start=seed_list[s], range_len=seed_list[s + 1]))

    return seeds


def get_range_map(input: list[str], key: str) -> list[RangeMap]:
    range_maps: list[RangeMap] = []

    found_map = False

    for line in input:
        if key in line:
            found_map = True
            continue

        if found_map and line == "\n":
            break

        if found_map:
            map_split = line.strip().split(" ")
            range_maps.append(
                RangeMap(
                    dst_start=int(map_split[0].strip()),
                    src_start=int(map_split[1].strip()),
                    range_len=int(map_split[2].strip()),
                )
            )

    return range_maps
