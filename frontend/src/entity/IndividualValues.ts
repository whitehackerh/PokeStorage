export interface IndividualValues {
    id: string;
    hitPoints: number;
    attack: number;
    defense: number;
    specialAttack: number;
    specialDefense: number;
    speed: number;
}

export const NewIndividualValuesFromJson = (individualValues: any): IndividualValues => {
    return {
        id: individualValues.id,
        hitPoints: individualValues.hit_points,
        attack: individualValues.attack,
        defense: individualValues.defense,
        specialAttack: individualValues.special_attack,
        specialDefense: individualValues.special_defense,
        speed: individualValues.speed,
    };
}