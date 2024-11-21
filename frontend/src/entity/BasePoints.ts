export interface BasePoints {
    id: string;
    hitPoints: number;
    attack: number;
    defense: number;
    specialAttack: number;
    specialDefense: number;
    speed: number;
}

export const NewBasePointsFromJson = (basePoints: any): BasePoints => {
    return {
        id: basePoints.id,
        hitPoints: basePoints.hit_points,
        attack: basePoints.attack,
        defense: basePoints.defense,
        specialAttack: basePoints.special_attack,
        specialDefense: basePoints.special_defense,
        speed: basePoints.speed,
    };
}