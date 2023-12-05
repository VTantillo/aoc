from pydantic import BaseModel


class RangeMap(BaseModel):
    dst_start: int
    src_start: int
    range_len: int


class SeedRange(BaseModel):
    start: int
    range_len: int


def take_2(input: list[str]) -> int:
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

    found_seed = False
    location = 0

    while not found_seed:
        humidity = get_src(location, humidity_to_location)
        temp = get_src(humidity, temp_to_humidity)
        light = get_src(temp, light_to_temp)
        water = get_src(light, water_to_light)
        fertilizer = get_src(water, fertilizer_to_water)
        soil = get_src(fertilizer, soil_to_fertilizer)
        seed = get_src(soil, seed_to_soil)

        found_seed = check_seed(seed, seeds)

        location = location + 1 if not found_seed else location

    return location


def check_seed(seed: int, seeds: list[SeedRange]) -> bool:
    if seed < seeds[0].start:
        return False

    for s in seeds:
        if seed >= s.start and seed < s.start + s.range_len:
            return True

    return False


def get_dst(src: int, range_map: list[RangeMap]) -> int:
    for m in range_map:
        if src >= m.src_start and src < m.src_start + m.range_len:
            diff = m.src_start - m.dst_start
            return src - diff

    return src


def get_src(dst: int, range_map: list[RangeMap]) -> int:
    for m in range_map:
        if dst >= m.dst_start and dst < m.dst_start + m.range_len:
            diff = m.dst_start - m.src_start
            return dst - diff

    return dst


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

    seeds.sort(key=lambda x: x.start)
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
