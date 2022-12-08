def get_coin_combo(total, coins):
    if total == 0:
        return []
    combinations = []
    for coin in coins:
        if coin <= total:
            remaining = get_coin_combo(total-coin, coins)
            for combination in remaining:
                combinations.append([coin] + combination)
    return combinations


coins = [1,2,5]
con = get_coin_combo(10, coins)
print(con)