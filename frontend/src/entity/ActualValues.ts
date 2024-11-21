export interface ActualValues {
    id: string;
    hitPoints: number;
    attack: number;
    defense: number;
    specialAttack: number;
    specialDefense: number;
    speed: number;
}

export const NewActualValuesFromJson = (actualValues: any): ActualValues => {
    return {
        id: actualValues.id,
        hitPoints: actualValues.hit_points,
        attack: actualValues.attack,
        defense: actualValues.defense,
        specialAttack: actualValues.special_attack,
        specialDefense: actualValues.special_defense,
        speed: actualValues.speed,
    };
}