from copy import deepcopy


def game(decks):
    decks_played = set()
    while len(decks[0]) > 0 and len(decks[1]) > 0:
        # player 1 automatically wins if the same deck is already played
        if (tuple(decks[0]), tuple(decks[1])) in decks_played:
            return 0
        # add the two decks to already played
        decks_played.add((tuple(decks[0]), tuple(decks[1])))
        # one player does not have enough card to recurse, winner is simply one with higher card
        if decks[0][0] >= len(decks[0]) or decks[1][0] >= len(decks[1]):
            winner = 0 if decks[0][0] > decks[1][0] else 1
        else:  # need to recurse, subdecks will be from next card, plus number of cards required
            winner = game([deck[1:deck[0]+1] for deck in decks])
        # add the 2 cards to the winner deck
        decks[winner].extend(
            [decks[winner].pop(0), decks[(winner+1) % 2].pop(0)])

    return 0 if len(decks[0]) > 0 else 1


players = open(f'22/input', 'r').read().split('\n\n')

original = [[int(x) for x in player.split('\n')[1:]]for player in players]

# part 1
decks = deepcopy(original)
while len(decks[0]) > 0 and len(decks[1]) > 0:
    winner = 0 if decks[0][0] > decks[1][0] else 1
    decks[winner].extend(
        [decks[winner].pop(0), decks[(winner+1) % 2].pop(0)])

ans = sum([(i+1) * val for x in decks for i, val in enumerate(x[::-1])])
print(ans)

# part 2
decks = deepcopy(original)
game(decks)
ans = sum([(i+1) * val for x in decks for i, val in enumerate(x[::-1])])
print(ans)
