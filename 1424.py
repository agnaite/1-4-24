#!/usr/bin/env python3

import random
from typing import Iterator

def roll(d: int = 6) -> int:
    """

    Rolls the die

    >>> roll(6) <= 6
    True
    """

    return random.choice(range(1,d))


def is_qualified(dice: Iterator[int]) -> bool:
    """

    Checks if roll qualifies (has a 1 and a 4)

    >>> is_qualified([1, 4, 6, 6, 6, 6])
    True

    >>> is_qualified([1, 6, 6, 6, 6, 6])
    False

    >>> is_qualified([1, 4])
    True

    >>> is_qualified([6, 6])
    False

    """

    return (1 in dice) and (4 in dice)


def main():

    dice = []

    while len(dice) < 6:
        dice.append(roll())

    print(is_qualified(dice))
    print(dice)


if __name__ == '__main__':
    main()
