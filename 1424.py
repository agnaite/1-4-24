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

        >>> Die(6)
        <An unrolled, 6-sided Die>
        """

        self.sides = sides
        self.choices = list(range(1, self.sides))
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


    def __repr__(self) -> str:
        if self.choice != -1:
            return f'<A {self.sides}-sided Die with {self.choices} sides that rolled f{self.choice}>'
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
