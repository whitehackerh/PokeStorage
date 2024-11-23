export interface BaseStats {
    id: number;
    hitPoints: number;
    attack: number;
    defense: number;
    specialAttack: number;
    specialDefense: number;
    speed: number;
    average: number;
    total: number;
}

export const NewBaseStatsFromJson = (baseStats: any) : BaseStats => {
    return {
        id: baseStats.id,
        hitPoints: baseStats.hit_points,
        attack: baseStats.attack,
        defense: baseStats.defense,
        specialAttack: baseStats.special_attack,
        specialDefense: baseStats.special_defense,
        speed: baseStats.speed,
        average: baseStats.average,
        total: baseStats.total,
    };
}