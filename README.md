# advent-of-code

## 2023

Link to site: [https://adventofcode.com/2023/](https://adventofcode.com/2023/)

## 2023 Table of Contents

[1. Background 2023](#background-2023)\
[2. Day 1: Trebuchet?!](#2023-day-1-trebuchet)\
[3. Day 1 - part 2:](#2023-day-1-part-2)\
[4. Day 2: Cube Conundrum](#2023-day-2-cube-conundrum)\
[5. Day 2 - part 2:](#2023-day-2-part-2)\
[6. Day 3:](#2023-day-3)\
[7. Day 3 - part 2:](#2023-day-3-part-2)\
[6. Day 4:](#2023-day-4)\
[7. Day 4 - part 2:](#2023-day-4-part-2)\
[6. Day 5:](#2023-day-5)\
[7. Day 5 - part 2:](#2023-day-5-part-2)\
[6. Day 6:](#2023-day-6)\
[7. Day 6 - part 2:](#2023-day-6-part-2)\
[6. Day 7:](#2023-day-7)\
[7. Day 7 - part 2:](#2023-day-7-part-2)\
[6. Day 8:](#2023-day-8)\
[7. Day 8 - part 2:](#2023-day-8-part-2)\
[6. Day 9:](#2023-day-9)\
[7. Day 9 - part 2:](#2023-day-9-part-2)\
[6. Day 10:](#2023-day-10)\
[7. Day 10 - part 2:](#2023-day-10-part-2)\
[6. Day 11:](#2023-day-11)\
[7. Day 11 - part 2:](#2023-day-11-part-2)\
[6. Day 12:](#2023-day-12)\
[7. Day 12 - part 2:](#2023-day-12-part-2)\
[6. Day 13:](#2023-day-13)\
[7. Day 13 - part 2:](#2023-day-13-part-2)\
[6. Day 14:](#2023-day-14)\
[7. Day 14 - part 2:](#2023-day-14-part-2)\
[6. Day 15:](#2023-day-15)\
[7. Day 15 - part 2:](#2023-day-15-part-2)\
[6. Day 16:](#2023-day-16)\
[7. Day 16 - part 2:](#2023-day-16-part-2)\
[6. Day 17:](#2023-day-17)\
[7. Day 17 - part 2:](#2023-day-17-part-2)\
[6. Day 18:](#2023-day-18)\
[7. Day 18 - part 2:](#2023-day-18-part-2)\
[6. Day 19:](#2023-day-19)\
[7. Day 19 - part 2:](#2023-day-19-part-2)\
[6. Day 20:](#2023-day-20)\
[7. Day 20 - part 2:](#2023-day-20-part-2)\
[6. Day 21:](#2023-day-21)\
[7. Day 21 - part 2:](#2023-day-21-part-2)\
[6. Day 22:](#2023-day-22)\
[7. Day 22 - part 2:](#2023-day-22-part-2)\
[6. Day 23:](#2023-day-23)\
[7. Day 23 - part 2:](#2023-day-23-part-2)\
[6. Day 24:](#2023-day-24)\
[7. Day 24 - part 2:](#2023-day-24-part-2)\
[6. Day 25:](#2023-day-25)\
[7. Day 25 - part 2:](#2023-day-25-part-2)


------------------------------------------------------------------------------------


### Background 2023

Something is wrong with global snow production, and you've been selected to take a look. The Elves have even given you a map; on it, they've used stars to mark the top fifty locations that are likely to be having problems.

You've been doing this long enough to know that to restore snow operations, you need to check all fifty stars by December 25th.

Collect stars by solving puzzles. Two puzzles will be made available on each day in the Advent calendar; the second puzzle is unlocked when you complete the first. Each puzzle grants one star. Good luck!

### 2023 Day 1: Trebuchet?!
[back to TOC](#2023-table-of-contents)

You try to ask the following questions:

- why they can't just use a weather machine ("not powerful enough")
- where they're even sending you ("the sky")
- why your map looks mostly blank ("you sure ask a lot of questions")
- hang on did you just say the sky ("of course, where do you think snow comes from")

You realize that the Elves are already loading you into a trebuchet ("please hold still, we need to strap you in").

As they're making the final adjustments, they discover that their calibration document (your puzzle input) has been amended by a very young Elf who was apparently just excited to show off her art skills. Consequently, the Elves are having trouble reading the values on the document.

The newly-improved calibration document consists of lines of text; each line originally contained a specific calibration value that the Elves now need to recover. On each line, the calibration value can be found by combining the first digit and the last digit (in that order) to form a single two-digit number.

```go
1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
```

In this example, the calibration values of these four lines are `12`, `38`, `15`, and `77`. Adding these together produces `142`.

Consider your entire calibration document. What is the sum of all of the calibration values? `55386`


### 2023 Day 1: part 2:
[back to TOC](#2023-table-of-contents)

Your calculation isn't quite right. It looks like some of the digits are actually spelled out with letters: one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and last digit on each line. For example:

```go
two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen
```

In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.

What is the sum of all of the calibration values? `54824`


### 2023 Day 2: Cube Conundrum
[back to TOC](#2023-table-of-contents)

You're launched high into the atmosphere! The apex of your trajectory just barely reaches the surface of a large island floating in the sky. You gently land in a fluffy pile of leaves. It's quite cold, but you don't see much snow. An Elf runs over to greet you.

The Elf explains that you've arrived at Snow Island and apologizes for the lack of snow. He'll be happy to explain the situation, but it's a bit of a walk, so you have some time. They don't get many visitors up here; would you like to play a game in the meantime?

As you walk, the Elf shows you a small bag and some `cubes` which are either `red`, `green`, or `blue`. Each time you play this game, he will hide a secret number of `cubes` of each color in the bag, and your goal is to figure out information about the number of `cubes`.

To get information, once a bag has been loaded with `cubes`, the Elf will reach into the bag, grab a handful of random `cubes`, show them to you, and then put them back in the bag. He'll do this a few times per game.

You play several games and record the information from each game (your puzzle input). Each game is listed with its `ID number` (like the 11 in Game 11: ...) followed by a semicolon-separated list of subsets of `cubes` that were revealed from the bag (like 3 red, 5 green, 4 blue).

For example, the record of a few games might look like this:

```go
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
```

In game 1, three sets of cubes are revealed from the bag (and then put back again). The first set is `3 blue cubes and 4 red cubes`; the second set is `1 red cube, 2 green cubes, and 6 blue cubes`; the third set is only `2 green cubes`.

The Elf would first like to know which games would have been possible if the bag contained only `12 red cubes`, `13 green cubes`, and `14 blue cubes`?

In the example above, games 1, 2, and 5 would have been possible if the bag had been loaded with that configuration. However, game 3 would have been impossible because at one point the Elf showed you `20 red cubes` at once; similarly, game 4 would also have been impossible because the Elf showed you `15 blue cubes` at once. If you add up the `IDs` of the games that would have been possible, you get `8`.

Determine which games would have been possible if the bag had been loaded with only `12 red cubes`, `13 green cubes`, and `14 blue cubes`.

`What is the sum of the IDs of those games?`: `2632`


### 2023 Day 2: part 2:
[back to TOC](#2023-table-of-contents)

The Elf says they've stopped producing snow because they aren't getting any `water`! He isn't sure why the water stopped; however, he can show you how to get to the water source to check it out for yourself. It's just up ahead!

As you continue your walk, the Elf poses a second question: in each game you played, what is the `fewest number of cubes of each color` that could have been in the bag to make the game possible?

Again consider the example games from earlier:

```go
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
```

In game 1, the game could have been played with as few as 4 red, 2 green, and 6 blue cubes. If any color had even one fewer cube, the game would have been impossible.
Game 2 could have been played with a minimum of 1 red, 3 green, and 4 blue cubes.
Game 3 must have been played with at least 20 red, 13 green, and 6 blue cubes.
Game 4 required at least 14 red, 3 green, and 15 blue cubes.
Game 5 needed no fewer than 6 red, 3 green, and 2 blue cubes in the bag.

The `power` of a set of cubes is equal to the numbers of red, green, and blue cubes multiplied together. The power of the minimum set of cubes in game 1 is 48. In games 2-5 it was 12, 1560, 630, and 36, respectively. Adding up these five powers produces the sum `2286`.

For each game, find the minimum set of cubes that must have been present.

`What is the sum of the power of these sets?`: `69629`


### 2023 Day 3:
[back to TOC](#2023-table-of-contents)



### 2023 Day 3. part 2:
[back to TOC](#2023-table-of-contents)



### 2023 Day 4:
[back to TOC](#2023-table-of-contents)



### 2023 Day 4. part 2:
[back to TOC](#2023-table-of-contents)



### 2023 Day 5:
[back to TOC](#2023-table-of-contents)



### 2023 Day 5. part 2:
[back to TOC](#2023-table-of-contents)



### 2023 Day 6:
[back to TOC](#2023-table-of-contents)



### 2023 Day 6. part 2:
[back to TOC](#2023-table-of-contents)



### 2023 Day 7.
[back to TOC](#2023-table-of-contents)



### 2023 Day 7. part 2:
[back to TOC](#2023-table-of-contents)



### 2023 Day 8.
[back to TOC](#2023-table-of-contents)



### 2023 Day 8. part 2:
[back to TOC](#2023-table-of-contents)



### 2023 Day 9.
[back to TOC](#2023-table-of-contents)



### 2023 Day 9. part 2:
[back to TOC](#2023-table-of-contents)



### 2023 Day 10.
[back to TOC](#2023-table-of-contents)



### 2023 Day 10. part 2:
[back to TOC](#2023-table-of-contents)



### 2023 Day 11.
[back to TOC](#2023-table-of-contents)



### 2023 Day 11. part 2:
[back to TOC](#2023-table-of-contents)



### 2023 Day 12.
[back to TOC](#2023-table-of-contents)



### 2023 Day 12. part 2:
[back to TOC](#2023-table-of-contents)



### 2023 Day 13.
[back to TOC](#2023-table-of-contents)



### 2023 Day 13. part 2:
[back to TOC](#2023-table-of-contents)



### 2023 Day 14.
[back to TOC](#2023-table-of-contents)



### 2023 Day 14. part 2:
[back to TOC](#2023-table-of-contents)



### 2023 Day 15.
[back to TOC](#2023-table-of-contents)



### 2023 Day 15. part 2:
[back to TOC](#2023-table-of-contents)



### 2023 Day 16.
[back to TOC](#2023-table-of-contents)



### 2023 Day 16. part 2:
[back to TOC](#2023-table-of-contents)



### 2023 Day 17.
[back to TOC](#2023-table-of-contents)



### 2023 Day 17. part 2:
[back to TOC](#2023-table-of-contents)



### 2023 Day 18.
[back to TOC](#2023-table-of-contents)



### 2023 Day 18. part 2:
[back to TOC](#2023-table-of-contents)



### 2023 Day 19.
[back to TOC](#2023-table-of-contents)



### 2023 Day 19. part 2:
[back to TOC](#2023-table-of-contents)



### 2023 Day 20.
[back to TOC](#2023-table-of-contents)



### 2023 Day 20. part 2:
[back to TOC](#2023-table-of-contents)



### 2023 Day 21.
[back to TOC](#2023-table-of-contents)



### 2023 Day 21. part 2:
[back to TOC](#2023-table-of-contents)



### 2023 Day 22.
[back to TOC](#2023-table-of-contents)



### 2023 Day 22. part 2:
[back to TOC](#2023-table-of-contents)



### 2023 Day 23.
[back to TOC](#2023-table-of-contents)



### 2023 Day 23. part 2:
[back to TOC](#2023-table-of-contents)



### 2023 Day 24.
[back to TOC](#2023-table-of-contents)



### 2023 Day 24. part 2:
[back to TOC](#2023-table-of-contents)



### 2023 Day 25.
[back to TOC](#2023-table-of-contents)



### 2023 Day 25. part 2:
[back to TOC](#2023-table-of-contents)



------------------------------------------------------------------------------------