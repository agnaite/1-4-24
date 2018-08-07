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
        <Die with 6 sides>
        """

        self.sides = sides

    def roll(self) -> int:
        """
        Rolls the die

        >>> die = Die(6)
        >>> die.roll() <= 6
        True
        """

        return random.choice(range(1, self.sides))

    def __repr__(self) -> str:
        return f'<Die with {self.sides} sides>'


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
