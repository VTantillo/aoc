from functools import cmp_to_key

from pydantic import BaseModel
from rich import print


class Hand(BaseModel):
    cards: str
    bid: int
    card_map: dict[str, int] = {}
    rank: int = 0
    score: int = 0


cards = ["J", "2", "3", "4", "5", "6", "7", "8", "9", "T", "Q", "K", "A"]


def day_7_pt_1(input: list[str]) -> int:
    hands = parse_input(input)
    five_hands: list[Hand] = []
    four_hands: list[Hand] = []
    full_hands: list[Hand] = []
    three_hands: list[Hand] = []
    two_pair_hands: list[Hand] = []
    pair_hands: list[Hand] = []
    high_card_hands: list[Hand] = []

    for hand in hands:
        hand.card_map = count_cards(hand.cards)

    for hand in hands:
        if check_five_of_a_kind(hand.card_map):
            five_hands.append(hand)
        elif check_four_of_a_kind(hand.card_map):
            four_hands.append(hand)
        elif check_full_house(hand.card_map):
            full_hands.append(hand)
        elif check_three_of_a_kind(hand.card_map):
            three_hands.append(hand)
        elif check_two_pair(hand.card_map):
            two_pair_hands.append(hand)
        elif check_pair(hand.card_map):
            pair_hands.append(hand)
        else:
            high_card_hands.append(hand)

    sorted_high_hands = sorted(high_card_hands, key=cmp_to_key(compare_hands))
    sorted_pair_hands = sorted(pair_hands, key=cmp_to_key(compare_hands))
    sorted_two_pair_hands = sorted(two_pair_hands, key=cmp_to_key(compare_hands))
    sorted_three_hands = sorted(three_hands, key=cmp_to_key(compare_hands))
    sorted_full_hands = sorted(full_hands, key=cmp_to_key(compare_hands))
    sorted_four_hands = sorted(four_hands, key=cmp_to_key(compare_hands))
    sorted_five_hands = sorted(five_hands, key=cmp_to_key(compare_hands))

    print(f"{len(sorted_five_hands)} five of a kind", sorted_five_hands)
    print(f"{len(sorted_four_hands)} four of a kind", sorted_four_hands)
    print(f"{len(sorted_full_hands)} full house", sorted_full_hands)
    print(f"{len(sorted_three_hands)} three of a kind", sorted_three_hands)
    print(f"{len(sorted_two_pair_hands)} two pair", sorted_two_pair_hands)
    print(f"{len(sorted_pair_hands)} pair", sorted_pair_hands)
    print(f"{len(sorted_high_hands)} high card", sorted_high_hands)

    sorted_hands: list[Hand] = []
    sorted_hands.extend(sorted_high_hands)
    sorted_hands.extend(sorted_pair_hands)
    sorted_hands.extend(sorted_two_pair_hands)
    sorted_hands.extend(sorted_three_hands)
    sorted_hands.extend(sorted_full_hands)
    sorted_hands.extend(sorted_four_hands)
    sorted_hands.extend(sorted_five_hands)

    for i, hand in enumerate(sorted_hands):
        hand.rank = i + 1

    for hand in sorted_hands:
        hand.score = hand.rank * hand.bid

    score = 0
    for hand in sorted_hands:
        score += hand.score

    return score


def compare_hands(a: Hand, b: Hand) -> int:
    for idx in range(len(a.cards)):
        rank_a = cards.index(a.cards[idx])
        rank_b = cards.index(b.cards[idx])

        if rank_a == rank_b:
            continue

        return rank_a - rank_b

    return 0


def count_cards(cards: str) -> dict[str, int]:
    hand_set = set()
    for card in cards:
        hand_set.add(card)

    card_map: dict[str, int] = {}
    for card in hand_set:
        card_map[card] = cards.count(card)

    return card_map


def check_pair(card_map: dict[str, int]) -> bool:
    jokers = card_map.get("J") or 0

    for k, v in card_map.items():
        count = v + jokers
        if k != "J" and count == 2:
            return True

    return jokers == 2


def check_two_pair(card_map: dict[str, int]) -> bool:
    jokers = card_map.get("J") or 0

    one_pair = False

    if jokers >= 2:
        one_pair = True
        jokers -= 2

    for _, v in card_map.items():
        count = v + jokers
        if not one_pair and count == 2:
            one_pair = True
            jokers = 0
            continue

        if one_pair and count == 2:
            return True

    return False


def check_three_of_a_kind(card_map: dict[str, int]) -> bool:
    jokers = card_map.get("J") or 0

    for k, v in card_map.items():
        count = v + jokers
        if k != "J" and count == 3:
            return True

    return jokers == 3


def check_full_house(card_map: dict[str, int]) -> bool:
    jokers = card_map.get("J") or 0

    found_pair = False
    found_three = False

    for k, v in card_map.items():
        count = v + jokers
        if k != "J" and v == 2 and not found_pair:
            found_pair = True

        elif k != "J" and v == 3 and not found_three:
            found_three = True

        elif k != "J" and count == 2 and not found_pair:
            found_pair = True
            jokers = 0

        elif k != "J" and count == 3 and not found_three:
            found_three = True
            jokers = 0

    return found_pair and found_three


def check_four_of_a_kind(card_map: dict[str, int]) -> bool:
    jokers = card_map.get("J") or 0

    for k, v in card_map.items():
        count = v + jokers
        if k != "J" and count == 4:
            return True

    return jokers == 4


def check_five_of_a_kind(card_map: dict[str, int]) -> bool:
    jokers = card_map.get("J") or 0

    for k, v in card_map.items():
        count = v + jokers
        if k != "J" and count == 5:
            return True

    return jokers == 5


def parse_input(input: list[str]) -> list[Hand]:
    hands: list[Hand] = []

    for line in input:
        line_split = line.split(" ")
        hands.append(Hand(cards=line_split[0], bid=int(line_split[1])))

    return hands
