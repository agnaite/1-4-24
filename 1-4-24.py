#!/usr/bin/env python3

import random
from typing import Iterator

class Die:
    """
    Represents a die
    """

    def __init__(self, sides: int = 6) -> None:
        """
        Initiates a new die

        # sets seed for testing
        >>> random.seed(1424)

        >>> Die(6)
        <An unrolled, 6-sided Die>

        # it contains state
        >>> die = Die(6)
        >>> die.roll()
        1

        >>> die
        <A 6-sided Die with [1, 2, 3, 4, 5, 6] sides that rolled 1>

        >>> Die(6).roll() <= 6
        True
        """

        self.sides = sides
        self.choices = list(range(1, self.sides+1))
        self.choice = -1


    def roll(self) -> int:
        """
        Rolls the die

        >>> die = Die(6)
        >>> die.roll() == die.roll()
        True
        """

        if self.choice == -1:
            self.choice = random.choice(self.choices)
        return self.choice


    def bottoms(self) -> int:
        """

        Checks bottoms of a die (if rolled). Otherwise, raises an Exception

        >>> die = Die(6)
        >>> die.choices
        [1, 2, 3, 4, 5, 6]

        >>> die.choice = 6
        >>> die.bottoms()
        1

        >>> die.choice = 5
        >>> die.bottoms()
        2

        >>> die.choice = 4
        >>> die.bottoms()
        3

        >>> die.choice = 3
        >>> die.bottoms()
        4

        >>> die.choice = 2
        >>> die.bottoms()
        5

        >>> die.choice = 1
        >>> die.bottoms()
        6

        """

        return self.choices[::-1][self.choices.index(self.choice)]


    def __repr__(self) -> str:
        if self.choice != -1:
            return f'<A {self.sides}-sided Die with {self.choices} sides that rolled {self.choice}>'
        else:
            return f'<An unrolled, {self.sides}-sided Die>'


class Roll:
    """
    Represents a roll in 1-4-24
    """

    def __init__(self, rolls: Iterator[int]) -> None:
        """
        Initiate the roll
        >>> Roll([1, 4, 6, 6, 6, 6])
        <Roll [1, 4, 6, 6, 6, 6]>
        """

        self.rolls = rolls

    def is_qualified(self) -> bool:
        """
        Checks if roll qualifies (has a 1 and a 4)

        >>> Roll([1, 4, 6, 6, 6, 6]).is_qualified()
        True
        """

        return (1 in self.rolls) and (4 in self.rolls)

    def __repr__(self) -> str:

        return f'<Roll {self.rolls}>'
